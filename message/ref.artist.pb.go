// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ref.artist.proto

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

type RefArtist struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Uri                  string   `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefArtist) Reset()         { *m = RefArtist{} }
func (m *RefArtist) String() string { return proto.CompactTextString(m) }
func (*RefArtist) ProtoMessage()    {}
func (*RefArtist) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a3b8eee6ad1298d, []int{0}
}

func (m *RefArtist) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefArtist.Unmarshal(m, b)
}
func (m *RefArtist) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefArtist.Marshal(b, m, deterministic)
}
func (m *RefArtist) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefArtist.Merge(m, src)
}
func (m *RefArtist) XXX_Size() int {
	return xxx_messageInfo_RefArtist.Size(m)
}
func (m *RefArtist) XXX_DiscardUnknown() {
	xxx_messageInfo_RefArtist.DiscardUnknown(m)
}

var xxx_messageInfo_RefArtist proto.InternalMessageInfo

func (m *RefArtist) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RefArtist) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RefArtist) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func init() {
	proto.RegisterType((*RefArtist)(nil), "message.RefArtist")
}

func init() {
	proto.RegisterFile("ref.artist.proto", fileDescriptor_0a3b8eee6ad1298d)
}

var fileDescriptor_0a3b8eee6ad1298d = []byte{
	// 108 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x4a, 0x4d, 0xd3,
	0x4b, 0x2c, 0x2a, 0xc9, 0x2c, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcf, 0x4d,
	0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x55, 0x72, 0xe4, 0xe2, 0x0c, 0x4a, 0x4d, 0x73, 0x04, 0xcb, 0x09,
	0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9,
	0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x12, 0x4c, 0x60, 0x11, 0xa6, 0xcc, 0x14, 0x21, 0x01, 0x2e,
	0xe6, 0xd2, 0xa2, 0x4c, 0x09, 0x66, 0xb0, 0x00, 0x88, 0x99, 0xc4, 0x06, 0x36, 0xd2, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x7f, 0x4a, 0xf6, 0xa8, 0x66, 0x00, 0x00, 0x00,
}
