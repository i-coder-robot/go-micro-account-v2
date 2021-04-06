// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/account.proto

package account

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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
	Register(ctx context.Context, in *AccountRequest, opts ...client.CallOption) (*AccountResponse, error)
	Login(ctx context.Context, in *AccountRequest, opts ...client.CallOption) (*AccountResponse, error)
	GetUserInfo(ctx context.Context, in *AccountRequest, opts ...client.CallOption) (*AccountResponse, error)
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

func (c *accountService) Register(ctx context.Context, in *AccountRequest, opts ...client.CallOption) (*AccountResponse, error) {
	req := c.c.NewRequest(c.name, "Account.Register", in)
	out := new(AccountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) Login(ctx context.Context, in *AccountRequest, opts ...client.CallOption) (*AccountResponse, error) {
	req := c.c.NewRequest(c.name, "Account.Login", in)
	out := new(AccountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountService) GetUserInfo(ctx context.Context, in *AccountRequest, opts ...client.CallOption) (*AccountResponse, error) {
	req := c.c.NewRequest(c.name, "Account.GetUserInfo", in)
	out := new(AccountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Account service

type AccountHandler interface {
	Register(context.Context, *AccountRequest, *AccountResponse) error
	Login(context.Context, *AccountRequest, *AccountResponse) error
	GetUserInfo(context.Context, *AccountRequest, *AccountResponse) error
}

func RegisterAccountHandler(s server.Server, hdlr AccountHandler, opts ...server.HandlerOption) error {
	type account interface {
		Register(ctx context.Context, in *AccountRequest, out *AccountResponse) error
		Login(ctx context.Context, in *AccountRequest, out *AccountResponse) error
		GetUserInfo(ctx context.Context, in *AccountRequest, out *AccountResponse) error
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

func (h *accountHandler) Register(ctx context.Context, in *AccountRequest, out *AccountResponse) error {
	return h.AccountHandler.Register(ctx, in, out)
}

func (h *accountHandler) Login(ctx context.Context, in *AccountRequest, out *AccountResponse) error {
	return h.AccountHandler.Login(ctx, in, out)
}

func (h *accountHandler) GetUserInfo(ctx context.Context, in *AccountRequest, out *AccountResponse) error {
	return h.AccountHandler.GetUserInfo(ctx, in, out)
}