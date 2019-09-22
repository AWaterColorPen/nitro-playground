// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: k8s/proto/shunmu.proto

package tencho

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ShunMu service

type ShunMuService interface {
	A(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type shunMuService struct {
	c    client.Client
	name string
}

func NewShunMuService(name string, c client.Client) ShunMuService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "tencho"
	}
	return &shunMuService{
		c:    c,
		name: name,
	}
}

func (c *shunMuService) A(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ShunMu.A", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ShunMu service

type ShunMuHandler interface {
	A(context.Context, *Request, *Response) error
}

func RegisterShunMuHandler(s server.Server, hdlr ShunMuHandler, opts ...server.HandlerOption) error {
	type shunMu interface {
		A(ctx context.Context, in *Request, out *Response) error
	}
	type ShunMu struct {
		shunMu
	}
	h := &shunMuHandler{hdlr}
	return s.Handle(s.NewHandler(&ShunMu{h}, opts...))
}

type shunMuHandler struct {
	ShunMuHandler
}

func (h *shunMuHandler) A(ctx context.Context, in *Request, out *Response) error {
	return h.ShunMuHandler.A(ctx, in, out)
}
