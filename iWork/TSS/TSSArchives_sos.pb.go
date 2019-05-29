// Code generated by protoc-gen-go. DO NOT EDIT.
// source: TSSArchives_sos.proto

package TSS

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	TSP "github.com/zimbabao/docconv/iWork/TSP"
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

type SpecSetBoolArchive struct {
	Value                *bool    `protobuf:"varint,1,req,name=value" json:"value,omitempty"`
	Unset                *bool    `protobuf:"varint,2,req,name=unset" json:"unset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpecSetBoolArchive) Reset()         { *m = SpecSetBoolArchive{} }
func (m *SpecSetBoolArchive) String() string { return proto.CompactTextString(m) }
func (*SpecSetBoolArchive) ProtoMessage()    {}
func (*SpecSetBoolArchive) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7305b92c4b870a3, []int{0}
}

func (m *SpecSetBoolArchive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpecSetBoolArchive.Unmarshal(m, b)
}
func (m *SpecSetBoolArchive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpecSetBoolArchive.Marshal(b, m, deterministic)
}
func (m *SpecSetBoolArchive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpecSetBoolArchive.Merge(m, src)
}
func (m *SpecSetBoolArchive) XXX_Size() int {
	return xxx_messageInfo_SpecSetBoolArchive.Size(m)
}
func (m *SpecSetBoolArchive) XXX_DiscardUnknown() {
	xxx_messageInfo_SpecSetBoolArchive.DiscardUnknown(m)
}

var xxx_messageInfo_SpecSetBoolArchive proto.InternalMessageInfo

func (m *SpecSetBoolArchive) GetValue() bool {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return false
}

func (m *SpecSetBoolArchive) GetUnset() bool {
	if m != nil && m.Unset != nil {
		return *m.Unset
	}
	return false
}

type SpecSetColorArchive struct {
	Color                *TSP.Color `protobuf:"bytes,1,opt,name=color" json:"color,omitempty"`
	Unset                *bool      `protobuf:"varint,2,req,name=unset" json:"unset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SpecSetColorArchive) Reset()         { *m = SpecSetColorArchive{} }
func (m *SpecSetColorArchive) String() string { return proto.CompactTextString(m) }
func (*SpecSetColorArchive) ProtoMessage()    {}
func (*SpecSetColorArchive) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7305b92c4b870a3, []int{1}
}

func (m *SpecSetColorArchive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpecSetColorArchive.Unmarshal(m, b)
}
func (m *SpecSetColorArchive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpecSetColorArchive.Marshal(b, m, deterministic)
}
func (m *SpecSetColorArchive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpecSetColorArchive.Merge(m, src)
}
func (m *SpecSetColorArchive) XXX_Size() int {
	return xxx_messageInfo_SpecSetColorArchive.Size(m)
}
func (m *SpecSetColorArchive) XXX_DiscardUnknown() {
	xxx_messageInfo_SpecSetColorArchive.DiscardUnknown(m)
}

var xxx_messageInfo_SpecSetColorArchive proto.InternalMessageInfo

func (m *SpecSetColorArchive) GetColor() *TSP.Color {
	if m != nil {
		return m.Color
	}
	return nil
}

func (m *SpecSetColorArchive) GetUnset() bool {
	if m != nil && m.Unset != nil {
		return *m.Unset
	}
	return false
}

