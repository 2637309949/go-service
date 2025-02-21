package handler

import (
	"comm/logger"
	"comm/mark"
	"context"
	"net/http"
)

func (h *Handler) Activitys(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var timemark mark.TimeMark
	logger := logger.Extract(ctx)
	defer timemark.Init(ctx, "Activitys")()
	timemark.Mark("Activitys")
	logger.Info("Activitys")
}
