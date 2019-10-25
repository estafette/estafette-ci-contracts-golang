// Code generated by protoc-gen-go. DO NOT EDIT.
// source: manifest.v1/estafette_manual_event.proto

package manifest_v1

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

type EstafetteManualEvent struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EstafetteManualEvent) Reset()         { *m = EstafetteManualEvent{} }
func (m *EstafetteManualEvent) String() string { return proto.CompactTextString(m) }
func (*EstafetteManualEvent) ProtoMessage()    {}
func (*EstafetteManualEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_02f49e5e44b15ddc, []int{0}
}

func (m *EstafetteManualEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EstafetteManualEvent.Unmarshal(m, b)
}
func (m *EstafetteManualEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EstafetteManualEvent.Marshal(b, m, deterministic)
}
func (m *EstafetteManualEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EstafetteManualEvent.Merge(m, src)
}
func (m *EstafetteManualEvent) XXX_Size() int {
	return xxx_messageInfo_EstafetteManualEvent.Size(m)
}
func (m *EstafetteManualEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_EstafetteManualEvent.DiscardUnknown(m)
}

var xxx_messageInfo_EstafetteManualEvent proto.InternalMessageInfo

func (m *EstafetteManualEvent) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func init() {
	proto.RegisterType((*EstafetteManualEvent)(nil), "manifest.v1.EstafetteManualEvent")
}

func init() {
	proto.RegisterFile("manifest.v1/estafette_manual_event.proto", fileDescriptor_02f49e5e44b15ddc)
}

var fileDescriptor_02f49e5e44b15ddc = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xc8, 0x4d, 0xcc, 0xcb,
	0x4c, 0x4b, 0x2d, 0x2e, 0xd1, 0x2b, 0x33, 0xd4, 0x4f, 0x2d, 0x2e, 0x49, 0x4c, 0x4b, 0x2d, 0x29,
	0x49, 0x8d, 0xcf, 0x4d, 0xcc, 0x2b, 0x4d, 0xcc, 0x89, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x46, 0x52, 0xa9, 0xa4, 0xcf, 0x25, 0xe2, 0x0a, 0x53, 0xec,
	0x0b, 0x56, 0xeb, 0x0a, 0x52, 0x2a, 0x24, 0xce, 0xc5, 0x5e, 0x5a, 0x9c, 0x5a, 0x14, 0x9f, 0x99,
	0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0xc4, 0x06, 0xe2, 0x7a, 0xa6, 0x38, 0x95, 0x72, 0x49,
	0x67, 0xe6, 0xeb, 0xc1, 0x2d, 0xd0, 0x4b, 0xce, 0xd4, 0x43, 0x32, 0x2f, 0xca, 0x3a, 0x3d, 0xb3,
	0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x17, 0xe1, 0x0a, 0x04, 0x4b, 0x37, 0x39, 0x53, 0x17,
	0xec, 0x8a, 0x62, 0xdd, 0xf4, 0xfc, 0x9c, 0xc4, 0xbc, 0x74, 0x7d, 0x98, 0xe6, 0xf8, 0x32, 0xc3,
	0x55, 0x4c, 0x12, 0x70, 0xb7, 0xe8, 0x39, 0x7b, 0xea, 0xf9, 0xc2, 0xcc, 0x0d, 0x33, 0x4c, 0x62,
	0x03, 0xeb, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xff, 0xf2, 0xb5, 0x18, 0xe7, 0x00, 0x00,
	0x00,
}