type SpecSetDoubleArchive struct {
	Value                *float64 `protobuf:"fixed64,1,req,name=value" json:"value,omitempty"`
	Unset                *bool    `protobuf:"varint,2,req,name=unset" json:"unset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpecSetDoubleArchive) Reset()         { *m = SpecSetDoubleArchive{} }
func (m *SpecSetDoubleArchive) String() string { return proto.CompactTextString(m) }
func (*SpecSetDoubleArchive) ProtoMessage()    {}
func (*SpecSetDoubleArchive) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7305b92c4b870a3, []int{2}
}

func (m *SpecSetDoubleArchive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpecSetDoubleArchive.Unmarshal(m, b)
}
func (m *SpecSetDoubleArchive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpecSetDoubleArchive.Marshal(b, m, deterministic)
}
func (m *SpecSetDoubleArchive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpecSetDoubleArchive.Merge(m, src)
}
func (m *SpecSetDoubleArchive) XXX_Size() int {
	return xxx_messageInfo_SpecSetDoubleArchive.Size(m)
}
func (m *SpecSetDoubleArchive) XXX_DiscardUnknown() {
	xxx_messageInfo_SpecSetDoubleArchive.DiscardUnknown(m)
}

var xxx_messageInfo_SpecSetDoubleArchive proto.InternalMessageInfo

func (m *SpecSetDoubleArchive) GetValue() float64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *SpecSetDoubleArchive) GetUnset() bool {
	if m != nil && m.Unset != nil {
		return *m.Unset
	}
	return false
}

type SpecSetIntegerArchive struct {
	Value                *int32   `protobuf:"varint,1,req,name=value" json:"value,omitempty"`
	Unset                *bool    `protobuf:"varint,2,req,name=unset" json:"unset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpecSetIntegerArchive) Reset()         { *m = SpecSetIntegerArchive{} }
func (m *SpecSetIntegerArchive) String() string { return proto.CompactTextString(m) }
func (*SpecSetIntegerArchive) ProtoMessage()    {}
func (*SpecSetIntegerArchive) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7305b92c4b870a3, []int{3}
}

func (m *SpecSetIntegerArchive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpecSetIntegerArchive.Unmarshal(m, b)
}
func (m *SpecSetIntegerArchive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpecSetIntegerArchive.Marshal(b, m, deterministic)
}
func (m *SpecSetIntegerArchive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpecSetIntegerArchive.Merge(m, src)
}
func (m *SpecSetIntegerArchive) XXX_Size() int {
	return xxx_messageInfo_SpecSetIntegerArchive.Size(m)
}
func (m *SpecSetIntegerArchive) XXX_DiscardUnknown() {
	xxx_messageInfo_SpecSetIntegerArchive.DiscardUnknown(m)
}

var xxx_messageInfo_SpecSetIntegerArchive proto.InternalMessageInfo

func (m *SpecSetIntegerArchive) GetValue() int32 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

func (m *SpecSetIntegerArchive) GetUnset() bool {
	if m != nil && m.Unset != nil {
		return *m.Unset
	}
	return false
}

type SpecSetStringArchive struct {
	Value                *string  `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	Unset                *bool    `protobuf:"varint,2,req,name=unset" json:"unset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpecSetStringArchive) Reset()         { *m = SpecSetStringArchive{} }
func (m *SpecSetStringArchive) String() string { return proto.CompactTextString(m) }
func (*SpecSetStringArchive) ProtoMessage()    {}
func (*SpecSetStringArchive) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7305b92c4b870a3, []int{4}
}

func (m *SpecSetStringArchive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpecSetStringArchive.Unmarshal(m, b)
}
func (m *SpecSetStringArchive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpecSetStringArchive.Marshal(b, m, deterministic)
}
func (m *SpecSetStringArchive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpecSetStringArchive.Merge(m, src)
}
func (m *SpecSetStringArchive) XXX_Size() int {
	return xxx_messageInfo_SpecSetStringArchive.Size(m)
}
func (m *SpecSetStringArchive) XXX_DiscardUnknown() {
	xxx_messageInfo_SpecSetStringArchive.DiscardUnknown(m)
}

var xxx_messageInfo_SpecSetStringArchive proto.InternalMessageInfo

func (m *SpecSetStringArchive) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func (m *SpecSetStringArchive) GetUnset() bool {
	if m != nil && m.Unset != nil {
		return *m.Unset
	}
	return false
}

