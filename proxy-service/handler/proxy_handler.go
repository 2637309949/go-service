package handler

import (
	"context"
	pbProxy "proto/proxy"
)

func (h *Handler) Home(ctx context.Context, req *pbProxy.MicroFilter, rsp *pbProxy.Micro) error {
	return nil
}

func (h *Handler) ListWeather(ctx context.Context, req *pbProxy.WeatherFilter, rsp *pbProxy.WeatherList) error {
	return nil
}
