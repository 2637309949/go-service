package handler

import (
	pbCache "proto/cache"
	pbProxy "proto/proxy"
)

type Handler struct {
	CacheService pbCache.CacheService
	MeiZuService pbProxy.MeiZuService
}
