// Code generated by protoc-gen-go. DO NOT EDIT.
// source: asim/go-micro/examples/greeter/srv/proto/hello/hello.proto

package go_micro_srv_greeter

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

type Request struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9c4d46e1c0507ba, []int{0}
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

type Response struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9c4d46e1c0507ba, []int{1}
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

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.greeter.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.greeter.Response")
}

func init() {
	proto.RegisterFile("asim/go-micro/examples/greeter/srv/proto/hello/hello.proto", fileDescriptor_f9c4d46e1c0507ba)
}

var fileDescriptor_f9c4d46e1c0507ba = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0xca, 0xcd, 0x4c, 0x2e,
	0xca, 0xd7, 0x4f, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0x2d, 0xd6, 0x4f, 0x2f, 0x4a, 0x4d, 0x2d,
	0x49, 0x2d, 0xd2, 0x2f, 0x2e, 0x2a, 0xd3, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0xcf, 0x48, 0xcd,
	0xc9, 0x81, 0x92, 0x7a, 0x60, 0x11, 0x21, 0x91, 0xf4, 0x7c, 0x3d, 0xb0, 0x36, 0xbd, 0xe2, 0xa2,
	0x32, 0x3d, 0xa8, 0x0e, 0x25, 0x59, 0x2e, 0xf6, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21,
	0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b,
	0x49, 0x86, 0x8b, 0x23, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x80, 0x8b, 0x39,
	0xb7, 0x38, 0x1d, 0x2a, 0x0d, 0x62, 0x1a, 0xf9, 0x73, 0x31, 0x07, 0x27, 0x56, 0x0a, 0x79, 0x70,
	0xb1, 0x7a, 0x80, 0x2c, 0x12, 0x92, 0xd5, 0xc3, 0x66, 0x87, 0x1e, 0xd4, 0x02, 0x29, 0x39, 0x5c,
	0xd2, 0x10, 0x0b, 0x94, 0x18, 0x92, 0xd8, 0xc0, 0x4e, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x0c, 0xa9, 0xad, 0x53, 0xe0, 0x00, 0x00, 0x00,
}
