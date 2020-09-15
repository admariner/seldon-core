// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow_serving/util/class_registration_test.proto

package serving

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

type Config1 struct {
	StringField          string   `protobuf:"bytes,1,opt,name=string_field,json=stringField,proto3" json:"string_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config1) Reset()         { *m = Config1{} }
func (m *Config1) String() string { return proto.CompactTextString(m) }
func (*Config1) ProtoMessage()    {}
func (*Config1) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a1690d17e90f318, []int{0}
}

func (m *Config1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config1.Unmarshal(m, b)
}
func (m *Config1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config1.Marshal(b, m, deterministic)
}
func (m *Config1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config1.Merge(m, src)
}
func (m *Config1) XXX_Size() int {
	return xxx_messageInfo_Config1.Size(m)
}
func (m *Config1) XXX_DiscardUnknown() {
	xxx_messageInfo_Config1.DiscardUnknown(m)
}

var xxx_messageInfo_Config1 proto.InternalMessageInfo

func (m *Config1) GetStringField() string {
	if m != nil {
		return m.StringField
	}
	return ""
}

type Config2 struct {
	StringField          string   `protobuf:"bytes,1,opt,name=string_field,json=stringField,proto3" json:"string_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config2) Reset()         { *m = Config2{} }
func (m *Config2) String() string { return proto.CompactTextString(m) }
func (*Config2) ProtoMessage()    {}
func (*Config2) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a1690d17e90f318, []int{1}
}

func (m *Config2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config2.Unmarshal(m, b)
}
func (m *Config2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config2.Marshal(b, m, deterministic)
}
func (m *Config2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config2.Merge(m, src)
}
func (m *Config2) XXX_Size() int {
	return xxx_messageInfo_Config2.Size(m)
}
func (m *Config2) XXX_DiscardUnknown() {
	xxx_messageInfo_Config2.DiscardUnknown(m)
}

var xxx_messageInfo_Config2 proto.InternalMessageInfo

func (m *Config2) GetStringField() string {
	if m != nil {
		return m.StringField
	}
	return ""
}

type MessageWithAny struct {
	AnyField             *any.Any `protobuf:"bytes,1,opt,name=any_field,json=anyField,proto3" json:"any_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageWithAny) Reset()         { *m = MessageWithAny{} }
func (m *MessageWithAny) String() string { return proto.CompactTextString(m) }
func (*MessageWithAny) ProtoMessage()    {}
func (*MessageWithAny) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a1690d17e90f318, []int{2}
}

func (m *MessageWithAny) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageWithAny.Unmarshal(m, b)
}
func (m *MessageWithAny) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageWithAny.Marshal(b, m, deterministic)
}
func (m *MessageWithAny) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageWithAny.Merge(m, src)
}
func (m *MessageWithAny) XXX_Size() int {
	return xxx_messageInfo_MessageWithAny.Size(m)
}
func (m *MessageWithAny) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageWithAny.DiscardUnknown(m)
}

var xxx_messageInfo_MessageWithAny proto.InternalMessageInfo

func (m *MessageWithAny) GetAnyField() *any.Any {
	if m != nil {
		return m.AnyField
	}
	return nil
}

func init() {
	proto.RegisterType((*Config1)(nil), "tensorflow.serving.Config1")
	proto.RegisterType((*Config2)(nil), "tensorflow.serving.Config2")
	proto.RegisterType((*MessageWithAny)(nil), "tensorflow.serving.MessageWithAny")
}

func init() {
	proto.RegisterFile("tensorflow_serving/util/class_registration_test.proto", fileDescriptor_2a1690d17e90f318)
}

var fileDescriptor_2a1690d17e90f318 = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x8e, 0x31, 0x4b, 0x04, 0x31,
	0x10, 0x46, 0xb9, 0x46, 0xbd, 0x9c, 0x58, 0x04, 0x0b, 0xb5, 0x52, 0x2b, 0x0b, 0x49, 0xb8, 0x13,
	0x7f, 0xc0, 0xb1, 0x60, 0x67, 0xb3, 0x8d, 0x65, 0xc8, 0xea, 0x24, 0x0e, 0x84, 0x19, 0xc9, 0xcc,
	0x2a, 0xf9, 0xf7, 0xe2, 0xae, 0xb2, 0x96, 0xd7, 0x3e, 0xde, 0xfb, 0xf8, 0xcc, 0xa3, 0x02, 0x09,
	0xd7, 0x54, 0xf8, 0x2b, 0x08, 0xd4, 0x4f, 0xa4, 0xec, 0x47, 0xc5, 0xe2, 0x5f, 0x4b, 0x14, 0x09,
	0x15, 0x32, 0x8a, 0xd6, 0xa8, 0xc8, 0x14, 0x14, 0x44, 0xdd, 0x47, 0x65, 0x65, 0x6b, 0x97, 0xcc,
	0xfd, 0x66, 0x57, 0x97, 0x99, 0x39, 0x17, 0xf0, 0x93, 0x31, 0x8c, 0xc9, 0x47, 0x6a, 0xb3, 0x7e,
	0x7b, 0x6f, 0x8e, 0x3b, 0xa6, 0x84, 0x79, 0x6b, 0x6f, 0xcc, 0xa9, 0x68, 0x45, 0xca, 0x21, 0x21,
	0x94, 0xb7, 0x8b, 0xd5, 0xf5, 0xea, 0x6e, 0xdd, 0x6f, 0x66, 0xf6, 0xf4, 0x83, 0x16, 0x7b, 0x77,
	0x88, 0xdd, 0x99, 0xb3, 0x67, 0x10, 0x89, 0x19, 0x5e, 0x50, 0xdf, 0xf7, 0xd4, 0xec, 0xd6, 0xac,
	0x23, 0xb5, 0x7f, 0xc5, 0x66, 0x77, 0xee, 0xe6, 0x73, 0xee, 0xef, 0x9c, 0xdb, 0x53, 0xeb, 0x4f,
	0x22, 0xb5, 0x69, 0x64, 0x38, 0x9a, 0xf8, 0xc3, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x42, 0xfa,
	0xee, 0xbf, 0x0f, 0x01, 0x00, 0x00,
}