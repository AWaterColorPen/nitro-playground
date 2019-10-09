// Code generated by protoc-gen-go. DO NOT EDIT.
// source: k8s/proto/akin.proto

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

func init() { proto.RegisterFile("k8s/proto/akin.proto", fileDescriptor_9ea397a364721dfd) }

var fileDescriptor_9ea397a364721dfd = []byte{
	// 104 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc9, 0xb6, 0x28, 0xd6,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4f, 0xcc, 0xce, 0xcc, 0xd3, 0x03, 0x33, 0x85, 0xd8, 0x4a,
	0x52, 0xf3, 0x92, 0x33, 0xf2, 0xa5, 0xc4, 0x10, 0xb2, 0xc9, 0xf9, 0xb9, 0xb9, 0xf9, 0x50, 0x79,
	0x23, 0x03, 0x2e, 0x16, 0xc7, 0xec, 0xcc, 0x3c, 0x21, 0x0d, 0x2e, 0x46, 0x47, 0x21, 0x7e, 0x3d,
	0x88, 0x6a, 0xbd, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x29, 0x01, 0x84, 0x40, 0x71, 0x41,
	0x7e, 0x5e, 0x71, 0xaa, 0x12, 0x43, 0x12, 0x1b, 0x58, 0xa3, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff,
	0x57, 0x92, 0x1c, 0x95, 0x70, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AkinClient is the client API for Akin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AkinClient interface {
	A(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type akinClient struct {
	cc *grpc.ClientConn
}

func NewAkinClient(cc *grpc.ClientConn) AkinClient {
	return &akinClient{cc}
}

func (c *akinClient) A(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/tencho.Akin/A", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AkinServer is the server API for Akin service.
type AkinServer interface {
	A(context.Context, *Request) (*Response, error)
}

// UnimplementedAkinServer can be embedded to have forward compatible implementations.
type UnimplementedAkinServer struct {
}

func (*UnimplementedAkinServer) A(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method A not implemented")
}

func RegisterAkinServer(s *grpc.Server, srv AkinServer) {
	s.RegisterService(&_Akin_serviceDesc, srv)
}

func _Akin_A_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AkinServer).A(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tencho.Akin/A",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AkinServer).A(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Akin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tencho.Akin",
	HandlerType: (*AkinServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "A",
			Handler:    _Akin_A_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "k8s/proto/akin.proto",
}