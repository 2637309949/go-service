package main

import (
	"comm/service/web"
	"download/handler"
	pbDownload "proto/web/download"
)

func main() {
	service := web.NewService(
		web.Name("download"),
	)
	pbDownload.RegisterDownloadServiceHandler(service, handler.NewHandler())
	service.Run()
}