type SpecBoolArchive struct {
	SpecSetBool          *SpecSetBoolArchive `protobuf:"bytes,1,opt,name=spec_set_bool,json=specSetBool" json:"spec_set_bool,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SpecBoolArchive) Reset()         { *m = SpecBoolArchive{} }
func (m *SpecBoolArchive) String() string { return proto.CompactTextString(m) }
func (*SpecBoolArchive) ProtoMessage()    {}
func (*SpecBoolArchive) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7305b92c4b870a3, []int{5}
}

func (m *SpecBoolArchive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpecBoolArchive.Unmarshal(m, b)
}
func (m *SpecBoolArchive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpecBoolArchive.Marshal(b, m, deterministic)
}
func (m *SpecBoolArchive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpecBoolArchive.Merge(m, src)
}
func (m *SpecBoolArchive) XXX_Size() int {
	return xxx_messageInfo_SpecBoolArchive.Size(m)
}
func (m *SpecBoolArchive) XXX_DiscardUnknown() {
	xxx_messageInfo_SpecBoolArchive.DiscardUnknown(m)
}

var xxx_messageInfo_SpecBoolArchive proto.InternalMessageInfo

func (m *SpecBoolArchive) GetSpecSetBool() *SpecSetBoolArchive {
	if m != nil {
		return m.SpecSetBool
	}
	return nil
}

type SpecColorArchive struct {
	SpecSetColor         *SpecSetColorArchive `protobuf:"bytes,1,opt,name=spec_set_color,json=specSetColor" json:"spec_set_color,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SpecColorArchive) Reset()         { *m = SpecColorArchive{} }
func (m *SpecColorArchive) String() string { return proto.CompactTextString(m) }
func (*SpecColorArchive) ProtoMessage()    {}
func (*SpecColorArchive) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7305b92c4b870a3, []int{6}
}

func (m *SpecColorArchive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpecColorArchive.Unmarshal(m, b)
}
func (m *SpecColorArchive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpecColorArchive.Marshal(b, m, deterministic)
}
func (m *SpecColorArchive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpecColorArchive.Merge(m, src)
}
func (m *SpecColorArchive) XXX_Size() int {
	return xxx_messageInfo_SpecColorArchive.Size(m)
}
func (m *SpecColorArchive) XXX_DiscardUnknown() {
	xxx_messageInfo_SpecColorArchive.DiscardUnknown(m)
}

var xxx_messageInfo_SpecColorArchive proto.InternalMessageInfo

func (m *SpecColorArchive) GetSpecSetColor() *SpecSetColorArchive {
	if m != nil {
		return m.SpecSetColor
	}
	return nil
}

type SpecDoubleArchive struct {
	SpecSetDouble        *SpecSetDoubleArchive `protobuf:"bytes,1,opt,name=spec_set_double,json=specSetDouble" json:"spec_set_double,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *SpecDoubleArchive) Reset()         { *m = SpecDoubleArchive{} }
func (m *SpecDoubleArchive) String() string { return proto.CompactTextString(m) }
func (*SpecDoubleArchive) ProtoMessage()    {}
func (*SpecDoubleArchive) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7305b92c4b870a3, []int{7}
}

func (m *SpecDoubleArchive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpecDoubleArchive.Unmarshal(m, b)
}
func (m *SpecDoubleArchive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpecDoubleArchive.Marshal(b, m, deterministic)
}
func (m *SpecDoubleArchive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpecDoubleArchive.Merge(m, src)
}
func (m *SpecDoubleArchive) XXX_Size() int {
	return xxx_messageInfo_SpecDoubleArchive.Size(m)
}
func (m *SpecDoubleArchive) XXX_DiscardUnknown() {
	xxx_messageInfo_SpecDoubleArchive.DiscardUnknown(m)
}

var xxx_messageInfo_SpecDoubleArchive proto.InternalMessageInfo

func (m *SpecDoubleArchive) GetSpecSetDouble() *SpecSetDoubleArchive {
	if m != nil {
		return m.SpecSetDouble
	}
	return nil
}

type SpecIntegerArchive struct {
	SpecSetInteger       *SpecSetIntegerArchive `protobuf:"bytes,1,opt,name=spec_set_integer,json=specSetInteger" json:"spec_set_integer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SpecIntegerArchive) Reset()         { *m = SpecIntegerArchive{} }
func (m *SpecIntegerArchive) String() string { return proto.CompactTextString(m) }
func (*SpecIntegerArchive) ProtoMessage()    {}
func (*SpecIntegerArchive) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7305b92c4b870a3, []int{8}
}

