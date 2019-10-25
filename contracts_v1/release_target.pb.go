// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contracts.v1/release_target.proto

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
	return fileDescriptor_77974017ed309ad7, []int{0}
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

func init() { proto.RegisterFile("contracts.v1/release_target.proto", fileDescriptor_77974017ed309ad7) }

var fileDescriptor_77974017ed309ad7 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x4f, 0x4b, 0x03, 0x31,
	0x14, 0xc4, 0xd9, 0x56, 0x14, 0xe3, 0x3f, 0x08, 0x08, 0xeb, 0x9e, 0xaa, 0xbd, 0x14, 0x61, 0x5f,
	0x59, 0xbd, 0xaa, 0xa0, 0xc5, 0x83, 0xd7, 0x45, 0x3c, 0x78, 0x59, 0xd2, 0xf0, 0xba, 0x06, 0xba,
	0x89, 0x24, 0xcf, 0xbd, 0xfa, 0x5d, 0x3c, 0xfa, 0x29, 0xa5, 0xc9, 0x6e, 0xba, 0x42, 0x6f, 0x13,
	0x32, 0xbf, 0x99, 0x37, 0xec, 0x52, 0x1a, 0x4d, 0x56, 0x48, 0x72, 0xd0, 0x16, 0x73, 0x8b, 0x6b,
	0x14, 0x0e, 0x2b, 0x12, 0xb6, 0x46, 0x82, 0x4f, 0x6b, 0xc8, 0xf0, 0xe3, 0xa1, 0x25, 0xcb, 0x76,
	0x01, 0xc1, 0x99, 0x5d, 0x37, 0x42, 0xab, 0x15, 0x3a, 0xda, 0x7c, 0xa1, 0x23, 0xb1, 0x42, 0x22,
	0xac, 0xfa, 0x54, 0x21, 0x49, 0x19, 0x1d, 0xbc, 0x57, 0x3f, 0x09, 0x3b, 0x29, 0xc3, 0xc7, 0xab,
	0x6f, 0xe3, 0x9c, 0xed, 0x69, 0xd1, 0x60, 0x9a, 0x4c, 0x92, 0xd9, 0x61, 0xe9, 0x35, 0xbf, 0x67,
	0x07, 0x81, 0x72, 0xe9, 0x68, 0x32, 0x9e, 0x1d, 0xdd, 0x4c, 0x61, 0xd0, 0x01, 0xcf, 0x7d, 0x47,
	0x97, 0xf4, 0xe8, 0xbd, 0x65, 0xcf, 0xf0, 0x07, 0x76, 0xb6, 0x91, 0x6d, 0xbc, 0xc1, 0xa5, 0x63,
	0x1f, 0x73, 0x0e, 0xc3, 0x19, 0xd0, 0xe1, 0xe5, 0x69, 0x70, 0x77, 0x4f, 0xf7, 0xf4, 0xcd, 0xa6,
	0xca, 0x40, 0x5c, 0x02, 0x52, 0x85, 0xeb, 0xdd, 0x3f, 0xfc, 0xfd, 0xae, 0x56, 0xf4, 0xf1, 0xb5,
	0x04, 0x69, 0x9a, 0xed, 0xec, 0xad, 0xca, 0xa5, 0xca, 0x03, 0x96, 0xd7, 0x66, 0x2d, 0x74, 0x3d,
	0x8f, 0x74, 0xd5, 0x16, 0xbf, 0xa3, 0x8b, 0x38, 0x03, 0x16, 0x2f, 0xb0, 0x88, 0xc9, 0x6f, 0xc5,
	0x72, 0xdf, 0x73, 0xb7, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x98, 0x15, 0xc7, 0xdf, 0xa7, 0x01,
	0x00, 0x00,
}
