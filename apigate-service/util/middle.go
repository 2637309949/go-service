package util

import (
	"comm/service/web"
	"net/http"
)

func WrapCors() web.HandlerWrapper {
	return func(h web.HandlerFunc) web.HandlerFunc {
		set := func(w http.ResponseWriter, k, v string) {
			if v := w.Header().Get(k); len(v) > 0 {
				return
			}
			w.Header().Set(k, v)
		}
		return func(w http.ResponseWriter, r *http.Request) {
			if origin := r.Header.Get("Origin"); len(origin) > 0 {
				set(w, "Access-Control-Allow-Origin", origin)
			} else {
				set(w, "Access-Control-Allow-Origin", "*")
			}
			set(w, "Access-Control-Allow-Credentials", "true")
			set(w, "Access-Control-Allow-Methods", "POST, PATCH, GET, OPTIONS, PUT, DELETE")
			set(w, "Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Micro-Namespace")

			if r.Method == "OPTIONS" {
				return
			}
			h(w, r)
		}
	}
}
