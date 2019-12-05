// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/manifest.v1/estafette_docker_event.proto

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

type EstafetteDockerEvent struct {
	Event                string   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Image                string   `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Tag                  string   `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EstafetteDockerEvent) Reset()         { *m = EstafetteDockerEvent{} }
func (m *EstafetteDockerEvent) String() string { return proto.CompactTextString(m) }
func (*EstafetteDockerEvent) ProtoMessage()    {}
func (*EstafetteDockerEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfff9dde346070f9, []int{0}
}

func (m *EstafetteDockerEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EstafetteDockerEvent.Unmarshal(m, b)
}
func (m *EstafetteDockerEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EstafetteDockerEvent.Marshal(b, m, deterministic)
}
func (m *EstafetteDockerEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EstafetteDockerEvent.Merge(m, src)
}
func (m *EstafetteDockerEvent) XXX_Size() int {
	return xxx_messageInfo_EstafetteDockerEvent.Size(m)
}
func (m *EstafetteDockerEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_EstafetteDockerEvent.DiscardUnknown(m)
}

var xxx_messageInfo_EstafetteDockerEvent proto.InternalMessageInfo

func (m *EstafetteDockerEvent) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *EstafetteDockerEvent) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *EstafetteDockerEvent) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func init() {
	proto.RegisterType((*EstafetteDockerEvent)(nil), "manifest.v1.EstafetteDockerEvent")
}

func init() {
	proto.RegisterFile("protos/manifest.v1/estafette_docker_event.proto", fileDescriptor_bfff9dde346070f9)
}

var fileDescriptor_bfff9dde346070f9 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0xcf, 0x4d, 0xcc, 0xcb, 0x4c, 0x4b, 0x2d, 0x2e, 0xd1, 0x2b, 0x33, 0xd4, 0x4f,
	0x2d, 0x2e, 0x49, 0x4c, 0x4b, 0x2d, 0x29, 0x49, 0x8d, 0x4f, 0xc9, 0x4f, 0xce, 0x4e, 0x2d, 0x8a,
	0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0xd1, 0x03, 0xab, 0x14, 0xe2, 0x46, 0x52, 0xa9, 0x14, 0xc2, 0x25,
	0xe2, 0x0a, 0x53, 0xec, 0x02, 0x56, 0xeb, 0x0a, 0x52, 0x2a, 0x24, 0xc2, 0xc5, 0x0a, 0xd6, 0x23,
	0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe1, 0x80, 0x44, 0x33, 0x73, 0x13, 0xd3, 0x53, 0x25,
	0x98, 0x20, 0xa2, 0x60, 0x8e, 0x90, 0x00, 0x17, 0x73, 0x49, 0x62, 0xba, 0x04, 0x33, 0x58, 0x0c,
	0xc4, 0x74, 0xaa, 0xe1, 0x52, 0xca, 0xcc, 0xd7, 0x83, 0xbb, 0x42, 0x2f, 0x39, 0x13, 0x62, 0x77,
	0xb1, 0x1e, 0x92, 0xdd, 0x51, 0xd6, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9,
	0x08, 0x17, 0x23, 0x58, 0xba, 0xc9, 0x99, 0xba, 0x10, 0x5d, 0xba, 0xe9, 0xf9, 0x39, 0x89, 0x79,
	0xe9, 0x70, 0x2f, 0xc6, 0x97, 0x19, 0xae, 0x62, 0x92, 0x80, 0xbb, 0x5b, 0xcf, 0xd9, 0x53, 0xcf,
	0x17, 0x66, 0x6e, 0x98, 0x61, 0x12, 0x1b, 0x58, 0x97, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x6b,
	0xbe, 0x62, 0xac, 0x1a, 0x01, 0x00, 0x00,
}
