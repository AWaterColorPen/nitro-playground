// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/common/common.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type SERVICE_NAME int32

const (
	SERVICE_NAME_UNKNOWN SERVICE_NAME = 0
	SERVICE_NAME_AKIN    SERVICE_NAME = 1
	SERVICE_NAME_TOSUI   SERVICE_NAME = 2
)

var SERVICE_NAME_name = map[int32]string{
	0: "UNKNOWN",
	1: "AKIN",
	2: "TOSUI",
}

var SERVICE_NAME_value = map[string]int32{
	"UNKNOWN": 0,
	"AKIN":    1,
	"TOSUI":   2,
}

func (x SERVICE_NAME) String() string {
	return proto.EnumName(SERVICE_NAME_name, int32(x))
}

func (SERVICE_NAME) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_856864ca6250a18a, []int{0}
}

type Request struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Query                string   `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_856864ca6250a18a, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Request) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type Response struct {
	Code                 uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_856864ca6250a18a, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterEnum("common.SERVICE_NAME", SERVICE_NAME_name, SERVICE_NAME_value)
	proto.RegisterType((*Request)(nil), "common.Request")
	proto.RegisterType((*Response)(nil), "common.Response")
}

func init() { proto.RegisterFile("proto/common/common.proto", fileDescriptor_856864ca6250a18a) }

var fileDescriptor_856864ca6250a18a = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0x83, 0x52, 0x7a, 0x60, 0x31, 0x21, 0x36, 0x08,
	0x4f, 0xc9, 0x98, 0x8b, 0x3d, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x88, 0x8b, 0x25,
	0x2f, 0x31, 0x37, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x16, 0x12, 0xe1, 0x62,
	0x2d, 0x2c, 0x4d, 0x2d, 0xaa, 0x94, 0x60, 0x02, 0x0b, 0x42, 0x38, 0x4a, 0x72, 0x5c, 0x1c, 0x41,
	0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x20, 0x5d, 0xc9, 0xf9, 0x29, 0x10, 0x5d, 0xbc, 0x41,
	0x60, 0xb6, 0x96, 0x01, 0x17, 0x4f, 0xb0, 0x6b, 0x50, 0x98, 0xa7, 0xb3, 0x6b, 0xbc, 0x9f, 0xa3,
	0xaf, 0xab, 0x10, 0x37, 0x17, 0x7b, 0xa8, 0x9f, 0xb7, 0x9f, 0x7f, 0xb8, 0x9f, 0x00, 0x83, 0x10,
	0x07, 0x17, 0x8b, 0xa3, 0xb7, 0xa7, 0x9f, 0x00, 0xa3, 0x10, 0x27, 0x17, 0x6b, 0x88, 0x7f, 0x70,
	0xa8, 0xa7, 0x00, 0x93, 0x93, 0x65, 0x94, 0x79, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72,
	0x7e, 0xae, 0x7e, 0x62, 0x79, 0x62, 0x49, 0x6a, 0x51, 0x72, 0x7e, 0x4e, 0x7e, 0x51, 0x41, 0x6a,
	0x9e, 0x7e, 0x5e, 0x66, 0x49, 0x51, 0xbe, 0x6e, 0x41, 0x4e, 0x62, 0x65, 0x7a, 0x51, 0x7e, 0x69,
	0x5e, 0x8a, 0x3e, 0xb2, 0xb7, 0x92, 0xd8, 0xc0, 0x3c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x62, 0x80, 0x7d, 0xeb, 0xed, 0x00, 0x00, 0x00,
}
