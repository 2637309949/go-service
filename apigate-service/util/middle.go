package util

import (
	"comm/service/web"
	"net/http"
)

func setHeader(w http.ResponseWriter, k, v string) {
	if v := w.Header().Get(k); len(v) > 0 {
		return
	}
	w.Header().Set(k, v)
}

func WrapCors() web.HandlerWrapper {
	return func(h web.HandlerFunc) web.HandlerFunc {
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
}