func (m *SpecIntegerArchive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpecIntegerArchive.Unmarshal(m, b)
}
func (m *SpecIntegerArchive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpecIntegerArchive.Marshal(b, m, deterministic)
}
func (m *SpecIntegerArchive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpecIntegerArchive.Merge(m, src)
}
func (m *SpecIntegerArchive) XXX_Size() int {
	return xxx_messageInfo_SpecIntegerArchive.Size(m)
}
func (m *SpecIntegerArchive) XXX_DiscardUnknown() {
	xxx_messageInfo_SpecIntegerArchive.DiscardUnknown(m)
}

var xxx_messageInfo_SpecIntegerArchive proto.InternalMessageInfo

func (m *SpecIntegerArchive) GetSpecSetInteger() *SpecSetIntegerArchive {
	if m != nil {
		return m.SpecSetInteger
	}
	return nil
}

type SpecStringArchive struct {
	SpecSetString        *SpecSetStringArchive `protobuf:"bytes,1,opt,name=spec_set_string,json=specSetString" json:"spec_set_string,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *SpecStringArchive) Reset()         { *m = SpecStringArchive{} }
func (m *SpecStringArchive) String() string { return proto.CompactTextString(m) }
func (*SpecStringArchive) ProtoMessage()    {}
func (*SpecStringArchive) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7305b92c4b870a3, []int{9}
}

func (m *SpecStringArchive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpecStringArchive.Unmarshal(m, b)
}
func (m *SpecStringArchive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpecStringArchive.Marshal(b, m, deterministic)
}
func (m *SpecStringArchive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpecStringArchive.Merge(m, src)
}
func (m *SpecStringArchive) XXX_Size() int {
	return xxx_messageInfo_SpecStringArchive.Size(m)
}
func (m *SpecStringArchive) XXX_DiscardUnknown() {
	xxx_messageInfo_SpecStringArchive.DiscardUnknown(m)
}

var xxx_messageInfo_SpecStringArchive proto.InternalMessageInfo

func (m *SpecStringArchive) GetSpecSetString() *SpecSetStringArchive {
	if m != nil {
		return m.SpecSetString
	}
	return nil
}

func init() {
	proto.RegisterType((*SpecSetBoolArchive)(nil), "TSS.SpecSetBoolArchive")
	proto.RegisterType((*SpecSetColorArchive)(nil), "TSS.SpecSetColorArchive")
	proto.RegisterType((*SpecSetDoubleArchive)(nil), "TSS.SpecSetDoubleArchive")
	proto.RegisterType((*SpecSetIntegerArchive)(nil), "TSS.SpecSetIntegerArchive")
	proto.RegisterType((*SpecSetStringArchive)(nil), "TSS.SpecSetStringArchive")
	proto.RegisterType((*SpecBoolArchive)(nil), "TSS.SpecBoolArchive")
	proto.RegisterType((*SpecColorArchive)(nil), "TSS.SpecColorArchive")
	proto.RegisterType((*SpecDoubleArchive)(nil), "TSS.SpecDoubleArchive")
	proto.RegisterType((*SpecIntegerArchive)(nil), "TSS.SpecIntegerArchive")
	proto.RegisterType((*SpecStringArchive)(nil), "TSS.SpecStringArchive")
}

func init() { proto.RegisterFile("TSSArchives_sos.proto", fileDescriptor_e7305b92c4b870a3) }

var fileDescriptor_e7305b92c4b870a3 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcd, 0x6b, 0xea, 0x40,
	0x14, 0xc5, 0x89, 0x0f, 0xe1, 0xbd, 0xeb, 0xf3, 0xf3, 0x29, 0x2f, 0x75, 0x25, 0x81, 0x52, 0x57,
	0x09, 0x74, 0x5b, 0x28, 0xf5, 0x63, 0xd3, 0x85, 0x45, 0x32, 0xa1, 0x05, 0x37, 0x92, 0x8c, 0x43,
	0x0c, 0x8d, 0xb9, 0x92, 0x99, 0xb8, 0xe8, 0x5f, 0x5f, 0x92, 0x8c, 0x71, 0xa6, 0x4d, 0x5d, 0xce,
	0xb9, 0xf7, 0xfc, 0x38, 0x9c, 0x3b, 0x30, 0xf2, 0x08, 0x99, 0xa5, 0x74, 0x1f, 0x9d, 0x18, 0xdf,
	0x72, 0xe4, 0xf6, 0x31, 0x45, 0x81, 0x83, 0x5f, 0x1e, 0x21, 0xe3, 0xbe, 0x47, 0xd6, 0x2b, 0xc6,
	0xb9, 0x1f, 0x32, 0xa9, 0x5b, 0x4f, 0x30, 0x20, 0x47, 0x46, 0x09, 0x13, 0x73, 0xc4, 0x58, 0x1a,
	0x07, 0x43, 0x68, 0x9e, 0xfc, 0x38, 0x63, 0xa6, 0x31, 0x69, 0x4c, 0x7f, 0xbb, 0xe5, 0x23, 0x57,
	0xb3, 0x84, 0x33, 0x61, 0x36, 0x4a, 0xb5, 0x78, 0x58, 0x2b, 0xf8, 0x27, 0x09, 0x0b, 0x8c, 0x31,
	0x3d, 0x23, 0x26, 0xd0, 0xa4, 0xf9, 0xdb, 0x34, 0x26, 0xc6, 0xb4, 0x75, 0x0f, 0xb6, 0x47, 0xd6,
	0x76, 0xb1, 0xe1, 0x96, 0x83, 0x1f, 0x70, 0x73, 0x18, 0x4a, 0xdc, 0x12, 0xb3, 0x20, 0x66, 0xb5,
	0x91, 0x8c, 0xeb, 0x91, 0x16, 0x30, 0x92, 0x8c, 0xe7, 0x44, 0xb0, 0x90, 0xa5, 0xb5, 0x90, 0xe6,
	0x75, 0xc8, 0x25, 0x08, 0x11, 0x69, 0x94, 0x84, 0x35, 0x0c, 0x63, 0xfa, 0xe7, 0x3a, 0xe3, 0x05,
	0xba, 0x39, 0x43, 0xad, 0xf6, 0x01, 0xda, 0xfc, 0xc8, 0xe8, 0x96, 0x33, 0xb1, 0x0d, 0x10, 0x63,
	0xd9, 0xcf, 0x7f, 0xdb, 0x23, 0xc4, 0xfe, 0x7e, 0x0a, 0xb7, 0xc5, 0x2f, 0x9a, 0xe5, 0x42, 0x2f,
	0x5f, 0xd1, 0x8a, 0x7e, 0x84, 0x4e, 0x05, 0x54, 0x1b, 0x37, 0x55, 0xa2, 0xea, 0x70, 0xff, 0x72,
	0x45, 0xb4, 0x5e, 0xa1, 0x9f, 0x2f, 0xe9, 0x6d, 0xcf, 0xa0, 0x5b, 0x41, 0x77, 0xc5, 0x44, 0x52,
	0x6f, 0x54, 0xaa, 0xe6, 0x71, 0xdb, 0x5c, 0x55, 0xad, 0x4d, 0xf9, 0xb3, 0xbe, 0x5c, 0x60, 0x09,
	0xbd, 0x0a, 0x1c, 0x95, 0x23, 0x49, 0x1e, 0xab, 0x64, 0xdd, 0xe5, 0x76, 0xb8, 0x26, 0x9f, 0x33,
	0xeb, 0x87, 0x51, 0x33, 0xf3, 0x62, 0x52, 0x97, 0x59, 0xf3, 0x54, 0x99, 0x4b, 0x75, 0x7e, 0xb7,
	0xb9, 0x0d, 0x23, 0xb1, 0xcf, 0x02, 0x9b, 0xe2, 0xc1, 0xf9, 0x88, 0x0e, 0x81, 0x1f, 0xf8, 0xe8,
	0xec, 0x90, 0x52, 0x4c, 0x4e, 0x4e, 0xf4, 0x86, 0xe9, 0xbb, 0xe3, 0x11, 0xf2, 0x19, 0x00, 0x00,
	0xff, 0xff, 0x54, 0x16, 0x11, 0x9d, 0x66, 0x03, 0x00, 0x00,
}
