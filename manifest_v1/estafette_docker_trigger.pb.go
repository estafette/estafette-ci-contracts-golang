// Code generated by protoc-gen-go. DO NOT EDIT.
// source: manifest.v1/estafette_docker_trigger.proto

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

type EstafetteDockerTrigger struct {
	Event                string   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Image                string   `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Tag                  string   `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EstafetteDockerTrigger) Reset()         { *m = EstafetteDockerTrigger{} }
func (m *EstafetteDockerTrigger) String() string { return proto.CompactTextString(m) }
func (*EstafetteDockerTrigger) ProtoMessage()    {}
func (*EstafetteDockerTrigger) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ea232635fc43b60, []int{0}
}

func (m *EstafetteDockerTrigger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EstafetteDockerTrigger.Unmarshal(m, b)
}
func (m *EstafetteDockerTrigger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EstafetteDockerTrigger.Marshal(b, m, deterministic)
}
func (m *EstafetteDockerTrigger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EstafetteDockerTrigger.Merge(m, src)
}
func (m *EstafetteDockerTrigger) XXX_Size() int {
	return xxx_messageInfo_EstafetteDockerTrigger.Size(m)
}
func (m *EstafetteDockerTrigger) XXX_DiscardUnknown() {
	xxx_messageInfo_EstafetteDockerTrigger.DiscardUnknown(m)
}

var xxx_messageInfo_EstafetteDockerTrigger proto.InternalMessageInfo

func (m *EstafetteDockerTrigger) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *EstafetteDockerTrigger) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *EstafetteDockerTrigger) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func init() {
	proto.RegisterType((*EstafetteDockerTrigger)(nil), "manifest.v1.EstafetteDockerTrigger")
}

func init() {
	proto.RegisterFile("manifest.v1/estafette_docker_trigger.proto", fileDescriptor_3ea232635fc43b60)
}

var fileDescriptor_3ea232635fc43b60 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xca, 0x4d, 0xcc, 0xcb,
	0x4c, 0x4b, 0x2d, 0x2e, 0xd1, 0x2b, 0x33, 0xd4, 0x4f, 0x2d, 0x2e, 0x49, 0x4c, 0x4b, 0x2d, 0x29,
	0x49, 0x8d, 0x4f, 0xc9, 0x4f, 0xce, 0x4e, 0x2d, 0x8a, 0x2f, 0x29, 0xca, 0x4c, 0x4f, 0x4f, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x46, 0x52, 0xab, 0x14, 0xc6, 0x25, 0xe6, 0x0a,
	0x53, 0xee, 0x02, 0x56, 0x1d, 0x02, 0x51, 0x2c, 0x24, 0xc2, 0xc5, 0x9a, 0x5a, 0x96, 0x9a, 0x57,
	0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe1, 0x80, 0x44, 0x33, 0x73, 0x13, 0xd3, 0x53,
	0x25, 0x98, 0x20, 0xa2, 0x60, 0x8e, 0x90, 0x00, 0x17, 0x73, 0x49, 0x62, 0xba, 0x04, 0x33, 0x58,
	0x0c, 0xc4, 0x74, 0xaa, 0xe1, 0x52, 0xca, 0xcc, 0xd7, 0x83, 0xbb, 0x44, 0x2f, 0x39, 0x13, 0x62,
	0x7b, 0xb1, 0x1e, 0x92, 0xed, 0x51, 0xd6, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9,
	0xb9, 0x08, 0x57, 0x23, 0x58, 0xba, 0xc9, 0x99, 0xba, 0x10, 0x5d, 0xba, 0xe9, 0xf9, 0x39, 0x89,
	0x79, 0xe9, 0xfa, 0x30, 0xcd, 0xf1, 0x65, 0x86, 0xab, 0x98, 0x24, 0xe0, 0x2e, 0xd7, 0x73, 0xf6,
	0xd4, 0xf3, 0x85, 0x99, 0x1b, 0x66, 0x98, 0xc4, 0x06, 0xd6, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x12, 0xe9, 0x09, 0xee, 0x17, 0x01, 0x00, 0x00,
}
