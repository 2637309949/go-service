package handler

import (
	"context"
	pbUser "proto/user"
)

func (h *Handler) Hello(ctx context.Context, req *pbUser.Request, rsp *pbUser.Response) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}