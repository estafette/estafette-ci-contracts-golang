// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contracts.v1/build_log.proto

package contracts_v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type BuildLog struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RepoSource           string               `protobuf:"bytes,2,opt,name=repo_source,json=repoSource,proto3" json:"repo_source,omitempty"`
	RepoOwner            string               `protobuf:"bytes,3,opt,name=repo_owner,json=repoOwner,proto3" json:"repo_owner,omitempty"`
	RepoName             string               `protobuf:"bytes,4,opt,name=repo_name,json=repoName,proto3" json:"repo_name,omitempty"`
	RepoBranch           string               `protobuf:"bytes,5,opt,name=repo_branch,json=repoBranch,proto3" json:"repo_branch,omitempty"`
	RepoRevision         string               `protobuf:"bytes,6,opt,name=repo_revision,json=repoRevision,proto3" json:"repo_revision,omitempty"`
	BuildId              string               `protobuf:"bytes,7,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
	Steps                []*BuildLogStep      `protobuf:"bytes,8,rep,name=steps,proto3" json:"steps,omitempty"`
	InsertedAt           *timestamp.Timestamp `protobuf:"bytes,9,opt,name=inserted_at,json=insertedAt,proto3" json:"inserted_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *BuildLog) Reset()         { *m = BuildLog{} }
func (m *BuildLog) String() string { return proto.CompactTextString(m) }
func (*BuildLog) ProtoMessage()    {}
func (*BuildLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ecc9a558dc23734, []int{0}
}

func (m *BuildLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildLog.Unmarshal(m, b)
}
func (m *BuildLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildLog.Marshal(b, m, deterministic)
}
func (m *BuildLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildLog.Merge(m, src)
}
func (m *BuildLog) XXX_Size() int {
	return xxx_messageInfo_BuildLog.Size(m)
}
func (m *BuildLog) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildLog.DiscardUnknown(m)
}

var xxx_messageInfo_BuildLog proto.InternalMessageInfo

func (m *BuildLog) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *BuildLog) GetRepoSource() string {
	if m != nil {
		return m.RepoSource
	}
	return ""
}

func (m *BuildLog) GetRepoOwner() string {
	if m != nil {
		return m.RepoOwner
	}
	return ""
}

func (m *BuildLog) GetRepoName() string {
	if m != nil {
		return m.RepoName
	}
	return ""
}

func (m *BuildLog) GetRepoBranch() string {
	if m != nil {
		return m.RepoBranch
	}
	return ""
}

func (m *BuildLog) GetRepoRevision() string {
	if m != nil {
		return m.RepoRevision
	}
	return ""
}

func (m *BuildLog) GetBuildId() string {
	if m != nil {
		return m.BuildId
	}
	return ""
}

func (m *BuildLog) GetSteps() []*BuildLogStep {
	if m != nil {
		return m.Steps
	}
	return nil
}

func (m *BuildLog) GetInsertedAt() *timestamp.Timestamp {
	if m != nil {
		return m.InsertedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*BuildLog)(nil), "contracts.v1.BuildLog")
}

func init() { proto.RegisterFile("contracts.v1/build_log.proto", fileDescriptor_5ecc9a558dc23734) }

var fileDescriptor_5ecc9a558dc23734 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x4e, 0xf2, 0x40,
	0x14, 0xc5, 0x43, 0xf9, 0x80, 0x32, 0xf0, 0xb9, 0x98, 0xd5, 0x88, 0x1a, 0x50, 0x37, 0x6c, 0x68,
	0x45, 0x97, 0x2e, 0x0c, 0xec, 0x4c, 0x8c, 0x26, 0xc5, 0x95, 0x9b, 0xa6, 0x7f, 0x2e, 0xc3, 0x24,
	0xed, 0x4c, 0x33, 0x33, 0xc5, 0x47, 0xf4, 0xb5, 0xcc, 0xdc, 0x52, 0x60, 0xe1, 0xae, 0xe7, 0x77,
	0xce, 0xbd, 0x9d, 0x7b, 0xc8, 0x75, 0xa6, 0xa4, 0xd5, 0x49, 0x66, 0x4d, 0xb0, 0x5f, 0x86, 0x69,
	0x2d, 0x8a, 0x3c, 0x2e, 0x14, 0x0f, 0x2a, 0xad, 0xac, 0xa2, 0xe3, 0x73, 0x77, 0x72, 0xfb, 0x77,
	0x36, 0x36, 0x16, 0xaa, 0x66, 0x60, 0x32, 0xe5, 0x4a, 0xf1, 0x02, 0x42, 0x54, 0x69, 0xbd, 0x0d,
	0xad, 0x28, 0xc1, 0xd8, 0xa4, 0x3c, 0x04, 0xee, 0x7e, 0x3c, 0xe2, 0xaf, 0xdd, 0xe4, 0x9b, 0xe2,
	0xf4, 0x82, 0x78, 0x22, 0x67, 0x9d, 0x59, 0x67, 0x3e, 0x8c, 0x3c, 0x91, 0xd3, 0x29, 0x19, 0x69,
	0xa8, 0x54, 0x6c, 0x54, 0xad, 0x33, 0x60, 0x1e, 0x1a, 0xc4, 0xa1, 0x0d, 0x12, 0x7a, 0x43, 0x50,
	0xc5, 0xea, 0x5b, 0x82, 0x66, 0x5d, 0xf4, 0x87, 0x8e, 0x7c, 0x38, 0x40, 0xaf, 0x08, 0x8a, 0x58,
	0x26, 0x25, 0xb0, 0x7f, 0xe8, 0xfa, 0x0e, 0xbc, 0x27, 0x25, 0x1c, 0x97, 0xa7, 0x3a, 0x91, 0xd9,
	0x8e, 0xf5, 0x4e, 0xcb, 0xd7, 0x48, 0xe8, 0x3d, 0xf9, 0x8f, 0x01, 0x0d, 0x7b, 0x61, 0x84, 0x92,
	0xac, 0x8f, 0x91, 0xb1, 0x83, 0xd1, 0x81, 0xd1, 0x4b, 0xe2, 0x37, 0x87, 0x8b, 0x9c, 0x0d, 0xd0,
	0x1f, 0xa0, 0x7e, 0xcd, 0xe9, 0x03, 0xe9, 0xb9, 0x26, 0x0c, 0xf3, 0x67, 0xdd, 0xf9, 0xe8, 0x71,
	0x12, 0x9c, 0xd7, 0x15, 0xb4, 0x47, 0x6f, 0x2c, 0x54, 0x51, 0x13, 0xa4, 0xcf, 0x64, 0x24, 0xa4,
	0x01, 0x6d, 0x21, 0x8f, 0x13, 0xcb, 0x86, 0xb3, 0x0e, 0xce, 0x35, 0x1d, 0x06, 0x6d, 0x87, 0xc1,
	0x67, 0xdb, 0x61, 0x44, 0xda, 0xf8, 0xca, 0xae, 0x57, 0x5f, 0x2f, 0x5c, 0xd8, 0x5d, 0x9d, 0x06,
	0x99, 0x2a, 0x43, 0x17, 0xd8, 0x82, 0xb5, 0x70, 0xfa, 0x5a, 0x64, 0x62, 0x71, 0x7c, 0xc2, 0x82,
	0xab, 0x22, 0x91, 0x3c, 0x3c, 0x82, 0x78, 0xbf, 0x4c, 0xfb, 0xf8, 0x8b, 0xa7, 0xdf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xe6, 0x3d, 0x53, 0xef, 0x05, 0x02, 0x00, 0x00,
}