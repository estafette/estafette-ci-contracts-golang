// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/contracts.v1/label.proto

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

type Label struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Label) Reset()         { *m = Label{} }
func (m *Label) String() string { return proto.CompactTextString(m) }
func (*Label) ProtoMessage()    {}
func (*Label) Descriptor() ([]byte, []int) {
	return fileDescriptor_719bc61106ba793b, []int{0}
}

func (m *Label) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Label.Unmarshal(m, b)
}
func (m *Label) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Label.Marshal(b, m, deterministic)
}
func (m *Label) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Label.Merge(m, src)
}
func (m *Label) XXX_Size() int {
	return xxx_messageInfo_Label.Size(m)
}
func (m *Label) XXX_DiscardUnknown() {
	xxx_messageInfo_Label.DiscardUnknown(m)
}

var xxx_messageInfo_Label proto.InternalMessageInfo

func (m *Label) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Label) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Label)(nil), "contracts.v1.Label")
}

func init() { proto.RegisterFile("protos/contracts.v1/label.proto", fileDescriptor_719bc61106ba793b) }

var fileDescriptor_719bc61106ba793b = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x4f, 0xce, 0xcf, 0x2b, 0x29, 0x4a, 0x4c, 0x2e, 0x29, 0xd6, 0x2b, 0x33, 0xd4,
	0xcf, 0x49, 0x4c, 0x4a, 0xcd, 0xd1, 0x03, 0xcb, 0x08, 0xf1, 0x20, 0xcb, 0x28, 0xe9, 0x73, 0xb1,
	0xfa, 0x80, 0x24, 0x85, 0x04, 0xb8, 0x98, 0xb3, 0x53, 0x2b, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38,
	0x83, 0x40, 0x4c, 0x21, 0x11, 0x2e, 0xd6, 0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x26, 0xb0, 0x18,
	0x84, 0xe3, 0x54, 0xcf, 0xa5, 0x9c, 0x99, 0xaf, 0x97, 0x5a, 0x5c, 0x92, 0x98, 0x96, 0x5a, 0x52,
	0x92, 0xaa, 0x97, 0x9c, 0x09, 0x31, 0xb7, 0x58, 0x0f, 0xd9, 0xdc, 0x28, 0x9b, 0xf4, 0xcc, 0x92,
	0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0xb8, 0x62, 0x04, 0x4b, 0x37, 0x39, 0x53, 0x17,
	0xa2, 0x4d, 0x37, 0x3d, 0x3f, 0x27, 0x31, 0x2f, 0x1d, 0xe1, 0xde, 0xf8, 0x32, 0xc3, 0x55, 0x4c,
	0x92, 0xae, 0x70, 0x0b, 0x9c, 0x3d, 0xf5, 0x9c, 0xe1, 0x26, 0x87, 0x19, 0x26, 0xb1, 0x81, 0xf5,
	0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x13, 0xbc, 0x83, 0xe9, 0x00, 0x00, 0x00,
}
