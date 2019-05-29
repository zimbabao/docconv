// Code generated by protoc-gen-go. DO NOT EDIT.
// source: KNArchives_sos.proto

package KN

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	TSD "github.com/zimbabao/docconv/iWork/TSD"
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

type SpecSetTransitionAttributesArchive struct {
	TransitionAttributes *TransitionAttributesArchive `protobuf:"bytes,1,opt,name=transition_attributes,json=transitionAttributes" json:"transition_attributes,omitempty"`
	Unset                *bool                        `protobuf:"varint,2,req,name=unset" json:"unset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *SpecSetTransitionAttributesArchive) Reset()         { *m = SpecSetTransitionAttributesArchive{} }
func (m *SpecSetTransitionAttributesArchive) String() string { return proto.CompactTextString(m) }
func (*SpecSetTransitionAttributesArchive) ProtoMessage()    {}
func (*SpecSetTransitionAttributesArchive) Descriptor() ([]byte, []int) {
	return fileDescriptor_1933f80b989c46e2, []int{0}
}

func (m *SpecSetTransitionAttributesArchive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpecSetTransitionAttributesArchive.Unmarshal(m, b)
}
func (m *SpecSetTransitionAttributesArchive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpecSetTransitionAttributesArchive.Marshal(b, m, deterministic)
}
func (m *SpecSetTransitionAttributesArchive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpecSetTransitionAttributesArchive.Merge(m, src)
}
func (m *SpecSetTransitionAttributesArchive) XXX_Size() int {
	return xxx_messageInfo_SpecSetTransitionAttributesArchive.Size(m)
}
func (m *SpecSetTransitionAttributesArchive) XXX_DiscardUnknown() {
	xxx_messageInfo_SpecSetTransitionAttributesArchive.DiscardUnknown(m)
}

var xxx_messageInfo_SpecSetTransitionAttributesArchive proto.InternalMessageInfo

func (m *SpecSetTransitionAttributesArchive) GetTransitionAttributes() *TransitionAttributesArchive {
	if m != nil {
		return m.TransitionAttributes
	}
	return nil
}

func (m *SpecSetTransitionAttributesArchive) GetUnset() bool {
	if m != nil && m.Unset != nil {
		return *m.Unset
	}
	return false
}

type SpecTransitionAttributesArchive struct {
	SpecSetTransitionAttributes *SpecSetTransitionAttributesArchive `protobuf:"bytes,1,opt,name=spec_set_transition_attributes,json=specSetTransitionAttributes" json:"spec_set_transition_attributes,omitempty"`
	XXX_NoUnkeyedLiteral        struct{}                            `json:"-"`
	XXX_unrecognized            []byte                              `json:"-"`
	XXX_sizecache               int32                               `json:"-"`
}

func (m *SpecTransitionAttributesArchive) Reset()         { *m = SpecTransitionAttributesArchive{} }
func (m *SpecTransitionAttributesArchive) String() string { return proto.CompactTextString(m) }
func (*SpecTransitionAttributesArchive) ProtoMessage()    {}
func (*SpecTransitionAttributesArchive) Descriptor() ([]byte, []int) {
	return fileDescriptor_1933f80b989c46e2, []int{1}
}

func (m *SpecTransitionAttributesArchive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpecTransitionAttributesArchive.Unmarshal(m, b)
}
func (m *SpecTransitionAttributesArchive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpecTransitionAttributesArchive.Marshal(b, m, deterministic)
}
func (m *SpecTransitionAttributesArchive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpecTransitionAttributesArchive.Merge(m, src)
}
func (m *SpecTransitionAttributesArchive) XXX_Size() int {
	return xxx_messageInfo_SpecTransitionAttributesArchive.Size(m)
}
func (m *SpecTransitionAttributesArchive) XXX_DiscardUnknown() {
	xxx_messageInfo_SpecTransitionAttributesArchive.DiscardUnknown(m)
}

var xxx_messageInfo_SpecTransitionAttributesArchive proto.InternalMessageInfo

func (m *SpecTransitionAttributesArchive) GetSpecSetTransitionAttributes() *SpecSetTransitionAttributesArchive {
	if m != nil {
		return m.SpecSetTransitionAttributes
	}
	return nil
}

type SlideStylePropertyChangeSetArchive struct {
	Fill                     *TSD.SpecFillArchive             `protobuf:"bytes,1,opt,name=fill" json:"fill,omitempty"`
	FillUndefined            *bool                            `protobuf:"varint,2,opt,name=fill_undefined,json=fillUndefined" json:"fill_undefined,omitempty"`
	SlideTransition          *SpecTransitionAttributesArchive `protobuf:"bytes,3,opt,name=slide_transition,json=slideTransition" json:"slide_transition,omitempty"`
	SlideTransitionUndefined *bool                            `protobuf:"varint,4,opt,name=slide_transition_undefined,json=slideTransitionUndefined" json:"slide_transition_undefined,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                         `json:"-"`
	XXX_unrecognized         []byte                           `json:"-"`
	XXX_sizecache            int32                            `json:"-"`
}

