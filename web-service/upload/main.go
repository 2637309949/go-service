package main

import (
	"comm/service/web"
	pbUpload "proto/web/upload"
	"upload/handler"
)

func main() {
	service := web.NewService(
		web.Name("upload"),
	)
	pbUpload.RegisterUploadServiceHandler(service, handler.NewHandler())
	service.Run()
}
