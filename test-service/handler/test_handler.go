package handler

import (
	"comm/logger"
	"comm/mark"
	"context"
	pbProxy "proto/proxy"
	pbTest "proto/test"
)

func (h *Handler) ListWeather(ctx context.Context, req *pbTest.Empty, rsp *pbTest.Empty) error {
	var err error
	var timemark mark.TimeMark
	logger := logger.Extract(ctx)
	defer timemark.Init(ctx, "ListWeather")()
	v, err := h.MeiZuService.ListWeather(ctx, &pbProxy.WeatherFilter{CityIds: "101240101"})
	if err != nil {
		logger.Debugf("ListWeather failed. [%s]", err.Error())
		return err
	}
	logger.Infof("%+v", v)
	return nil
}
