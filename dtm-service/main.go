package main

import (
	"comm/service"
	"dtm/handler"
	pbDtm "proto/dtm"
)

func main() {
	service := service.NewService(
		service.Name("dtm"),
	)
	h := handler.Handler{}
	pbDtm.RegisterDtmServiceHandler(service.Server(), &h)
	pbDtm.RegisterWorkflowServiceHandler(service.Server(), &h)
	service.Run()
}
