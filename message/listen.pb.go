// Code generated by protoc-gen-go. DO NOT EDIT.
// source: listen.proto

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

type ListenResponse_Playback int32

const (
	ListenResponse_LISTENING ListenResponse_Playback = 0
	ListenResponse_ACTIVE    ListenResponse_Playback = 1
	ListenResponse_PROGRESS  ListenResponse_Playback = 2
	ListenResponse_CURRENT   ListenResponse_Playback = 3
)

var ListenResponse_Playback_name = map[int32]string{
	0: "LISTENING",
	1: "ACTIVE",
	2: "PROGRESS",
	3: "CURRENT",
}

var ListenResponse_Playback_value = map[string]int32{
	"LISTENING": 0,
	"ACTIVE":    1,
	"PROGRESS":  2,
	"CURRENT":   3,
}

func (x ListenResponse_Playback) String() string {
	return proto.EnumName(ListenResponse_Playback_name, int32(x))
}

func (ListenResponse_Playback) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f05da38a4a3e5177, []int{1, 0}
}

type ListenResponse_Change int32

const (
	ListenResponse_ACTION   ListenResponse_Change = 0
	ListenResponse_PLAYBACK ListenResponse_Change = 1
)

var ListenResponse_Change_name = map[int32]string{
	0: "ACTION",
	1: "PLAYBACK",
}

var ListenResponse_Change_value = map[string]int32{
	"ACTION":   0,
	"PLAYBACK": 1,
}

func (x ListenResponse_Change) String() string {
	return proto.EnumName(ListenResponse_Change_name, int32(x))
}

func (ListenResponse_Change) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f05da38a4a3e5177, []int{1, 1}
}

type Listen struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Listen) Reset()         { *m = Listen{} }
func (m *Listen) String() string { return proto.CompactTextString(m) }
func (*Listen) ProtoMessage()    {}
func (*Listen) Descriptor() ([]byte, []int) {
	return fileDescriptor_f05da38a4a3e5177, []int{0}
}

func (m *Listen) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Listen.Unmarshal(m, b)
}
func (m *Listen) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Listen.Marshal(b, m, deterministic)
}
func (m *Listen) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Listen.Merge(m, src)
}
func (m *Listen) XXX_Size() int {
	return xxx_messageInfo_Listen.Size(m)
}
func (m *Listen) XXX_DiscardUnknown() {
	xxx_messageInfo_Listen.DiscardUnknown(m)
}

var xxx_messageInfo_Listen proto.InternalMessageInfo

func (m *Listen) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type ListenResponse struct {
	UserId               string                `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Change               ListenResponse_Change `protobuf:"varint,2,opt,name=change,proto3,enum=message.ListenResponse_Change" json:"change,omitempty"`
	Meta                 []byte                `protobuf:"bytes,3,opt,name=meta,proto3" json:"meta,omitempty"`
	Ok                   bool                  `protobuf:"varint,4,opt,name=ok,proto3" json:"ok,omitempty"`
	Error                string                `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListenResponse) Reset()         { *m = ListenResponse{} }
func (m *ListenResponse) String() string { return proto.CompactTextString(m) }
func (*ListenResponse) ProtoMessage()    {}
func (*ListenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f05da38a4a3e5177, []int{1}
}

func (m *ListenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenResponse.Unmarshal(m, b)
}
func (m *ListenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenResponse.Marshal(b, m, deterministic)
}
func (m *ListenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenResponse.Merge(m, src)
}
func (m *ListenResponse) XXX_Size() int {
	return xxx_messageInfo_ListenResponse.Size(m)
}
func (m *ListenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListenResponse proto.InternalMessageInfo

func (m *ListenResponse) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ListenResponse) GetChange() ListenResponse_Change {
	if m != nil {
		return m.Change
	}
	return ListenResponse_ACTION
}

