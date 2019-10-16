// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contracts.v1/list_response.proto

package contracts_v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type ListResponse struct {
	Items                []*any.Any  `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Pagination           *Pagination `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2190169add2d2842, []int{0}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetItems() []*any.Any {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *ListResponse) GetPagination() *Pagination {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*ListResponse)(nil), "contracts.v1.ListResponse")
}

func init() { proto.RegisterFile("contracts.v1/list_response.proto", fileDescriptor_2190169add2d2842) }

var fileDescriptor_2190169add2d2842 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8e, 0xb1, 0x6b, 0x86, 0x30,
	0x10, 0xc5, 0xf9, 0x5a, 0xda, 0x21, 0x3a, 0x49, 0x07, 0x2b, 0x14, 0xa4, 0x93, 0x14, 0xbc, 0xa0,
	0x5d, 0xba, 0x15, 0x3b, 0x77, 0x28, 0x8e, 0x5d, 0x24, 0x86, 0x98, 0x06, 0x34, 0x27, 0xe6, 0x14,
	0xfc, 0xef, 0x0b, 0xa6, 0x55, 0xbf, 0xed, 0xee, 0xde, 0xef, 0xbd, 0x7b, 0x2c, 0x95, 0x68, 0x69,
	0x12, 0x92, 0x1c, 0x2c, 0x05, 0xef, 0x8d, 0xa3, 0x66, 0x52, 0x6e, 0x44, 0xeb, 0x14, 0x8c, 0x13,
	0x12, 0x46, 0xe1, 0x99, 0x48, 0x9e, 0xae, 0xf8, 0x51, 0x68, 0x63, 0x05, 0x19, 0xb4, 0x1e, 0x4e,
	0x1e, 0x35, 0xa2, 0xee, 0x15, 0xdf, 0xb6, 0x76, 0xee, 0xb8, 0xb0, 0xab, 0x97, 0x9e, 0x89, 0x85,
	0x9f, 0xc6, 0x51, 0xfd, 0x97, 0x1e, 0xbd, 0xb0, 0x3b, 0x43, 0x6a, 0x70, 0xf1, 0x25, 0xbd, 0xcd,
	0x82, 0xf2, 0x01, 0xbc, 0x15, 0xfe, 0xad, 0x50, 0xd9, 0xb5, 0xf6, 0x48, 0xf4, 0xc6, 0xd8, 0xf1,
	0x2a, 0xbe, 0x49, 0x2f, 0x59, 0x50, 0xc6, 0x70, 0xae, 0x02, 0x5f, 0xbb, 0x5e, 0x9f, 0xd8, 0x8f,
	0xea, 0xfb, 0x5d, 0x1b, 0xfa, 0x99, 0x5b, 0x90, 0x38, 0x70, 0xe5, 0x48, 0x74, 0x8a, 0x48, 0x1d,
	0x53, 0x2e, 0x4d, 0xbe, 0x07, 0xe5, 0x1a, 0x7b, 0x61, 0x35, 0xdf, 0x0f, 0xcd, 0x52, 0xb4, 0xf7,
	0x5b, 0xa3, 0xd7, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x84, 0x48, 0x2b, 0xe8, 0x2b, 0x01, 0x00,
	0x00,
}
