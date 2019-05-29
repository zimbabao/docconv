// Code generated by protoc-gen-go. DO NOT EDIT.
// source: KNCommandArchives_sos.proto

package KN

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	TSA "github.com/zimbabao/docconv/iWork/TSA"
	TSK "github.com/zimbabao/docconv/iWork/TSK"
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

type InducedVerifyDocumentWithServerCommandArchive struct {
	Super                    *TSK.CommandArchive `protobuf:"bytes,1,req,name=super" json:"super,omitempty"`
	SlideNodeIdList          []*TSP.UUID         `protobuf:"bytes,2,rep,name=slide_node_id_list,json=slideNodeIdList" json:"slide_node_id_list,omitempty"`
	SlideNodeIdListUndefined *bool               `protobuf:"varint,3,opt,name=slide_node_id_list_undefined,json=slideNodeIdListUndefined" json:"slide_node_id_list_undefined,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}            `json:"-"`
	XXX_unrecognized         []byte              `json:"-"`
	XXX_sizecache            int32               `json:"-"`
}

func (m *InducedVerifyDocumentWithServerCommandArchive) Reset() {
	*m = InducedVerifyDocumentWithServerCommandArchive{}
}
func (m *InducedVerifyDocumentWithServerCommandArchive) String() string {
	return proto.CompactTextString(m)
}
func (*InducedVerifyDocumentWithServerCommandArchive) ProtoMessage() {}
func (*InducedVerifyDocumentWithServerCommandArchive) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4661028b68bb679, []int{0}
}

func (m *InducedVerifyDocumentWithServerCommandArchive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InducedVerifyDocumentWithServerCommandArchive.Unmarshal(m, b)
}
func (m *InducedVerifyDocumentWithServerCommandArchive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InducedVerifyDocumentWithServerCommandArchive.Marshal(b, m, deterministic)
}
func (m *InducedVerifyDocumentWithServerCommandArchive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InducedVerifyDocumentWithServerCommandArchive.Merge(m, src)
}
func (m *InducedVerifyDocumentWithServerCommandArchive) XXX_Size() int {
	return xxx_messageInfo_InducedVerifyDocumentWithServerCommandArchive.Size(m)
}
func (m *InducedVerifyDocumentWithServerCommandArchive) XXX_DiscardUnknown() {
	xxx_messageInfo_InducedVerifyDocumentWithServerCommandArchive.DiscardUnknown(m)
}

var xxx_messageInfo_InducedVerifyDocumentWithServerCommandArchive proto.InternalMessageInfo

func (m *InducedVerifyDocumentWithServerCommandArchive) GetSuper() *TSK.CommandArchive {
	if m != nil {
		return m.Super
	}
	return nil
}

func (m *InducedVerifyDocumentWithServerCommandArchive) GetSlideNodeIdList() []*TSP.UUID {
	if m != nil {
		return m.SlideNodeIdList
	}
	return nil
}

func (m *InducedVerifyDocumentWithServerCommandArchive) GetSlideNodeIdListUndefined() bool {
	if m != nil && m.SlideNodeIdListUndefined != nil {
		return *m.SlideNodeIdListUndefined
	}
	return false
}

type InducedVerifyDrawableZOrdersWithServerCommandArchive struct {
	Super                *TSA.InducedVerifyDrawableZOrdersWithServerCommandArchive `protobuf:"bytes,1,req,name=super" json:"super,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                  `json:"-"`
	XXX_unrecognized     []byte                                                    `json:"-"`
	XXX_sizecache        int32                                                     `json:"-"`
}

func (m *InducedVerifyDrawableZOrdersWithServerCommandArchive) Reset() {
	*m = InducedVerifyDrawableZOrdersWithServerCommandArchive{}
}
func (m *InducedVerifyDrawableZOrdersWithServerCommandArchive) String() string {
	return proto.CompactTextString(m)
}
func (*InducedVerifyDrawableZOrdersWithServerCommandArchive) ProtoMessage() {}
func (*InducedVerifyDrawableZOrdersWithServerCommandArchive) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4661028b68bb679, []int{1}
}

func (m *InducedVerifyDrawableZOrdersWithServerCommandArchive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InducedVerifyDrawableZOrdersWithServerCommandArchive.Unmarshal(m, b)
}
func (m *InducedVerifyDrawableZOrdersWithServerCommandArchive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InducedVerifyDrawableZOrdersWithServerCommandArchive.Marshal(b, m, deterministic)
}
func (m *InducedVerifyDrawableZOrdersWithServerCommandArchive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InducedVerifyDrawableZOrdersWithServerCommandArchive.Merge(m, src)
}
func (m *InducedVerifyDrawableZOrdersWithServerCommandArchive) XXX_Size() int {
	return xxx_messageInfo_InducedVerifyDrawableZOrdersWithServerCommandArchive.Size(m)
}
func (m *InducedVerifyDrawableZOrdersWithServerCommandArchive) XXX_DiscardUnknown() {
	xxx_messageInfo_InducedVerifyDrawableZOrdersWithServerCommandArchive.DiscardUnknown(m)
}

