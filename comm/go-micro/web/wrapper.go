package web

import "net/http"

type HandlerFunc func(http.ResponseWriter, *http.Request)

// HandlerWrapper wraps the HandlerFunc and returns the equivalent.
type HandlerWrapper func(HandlerFunc) HandlerFunc
