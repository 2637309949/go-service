package web

import (
	"comm/logger"
	"net/http"
	"strings"

	trace "github.com/micro/plugins/v5/wrapper/trace/opentracing"
	uhttp "go-micro.dev/v5/util/http"

	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v5/metadata"
	"go-micro.dev/v5/web"
)

func loggerHandler(h web.HandlerFunc) web.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cx := uhttp.FromRequest(r)
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
			cx := uhttp.FromRequest(r)
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
			cx := uhttp.FromRequest(r)
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

func corsWrapper(h web.HandlerFunc) web.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); len(origin) > 0 {
			setHeader(w, "Access-Control-Allow-Origin", origin)
		} else {
			setHeader(w, "Access-Control-Allow-Origin", "*")
		}
		setHeader(w, "Access-Control-Allow-Credentials", "true")
		setHeader(w, "Access-Control-Allow-Methods", "POST, PATCH, GET, OPTIONS, PUT, DELETE")
		setHeader(w, "Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Micro-Namespace")

		if r.Method == "OPTIONS" {
			return
		}

		h(w, r)
	}
}

func setHeader(w http.ResponseWriter, k, v string) {
	if v := w.Header().Get(k); len(v) > 0 {
		return
	}
	w.Header().Set(k, v)
}
