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
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd1, 0xd1, 0x4b, 0x02, 0x31,
	0x1c, 0xc0, 0x71, 0x4c, 0xd3, 0x98, 0xd0, 0xc3, 0x9e, 0x86, 0x0f, 0x95, 0x46, 0x61, 0x81, 0x67,
	0xe6, 0x53, 0x11, 0x15, 0x95, 0xf4, 0x14, 0x88, 0xbd, 0xf5, 0x22, 0xbb, 0x35, 0xaf, 0xd1, 0xb9,
	0x1d, 0xbb, 0xdf, 0xfc, 0x9b, 0xfa, 0x33, 0xe3, 0xb6, 0x9d, 0x7a, 0x30, 0xc4, 0x37, 0xe5, 0xbe,
	0x9f, 0x6d, 0xfc, 0x7e, 0xa8, 0xbb, 0xa4, 0x52, 0x2c, 0x78, 0x0e, 0xd1, 0x6a, 0x34, 0xe4, 0x39,
	0xd0, 0x05, 0x07, 0xe0, 0x73, 0xbe, 0xe2, 0x12, 0xa2, 0x4c, 0x2b, 0x50, 0xb8, 0xbd, 0x95, 0x74,
	0xae, 0xc3, 0x7d, 0x26, 0x32, 0x9e, 0x0a, 0x59, 0x81, 0x9d, 0xab, 0x70, 0xab, 0x79, 0xca, 0x69,
	0x5e, 0x4d, 0x2f, 0xc2, 0x69, 0x22, 0xa0, 0x92, 0xf5, 0xc3, 0xd9, 0xb7, 0x62, 0xbf, 0x5c, 0x57,
	0xca, 0xcb, 0x70, 0xc9, 0xb4, 0x92, 0xfb, 0xbc, 0x31, 0x33, 0xf1, 0x3c, 0x37, 0xf1, 0x3e, 0x97,
	0x2f, 0xa9, 0x34, 0x34, 0xdd, 0x2e, 0x7b, 0x7f, 0x75, 0x74, 0x3c, 0x29, 0x83, 0x49, 0xf1, 0x01,
	0x3f, 0xa1, 0xa3, 0x72, 0x46, 0xa4, 0x76, 0x56, 0xeb, 0xb7, 0x6f, 0xcf, 0xa3, 0xad, 0xf3, 0xa2,
	0x75, 0x3e, 0xf5, 0x95, 0x65, 0xb3, 0x35, 0xc2, 0x0f, 0xa8, 0xe5, 0x07, 0x47, 0x0e, 0xac, 0xef,
	0x85, 0xfd, 0xcc, 0x45, 0x8e, 0x97, 0x04, 0xdf, 0xa0, 0x7a, 0x22, 0x80, 0xd4, 0xad, 0x3c, 0x09,
	0xcb, 0x77, 0x01, 0x4e, 0x15, 0x29, 0xbe, 0x43, 0x4d, 0x37, 0x56, 0xd2, 0xb0, 0xa8, 0x1b, 0x46,
	0x6f, 0xb6, 0x71, 0xce, 0x03, 0x3c, 0x46, 0x8d, 0x62, 0xce, 0xe4, 0xd0, 0xc2, 0xd3, 0x30, 0x7c,
	0xd5, 0x4a, 0x3a, 0x66, 0x63, 0x7c, 0x8f, 0x5a, 0x7e, 0xe8, 0xa4, 0xb9, 0xeb, 0xc2, 0xa9, 0x89,
	0x3f, 0x4d, 0xec, 0x2f, 0xcc, 0xec, 0x9f, 0xe2, 0xad, 0x6e, 0x0b, 0xa4, 0xb5, 0x8b, 0x7e, 0xd8,
	0xc6, 0x53, 0x07, 0x5e, 0x9e, 0xbf, 0x1e, 0x13, 0x01, 0x3f, 0x26, 0x8e, 0x98, 0x5a, 0x6e, 0xb6,
	0xba, 0xf9, 0x35, 0x60, 0x62, 0xc0, 0x94, 0x04, 0x4d, 0x19, 0xe4, 0x83, 0x44, 0xa5, 0x54, 0x26,
	0xc3, 0xf2, 0xf8, 0xf9, 0x6a, 0x14, 0x37, 0xed, 0xce, 0xc7, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x4c, 0xc7, 0x8e, 0x49, 0x4a, 0x03, 0x00, 0x00,
}
