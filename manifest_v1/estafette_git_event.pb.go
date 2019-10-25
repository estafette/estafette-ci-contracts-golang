// Code generated by protoc-gen-go. DO NOT EDIT.
// source: manifest.v1/estafette_git_event.proto

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

type EstafetteGitEvent struct {
	Event                string   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Repository           string   `protobuf:"bytes,2,opt,name=repository,proto3" json:"repository,omitempty"`
	Branch               string   `protobuf:"bytes,3,opt,name=branch,proto3" json:"branch,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EstafetteGitEvent) Reset()         { *m = EstafetteGitEvent{} }
func (m *EstafetteGitEvent) String() string { return proto.CompactTextString(m) }
func (*EstafetteGitEvent) ProtoMessage()    {}
func (*EstafetteGitEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_97f91b9bdb442cc6, []int{0}
}

func (m *EstafetteGitEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EstafetteGitEvent.Unmarshal(m, b)
}
func (m *EstafetteGitEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EstafetteGitEvent.Marshal(b, m, deterministic)
}
func (m *EstafetteGitEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EstafetteGitEvent.Merge(m, src)
}
func (m *EstafetteGitEvent) XXX_Size() int {
	return xxx_messageInfo_EstafetteGitEvent.Size(m)
}
func (m *EstafetteGitEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_EstafetteGitEvent.DiscardUnknown(m)
}

var xxx_messageInfo_EstafetteGitEvent proto.InternalMessageInfo

func (m *EstafetteGitEvent) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *EstafetteGitEvent) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

func (m *EstafetteGitEvent) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

func init() {
	proto.RegisterType((*EstafetteGitEvent)(nil), "manifest.v1.EstafetteGitEvent")
}

func init() {
	proto.RegisterFile("manifest.v1/estafette_git_event.proto", fileDescriptor_97f91b9bdb442cc6)
}

var fileDescriptor_97f91b9bdb442cc6 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcd, 0x4d, 0xcc, 0xcb,
	0x4c, 0x4b, 0x2d, 0x2e, 0xd1, 0x2b, 0x33, 0xd4, 0x4f, 0x2d, 0x2e, 0x49, 0x4c, 0x4b, 0x2d, 0x29,
	0x49, 0x8d, 0x4f, 0xcf, 0x2c, 0x89, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x46, 0x52, 0xa6, 0x94, 0xc8, 0x25, 0xe8, 0x0a, 0x53, 0xe9, 0x9e, 0x59, 0xe2,
	0x0a, 0x52, 0x27, 0x24, 0xc2, 0xc5, 0x0a, 0xd6, 0x20, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04,
	0xe1, 0x08, 0xc9, 0x71, 0x71, 0x15, 0xa5, 0x16, 0xe4, 0x17, 0x67, 0x96, 0xe4, 0x17, 0x55, 0x4a,
	0x30, 0x81, 0xa5, 0x90, 0x44, 0x84, 0xc4, 0xb8, 0xd8, 0x92, 0x8a, 0x12, 0xf3, 0x92, 0x33, 0x24,
	0x98, 0xc1, 0x72, 0x50, 0x9e, 0x53, 0x0d, 0x97, 0x52, 0x66, 0xbe, 0x1e, 0xdc, 0x3d, 0x7a, 0xc9,
	0x99, 0x10, 0x87, 0x14, 0xeb, 0x21, 0x39, 0x24, 0xca, 0x3a, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49,
	0x2f, 0x39, 0x3f, 0x17, 0xe1, 0x76, 0x04, 0x4b, 0x37, 0x39, 0x53, 0x17, 0xa2, 0x4b, 0x37, 0x3d,
	0x3f, 0x27, 0x31, 0x2f, 0x5d, 0x1f, 0xa6, 0x39, 0xbe, 0xcc, 0x70, 0x15, 0x93, 0x04, 0xdc, 0x13,
	0x7a, 0xce, 0x9e, 0x7a, 0xbe, 0x30, 0x73, 0xc3, 0x0c, 0x93, 0xd8, 0xc0, 0xba, 0x8c, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xe1, 0xcc, 0x25, 0xe2, 0x1d, 0x01, 0x00, 0x00,
}
