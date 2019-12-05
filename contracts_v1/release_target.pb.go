// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/contracts.v1/release_target.proto

package contracts_v1

import (
	fmt "fmt"
	manifest_v1 "github.com/estafette/estafette-ci-protos-golang/manifest_v1"
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

type ReleaseTarget struct {
	Name                 string                                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Actions              []*manifest_v1.EstafetteReleaseAction `protobuf:"bytes,2,rep,name=actions,proto3" json:"actions,omitempty"`
	ActiveReleases       []*Release                            `protobuf:"bytes,3,rep,name=active_releases,json=activeReleases,proto3" json:"active_releases,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *ReleaseTarget) Reset()         { *m = ReleaseTarget{} }
func (m *ReleaseTarget) String() string { return proto.CompactTextString(m) }
func (*ReleaseTarget) ProtoMessage()    {}
func (*ReleaseTarget) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e75a91d5ab21eba, []int{0}
}

func (m *ReleaseTarget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleaseTarget.Unmarshal(m, b)
}
func (m *ReleaseTarget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleaseTarget.Marshal(b, m, deterministic)
}
func (m *ReleaseTarget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleaseTarget.Merge(m, src)
}
func (m *ReleaseTarget) XXX_Size() int {
	return xxx_messageInfo_ReleaseTarget.Size(m)
}
func (m *ReleaseTarget) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleaseTarget.DiscardUnknown(m)
}

var xxx_messageInfo_ReleaseTarget proto.InternalMessageInfo

func (m *ReleaseTarget) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReleaseTarget) GetActions() []*manifest_v1.EstafetteReleaseAction {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *ReleaseTarget) GetActiveReleases() []*Release {
	if m != nil {
		return m.ActiveReleases
	}
	return nil
}

func init() {
	proto.RegisterType((*ReleaseTarget)(nil), "contracts.v1.ReleaseTarget")
}

func init() {
	proto.RegisterFile("protos/contracts.v1/release_target.proto", fileDescriptor_7e75a91d5ab21eba)
}

var fileDescriptor_7e75a91d5ab21eba = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x3d, 0x4b, 0xc4, 0x40,
	0x10, 0x86, 0xc9, 0x9d, 0x28, 0xae, 0x5f, 0xb0, 0x20, 0xc4, 0xab, 0x4e, 0xaf, 0x49, 0x93, 0x39,
	0xa2, 0xad, 0x0a, 0x7a, 0x58, 0xd8, 0x06, 0xb1, 0xb0, 0x09, 0x7b, 0xcb, 0x5c, 0x5c, 0xb8, 0xec,
	0x4a, 0x66, 0x4c, 0xeb, 0x7f, 0xb1, 0xf4, 0x57, 0xca, 0xed, 0xe6, 0xcb, 0xc6, 0x6e, 0x42, 0x9e,
	0x67, 0xde, 0x79, 0x57, 0x24, 0x1f, 0xb5, 0x63, 0x47, 0x4b, 0xed, 0x2c, 0xd7, 0x4a, 0x33, 0x41,
	0x93, 0x2d, 0x6b, 0xdc, 0xa2, 0x22, 0x2c, 0x58, 0xd5, 0x25, 0x32, 0x78, 0x44, 0x1e, 0x8f, 0x91,
	0xd9, 0xe5, 0x3f, 0x5e, 0x10, 0x66, 0x59, 0x8b, 0x54, 0xca, 0x9a, 0x0d, 0x12, 0xef, 0x08, 0x24,
	0x56, 0x1b, 0x64, 0xc6, 0xa2, 0xcb, 0x50, 0x9a, 0x8d, 0xb3, 0x41, 0xb9, 0xfa, 0x8e, 0xc4, 0x49,
	0x1e, 0x7e, 0xbc, 0xf8, 0x6c, 0x29, 0xc5, 0x9e, 0x55, 0x15, 0xc6, 0xd1, 0x3c, 0x4a, 0x0e, 0x73,
	0x3f, 0xcb, 0x3b, 0x71, 0x10, 0x2c, 0x8a, 0x27, 0xf3, 0x69, 0x72, 0x74, 0xbd, 0x80, 0x51, 0x06,
	0x3c, 0x75, 0x19, 0xed, 0xa6, 0x07, 0xcf, 0xe6, 0x9d, 0x23, 0xef, 0xc5, 0xd9, 0x6e, 0x6c, 0xfa,
	0x1b, 0x28, 0x9e, 0xfa, 0x35, 0xe7, 0x30, 0x6e, 0x03, 0xad, 0x9e, 0x9f, 0x06, 0xba, 0xfd, 0xa4,
	0xc7, 0x2f, 0xb1, 0x30, 0x0e, 0xfa, 0x26, 0xa0, 0x4d, 0xb8, 0x9e, 0xfe, 0xe8, 0x6f, 0xb7, 0xa5,
	0xe1, 0xf7, 0xcf, 0x35, 0x68, 0x57, 0x0d, 0xb5, 0x87, 0x29, 0xd5, 0x26, 0x0d, 0x5a, 0x5a, 0xba,
	0xad, 0xb2, 0xe5, 0xf0, 0x94, 0x45, 0x93, 0xfd, 0x4c, 0x2e, 0xfa, 0x1a, 0xb0, 0x7a, 0x86, 0x55,
	0xbf, 0xf9, 0x35, 0x5b, 0xef, 0x7b, 0xef, 0xe6, 0x37, 0x00, 0x00, 0xff, 0xff, 0x9a, 0xd1, 0x90,
	0x59, 0xbc, 0x01, 0x00, 0x00,
}