func (m *ListenResponse) GetMeta() []byte {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *ListenResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *ListenResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterEnum("message.ListenResponse_Playback", ListenResponse_Playback_name, ListenResponse_Playback_value)
	proto.RegisterEnum("message.ListenResponse_Change", ListenResponse_Change_name, ListenResponse_Change_value)
	proto.RegisterType((*Listen)(nil), "message.Listen")
	proto.RegisterType((*ListenResponse)(nil), "message.ListenResponse")
}

func init() {
	proto.RegisterFile("listen.proto", fileDescriptor_f05da38a4a3e5177)
}

var fileDescriptor_f05da38a4a3e5177 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4f, 0x83, 0x40,
	0x10, 0x85, 0xbb, 0xb4, 0x5d, 0xe8, 0x88, 0x64, 0x33, 0x31, 0x91, 0x93, 0x41, 0x4e, 0x9c, 0x38,
	0x68, 0xe2, 0x59, 0x24, 0xa4, 0x21, 0x12, 0xda, 0x2c, 0x68, 0xe2, 0xc9, 0xd0, 0x76, 0x53, 0x1b,
	0x5a, 0xb6, 0xd9, 0xc5, 0x83, 0x47, 0xff, 0xb9, 0x71, 0xa9, 0x07, 0x0f, 0xbd, 0xcd, 0xcb, 0x7c,
	0x6f, 0xde, 0xe4, 0x81, 0xbb, 0xdf, 0xe9, 0x5e, 0x74, 0xf1, 0x51, 0xc9, 0x5e, 0xa2, 0x7d, 0x10,
	0x5a, 0x37, 0x5b, 0x11, 0xde, 0x02, 0x2d, 0xcc, 0x02, 0xaf, 0xc1, 0xfe, 0xd4, 0x42, 0xbd, 0xef,
	0x36, 0x3e, 0x09, 0x48, 0x34, 0xe3, 0xf4, 0x57, 0xe6, 0x9b, 0xf0, 0xdb, 0x02, 0x6f, 0x60, 0xb8,
	0xd0, 0x47, 0xd9, 0x69, 0x71, 0x96, 0xc5, 0x07, 0xa0, 0xeb, 0x8f, 0xa6, 0xdb, 0x0a, 0xdf, 0x0a,
	0x48, 0xe4, 0xdd, 0xdd, 0xc4, 0xa7, 0xa0, 0xf8, 0xff, 0x85, 0x38, 0x35, 0x14, 0x3f, 0xd1, 0x88,
	0x30, 0x39, 0x88, 0xbe, 0xf1, 0xc7, 0x01, 0x89, 0x5c, 0x6e, 0x66, 0xf4, 0xc0, 0x92, 0xad, 0x3f,
	0x09, 0x48, 0xe4, 0x70, 0x4b, 0xb6, 0x78, 0x05, 0x53, 0xa1, 0x94, 0x54, 0xfe, 0xd4, 0x44, 0x0e,
	0x22, 0x7c, 0x04, 0x67, 0xb9, 0x6f, 0xbe, 0x56, 0xcd, 0xba, 0xc5, 0x4b, 0x98, 0x15, 0x79, 0x55,
	0x67, 0x65, 0x5e, 0xce, 0xd9, 0x08, 0x01, 0x68, 0x92, 0xd6, 0xf9, 0x6b, 0xc6, 0x08, 0xba, 0xe0,
	0x2c, 0xf9, 0x62, 0xce, 0xb3, 0xaa, 0x62, 0x16, 0x5e, 0x80, 0x9d, 0xbe, 0x70, 0x9e, 0x95, 0x35,
	0x1b, 0x87, 0x21, 0xd0, 0xe1, 0x9b, 0x3f, 0xc3, 0xa2, 0x64, 0x23, 0x63, 0x28, 0x92, 0xb7, 0xa7,
	0x24, 0x7d, 0x66, 0x64, 0x45, 0x4d, 0x6d, 0xf7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x15, 0x0c,
	0x76, 0xdc, 0x46, 0x01, 0x00, 0x00,
}
