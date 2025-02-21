package handler

import (
	"comm/logger"
	"comm/mark"
	"context"
	"net/http"
)

func (h *Handler) Download(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var timemark mark.TimeMark
	logger := logger.Extract(ctx)
	defer timemark.Init(ctx, "Download")()
	timemark.Mark("Download")
	logger.Info("Download")
	w.Write([]byte(`{"hello": "world"}`))
}
