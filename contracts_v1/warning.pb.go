// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contracts.v1/warning.proto

package contracts_v1

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

type Warning struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Warning) Reset()         { *m = Warning{} }
func (m *Warning) String() string { return proto.CompactTextString(m) }
func (*Warning) ProtoMessage()    {}
func (*Warning) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c645fc48882bdc2, []int{0}
}

func (m *Warning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Warning.Unmarshal(m, b)
}
func (m *Warning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Warning.Marshal(b, m, deterministic)
}
func (m *Warning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Warning.Merge(m, src)
}
func (m *Warning) XXX_Size() int {
	return xxx_messageInfo_Warning.Size(m)
}
func (m *Warning) XXX_DiscardUnknown() {
	xxx_messageInfo_Warning.DiscardUnknown(m)
}

var xxx_messageInfo_Warning proto.InternalMessageInfo

func (m *Warning) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Warning) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Warning)(nil), "contracts.v1.Warning")
}

func init() { proto.RegisterFile("contracts.v1/warning.proto", fileDescriptor_1c645fc48882bdc2) }

var fileDescriptor_1c645fc48882bdc2 = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0xce, 0xcf, 0x2b,
	0x29, 0x4a, 0x4c, 0x2e, 0x29, 0xd6, 0x2b, 0x33, 0xd4, 0x2f, 0x4f, 0x2c, 0xca, 0xcb, 0xcc, 0x4b,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x41, 0x96, 0x53, 0xb2, 0xe6, 0x62, 0x0f, 0x87,
	0x48, 0x0b, 0x89, 0x71, 0xb1, 0x15, 0x97, 0x24, 0x96, 0x94, 0x16, 0x4b, 0x30, 0x2a, 0x30, 0x6a,
	0x70, 0x06, 0x41, 0x79, 0x42, 0x12, 0x5c, 0xec, 0xb9, 0xa9, 0xc5, 0xc5, 0x89, 0xe9, 0xa9, 0x12,
	0x4c, 0x60, 0x09, 0x18, 0xd7, 0xa9, 0x82, 0x4b, 0x26, 0x33, 0x5f, 0x2f, 0xb5, 0xb8, 0x24, 0x31,
	0x2d, 0xb5, 0xa4, 0x24, 0x55, 0x2f, 0x39, 0x53, 0x0f, 0xd9, 0xf0, 0x28, 0x9b, 0xf4, 0xcc, 0x92,
	0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0xb8, 0x2a, 0x04, 0x4b, 0x37, 0x39, 0x53, 0x17,
	0xec, 0xa6, 0x62, 0xdd, 0xf4, 0xfc, 0x9c, 0xc4, 0xbc, 0x74, 0x7d, 0xb8, 0xee, 0xf8, 0x32, 0xc3,
	0x55, 0x4c, 0x92, 0xae, 0x70, 0x93, 0x9d, 0x3d, 0xf5, 0x9c, 0xe1, 0x26, 0x87, 0x19, 0x26, 0xb1,
	0x81, 0xf5, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x79, 0x36, 0x07, 0xe9, 0x00, 0x00,
	0x00,
}
