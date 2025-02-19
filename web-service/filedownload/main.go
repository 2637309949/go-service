package main

import (
	"comm/service/web"
	"filedownload/handler"
	pb "proto/web/filedownload"
)

func main() {
	service := web.NewService(
		web.Name("filedownload"),
	)
	pb.RegisterFileDownloadServiceHandler(service, handler.NewHandler())
	service.Run()
}
