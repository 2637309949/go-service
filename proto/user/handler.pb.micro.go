// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/user/handler.proto

package user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/2637309949/micro/v3/service/api"
	client "github.com/2637309949/micro/v3/service/client"
	server "github.com/2637309949/micro/v3/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Account service

func NewAccountEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Account service

type AccountService interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
	UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...client.CallOption) (*UpdatePasswordResponse, error)
	VerifyEmail(ctx context.Context, in *VerifyEmailRequest, opts ...client.CallOption) (*VerifyEmailResponse, error)
	SendVerificationEmail(ctx context.Context, in *SendVerificationEmailRequest, opts ...client.CallOption) (*SendVerificationEmailResponse, error)
	SendPasswordResetEmail(ctx context.Context, in *SendPasswordResetEmailRequest, opts ...client.CallOption) (*SendPasswordResetEmailResponse, error)
	ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...client.CallOption) (*ResetPasswordResponse, error)
	ValidPassword(ctx context.Context, in *ValidPasswordRequest, opts ...client.CallOption) (*ValidPasswordResponse, error)
}

type accountService struct {
	c    client.Client
	name string
}

func NewAccountService(name string, c client.Client) AccountService {
	return &accountService{
		c:    c,
		name: name,
	}
}

func (c *accountService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.name, "Account.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.name, "Account.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.name, "Account.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "Account.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Account.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, opts ...client.CallOption) (*UpdatePasswordResponse, error) {
	req := c.c.NewRequest(c.name, "Account.UpdatePassword", in)
	out := new(UpdatePasswordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) VerifyEmail(ctx context.Context, in *VerifyEmailRequest, opts ...client.CallOption) (*VerifyEmailResponse, error) {
	req := c.c.NewRequest(c.name, "Account.VerifyEmail", in)
	out := new(VerifyEmailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) SendVerificationEmail(ctx context.Context, in *SendVerificationEmailRequest, opts ...client.CallOption) (*SendVerificationEmailResponse, error) {
	req := c.c.NewRequest(c.name, "Account.SendVerificationEmail", in)
	out := new(SendVerificationEmailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) SendPasswordResetEmail(ctx context.Context, in *SendPasswordResetEmailRequest, opts ...client.CallOption) (*SendPasswordResetEmailResponse, error) {
	req := c.c.NewRequest(c.name, "Account.SendPasswordResetEmail", in)
	out := new(SendPasswordResetEmailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...client.CallOption) (*ResetPasswordResponse, error) {
	req := c.c.NewRequest(c.name, "Account.ResetPassword", in)
	out := new(ResetPasswordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) ValidPassword(ctx context.Context, in *ValidPasswordRequest, opts ...client.CallOption) (*ValidPasswordResponse, error) {
	req := c.c.NewRequest(c.name, "Account.ValidPassword", in)
	out := new(ValidPasswordResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Account service

type AccountHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Read(context.Context, *ReadRequest, *ReadResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	List(context.Context, *ListRequest, *ListResponse) error
	UpdatePassword(context.Context, *UpdatePasswordRequest, *UpdatePasswordResponse) error
	VerifyEmail(context.Context, *VerifyEmailRequest, *VerifyEmailResponse) error
	SendVerificationEmail(context.Context, *SendVerificationEmailRequest, *SendVerificationEmailResponse) error
	SendPasswordResetEmail(context.Context, *SendPasswordResetEmailRequest, *SendPasswordResetEmailResponse) error
	ResetPassword(context.Context, *ResetPasswordRequest, *ResetPasswordResponse) error
	ValidPassword(context.Context, *ValidPasswordRequest, *ValidPasswordResponse) error
}

func RegisterAccountHandler(s server.Server, hdlr AccountHandler, opts ...server.HandlerOption) error {
	type account interface {
		Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error
		Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error
		Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
		List(ctx context.Context, in *ListRequest, out *ListResponse) error
		UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, out *UpdatePasswordResponse) error
		VerifyEmail(ctx context.Context, in *VerifyEmailRequest, out *VerifyEmailResponse) error
		SendVerificationEmail(ctx context.Context, in *SendVerificationEmailRequest, out *SendVerificationEmailResponse) error
		SendPasswordResetEmail(ctx context.Context, in *SendPasswordResetEmailRequest, out *SendPasswordResetEmailResponse) error
		ResetPassword(ctx context.Context, in *ResetPasswordRequest, out *ResetPasswordResponse) error
		ValidPassword(ctx context.Context, in *ValidPasswordRequest, out *ValidPasswordResponse) error
	}
	type Account struct {
		account
	}
	h := &accountHandler{hdlr}
	return s.Handle(s.NewHandler(&Account{h}, opts...))
}

type accountHandler struct {
	AccountHandler
}

func (h *accountHandler) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.AccountHandler.Create(ctx, in, out)
}

func (h *accountHandler) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.AccountHandler.Read(ctx, in, out)
}

func (h *accountHandler) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.AccountHandler.Update(ctx, in, out)
}

func (h *accountHandler) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.AccountHandler.Delete(ctx, in, out)
}

func (h *accountHandler) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.AccountHandler.List(ctx, in, out)
}

func (h *accountHandler) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest, out *UpdatePasswordResponse) error {
	return h.AccountHandler.UpdatePassword(ctx, in, out)
}

func (h *accountHandler) VerifyEmail(ctx context.Context, in *VerifyEmailRequest, out *VerifyEmailResponse) error {
	return h.AccountHandler.VerifyEmail(ctx, in, out)
}

func (h *accountHandler) SendVerificationEmail(ctx context.Context, in *SendVerificationEmailRequest, out *SendVerificationEmailResponse) error {
	return h.AccountHandler.SendVerificationEmail(ctx, in, out)
}

func (h *accountHandler) SendPasswordResetEmail(ctx context.Context, in *SendPasswordResetEmailRequest, out *SendPasswordResetEmailResponse) error {
	return h.AccountHandler.SendPasswordResetEmail(ctx, in, out)
}

func (h *accountHandler) ResetPassword(ctx context.Context, in *ResetPasswordRequest, out *ResetPasswordResponse) error {
	return h.AccountHandler.ResetPassword(ctx, in, out)
}

func (h *accountHandler) ValidPassword(ctx context.Context, in *ValidPasswordRequest, out *ValidPasswordResponse) error {
	return h.AccountHandler.ValidPassword(ctx, in, out)
}
