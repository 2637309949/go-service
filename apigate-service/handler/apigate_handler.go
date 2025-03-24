package handler

import (
	"context"
	"net/http"
)

// Home 处理函数用于响应主页请求
func (a *h) Home(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"version": "5+"}`))
}

// Favicon 处理网站的 favicon 请求。
// 该方法响应了对网站图标（favicon）的请求，通过写入空的字节数据到响应体中。
// 主要目的是避免在访问网站时出现 404 错误，当 favicon 请求到来时，返回一个空的响应体。
func (a *h) Favicon(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(``))
}
