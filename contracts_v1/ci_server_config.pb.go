// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contracts.v1/ci_server_config.proto

package contracts_v1

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

type CIServerConfig struct {
	BaseUrl              string   `protobuf:"bytes,1,opt,name=base_url,json=baseUrl,proto3" json:"base_url,omitempty"`
	BuilderEventsUrl     string   `protobuf:"bytes,2,opt,name=builder_events_url,json=builderEventsUrl,proto3" json:"builder_events_url,omitempty"`
	PostLogsUrl          string   `protobuf:"bytes,3,opt,name=post_logs_url,json=postLogsUrl,proto3" json:"post_logs_url,omitempty"`
	ApiKey               string   `protobuf:"bytes,4,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CIServerConfig) Reset()         { *m = CIServerConfig{} }
func (m *CIServerConfig) String() string { return proto.CompactTextString(m) }
func (*CIServerConfig) ProtoMessage()    {}
func (*CIServerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_400d2efaa5f2b9b9, []int{0}
}

func (m *CIServerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CIServerConfig.Unmarshal(m, b)
}
func (m *CIServerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CIServerConfig.Marshal(b, m, deterministic)
}
func (m *CIServerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CIServerConfig.Merge(m, src)
}
func (m *CIServerConfig) XXX_Size() int {
	return xxx_messageInfo_CIServerConfig.Size(m)
}
func (m *CIServerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CIServerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CIServerConfig proto.InternalMessageInfo

func (m *CIServerConfig) GetBaseUrl() string {
	if m != nil {
		return m.BaseUrl
	}
	return ""
}

func (m *CIServerConfig) GetBuilderEventsUrl() string {
	if m != nil {
		return m.BuilderEventsUrl
	}
	return ""
}

func (m *CIServerConfig) GetPostLogsUrl() string {
	if m != nil {
		return m.PostLogsUrl
	}
	return ""
}

func (m *CIServerConfig) GetApiKey() string {
	if m != nil {
		return m.ApiKey
	}
	return ""
}

func init() {
	proto.RegisterType((*CIServerConfig)(nil), "contracts.v1.CIServerConfig")
}

func init() {
	proto.RegisterFile("contracts.v1/ci_server_config.proto", fileDescriptor_400d2efaa5f2b9b9)
}

var fileDescriptor_400d2efaa5f2b9b9 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xcf, 0xbf, 0x4a, 0x04, 0x31,
	0x10, 0x06, 0x70, 0x56, 0xe5, 0x4e, 0xe3, 0x1f, 0x24, 0x8d, 0x67, 0x27, 0x67, 0x63, 0xe1, 0xee,
	0x72, 0xf8, 0x00, 0xa2, 0x87, 0x85, 0x68, 0xa5, 0xd8, 0xd8, 0x84, 0x24, 0xce, 0xc5, 0x60, 0xcc,
	0x84, 0x64, 0x36, 0x70, 0x2f, 0xe2, 0xf3, 0xca, 0x8e, 0x72, 0xda, 0xcd, 0x7c, 0xdf, 0xaf, 0xf9,
	0xc4, 0xb9, 0xc5, 0x48, 0x59, 0x5b, 0x2a, 0x5d, 0x5d, 0xf4, 0xd6, 0xab, 0x02, 0xb9, 0x42, 0x56,
	0x16, 0xe3, 0xca, 0xbb, 0x2e, 0x65, 0x24, 0x94, 0x07, 0xff, 0xd1, 0xfc, 0xab, 0x11, 0x47, 0xcb,
	0xfb, 0x67, 0x76, 0x4b, 0x66, 0xf2, 0x54, 0xec, 0x1a, 0x5d, 0x40, 0x0d, 0x39, 0xcc, 0x9a, 0xb3,
	0xe6, 0x62, 0xef, 0x69, 0x3a, 0xfe, 0x2f, 0x39, 0xc8, 0x4b, 0x21, 0xcd, 0xe0, 0xc3, 0x1b, 0x64,
	0x05, 0x15, 0x22, 0x15, 0x46, 0x5b, 0x8c, 0x8e, 0x7f, 0x9b, 0x3b, 0x2e, 0x46, 0x3d, 0x17, 0x87,
	0x09, 0x0b, 0xa9, 0x80, 0xee, 0x07, 0x6e, 0x33, 0xdc, 0x1f, 0xc3, 0x47, 0x74, 0x6c, 0x4e, 0xc4,
	0x54, 0x27, 0xaf, 0x3e, 0x60, 0x3d, 0xdb, 0xe1, 0x76, 0xa2, 0x93, 0x7f, 0x80, 0xf5, 0xed, 0xcd,
	0xeb, 0xb5, 0xf3, 0xf4, 0x3e, 0x98, 0xce, 0xe2, 0x67, 0x0f, 0x85, 0xf4, 0x0a, 0x88, 0xe0, 0xef,
	0x6a, 0xad, 0x6f, 0x37, 0x53, 0x5a, 0x87, 0x41, 0x47, 0xd7, 0x6f, 0x02, 0x55, 0x17, 0x66, 0xc2,
	0x83, 0xaf, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0xce, 0xe7, 0x02, 0x41, 0x17, 0x01, 0x00, 0x00,
}
