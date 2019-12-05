// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/manifest.v1/estafette_stage.proto

package manifest_v1

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

type EstafetteStage struct {
	Name                 string              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ContainerImage       string              `protobuf:"bytes,2,opt,name=container_image,json=containerImage,proto3" json:"container_image,omitempty"`
	Shell                string              `protobuf:"bytes,3,opt,name=shell,proto3" json:"shell,omitempty"`
	WorkingDirectory     string              `protobuf:"bytes,4,opt,name=working_directory,json=workingDirectory,proto3" json:"working_directory,omitempty"`
	Commands             []string            `protobuf:"bytes,5,rep,name=commands,proto3" json:"commands,omitempty"`
	When                 string              `protobuf:"bytes,6,opt,name=when,proto3" json:"when,omitempty"`
	EnvVars              map[string]string   `protobuf:"bytes,7,rep,name=env_vars,json=envVars,proto3" json:"env_vars,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	AutoInjected         bool                `protobuf:"varint,8,opt,name=auto_injected,json=autoInjected,proto3" json:"auto_injected,omitempty"`
	Retries              int64               `protobuf:"varint,9,opt,name=retries,proto3" json:"retries,omitempty"`
	CustomProperties     map[string]*any.Any `protobuf:"bytes,10,rep,name=custom_properties,json=customProperties,proto3" json:"custom_properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *EstafetteStage) Reset()         { *m = EstafetteStage{} }
func (m *EstafetteStage) String() string { return proto.CompactTextString(m) }
func (*EstafetteStage) ProtoMessage()    {}
func (*EstafetteStage) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e3645acd16fad94, []int{0}
}

func (m *EstafetteStage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EstafetteStage.Unmarshal(m, b)
}
func (m *EstafetteStage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EstafetteStage.Marshal(b, m, deterministic)
}
func (m *EstafetteStage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EstafetteStage.Merge(m, src)
}
func (m *EstafetteStage) XXX_Size() int {
	return xxx_messageInfo_EstafetteStage.Size(m)
}
func (m *EstafetteStage) XXX_DiscardUnknown() {
	xxx_messageInfo_EstafetteStage.DiscardUnknown(m)
}

var xxx_messageInfo_EstafetteStage proto.InternalMessageInfo

func (m *EstafetteStage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EstafetteStage) GetContainerImage() string {
	if m != nil {
		return m.ContainerImage
	}
	return ""
}

func (m *EstafetteStage) GetShell() string {
	if m != nil {
		return m.Shell
	}
	return ""
}

func (m *EstafetteStage) GetWorkingDirectory() string {
	if m != nil {
		return m.WorkingDirectory
	}
	return ""
}

func (m *EstafetteStage) GetCommands() []string {
	if m != nil {
		return m.Commands
	}
	return nil
}

func (m *EstafetteStage) GetWhen() string {
	if m != nil {
		return m.When
	}
	return ""
}

func (m *EstafetteStage) GetEnvVars() map[string]string {
	if m != nil {
		return m.EnvVars
	}
	return nil
}

func (m *EstafetteStage) GetAutoInjected() bool {
	if m != nil {
		return m.AutoInjected
	}
	return false
}

func (m *EstafetteStage) GetRetries() int64 {
	if m != nil {
		return m.Retries
	}
	return 0
}

func (m *EstafetteStage) GetCustomProperties() map[string]*any.Any {
	if m != nil {
		return m.CustomProperties
	}
	return nil
}

func init() {
	proto.RegisterType((*EstafetteStage)(nil), "manifest.v1.EstafetteStage")
	proto.RegisterMapType((map[string]*any.Any)(nil), "manifest.v1.EstafetteStage.CustomPropertiesEntry")
	proto.RegisterMapType((map[string]string)(nil), "manifest.v1.EstafetteStage.EnvVarsEntry")
}

func init() {
	proto.RegisterFile("protos/manifest.v1/estafette_stage.proto", fileDescriptor_2e3645acd16fad94)
}

