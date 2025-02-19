package handler

import (
	"context"
	"fmt"
	"net/http"
)

func (h *Handler) Download(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}
