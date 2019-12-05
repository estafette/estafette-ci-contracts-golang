// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/contracts.v1/container_repository_credential_config.proto

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

type ContainerRepositoryCredentialConfig struct {
	Repository           string   `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContainerRepositoryCredentialConfig) Reset()         { *m = ContainerRepositoryCredentialConfig{} }
func (m *ContainerRepositoryCredentialConfig) String() string { return proto.CompactTextString(m) }
func (*ContainerRepositoryCredentialConfig) ProtoMessage()    {}
func (*ContainerRepositoryCredentialConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c6a84d91c7de7a2, []int{0}
}

func (m *ContainerRepositoryCredentialConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContainerRepositoryCredentialConfig.Unmarshal(m, b)
}
func (m *ContainerRepositoryCredentialConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContainerRepositoryCredentialConfig.Marshal(b, m, deterministic)
}
func (m *ContainerRepositoryCredentialConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContainerRepositoryCredentialConfig.Merge(m, src)
}
func (m *ContainerRepositoryCredentialConfig) XXX_Size() int {
	return xxx_messageInfo_ContainerRepositoryCredentialConfig.Size(m)
}
func (m *ContainerRepositoryCredentialConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ContainerRepositoryCredentialConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ContainerRepositoryCredentialConfig proto.InternalMessageInfo

func (m *ContainerRepositoryCredentialConfig) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

func (m *ContainerRepositoryCredentialConfig) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ContainerRepositoryCredentialConfig) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*ContainerRepositoryCredentialConfig)(nil), "contracts.v1.ContainerRepositoryCredentialConfig")
}

func init() {
	proto.RegisterFile("protos/contracts.v1/container_repository_credential_config.proto", fileDescriptor_8c6a84d91c7de7a2)
}

var fileDescriptor_8c6a84d91c7de7a2 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0xd5, 0x22, 0x21, 0xb0, 0x98, 0x32, 0x05, 0x06, 0x84, 0xe8, 0xc2, 0x12, 0x57, 0x11,
	0x2b, 0x03, 0xc2, 0x62, 0x60, 0xed, 0xc0, 0xc0, 0x12, 0xb9, 0xee, 0x35, 0x9c, 0xd4, 0xfa, 0xa2,
	0xbb, 0x6b, 0x11, 0x03, 0xe2, 0xff, 0xf0, 0x2b, 0x51, 0x1d, 0xc5, 0xc9, 0xe6, 0xe7, 0x77, 0xdf,
	0xbb, 0xa7, 0x33, 0xcf, 0x1d, 0x93, 0x92, 0x2c, 0x03, 0x45, 0x65, 0x1f, 0x54, 0xec, 0xb1, 0x4e,
	0xc2, 0x63, 0x04, 0x6e, 0x18, 0x3a, 0x12, 0x54, 0xe2, 0xef, 0x26, 0x30, 0x6c, 0x20, 0x2a, 0xfa,
	0x5d, 0x13, 0x28, 0x6e, 0xb1, 0xb5, 0x09, 0x2d, 0xae, 0xa6, 0xe8, 0xfd, 0x8f, 0x59, 0xb8, 0x81,
	0x5e, 0x65, 0xd8, 0x65, 0xd6, 0x25, 0xb4, 0xb8, 0x35, 0x66, 0x8c, 0x2e, 0x67, 0x77, 0xb3, 0x87,
	0xcb, 0xd5, 0xe4, 0xa7, 0xb8, 0x31, 0x17, 0x07, 0x01, 0x8e, 0x7e, 0x0f, 0xe5, 0x3c, 0xb9, 0x59,
	0x9f, 0xbc, 0xce, 0x8b, 0x7c, 0x11, 0x6f, 0xca, 0xb3, 0xde, 0x1b, 0xf4, 0xcb, 0xaf, 0x59, 0x20,
	0x59, 0x10, 0xf5, 0x5b, 0x50, 0x05, 0x1b, 0xb0, 0x6f, 0x29, 0x76, 0xda, 0xf2, 0xe3, 0xa9, 0x45,
	0xfd, 0x3c, 0xac, 0x6d, 0xa0, 0xfd, 0x32, 0x0f, 0x8f, 0xaf, 0x2a, 0x60, 0xd5, 0x63, 0x55, 0x4b,
	0x3b, 0x1f, 0xdb, 0xf1, 0x3c, 0xcd, 0xb1, 0xfe, 0x9b, 0x5f, 0xbf, 0xe6, 0x05, 0xee, 0xcd, 0xba,
	0x9c, 0xfc, 0x5e, 0xaf, 0xcf, 0x13, 0xf7, 0xf8, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x05, 0xdd, 0x43,
	0xd5, 0x58, 0x01, 0x00, 0x00,
}
