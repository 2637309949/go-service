package handler

import (
	"comm/logger"
	"comm/mark"
	"context"
	pbCache "proto/cache"
	"time"
)

func (h *Handler) Get(ctx context.Context, req *pbCache.CacheFilter, rsp *pbCache.Cache) error {
	var timemark mark.TimeMark
	logger := logger.Extract(ctx)
	defer timemark.Init(ctx, "Get")()
	logger.Infof("Received Handler.Get request: %+v", req)
	v, e, err := h.cache.Get(ctx, req.Key)
	if err != nil {
		return err
	}
	rsp.Value = v
	rsp.Expiration = e.Unix()
	rsp.Duration = time.Until(e).String()

	return nil
}

func (h *Handler) Put(ctx context.Context, req *pbCache.Cache, rsp *pbCache.Cache) error {
	var timemark mark.TimeMark
	logger := logger.Extract(ctx)
	defer timemark.Init(ctx, "Put")()
	logger.Infof("Received Handler.Put request: %+v", req)
	d, err := time.ParseDuration(req.Duration)
	if err != nil {
		return err
	}
	if err := h.cache.Put(ctx, req.Key, req.Value, d); err != nil {
		return err
	}
	rsp.Key = req.Key
	rsp.Value = req.Value
	rsp.Expiration = req.Expiration
	rsp.Duration = req.Duration

	return nil
}

func (h *Handler) Delete(ctx context.Context, req *pbCache.CacheFilter, rsp *pbCache.Cache) error {
	var timemark mark.TimeMark
	logger := logger.Extract(ctx)
	defer timemark.Init(ctx, "Delete")()
	logger.Infof("Received Handler.Delete request: %+v", req)
	if err := h.cache.Delete(ctx, req.Key); err != nil {
		return err
	}

	return nil
}
