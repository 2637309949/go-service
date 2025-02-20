package main

import (
	"apigate/api"
	"apigate/handler"
	"apigate/handler/rpc"
	"apigate/handler/web"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"go-micro.dev/v5/errors"
	"go-micro.dev/v5/logger"
	"go-micro.dev/v5/metadata"
)

var (
	defaultHandler = "rpc"
)

type apiHandler struct {
	opts handler.Options
}

// API handler is the default handler which takes api.Request and returns api.Response
func (a *apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var service *api.Service
	if a.opts.Service != nil {
		service = a.opts.Service
	} else if a.opts.Router != nil {
		s, err := a.opts.Router.Route(a.opts.ApiBase, r)
		if err != nil {
			writeError(w, r, errors.InternalServerError("go.micro.api", "%s", err.Error()))
			return
		}
		service = s
	} else {
		// we have no way of routing the request
		writeError(w, r, errors.InternalServerError("go.micro.api", "no route found"))
		return
	}

	md, _ := metadata.FromContext(r.Context())
	fmt.Printf("-----------%+v\n", md)
	handler := defaultHandler
	if len(service.Endpoint.Handler) > 0 {
		handler = service.Endpoint.Handler
	}

	if h, ok := a.opts.Handlers[handler]; ok {
		h.ServeHTTP(w, r)
		return
	}

	writeError(w, r, errors.InternalServerError("go.micro.api", "no handler found"))
}

func (a *apiHandler) String() string {
	return "api"
}

func writeError(w http.ResponseWriter, r *http.Request, err error) {
	// response content type
	w.Header().Set("Content-Type", "application/json")
	traceID := ""
	if md, ok := metadata.FromContext(r.Context()); ok {
		traceID = md["Uber-Trace-Id"]
		traceID = strings.Split(traceID, ":")[0]
	}
	// parse out the error code
	ce := errors.Parse(err.Error())
	switch ce.Code {
	case 0:
		ce.Code = 500
		ce.Detail = "error during request: " + ce.Detail
	}

	// Set trailers
	if strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
		w.Header().Set("Trailer", "grpc-status")
		w.Header().Set("Trailer", "grpc-message")
		w.Header().Set("grpc-status", "13")
		w.Header().Set("grpc-message", ce.Detail)
	}

	var data map[string]interface{}
	json.Unmarshal([]byte(ce.Error()), &data)
	delete(data, "id")
	delete(data, "status")
	if len(traceID) > 0 {
		data["request_id"] = traceID
	}
	updated, err := json.Marshal(data)
	if err != nil {
		if logger.V(logger.ErrorLevel, logger.DefaultLogger) {
			logger.Error(err)
		}
	}
	w.WriteHeader(200)
	_, err = w.Write(updated)
	if err != nil {
		if logger.V(logger.ErrorLevel, logger.DefaultLogger) {
			logger.Error(err)
		}
	}
}

func NewHandler(opts ...handler.Option) handler.Handler {
	opts = append(opts, handler.WithHandler(rpc.NewHandler(opts...)))
	opts = append(opts, handler.WithHandler(web.NewHandler(opts...)))
	options := handler.NewOptions(opts...)
	return &apiHandler{
		opts: options,
	}
}
