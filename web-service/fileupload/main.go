package main

import (
	"comm/service/web"
	"fileupload/handler"
	pb "proto/web/fileupload"
)

func main() {
	service := web.NewService(
		web.Name("fileupload"),
	)
	pb.RegisterFileUploadServiceHandler(service, handler.NewHandler())
	service.Run()
}
