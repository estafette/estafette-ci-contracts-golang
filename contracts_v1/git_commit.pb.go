// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contracts.v1/git_commit.proto

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

type GitCommit struct {
	Message              string     `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Author               *GitAuthor `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GitCommit) Reset()         { *m = GitCommit{} }
func (m *GitCommit) String() string { return proto.CompactTextString(m) }
func (*GitCommit) ProtoMessage()    {}
func (*GitCommit) Descriptor() ([]byte, []int) {
	return fileDescriptor_89b084dc1a152454, []int{0}
}

func (m *GitCommit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GitCommit.Unmarshal(m, b)
}
func (m *GitCommit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GitCommit.Marshal(b, m, deterministic)
}
func (m *GitCommit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GitCommit.Merge(m, src)
}
func (m *GitCommit) XXX_Size() int {
	return xxx_messageInfo_GitCommit.Size(m)
}
func (m *GitCommit) XXX_DiscardUnknown() {
	xxx_messageInfo_GitCommit.DiscardUnknown(m)
}

var xxx_messageInfo_GitCommit proto.InternalMessageInfo

func (m *GitCommit) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *GitCommit) GetAuthor() *GitAuthor {
	if m != nil {
		return m.Author
	}
	return nil
}

func init() {
	proto.RegisterType((*GitCommit)(nil), "contracts.v1.GitCommit")
}

func init() { proto.RegisterFile("contracts.v1/git_commit.proto", fileDescriptor_89b084dc1a152454) }

var fileDescriptor_89b084dc1a152454 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0xce, 0xcf, 0x2b,
	0x29, 0x4a, 0x4c, 0x2e, 0x29, 0xd6, 0x2b, 0x33, 0xd4, 0x4f, 0xcf, 0x2c, 0x89, 0x4f, 0xce, 0xcf,
	0xcd, 0xcd, 0x2c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x41, 0x96, 0x96, 0xc2, 0x54,
	0x9c, 0x58, 0x5a, 0x92, 0x91, 0x5f, 0x04, 0x51, 0xac, 0x14, 0xc6, 0xc5, 0xe9, 0x9e, 0x59, 0xe2,
	0x0c, 0xd6, 0x2f, 0x24, 0xc1, 0xc5, 0x9e, 0x9b, 0x5a, 0x5c, 0x9c, 0x98, 0x9e, 0x2a, 0xc1, 0xa8,
	0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe3, 0x0a, 0xe9, 0x73, 0xb1, 0x41, 0xb4, 0x49, 0x30, 0x29, 0x30,
	0x6a, 0x70, 0x1b, 0x89, 0xeb, 0x21, 0x1b, 0xab, 0xe7, 0x9e, 0x59, 0xe2, 0x08, 0x96, 0x0e, 0x82,
	0x2a, 0x73, 0xaa, 0xe7, 0x52, 0xce, 0xcc, 0xd7, 0x4b, 0x2d, 0x2e, 0x49, 0x4c, 0x4b, 0x2d, 0x29,
	0x49, 0xd5, 0x4b, 0xce, 0x84, 0x58, 0x58, 0x8c, 0xa2, 0x31, 0xca, 0x26, 0x3d, 0xb3, 0x24, 0xa3,
	0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f, 0xae, 0x18, 0xc1, 0xd2, 0x4d, 0xce, 0xd4, 0x85, 0x68,
	0xd3, 0x4d, 0xcf, 0xcf, 0x49, 0xcc, 0x4b, 0xd7, 0x87, 0xeb, 0x8e, 0x2f, 0x33, 0x5c, 0xc5, 0x24,
	0xe9, 0x0a, 0xb7, 0xc0, 0xd9, 0x53, 0xcf, 0x19, 0x6e, 0x72, 0x98, 0x61, 0x12, 0x1b, 0x58, 0x9f,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xa5, 0xeb, 0x72, 0xba, 0x2d, 0x01, 0x00, 0x00,
}
