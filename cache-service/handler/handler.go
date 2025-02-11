package handler

import "cache/models/api/cache"

type Handler struct {
	cache cache.Cache
}

func NewHandler(opts ...cache.Option) *Handler {
	c := cache.NewCache(opts...)
	return &Handler{c}
}
