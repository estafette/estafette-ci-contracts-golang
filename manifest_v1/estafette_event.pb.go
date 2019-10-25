// Code generated by protoc-gen-go. DO NOT EDIT.
// source: manifest.v1/estafette_event.proto

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

type EstafetteEvent struct {
	Pipeline             *EstafettePipelineEvent `protobuf:"bytes,1,opt,name=pipeline,proto3" json:"pipeline,omitempty"`
	Release              *EstafetteReleaseEvent  `protobuf:"bytes,2,opt,name=release,proto3" json:"release,omitempty"`
	Git                  *EstafetteGitEvent      `protobuf:"bytes,3,opt,name=git,proto3" json:"git,omitempty"`
	Docker               *EstafetteDockerEvent   `protobuf:"bytes,4,opt,name=docker,proto3" json:"docker,omitempty"`
	Cron                 *EstafetteCronEvent     `protobuf:"bytes,5,opt,name=cron,proto3" json:"cron,omitempty"`
	PubSub               *EstafettePubSubEvent   `protobuf:"bytes,6,opt,name=pub_sub,json=pubSub,proto3" json:"pub_sub,omitempty"`
	Manual               *EstafetteManualEvent   `protobuf:"bytes,7,opt,name=manual,proto3" json:"manual,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *EstafetteEvent) Reset()         { *m = EstafetteEvent{} }
func (m *EstafetteEvent) String() string { return proto.CompactTextString(m) }
func (*EstafetteEvent) ProtoMessage()    {}
func (*EstafetteEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6a0590b8cd7e52c, []int{0}
}

func (m *EstafetteEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EstafetteEvent.Unmarshal(m, b)
}
func (m *EstafetteEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EstafetteEvent.Marshal(b, m, deterministic)
}
func (m *EstafetteEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EstafetteEvent.Merge(m, src)
}
func (m *EstafetteEvent) XXX_Size() int {
	return xxx_messageInfo_EstafetteEvent.Size(m)
}
func (m *EstafetteEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_EstafetteEvent.DiscardUnknown(m)
}

var xxx_messageInfo_EstafetteEvent proto.InternalMessageInfo

func (m *EstafetteEvent) GetPipeline() *EstafettePipelineEvent {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

func (m *EstafetteEvent) GetRelease() *EstafetteReleaseEvent {
	if m != nil {
		return m.Release
	}
	return nil
}

func (m *EstafetteEvent) GetGit() *EstafetteGitEvent {
	if m != nil {
		return m.Git
	}
	return nil
}

func (m *EstafetteEvent) GetDocker() *EstafetteDockerEvent {
	if m != nil {
		return m.Docker
	}
	return nil
}

func (m *EstafetteEvent) GetCron() *EstafetteCronEvent {
	if m != nil {
		return m.Cron
	}
	return nil
}

func (m *EstafetteEvent) GetPubSub() *EstafettePubSubEvent {
	if m != nil {
		return m.PubSub
	}
	return nil
}

func (m *EstafetteEvent) GetManual() *EstafetteManualEvent {
	if m != nil {
		return m.Manual
	}
	return nil
}

func init() {
	proto.RegisterType((*EstafetteEvent)(nil), "manifest.v1.EstafetteEvent")
}

func init() { proto.RegisterFile("manifest.v1/estafette_event.proto", fileDescriptor_f6a0590b8cd7e52c) }

var fileDescriptor_f6a0590b8cd7e52c = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd2, 0x4d, 0x4b, 0xfb, 0x30,
	0x1c, 0xc0, 0x71, 0xf6, 0xf0, 0xdf, 0xfe, 0x64, 0xe0, 0x21, 0xa7, 0x30, 0x41, 0xdd, 0x44, 0x99,
	0xc2, 0x32, 0xe7, 0x4e, 0x3e, 0x80, 0xe0, 0x1c, 0xe2, 0x61, 0x30, 0x26, 0x78, 0xf0, 0x32, 0x9a,
	0x9a, 0xd5, 0x60, 0x97, 0x94, 0x36, 0xe9, 0x6b, 0x12, 0x5f, 0xa5, 0x34, 0x49, 0xbb, 0x16, 0xe2,
	0xd8, 0x6d, 0x63, 0xdf, 0x4f, 0x7e, 0x5d, 0x7f, 0x01, 0xbd, 0x8d, 0xc7, 0xd9, 0x9a, 0x26, 0x12,
	0xa7, 0xe3, 0x11, 0x4d, 0xa4, 0xb7, 0xa6, 0x52, 0xd2, 0x15, 0x4d, 0x29, 0x97, 0x38, 0x8a, 0x85,
	0x14, 0xb0, 0x53, 0x4a, 0xba, 0xe7, 0xee, 0xde, 0x8f, 0x05, 0x2f, 0xa3, 0xee, 0xc0, 0xdd, 0x7d,
	0x08, 0xff, 0x8b, 0xc6, 0x95, 0xf2, 0xcc, 0x5d, 0x06, 0x4c, 0xee, 0x73, 0xe0, 0xc6, 0xe3, 0xca,
	0x0b, 0x2b, 0xe5, 0xa5, 0xbb, 0x8c, 0x58, 0x44, 0x43, 0xc6, 0x2b, 0xff, 0xad, 0x7b, 0xf1, 0x47,
	0xab, 0xc8, 0x2a, 0x51, 0x64, 0x9f, 0x34, 0xa6, 0x21, 0xf5, 0x92, 0xca, 0xa9, 0xfd, 0xef, 0x06,
	0x38, 0x98, 0xe5, 0xc5, 0x2c, 0xfb, 0x01, 0x3e, 0x80, 0xff, 0xf9, 0x03, 0xa0, 0xda, 0x49, 0x6d,
	0xd0, 0xb9, 0x3e, 0xc5, 0xa5, 0x03, 0x71, 0x91, 0x2f, 0x6c, 0xa5, 0xd9, 0xb2, 0x40, 0xf0, 0x1e,
	0xb4, 0xed, 0x28, 0x54, 0xd7, 0xbe, 0xef, 0xf6, 0x4b, 0x13, 0x19, 0x9e, 0x13, 0x78, 0x05, 0x1a,
	0x01, 0x93, 0xa8, 0xa1, 0xe5, 0x91, 0x5b, 0x3e, 0x33, 0x69, 0x54, 0x96, 0xc2, 0x1b, 0xd0, 0x32,
	0xcb, 0x42, 0x4d, 0x8d, 0x7a, 0x6e, 0xf4, 0xa4, 0x1b, 0xe3, 0x2c, 0x80, 0x13, 0xd0, 0xcc, 0xee,
	0x03, 0xfa, 0xa7, 0xe1, 0xb1, 0x1b, 0x4e, 0x63, 0xc1, 0x0d, 0xd3, 0x31, 0xbc, 0x05, 0x6d, 0xfb,
	0xd6, 0x51, 0x6b, 0xd7, 0xc0, 0x85, 0x22, 0xaf, 0x8a, 0xd8, 0x81, 0x91, 0xfe, 0x92, 0x3d, 0xab,
	0xb9, 0x07, 0xa8, 0xbd, 0x8b, 0xce, 0x75, 0x63, 0xa9, 0x01, 0x8f, 0x0a, 0x1c, 0x32, 0x81, 0x8b,
	0x75, 0x62, 0x9f, 0x95, 0xed, 0xfb, 0x5d, 0xc0, 0xe4, 0xa7, 0x22, 0xd8, 0x17, 0x9b, 0xed, 0xce,
	0xb7, 0x9f, 0x86, 0x3e, 0x1b, 0xea, 0x95, 0x27, 0xc3, 0x40, 0x84, 0x1e, 0x0f, 0x46, 0x39, 0x5e,
	0xa5, 0xe3, 0x9f, 0x3a, 0x2a, 0x46, 0xe3, 0xe9, 0x0b, 0x9e, 0xe7, 0xe7, 0xbe, 0x8d, 0x49, 0x4b,
	0xab, 0xc9, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x8b, 0x2e, 0xce, 0x7f, 0x03, 0x00, 0x00,
}
