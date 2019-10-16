// Code generated by protoc-gen-go. DO NOT EDIT.
// source: manifest.v1/estafette_pub_sub_trigger.proto

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

type EstafettePubSubTrigger struct {
	Project              string   `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	Topic                string   `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EstafettePubSubTrigger) Reset()         { *m = EstafettePubSubTrigger{} }
func (m *EstafettePubSubTrigger) String() string { return proto.CompactTextString(m) }
func (*EstafettePubSubTrigger) ProtoMessage()    {}
func (*EstafettePubSubTrigger) Descriptor() ([]byte, []int) {
	return fileDescriptor_e87c69966d4acf70, []int{0}
}

func (m *EstafettePubSubTrigger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EstafettePubSubTrigger.Unmarshal(m, b)
}
func (m *EstafettePubSubTrigger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EstafettePubSubTrigger.Marshal(b, m, deterministic)
}
func (m *EstafettePubSubTrigger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EstafettePubSubTrigger.Merge(m, src)
}
func (m *EstafettePubSubTrigger) XXX_Size() int {
	return xxx_messageInfo_EstafettePubSubTrigger.Size(m)
}
func (m *EstafettePubSubTrigger) XXX_DiscardUnknown() {
	xxx_messageInfo_EstafettePubSubTrigger.DiscardUnknown(m)
}

var xxx_messageInfo_EstafettePubSubTrigger proto.InternalMessageInfo

func (m *EstafettePubSubTrigger) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *EstafettePubSubTrigger) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func init() {
	proto.RegisterType((*EstafettePubSubTrigger)(nil), "manifest.v1.EstafettePubSubTrigger")
}

func init() {
	proto.RegisterFile("manifest.v1/estafette_pub_sub_trigger.proto", fileDescriptor_e87c69966d4acf70)
}

var fileDescriptor_e87c69966d4acf70 = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0x31, 0x0b, 0xc2, 0x30,
	0x10, 0x46, 0xa9, 0xa0, 0x62, 0xdc, 0x8a, 0x48, 0x47, 0x71, 0x12, 0xa4, 0x2d, 0xc5, 0x5d, 0x44,
	0x10, 0x1c, 0x45, 0x9d, 0x5c, 0x42, 0x12, 0xae, 0x31, 0x62, 0x7b, 0x21, 0xb9, 0xf4, 0xf7, 0x0b,
	0x2d, 0xa1, 0x6e, 0xdf, 0xfb, 0x78, 0xc3, 0x63, 0xfb, 0x46, 0xb4, 0xa6, 0x06, 0x4f, 0x45, 0x57,
	0x95, 0xe0, 0x49, 0xd4, 0x40, 0x04, 0xdc, 0x06, 0xc9, 0x7d, 0x90, 0x9c, 0x9c, 0xd1, 0x1a, 0x5c,
	0x61, 0x1d, 0x12, 0xa6, 0xcb, 0x3f, 0x79, 0x7b, 0x65, 0xeb, 0x4b, 0xf4, 0x6f, 0x41, 0x3e, 0x82,
	0x7c, 0x0e, 0x72, 0x9a, 0xb1, 0xb9, 0x75, 0xf8, 0x01, 0x45, 0x59, 0xb2, 0x49, 0x76, 0x8b, 0x7b,
	0xc4, 0x74, 0xc5, 0xa6, 0x84, 0xd6, 0xa8, 0x6c, 0xd2, 0xff, 0x03, 0x9c, 0x4f, 0xaf, 0xa3, 0x36,
	0xf4, 0x0e, 0xb2, 0x50, 0xd8, 0x8c, 0x11, 0xe3, 0xca, 0x95, 0xc9, 0x15, 0xb6, 0xe4, 0x84, 0x22,
	0x9f, 0x6b, 0xfc, 0x8a, 0x56, 0x97, 0xb1, 0x85, 0x77, 0x95, 0x9c, 0xf5, 0x7d, 0x87, 0x5f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xde, 0xc9, 0x6d, 0xfa, 0xce, 0x00, 0x00, 0x00,
}
