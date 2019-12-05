// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/contracts.v1/ci_server_config.proto

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
	return fileDescriptor_7f08483e11d5a074, []int{0}
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
	proto.RegisterFile("protos/contracts.v1/ci_server_config.proto", fileDescriptor_7f08483e11d5a074)
}

var fileDescriptor_7f08483e11d5a074 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xd0, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0x06, 0x70, 0xb5, 0xa0, 0x16, 0xcc, 0x1f, 0xa1, 0x2c, 0xb4, 0x1b, 0x2a, 0x0b, 0x42, 0xc4,
	0x51, 0xc4, 0xca, 0x44, 0xd4, 0xa1, 0x82, 0x09, 0x04, 0x03, 0x8b, 0xe5, 0x98, 0xab, 0xb1, 0x30,
	0xb9, 0xc8, 0xbe, 0x44, 0xea, 0xc4, 0x5b, 0xf0, 0x10, 0x3c, 0x25, 0xca, 0x05, 0xa5, 0xdd, 0x6c,
	0x7f, 0xbf, 0xb3, 0x3f, 0x59, 0x5c, 0xd7, 0x01, 0x09, 0x63, 0x66, 0xb0, 0xa2, 0xa0, 0x0d, 0x45,
	0xd9, 0xe6, 0x99, 0x71, 0x2a, 0x42, 0x68, 0x21, 0x28, 0x83, 0xd5, 0xda, 0x59, 0xc9, 0x28, 0x39,
	0xde, 0x45, 0x8b, 0x9f, 0x91, 0x38, 0x2d, 0x56, 0xcf, 0xec, 0x0a, 0x66, 0xc9, 0x5c, 0x1c, 0x94,
	0x3a, 0x82, 0x6a, 0x82, 0x9f, 0x8d, 0x2e, 0x46, 0x57, 0x87, 0x4f, 0xd3, 0x6e, 0xff, 0x12, 0x7c,
	0x72, 0x23, 0x92, 0xb2, 0x71, 0xfe, 0x1d, 0x82, 0x82, 0x16, 0x2a, 0x8a, 0x8c, 0xc6, 0x8c, 0xce,
	0xfe, 0x93, 0x25, 0x07, 0x9d, 0x5e, 0x88, 0x93, 0x1a, 0x23, 0x29, 0x8f, 0xb6, 0x87, 0x7b, 0x0c,
	0x8f, 0xba, 0xc3, 0x47, 0xb4, 0x6c, 0xce, 0xc5, 0x54, 0xd7, 0x4e, 0x7d, 0xc2, 0x66, 0xb6, 0xcf,
	0xe9, 0x44, 0xd7, 0xee, 0x01, 0x36, 0xf7, 0xdf, 0xe2, 0xd2, 0xa1, 0x84, 0x48, 0x7a, 0x0d, 0x44,
	0x20, 0x8d, 0xeb, 0xfb, 0x47, 0xb9, 0xdb, 0xff, 0xed, 0xce, 0x3a, 0xfa, 0x68, 0x4a, 0x69, 0xf0,
	0x2b, 0x1b, 0xf0, 0x76, 0x95, 0x1a, 0x97, 0xf6, 0x63, 0xa9, 0x45, 0xaf, 0x2b, 0xbb, 0xfd, 0x22,
	0xd5, 0xe6, 0xbf, 0xe3, 0xf9, 0x72, 0x78, 0xa0, 0x58, 0xc9, 0x62, 0xb8, 0xf9, 0x35, 0x2f, 0x27,
	0x3c, 0x77, 0xfb, 0x17, 0x00, 0x00, 0xff, 0xff, 0x13, 0x0b, 0x32, 0x94, 0x5c, 0x01, 0x00, 0x00,
}
