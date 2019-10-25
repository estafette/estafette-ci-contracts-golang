// Code generated by protoc-gen-go. DO NOT EDIT.
// source: manifest.v1/estafette_pipeline_event.proto

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

type EstafettePipelineEvent struct {
	BuildVersion         string   `protobuf:"bytes,1,opt,name=build_version,json=buildVersion,proto3" json:"build_version,omitempty"`
	RepoSource           string   `protobuf:"bytes,2,opt,name=repo_source,json=repoSource,proto3" json:"repo_source,omitempty"`
	RepoOwner            string   `protobuf:"bytes,3,opt,name=repo_owner,json=repoOwner,proto3" json:"repo_owner,omitempty"`
	RepoName             string   `protobuf:"bytes,4,opt,name=repo_name,json=repoName,proto3" json:"repo_name,omitempty"`
	RepoBranch           string   `protobuf:"bytes,5,opt,name=repo_branch,json=repoBranch,proto3" json:"repo_branch,omitempty"`
	Status               string   `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	Event                string   `protobuf:"bytes,7,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EstafettePipelineEvent) Reset()         { *m = EstafettePipelineEvent{} }
func (m *EstafettePipelineEvent) String() string { return proto.CompactTextString(m) }
func (*EstafettePipelineEvent) ProtoMessage()    {}
func (*EstafettePipelineEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0cd8e4dfb4a4d7e, []int{0}
}

func (m *EstafettePipelineEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EstafettePipelineEvent.Unmarshal(m, b)
}
func (m *EstafettePipelineEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EstafettePipelineEvent.Marshal(b, m, deterministic)
}
func (m *EstafettePipelineEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EstafettePipelineEvent.Merge(m, src)
}
func (m *EstafettePipelineEvent) XXX_Size() int {
	return xxx_messageInfo_EstafettePipelineEvent.Size(m)
}
func (m *EstafettePipelineEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_EstafettePipelineEvent.DiscardUnknown(m)
}

var xxx_messageInfo_EstafettePipelineEvent proto.InternalMessageInfo

func (m *EstafettePipelineEvent) GetBuildVersion() string {
	if m != nil {
		return m.BuildVersion
	}
	return ""
}

func (m *EstafettePipelineEvent) GetRepoSource() string {
	if m != nil {
		return m.RepoSource
	}
	return ""
}

func (m *EstafettePipelineEvent) GetRepoOwner() string {
	if m != nil {
		return m.RepoOwner
	}
	return ""
}

func (m *EstafettePipelineEvent) GetRepoName() string {
	if m != nil {
		return m.RepoName
	}
	return ""
}

func (m *EstafettePipelineEvent) GetRepoBranch() string {
	if m != nil {
		return m.RepoBranch
	}
	return ""
}

func (m *EstafettePipelineEvent) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *EstafettePipelineEvent) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func init() {
	proto.RegisterType((*EstafettePipelineEvent)(nil), "manifest.v1.EstafettePipelineEvent")
}

func init() {
	proto.RegisterFile("manifest.v1/estafette_pipeline_event.proto", fileDescriptor_e0cd8e4dfb4a4d7e)
}

var fileDescriptor_e0cd8e4dfb4a4d7e = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0xe9, 0x74, 0xd5, 0x65, 0x7a, 0x09, 0x32, 0x02, 0x43, 0x14, 0xbd, 0x88, 0xd0, 0x8c,
	0xe2, 0xd1, 0xdb, 0x64, 0x07, 0x0f, 0xfe, 0x41, 0x61, 0x07, 0x2f, 0x21, 0xad, 0xef, 0xba, 0x40,
	0x9b, 0x94, 0x24, 0xad, 0xdf, 0xc9, 0x6f, 0xe6, 0xb7, 0x90, 0xbc, 0x5b, 0xdb, 0xdd, 0xf2, 0xfc,
	0x9e, 0xe7, 0x7d, 0x20, 0x0f, 0xb9, 0xaf, 0xa4, 0x56, 0x1b, 0x70, 0x9e, 0xb7, 0xe9, 0x02, 0x9c,
	0x97, 0x1b, 0xf0, 0x1e, 0x44, 0xad, 0x6a, 0x28, 0x95, 0x06, 0x01, 0x2d, 0x68, 0xcf, 0x6b, 0x6b,
	0xbc, 0xa1, 0xd3, 0x83, 0xec, 0xcd, 0x5f, 0x44, 0x66, 0xab, 0x2e, 0xff, 0xbe, 0x8f, 0xaf, 0x42,
	0x9a, 0xde, 0x92, 0xf3, 0xac, 0x51, 0xe5, 0xb7, 0x68, 0xc1, 0x3a, 0x65, 0x34, 0x8b, 0xae, 0xa3,
	0xbb, 0xc9, 0xc7, 0x19, 0xc2, 0xf5, 0x8e, 0xd1, 0x2b, 0x32, 0xb5, 0x50, 0x1b, 0xe1, 0x4c, 0x63,
	0x73, 0x60, 0x23, 0x8c, 0x90, 0x80, 0x3e, 0x91, 0xd0, 0x4b, 0x82, 0x4a, 0x98, 0x1f, 0x0d, 0x96,
	0x1d, 0xa1, 0x3f, 0x09, 0xe4, 0x2d, 0x00, 0x3a, 0x27, 0x28, 0x84, 0x96, 0x15, 0xb0, 0x63, 0x74,
	0x4f, 0x03, 0x78, 0x95, 0x15, 0xf4, 0xe5, 0x99, 0x95, 0x3a, 0xdf, 0xb2, 0xf1, 0x50, 0xbe, 0x44,
	0x42, 0x67, 0x24, 0x76, 0x5e, 0xfa, 0xc6, 0xb1, 0x18, 0xbd, 0xbd, 0xa2, 0x17, 0x64, 0x8c, 0x3f,
	0x66, 0x27, 0x88, 0x77, 0x62, 0xd9, 0x90, 0xb9, 0x32, 0xbc, 0x5f, 0x87, 0xe7, 0x8a, 0x1f, 0x4c,
	0xf1, 0xf5, 0x58, 0x28, 0xbf, 0x6d, 0x32, 0x9e, 0x9b, 0x6a, 0x98, 0x70, 0x78, 0x25, 0xb9, 0x4a,
	0x70, 0x40, 0x97, 0x14, 0xa6, 0x94, 0xba, 0x58, 0x74, 0xc7, 0xa2, 0x4d, 0x7f, 0x47, 0xac, 0x5f,
	0x91, 0x3f, 0x3d, 0xf3, 0x97, 0xae, 0x77, 0x9d, 0x66, 0x31, 0x5e, 0x3d, 0xfc, 0x07, 0x00, 0x00,
	0xff, 0xff, 0xf9, 0x5a, 0xff, 0xaf, 0xa4, 0x01, 0x00, 0x00,
}
