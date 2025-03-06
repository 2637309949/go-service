package handler

import (
	"net/http"
	"context"
)

func (a *h) Home(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"version": "5+"}`))
}

func (a *h) Favicon(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(``))
}
