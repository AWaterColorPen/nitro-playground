// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/tosui.proto

package tencho

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

func init() { proto.RegisterFile("proto/tosui.proto", fileDescriptor_581b90e70cc61f05) }

var fileDescriptor_581b90e70cc61f05 = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8d, 0xb1, 0x0a, 0xc2, 0x30,
	0x18, 0x84, 0x1d, 0xb4, 0x43, 0x17, 0x35, 0x63, 0x47, 0x57, 0x31, 0x01, 0x2d, 0xee, 0xda, 0x17,
	0x10, 0x15, 0x04, 0x17, 0x69, 0xe3, 0x4f, 0x1b, 0x48, 0x72, 0x35, 0xf9, 0x33, 0xf8, 0xf6, 0x42,
	0x33, 0xb8, 0x1d, 0xf7, 0x1d, 0xf7, 0x95, 0xeb, 0x31, 0x80, 0xa1, 0x18, 0x31, 0x19, 0x39, 0x65,
	0x51, 0x30, 0x79, 0x3d, 0xa0, 0x12, 0x19, 0x69, 0x38, 0x07, 0x9f, 0xd9, 0xbe, 0x2e, 0x17, 0x77,
	0xdc, 0x92, 0x11, 0xdb, 0x72, 0xde, 0xb4, 0xd6, 0x8a, 0xa5, 0xcc, 0x6b, 0x79, 0xa5, 0x4f, 0xa2,
	0xc8, 0xd5, 0xea, 0x5f, 0xc4, 0x11, 0x3e, 0xd2, 0x66, 0x76, 0x3e, 0x3e, 0xeb, 0xde, 0xf0, 0x90,
	0x3a, 0xa9, 0xe1, 0xd4, 0xe9, 0xd1, 0x32, 0x85, 0x06, 0x16, 0xe1, 0x42, 0x5e, 0xf5, 0xd8, 0x39,
	0xa3, 0x03, 0x5e, 0xa3, 0x6d, 0xbf, 0x7d, 0x40, 0xf2, 0x6f, 0x95, 0x2f, 0xba, 0x62, 0x92, 0x1e,
	0x7e, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x8f, 0x9b, 0xac, 0xa5, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ToSuiClient is the client API for ToSui service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ToSuiClient interface {
	Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type toSuiClient struct {
	cc *grpc.ClientConn
}

func NewToSuiClient(cc *grpc.ClientConn) ToSuiClient {
	return &toSuiClient{cc}
}

func (c *toSuiClient) Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/tencho.ToSui/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ToSuiServer is the server API for ToSui service.
type ToSuiServer interface {
	Call(context.Context, *Request) (*Response, error)
}

// UnimplementedToSuiServer can be embedded to have forward compatible implementations.
type UnimplementedToSuiServer struct {
}

func (*UnimplementedToSuiServer) Call(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}

func RegisterToSuiServer(s *grpc.Server, srv ToSuiServer) {
	s.RegisterService(&_ToSui_serviceDesc, srv)
}

func _ToSui_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ToSuiServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tencho.ToSui/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ToSuiServer).Call(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _ToSui_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tencho.ToSui",
	HandlerType: (*ToSuiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _ToSui_Call_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/tosui.proto",
}
