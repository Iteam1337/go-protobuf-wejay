// Code generated by protoc-gen-go. DO NOT EDIT.
// source: callbackURL.proto

package message

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

type CallbackURL struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallbackURL) Reset()         { *m = CallbackURL{} }
func (m *CallbackURL) String() string { return proto.CompactTextString(m) }
func (*CallbackURL) ProtoMessage()    {}
func (*CallbackURL) Descriptor() ([]byte, []int) {
	return fileDescriptor_172ffbc68f5a7fd2, []int{0}
}

func (m *CallbackURL) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallbackURL.Unmarshal(m, b)
}
func (m *CallbackURL) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallbackURL.Marshal(b, m, deterministic)
}
func (m *CallbackURL) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallbackURL.Merge(m, src)
}
func (m *CallbackURL) XXX_Size() int {
	return xxx_messageInfo_CallbackURL.Size(m)
}
func (m *CallbackURL) XXX_DiscardUnknown() {
	xxx_messageInfo_CallbackURL.DiscardUnknown(m)
}

var xxx_messageInfo_CallbackURL proto.InternalMessageInfo

func (m *CallbackURL) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type CallbackURLResponse struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Ok                   bool     `protobuf:"varint,3,opt,name=ok,proto3" json:"ok,omitempty"`
	Error                string   `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallbackURLResponse) Reset()         { *m = CallbackURLResponse{} }
func (m *CallbackURLResponse) String() string { return proto.CompactTextString(m) }
func (*CallbackURLResponse) ProtoMessage()    {}
func (*CallbackURLResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_172ffbc68f5a7fd2, []int{1}
}

func (m *CallbackURLResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallbackURLResponse.Unmarshal(m, b)
}
func (m *CallbackURLResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallbackURLResponse.Marshal(b, m, deterministic)
}
func (m *CallbackURLResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallbackURLResponse.Merge(m, src)
}
func (m *CallbackURLResponse) XXX_Size() int {
	return xxx_messageInfo_CallbackURLResponse.Size(m)
}
func (m *CallbackURLResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CallbackURLResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CallbackURLResponse proto.InternalMessageInfo

func (m *CallbackURLResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *CallbackURLResponse) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *CallbackURLResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *CallbackURLResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*CallbackURL)(nil), "message.CallbackURL")
	proto.RegisterType((*CallbackURLResponse)(nil), "message.CallbackURLResponse")
}

func init() {
	proto.RegisterFile("callbackURL.proto", fileDescriptor_172ffbc68f5a7fd2)
}

var fileDescriptor_172ffbc68f5a7fd2 = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x4e, 0xcc, 0xc9,
	0x49, 0x4a, 0x4c, 0xce, 0x0e, 0x0d, 0xf2, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcf,
	0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x55, 0x52, 0xe3, 0xe2, 0x76, 0x46, 0xc8, 0x0a, 0x89, 0x73,
	0xb1, 0x97, 0x16, 0xa7, 0x16, 0xc5, 0x67, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xb1,
	0x81, 0xb8, 0x9e, 0x29, 0x4a, 0x69, 0x5c, 0xc2, 0x48, 0xea, 0x82, 0x52, 0x8b, 0x0b, 0xf2, 0xf3,
	0x8a, 0x53, 0x71, 0xaa, 0x17, 0x12, 0xe0, 0x62, 0x2e, 0x2d, 0xca, 0x91, 0x60, 0x02, 0x0b, 0x82,
	0x98, 0x42, 0x7c, 0x5c, 0x4c, 0xf9, 0xd9, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0x1c, 0x41, 0x4c, 0xf9,
	0xd9, 0x42, 0x22, 0x5c, 0xac, 0xa9, 0x45, 0x45, 0xf9, 0x45, 0x12, 0x2c, 0x60, 0x35, 0x10, 0x4e,
	0x12, 0x1b, 0xd8, 0x7d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x5e, 0xe4, 0xfc, 0xb4,
	0x00, 0x00, 0x00,
}
