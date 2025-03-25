package rpc

import (
	"apigate/router"
	"apigate/handler"
	"apigate/util"
	bts "bytes"
	"encoding/json"
	"net/http"
	"strings"

	"go-micro.dev/v5/client"
	"go-micro.dev/v5/codec/bytes"
	"go-micro.dev/v5/errors"
	"go-micro.dev/v5/metadata"
)

var (
	// supported json codecs
	jsonCodecs = []string{
		"application/grpc+json",
		"application/json",
		"application/json-rpc",
	}

	// support proto codecs
	protoCodecs = []string{
		"application/grpc",
		"application/grpc+proto",
		"application/proto",
		"application/protobuf",
		"application/proto-rpc",
		"application/octet-stream",
	}
)

// see https://stackoverflow.com/questions/28595664/how-to-stop-json-marshal-from-escaping-and/28596225
func jsonMarshal(t interface{}) ([]byte, error) {
	buffer := &bts.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return bts.TrimRight(buffer.Bytes(), "\n"), err
}

type rpcHandler struct {
	opts handler.Options
}

func (h *rpcHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	bsize := handler.DefaultMaxRecvSize
	if h.opts.MaxRecvSize > 0 {
		bsize = h.opts.MaxRecvSize
	}

	r.Body = http.MaxBytesReader(w, r.Body, bsize)

	// create context
	cx := util.FromRequest(r)
	// strip headers grpc doesn't like
	md, _ := metadata.FromContext(cx)
	// delete websocket info
	delete(md, "Connection")
	cx = metadata.NewContext(cx, md)

	// set merged context to request
	*r = *r.Clone(cx)

	defer r.Body.Close()

	var service *router.Service
	if h.opts.Service != nil {
		service = h.opts.Service
	} else if h.opts.Router != nil {
		s, err := h.opts.Router.Route(h.opts.ApiBase, r)
		if err != nil {
			util.WriteError(w, r, errors.InternalServerError("go.micro.api", "%s", err.Error()))
			return
		}
		service = s
	} else {
		// we have no way of routing the request
		util.WriteError(w, r, errors.InternalServerError("go.micro.api", "no route found"))
		return
	}

	c := h.opts.Client
	ct := r.Header.Get("Content-Type")

	// Strip charset from Content-Type (like `application/json; charset=UTF-8`)
	if idx := strings.IndexRune(ct, ';'); idx >= 0 {
		ct = ct[:idx]
	}

	// create custom router
	var nodes []string
	for _, service := range service.Services {
		for _, node := range service.Nodes {
			nodes = append(nodes, node.Address)
		}
	}

	callOpt := []client.CallOption{}
	callOpt = append(callOpt, client.WithAddress(nodes...))

	// walk the standard call path
	// get payload
	br, err := util.RequestPayload(r)
	if err != nil {
		util.WriteError(w, r, err)
		return
	}

	var rsp []byte

	switch {
	// proto codecs
	case hasCodec(ct, protoCodecs):
		var request *bytes.Frame
		// if the extracted payload isn't empty lets use it
		if len(br) > 0 {
			request = &bytes.Frame{Data: br}
		}

		// create the request
		req := c.NewRequest(
			service.Name,
			service.Endpoint.Name,
			request,
			client.WithContentType(ct),
		)

		// make the call
		var response bytes.Frame
		if err := c.Call(cx, req, &response, callOpt...); err != nil {
			util.WriteError(w, r, err)
			return
		}
		rsp = response.Data
	default:
		// if json codec is not present set to json
		if !hasCodec(ct, jsonCodecs) {
			ct = "application/json"
		}

		// default to trying json
		var request json.RawMessage
		// if the extracted payload isn't empty lets use it
		if len(br) > 0 {
			request = json.RawMessage(br)
		}

		// create request/response
		var response map[string]interface{}

		req := c.NewRequest(
			service.Name,
			service.Endpoint.Name,
			&request,
			client.WithContentType(ct),
		)
		// make the call
		if err := c.Call(cx, req, &response, callOpt...); err != nil {
			util.WriteError(w, r, err)
			return
		}

		// marshall response
		// see https://play.golang.org/p/oBNxUjVTzus
		rsp, err = jsonMarshal(response)
		if err != nil {
			util.WriteError(w, r, err)
			return
		}
	}

	// write the response
	util.WriteResponse(w, r, rsp, ct)
}

func (rh *rpcHandler) String() string {
	return "rpc"
}

func hasCodec(ct string, codecs []string) bool {
	for _, codec := range codecs {
		if ct == codec {
			return true
		}
	}
	return false
}

func NewHandler(opts ...handler.Option) handler.Handler {
	options := handler.NewOptions(opts...)
	return &rpcHandler{
		opts: options,
	}
}