func (m *SlideStylePropertyChangeSetArchive) Reset()         { *m = SlideStylePropertyChangeSetArchive{} }
func (m *SlideStylePropertyChangeSetArchive) String() string { return proto.CompactTextString(m) }
func (*SlideStylePropertyChangeSetArchive) ProtoMessage()    {}
func (*SlideStylePropertyChangeSetArchive) Descriptor() ([]byte, []int) {
	return fileDescriptor_1933f80b989c46e2, []int{2}
}

func (m *SlideStylePropertyChangeSetArchive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SlideStylePropertyChangeSetArchive.Unmarshal(m, b)
}
func (m *SlideStylePropertyChangeSetArchive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SlideStylePropertyChangeSetArchive.Marshal(b, m, deterministic)
}
func (m *SlideStylePropertyChangeSetArchive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SlideStylePropertyChangeSetArchive.Merge(m, src)
}
func (m *SlideStylePropertyChangeSetArchive) XXX_Size() int {
	return xxx_messageInfo_SlideStylePropertyChangeSetArchive.Size(m)
}
func (m *SlideStylePropertyChangeSetArchive) XXX_DiscardUnknown() {
	xxx_messageInfo_SlideStylePropertyChangeSetArchive.DiscardUnknown(m)
}

var xxx_messageInfo_SlideStylePropertyChangeSetArchive proto.InternalMessageInfo

func (m *SlideStylePropertyChangeSetArchive) GetFill() *TSD.SpecFillArchive {
	if m != nil {
		return m.Fill
	}
	return nil
}

func (m *SlideStylePropertyChangeSetArchive) GetFillUndefined() bool {
	if m != nil && m.FillUndefined != nil {
		return *m.FillUndefined
	}
	return false
}

func (m *SlideStylePropertyChangeSetArchive) GetSlideTransition() *SpecTransitionAttributesArchive {
	if m != nil {
		return m.SlideTransition
	}
	return nil
}

func (m *SlideStylePropertyChangeSetArchive) GetSlideTransitionUndefined() bool {
	if m != nil && m.SlideTransitionUndefined != nil {
		return *m.SlideTransitionUndefined
	}
	return false
}

func init() {
	proto.RegisterType((*SpecSetTransitionAttributesArchive)(nil), "KN.SpecSetTransitionAttributesArchive")
	proto.RegisterType((*SpecTransitionAttributesArchive)(nil), "KN.SpecTransitionAttributesArchive")
	proto.RegisterType((*SlideStylePropertyChangeSetArchive)(nil), "KN.SlideStylePropertyChangeSetArchive")
}

func init() { proto.RegisterFile("KNArchives_sos.proto", fileDescriptor_1933f80b989c46e2) }