var xxx_messageInfo_InducedVerifyDrawableZOrdersWithServerCommandArchive proto.InternalMessageInfo

func (m *InducedVerifyDrawableZOrdersWithServerCommandArchive) GetSuper() *TSA.InducedVerifyDrawableZOrdersWithServerCommandArchive {
	if m != nil {
		return m.Super
	}
	return nil
}

func init() {
	proto.RegisterType((*InducedVerifyDocumentWithServerCommandArchive)(nil), "KN.InducedVerifyDocumentWithServerCommandArchive")
	proto.RegisterType((*InducedVerifyDrawableZOrdersWithServerCommandArchive)(nil), "KN.InducedVerifyDrawableZOrdersWithServerCommandArchive")
}

func init() { proto.RegisterFile("KNCommandArchives_sos.proto", fileDescriptor_b4661028b68bb679) }

var fileDescriptor_b4661028b68bb679 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0x8b, 0xa0, 0xe9, 0x41, 0x8c, 0x97, 0x50, 0x7b, 0x28, 0x45, 0xa4, 0x1e, 0xdc,
	0x40, 0x11, 0xc1, 0x8b, 0x50, 0xed, 0xa5, 0x44, 0xd3, 0xd2, 0xa4, 0x16, 0x7a, 0x09, 0x9b, 0x9d,
	0x69, 0xbb, 0xd8, 0xec, 0x94, 0xdd, 0x4d, 0x45, 0x5f, 0xc0, 0xf7, 0xf3, 0x89, 0x04, 0x35, 0x48,
	0x94, 0x5e, 0xbc, 0xfe, 0xf3, 0xfd, 0x33, 0x7c, 0xe3, 0x9e, 0x84, 0xd1, 0x1d, 0xe5, 0x39, 0x57,
	0xd0, 0xd7, 0x62, 0x25, 0xb7, 0x68, 0x52, 0x43, 0x86, 0x6d, 0x34, 0x59, 0xf2, 0x6a, 0x61, 0xd4,
	0x6c, 0x25, 0x71, 0x7f, 0x27, 0xd1, 0x3c, 0x4a, 0xe2, 0xb0, 0x8c, 0x7f, 0xa2, 0xf1, 0x03, 0x1a,
	0xc3, 0x97, 0x65, 0xd4, 0x79, 0x77, 0xdc, 0x8b, 0xa1, 0x82, 0x42, 0x20, 0x3c, 0xa2, 0x96, 0x8b,
	0x97, 0x01, 0x89, 0x22, 0x47, 0x65, 0x67, 0xd2, 0xae, 0x62, 0xd4, 0x5b, 0xd4, 0xd5, 0x13, 0xde,
	0xb9, 0xbb, 0x67, 0x8a, 0x0d, 0x6a, 0xdf, 0x69, 0xd7, 0xba, 0x8d, 0xde, 0x31, 0x4b, 0xe2, 0x90,
	0x55, 0x99, 0xc9, 0x17, 0xe1, 0x5d, 0xb9, 0x9e, 0x59, 0x4b, 0xc0, 0x54, 0x11, 0x60, 0x2a, 0x21,
	0x5d, 0x4b, 0x63, 0xfd, 0x5a, 0xbb, 0xde, 0x6d, 0xf4, 0x0e, 0x58, 0x12, 0x8f, 0xd9, 0x74, 0x3a,
	0x1c, 0x4c, 0x0e, 0x3f, 0xa1, 0x88, 0x00, 0x87, 0x70, 0x2f, 0x8d, 0xf5, 0x6e, 0xdc, 0xd6, 0xdf,
	0x5e, 0x5a, 0x28, 0xc0, 0x85, 0x54, 0x08, 0x7e, 0xbd, 0xed, 0x74, 0xf7, 0x27, 0xfe, 0xaf, 0xda,
	0xb4, 0x9c, 0x77, 0xde, 0x1c, 0xf7, 0xb2, 0x2a, 0xa5, 0xf9, 0x33, 0xcf, 0xd6, 0x38, 0x1f, 0x69,
	0x40, 0x6d, 0x76, 0xba, 0x8d, 0xaa, 0x6e, 0xd7, 0x2c, 0x89, 0xfb, 0xec, 0x3f, 0x9b, 0xbe, 0x3f,
	0x70, 0x7b, 0x36, 0x3f, 0x5d, 0x4a, 0xbb, 0x2a, 0x32, 0x26, 0x28, 0x0f, 0x5e, 0x65, 0x9e, 0xf1,
	0x8c, 0x53, 0x00, 0x24, 0x04, 0xa9, 0x6d, 0x20, 0x67, 0xa4, 0x9f, 0x82, 0x30, 0xfa, 0x08, 0x00,
	0x00, 0xff, 0xff, 0xbd, 0x17, 0xdd, 0xb1, 0xec, 0x01, 0x00, 0x00,
}
