// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/manifest.v1/estafette_trigger_release_action.proto

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

type EstafetteTriggerReleaseAction struct {
	Target               string   `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	Action               string   `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EstafetteTriggerReleaseAction) Reset()         { *m = EstafetteTriggerReleaseAction{} }
func (m *EstafetteTriggerReleaseAction) String() string { return proto.CompactTextString(m) }
func (*EstafetteTriggerReleaseAction) ProtoMessage()    {}
func (*EstafetteTriggerReleaseAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c06fb2e5bd80713, []int{0}
}

func (m *EstafetteTriggerReleaseAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EstafetteTriggerReleaseAction.Unmarshal(m, b)
}
func (m *EstafetteTriggerReleaseAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EstafetteTriggerReleaseAction.Marshal(b, m, deterministic)
}
func (m *EstafetteTriggerReleaseAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EstafetteTriggerReleaseAction.Merge(m, src)
}
func (m *EstafetteTriggerReleaseAction) XXX_Size() int {
	return xxx_messageInfo_EstafetteTriggerReleaseAction.Size(m)
}
func (m *EstafetteTriggerReleaseAction) XXX_DiscardUnknown() {
	xxx_messageInfo_EstafetteTriggerReleaseAction.DiscardUnknown(m)
}

var xxx_messageInfo_EstafetteTriggerReleaseAction proto.InternalMessageInfo

func (m *EstafetteTriggerReleaseAction) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *EstafetteTriggerReleaseAction) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *EstafetteTriggerReleaseAction) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterType((*EstafetteTriggerReleaseAction)(nil), "manifest.v1.EstafetteTriggerReleaseAction")
}

func init() {
	proto.RegisterFile("protos/manifest.v1/estafette_trigger_release_action.proto", fileDescriptor_8c06fb2e5bd80713)
}

var fileDescriptor_8c06fb2e5bd80713 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x3f, 0x4b, 0xc5, 0x30,
	0x14, 0xc5, 0x79, 0x4f, 0xa8, 0x18, 0xb7, 0x0e, 0x92, 0x45, 0x90, 0x4e, 0x2e, 0xbd, 0xa5, 0x38,
	0x89, 0x93, 0x8a, 0x83, 0x83, 0x4b, 0x11, 0x07, 0x97, 0x90, 0x86, 0xdb, 0x78, 0xa1, 0x6d, 0x24,
	0xb9, 0x76, 0xf2, 0x13, 0xf9, 0x29, 0xa5, 0xc9, 0xeb, 0x9f, 0x2d, 0xe7, 0x1e, 0x7e, 0x87, 0x73,
	0x22, 0xee, 0xbf, 0xbd, 0x63, 0x17, 0xaa, 0x41, 0x8f, 0xd4, 0x61, 0x60, 0x98, 0xea, 0x0a, 0x03,
	0xeb, 0x0e, 0x99, 0x51, 0xb1, 0x27, 0x6b, 0xd1, 0x2b, 0x8f, 0x3d, 0xea, 0x80, 0x4a, 0x1b, 0x26,
	0x37, 0x42, 0x64, 0xf2, 0xcb, 0x1d, 0x53, 0x90, 0xb8, 0x7e, 0x59, 0xb0, 0xf7, 0x44, 0x35, 0x09,
	0x7a, 0x8c, 0x4c, 0x7e, 0x25, 0x32, 0xd6, 0xde, 0x22, 0xcb, 0xc3, 0xcd, 0xe1, 0xf6, 0xa2, 0x39,
	0xa9, 0xf9, 0x9e, 0x52, 0xe5, 0x31, 0xdd, 0x93, 0xca, 0xa5, 0x38, 0x9f, 0xd0, 0x87, 0xd9, 0x38,
	0x8b, 0xc6, 0x22, 0x9f, 0x7e, 0x45, 0x41, 0x0e, 0xd6, 0x92, 0x60, 0x28, 0x15, 0x0a, 0xb0, 0x2b,
	0xf4, 0xf9, 0x60, 0x89, 0xbf, 0x7e, 0x5a, 0x30, 0x6e, 0xd8, 0x06, 0x6d, 0xaf, 0xd2, 0x50, 0x99,
	0xa8, 0xd2, 0xba, 0x5e, 0x8f, 0x76, 0xfd, 0x01, 0x35, 0xd5, 0x7f, 0x47, 0xb9, 0x8e, 0x81, 0xe7,
	0x57, 0x78, 0x5b, 0x72, 0x3f, 0xea, 0x36, 0x8b, 0xd4, 0xdd, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xc0, 0x84, 0xc9, 0xd8, 0x39, 0x01, 0x00, 0x00,
}