var fileDescriptor_1933f80b989c46e2 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0x4f, 0x6b, 0xab, 0x40,
	0x14, 0xc5, 0xd1, 0x97, 0x07, 0x61, 0x1e, 0xef, 0xbd, 0x20, 0x06, 0x24, 0x85, 0x26, 0xd8, 0x36,
	0xb8, 0x52, 0xe8, 0xba, 0x9b, 0xb4, 0xa1, 0x1b, 0x41, 0x8a, 0x5a, 0x0a, 0xdd, 0x88, 0x7f, 0x6e,
	0x92, 0x21, 0x66, 0x46, 0x66, 0xae, 0x81, 0x74, 0xdb, 0x7d, 0xe9, 0x47, 0x2e, 0x9a, 0x98, 0x84,
	0xb4, 0xb5, 0x2b, 0xf1, 0x9e, 0x73, 0xe6, 0xfc, 0x0e, 0xd1, 0x5d, 0x6f, 0x22, 0xd2, 0x05, 0x5d,
	0x83, 0x8c, 0x24, 0x97, 0x76, 0x21, 0x38, 0x72, 0x4d, 0x75, 0xbd, 0x41, 0xef, 0xa0, 0x6c, 0xaf,
	0x83, 0x7e, 0x18, 0x4c, 0x3f, 0x9b, 0xcd, 0x77, 0x85, 0x98, 0x41, 0x01, 0x69, 0x00, 0x18, 0x8a,
	0x98, 0x49, 0x8a, 0x94, 0xb3, 0x09, 0xa2, 0xa0, 0x49, 0x89, 0x20, 0x77, 0x09, 0x2d, 0x24, 0x7d,
	0xdc, 0xcb, 0x51, 0xbc, 0xd7, 0x0d, 0x65, 0xa4, 0x58, 0x7f, 0xae, 0x87, 0xb6, 0xeb, 0xd9, 0x2d,
	0x79, 0x5f, 0xc7, 0x2f, 0x44, 0x4d, 0x27, 0xbf, 0x4b, 0x26, 0x01, 0x0d, 0x75, 0xa4, 0x5a, 0x5d,
	0x7f, 0xfb, 0x63, 0xbe, 0x29, 0x64, 0x58, 0x21, 0xb5, 0xf1, 0x2c, 0xc9, 0xb9, 0x2c, 0x20, 0x8d,
	0x24, 0x60, 0xd4, 0x06, 0x36, 0xae, 0xc0, 0x7e, 0xde, 0xe7, 0x9f, 0xc9, 0xef, 0x3d, 0xe6, 0xab,
	0x4a, 0xcc, 0x20, 0xa7, 0x19, 0x04, 0xb8, 0xc9, 0xe1, 0x41, 0xf0, 0x02, 0x04, 0x6e, 0xee, 0x16,
	0x31, 0x9b, 0x43, 0x00, 0xd8, 0x30, 0x59, 0xa4, 0x33, 0xa3, 0x79, 0xbe, 0x6b, 0xd6, 0xed, 0x30,
	0x98, 0xd6, 0xd5, 0xf7, 0x34, 0xcf, 0x9b, 0x9e, 0xda, 0xa1, 0x5d, 0x91, 0x7f, 0xd5, 0x37, 0x2a,
	0x59, 0x06, 0x33, 0xca, 0x20, 0x33, 0xd4, 0x91, 0x62, 0x75, 0xfd, 0xbf, 0xd5, 0xf5, 0xb1, 0x39,
	0x6a, 0x1e, 0xe9, 0xc9, 0xaa, 0xf6, 0x68, 0xa1, 0xf1, 0xab, 0x7e, 0xfc, 0xa2, 0x99, 0xd5, 0xb6,
	0xe9, 0x7f, 0x1d, 0x3e, 0x38, 0xb4, 0x1b, 0x32, 0x38, 0x7d, 0xef, 0x08, 0xa1, 0x53, 0x23, 0x18,
	0x27, 0xa1, 0x3d, 0xcd, 0xed, 0xf8, 0xf9, 0x72, 0x4e, 0x71, 0x51, 0x26, 0x76, 0xca, 0x57, 0xce,
	0x0b, 0x5d, 0x25, 0x71, 0x12, 0x73, 0x27, 0xe3, 0x69, 0xca, 0xd9, 0xda, 0xa1, 0x4f, 0x5c, 0x2c,
	0x1d, 0xd7, 0xfb, 0x08, 0x00, 0x00, 0xff, 0xff, 0x15, 0xa2, 0xc7, 0xba, 0x95, 0x02, 0x00, 0x00,
}