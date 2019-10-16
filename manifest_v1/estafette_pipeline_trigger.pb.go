// Code generated by protoc-gen-go. DO NOT EDIT.
// source: manifest.v1/estafette_pipeline_trigger.proto

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

type EstafettePipelineTrigger struct {
	Event                string   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Branch               string   `protobuf:"bytes,4,opt,name=branch,proto3" json:"branch,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EstafettePipelineTrigger) Reset()         { *m = EstafettePipelineTrigger{} }
func (m *EstafettePipelineTrigger) String() string { return proto.CompactTextString(m) }
func (*EstafettePipelineTrigger) ProtoMessage()    {}
func (*EstafettePipelineTrigger) Descriptor() ([]byte, []int) {
	return fileDescriptor_dbd39a53672b20d9, []int{0}
}

func (m *EstafettePipelineTrigger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EstafettePipelineTrigger.Unmarshal(m, b)
}
func (m *EstafettePipelineTrigger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EstafettePipelineTrigger.Marshal(b, m, deterministic)
}
func (m *EstafettePipelineTrigger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EstafettePipelineTrigger.Merge(m, src)
}
func (m *EstafettePipelineTrigger) XXX_Size() int {
	return xxx_messageInfo_EstafettePipelineTrigger.Size(m)
}
func (m *EstafettePipelineTrigger) XXX_DiscardUnknown() {
	xxx_messageInfo_EstafettePipelineTrigger.DiscardUnknown(m)
}

var xxx_messageInfo_EstafettePipelineTrigger proto.InternalMessageInfo

func (m *EstafettePipelineTrigger) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *EstafettePipelineTrigger) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *EstafettePipelineTrigger) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EstafettePipelineTrigger) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

func init() {
	proto.RegisterType((*EstafettePipelineTrigger)(nil), "manifest.v1.EstafettePipelineTrigger")
}

func init() {
	proto.RegisterFile("manifest.v1/estafette_pipeline_trigger.proto", fileDescriptor_dbd39a53672b20d9)
}

var fileDescriptor_dbd39a53672b20d9 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x31, 0xab, 0xc2, 0x30,
	0x14, 0x85, 0xe9, 0x7b, 0xb5, 0x60, 0xdc, 0x82, 0x48, 0x46, 0x71, 0x72, 0xb0, 0x2d, 0xc5, 0x5d,
	0x44, 0x70, 0x17, 0x71, 0x72, 0x29, 0x69, 0xb8, 0x4d, 0x03, 0x6d, 0x52, 0x92, 0xdb, 0xfe, 0x7e,
	0x31, 0x69, 0xa9, 0xdb, 0xf9, 0x0e, 0xdf, 0x85, 0x73, 0xc9, 0xa9, 0xe3, 0x5a, 0xd5, 0xe0, 0x30,
	0x1b, 0x8b, 0x1c, 0x1c, 0xf2, 0x1a, 0x10, 0xa1, 0xec, 0x55, 0x0f, 0xad, 0xd2, 0x50, 0xa2, 0x55,
	0x52, 0x82, 0xcd, 0x7a, 0x6b, 0xd0, 0xd0, 0xcd, 0x8f, 0x7d, 0x40, 0xc2, 0xee, 0xf3, 0xc1, 0x63,
	0xf2, 0x5f, 0x41, 0xa7, 0x5b, 0xb2, 0x82, 0x11, 0x34, 0xb2, 0x68, 0x1f, 0x1d, 0xd7, 0xcf, 0x00,
	0x74, 0x47, 0x12, 0x87, 0x1c, 0x07, 0xc7, 0xfe, 0x7c, 0x3d, 0x11, 0xa5, 0x24, 0xd6, 0xbc, 0x03,
	0xf6, 0xef, 0x5b, 0x9f, 0xbf, 0x6e, 0x65, 0xb9, 0x16, 0x0d, 0x8b, 0x83, 0x1b, 0xe8, 0x76, 0x7d,
	0x5f, 0xa4, 0xc2, 0x66, 0xa8, 0x32, 0x61, 0xba, 0x65, 0xf1, 0x92, 0x52, 0xa1, 0x52, 0x61, 0x34,
	0x5a, 0x2e, 0xd0, 0xa5, 0xd2, 0xb4, 0x5c, 0xcb, 0x7c, 0xde, 0x5d, 0x8e, 0x45, 0x95, 0xf8, 0x5f,
	0xce, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x08, 0x76, 0xdc, 0xfb, 0x00, 0x00, 0x00,
}
