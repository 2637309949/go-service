package handler

import (
	"context"
	pbCache "proto/cache"
)

func (h *Handler) Get(ctx context.Context, req *pbCache.CacheFilter, rsp *pbCache.Cache) error {
	return nil
}

func (h *Handler) Put(ctx context.Context, req *pbCache.Cache, rsp *pbCache.Cache) error {
	return nil
}

func (h *Handler) Delete(ctx context.Context, req *pbCache.CacheFilter, rsp *pbCache.Cache) error {
	return nil
}