var fileDescriptor_2e3645acd16fad94 = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x56, 0x9a, 0x6e, 0x6d, 0xdd, 0x31, 0x3a, 0x6b, 0x48, 0xa6, 0xa7, 0x68, 0x1c, 0x88, 0x40,
	0x75, 0xd5, 0x71, 0x41, 0xe3, 0x04, 0xa5, 0x87, 0x1e, 0x90, 0x50, 0x90, 0x26, 0xc1, 0x81, 0xc8,
	0x4d, 0x5f, 0x53, 0xb3, 0xc4, 0xae, 0x6c, 0x27, 0x53, 0x24, 0x7e, 0x11, 0x7f, 0x81, 0x3f, 0x87,
	0x62, 0x37, 0x59, 0x40, 0x13, 0xb7, 0xf7, 0xbe, 0xf7, 0xbd, 0x2f, 0xdf, 0x8b, 0x3f, 0x14, 0x1e,
	0x94, 0x34, 0x52, 0xcf, 0x73, 0x26, 0xf8, 0x0e, 0xb4, 0xa1, 0xe5, 0x62, 0x0e, 0xda, 0xb0, 0x1d,
	0x18, 0x03, 0xb1, 0x36, 0x2c, 0x05, 0x6a, 0x29, 0x78, 0xdc, 0xa1, 0x4c, 0x9f, 0xa7, 0x52, 0xa6,
	0x19, 0xcc, 0xed, 0x68, 0x53, 0xec, 0xe6, 0x4c, 0x54, 0x8e, 0x77, 0xf5, 0xbb, 0x8f, 0xce, 0x57,
	0x8d, 0xc2, 0x97, 0x5a, 0x00, 0x63, 0xd4, 0x17, 0x2c, 0x07, 0xe2, 0x05, 0x5e, 0x38, 0x8a, 0x6c,
	0x8d, 0x5f, 0xa2, 0xa7, 0x89, 0x14, 0x86, 0x71, 0x01, 0x2a, 0xe6, 0x39, 0x4b, 0x81, 0xf4, 0xec,
	0xf8, 0xbc, 0x85, 0xd7, 0x35, 0x8a, 0x2f, 0xd1, 0x89, 0xde, 0x43, 0x96, 0x11, 0xdf, 0x8e, 0x5d,
	0x83, 0x5f, 0xa3, 0x8b, 0x7b, 0xa9, 0xee, 0xb8, 0x48, 0xe3, 0x2d, 0x57, 0x90, 0x18, 0xa9, 0x2a,
	0xd2, 0xb7, 0x8c, 0xc9, 0x71, 0xf0, 0xb1, 0xc1, 0xf1, 0x14, 0x0d, 0x13, 0x99, 0xe7, 0x4c, 0x6c,
	0x35, 0x39, 0x09, 0xfc, 0x70, 0x14, 0xb5, 0x7d, 0xed, 0xed, 0x7e, 0x0f, 0x82, 0x9c, 0x3a, 0x6f,
	0x75, 0x8d, 0x97, 0x68, 0x08, 0xa2, 0x8c, 0x4b, 0xa6, 0x34, 0x19, 0x04, 0x7e, 0x38, 0xbe, 0x0e,
	0x69, 0xe7, 0x7a, 0xfa, 0xf7, 0x79, 0x74, 0x25, 0xca, 0x5b, 0xa6, 0xf4, 0x4a, 0x18, 0x55, 0x45,
	0x03, 0x70, 0x1d, 0x7e, 0x81, 0x9e, 0xb0, 0xc2, 0xc8, 0x98, 0x8b, 0x1f, 0x90, 0x18, 0xd8, 0x92,
	0x61, 0xe0, 0x85, 0xc3, 0xe8, 0xac, 0x06, 0xd7, 0x47, 0x0c, 0x13, 0x34, 0x50, 0x60, 0x14, 0x07,
	0x4d, 0x46, 0x81, 0x17, 0xfa, 0x51, 0xd3, 0xe2, 0xef, 0xe8, 0x22, 0x29, 0xb4, 0x91, 0x79, 0x7c,
	0x50, 0xf2, 0x00, 0xca, 0xd4, 0x1c, 0x64, 0xcd, 0x2c, 0xfe, 0x67, 0x66, 0x69, 0x97, 0x3e, 0xb7,
	0x3b, 0xce, 0xd5, 0x24, 0xf9, 0x07, 0x9e, 0xde, 0xa0, 0xb3, 0xae, 0x6f, 0x3c, 0x41, 0xfe, 0x1d,
	0x54, 0xc7, 0x27, 0xaa, 0xcb, 0xfa, 0xc7, 0x97, 0x2c, 0x2b, 0x9a, 0x77, 0x71, 0xcd, 0x4d, 0xef,
	0xad, 0x37, 0xfd, 0x8a, 0x9e, 0x3d, 0xfa, 0x99, 0x47, 0x44, 0x5e, 0x75, 0x45, 0xc6, 0xd7, 0x97,
	0xd4, 0x05, 0x87, 0x36, 0xc1, 0xa1, 0xef, 0x45, 0xd5, 0x91, 0xfe, 0xf0, 0x13, 0x5d, 0x71, 0x49,
	0xdb, 0x04, 0xd2, 0x84, 0x3b, 0xaa, 0xee, 0xde, 0xfc, 0xed, 0x5d, 0xca, 0xcd, 0xbe, 0xd8, 0xd0,
	0x44, 0xe6, 0x0f, 0x69, 0x7d, 0xa8, 0x66, 0x09, 0x9f, 0xb9, 0xad, 0x59, 0x2a, 0x33, 0x26, 0xd2,
	0x36, 0xde, 0x71, 0xb9, 0xf8, 0xd5, 0x23, 0xed, 0x2f, 0xa3, 0xcb, 0x35, 0xfd, 0xd4, 0xe8, 0xde,
	0x2e, 0x36, 0xa7, 0x76, 0xeb, 0xcd, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x4f, 0x20, 0x94,
	0x16, 0x03, 0x00, 0x00,
}
