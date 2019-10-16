// Code generated by protoc-gen-go. DO NOT EDIT.
// source: label.proto

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
	return fileDescriptor_6100d2b5ae808f16, []int{0}
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

func init() { proto.RegisterFile("label.proto", fileDescriptor_6100d2b5ae808f16) }

var fileDescriptor_6100d2b5ae808f16 = []byte{
	// 93 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x49, 0x4c, 0x4a,
	0xcd, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x49, 0xce, 0xcf, 0x2b, 0x29, 0x4a, 0x4c,
	0x2e, 0x29, 0xd6, 0x2b, 0x33, 0x54, 0xd2, 0xe7, 0x62, 0xf5, 0x01, 0x49, 0x0a, 0x09, 0x70, 0x31,
	0x67, 0xa7, 0x56, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0x98, 0x42, 0x22, 0x5c, 0xac,
	0x65, 0x89, 0x39, 0xa5, 0xa9, 0x12, 0x4c, 0x60, 0x31, 0x08, 0x27, 0x89, 0x0d, 0x6c, 0x8a, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x3c, 0xf3, 0x31, 0x54, 0x00, 0x00, 0x00,
}
