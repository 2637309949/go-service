package web

import (
	"comm/logger"
	"context"
	"net/http"
	"net/textproto"
	"strings"

	trace "github.com/micro/plugins/v5/wrapper/trace/opentracing"

	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v5/metadata"
	"go-micro.dev/v5/web"
)

func FromRequest(r *http.Request) context.Context {
	ctx := r.Context()
	md, ok := metadata.FromContext(ctx)
	if !ok {
		md = make(metadata.Metadata)
	}
	for k, v := range r.Header {
		md[textproto.CanonicalMIMEHeaderKey(k)] = strings.Join(v, ",")
	}
	// pass http host
	md["Host"] = r.Host
	// pass http method
	md["Method"] = r.Method
	if r.URL != nil {
		md["URL"] = r.URL.String()
	}
	return metadata.NewContext(ctx, md)
}

func loggerHandler(h web.HandlerFunc) web.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cx := FromRequest(r)
		logger := logger.Extract(cx)
		fields := map[string]interface{}{}
		if md, ok := metadata.FromContext(cx); ok {
			if v, ok := md["Method"]; ok {
				fields["method"] = v
			}
			if v, ok := md["URL"]; ok {
				fields["url"] = v
			}
		}
		logger.WithFields(fields).Info("")
		h(w, r)
	}
}

// NewTracerWrapper accepts an opentracing Tracer and returns a Handler Wrapper.
func NewTracerWrapper(ot opentracing.Tracer) web.HandlerWrapper {
	return func(h web.HandlerFunc) web.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if ot == nil {
				ot = opentracing.GlobalTracer()
			}
			cx := FromRequest(r)
			logger := logger.Extract(cx)
			name := ""
			md, ok := metadata.FromContext(cx)
			if ok {
				name = md["URL"]
			}
			ctx, span, err := trace.StartSpanFromContext(cx, ot, name)
			if err != nil {
				logger.Error(err)
				return
			}
			*r = *r.Clone(ctx)
			defer span.Finish()
			h(w, r)
		}
	}
}

func loggerWrapper(l logger.Logger) web.HandlerWrapper {
	if l == nil {
		l = logger.DefaultLogger
	}
	return func(h web.HandlerFunc) web.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			cx := FromRequest(r)
			_, ok := logger.FromContext(cx)
			if !ok {
				if md, ok := metadata.FromContext(cx); ok {
					if traceID, exists := md["Uber-Trace-Id"]; exists {
						traceID = strings.Split(traceID, ":")[0]
						w.Header().Set("Uber-Trace-Id", traceID)
						l = l.Fields(map[string]interface{}{"traceid": traceID})
						cx = logger.NewContext(cx, l)
					}
				}
			}
			*r = *r.Clone(cx)
			h(w, r)
		}
	}
}
