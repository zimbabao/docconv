// Code generated by protoc-gen-go. DO NOT EDIT.
// source: TSPArchiveMessages.proto

package TSP

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

type FieldInfo_Type int32

const (
	FieldInfo_Value           FieldInfo_Type = 0
	FieldInfo_ObjectReference FieldInfo_Type = 1
	FieldInfo_DataReference   FieldInfo_Type = 2
	FieldInfo_Message         FieldInfo_Type = 3
)

var FieldInfo_Type_name = map[int32]string{
	0: "Value",
	1: "ObjectReference",
	2: "DataReference",
	3: "Message",
}

var FieldInfo_Type_value = map[string]int32{
	"Value":           0,
	"ObjectReference": 1,
	"DataReference":   2,
	"Message":         3,
}

func (x FieldInfo_Type) Enum() *FieldInfo_Type {
	p := new(FieldInfo_Type)
	*p = x
	return p
}

func (x FieldInfo_Type) String() string {
	return proto.EnumName(FieldInfo_Type_name, int32(x))
}

func (x *FieldInfo_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FieldInfo_Type_value, data, "FieldInfo_Type")
	if err != nil {
		return err
	}
	*x = FieldInfo_Type(value)
	return nil
}

func (FieldInfo_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_47802bcfa9d8cfa9, []int{2, 0}
}

type FieldInfo_Rule int32

const (
	FieldInfo_IgnoreAndDrop     FieldInfo_Rule = 0
	FieldInfo_IgnoreAndPreserve FieldInfo_Rule = 1
	FieldInfo_MustUnderstand    FieldInfo_Rule = 2
	FieldInfo_NotSupported      FieldInfo_Rule = -1
)

var FieldInfo_Rule_name = map[int32]string{
	0:  "IgnoreAndDrop",
	1:  "IgnoreAndPreserve",
	2:  "MustUnderstand",
	-1: "NotSupported",
}

var FieldInfo_Rule_value = map[string]int32{
	"IgnoreAndDrop":     0,
	"IgnoreAndPreserve": 1,
	"MustUnderstand":    2,
	"NotSupported":      -1,
}

func (x FieldInfo_Rule) Enum() *FieldInfo_Rule {
	p := new(FieldInfo_Rule)
	*p = x
	return p
}

func (x FieldInfo_Rule) String() string {
	return proto.EnumName(FieldInfo_Rule_name, int32(x))
}

func (x *FieldInfo_Rule) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FieldInfo_Rule_value, data, "FieldInfo_Rule")
	if err != nil {
		return err
	}
	*x = FieldInfo_Rule(value)
	return nil
}

func (FieldInfo_Rule) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_47802bcfa9d8cfa9, []int{2, 1}
}

type ArchiveInfo struct {
	Identifier           *uint64        `protobuf:"varint,1,opt,name=identifier" json:"identifier,omitempty"`
	MessageInfos         []*MessageInfo `protobuf:"bytes,2,rep,name=message_infos,json=messageInfos" json:"message_infos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ArchiveInfo) Reset()         { *m = ArchiveInfo{} }
func (m *ArchiveInfo) String() string { return proto.CompactTextString(m) }
func (*ArchiveInfo) ProtoMessage()    {}
func (*ArchiveInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_47802bcfa9d8cfa9, []int{0}
}

func (m *ArchiveInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArchiveInfo.Unmarshal(m, b)
}
func (m *ArchiveInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArchiveInfo.Marshal(b, m, deterministic)
}
func (m *ArchiveInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArchiveInfo.Merge(m, src)
}
func (m *ArchiveInfo) XXX_Size() int {
	return xxx_messageInfo_ArchiveInfo.Size(m)
}
func (m *ArchiveInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ArchiveInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ArchiveInfo proto.InternalMessageInfo

func (m *ArchiveInfo) GetIdentifier() uint64 {
	if m != nil && m.Identifier != nil {
		return *m.Identifier
	}
	return 0
}

func (m *ArchiveInfo) GetMessageInfos() []*MessageInfo {
	if m != nil {
		return m.MessageInfos
	}
	return nil
}

type MessageInfo struct {
	Type                 *uint32      `protobuf:"varint,1,req,name=type" json:"type,omitempty"`
	Version              []uint32     `protobuf:"varint,2,rep,packed,name=version" json:"version,omitempty"`
	Length               *uint32      `protobuf:"varint,3,req,name=length" json:"length,omitempty"`
	FieldInfos           []*FieldInfo `protobuf:"bytes,4,rep,name=field_infos,json=fieldInfos" json:"field_infos,omitempty"`
	ObjectReferences     []uint64     `protobuf:"varint,5,rep,packed,name=object_references,json=objectReferences" json:"object_references,omitempty"`
	DataReferences       []uint64     `protobuf:"varint,6,rep,packed,name=data_references,json=dataReferences" json:"data_references,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *MessageInfo) Reset()         { *m = MessageInfo{} }
func (m *MessageInfo) String() string { return proto.CompactTextString(m) }
func (*MessageInfo) ProtoMessage()    {}
func (*MessageInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_47802bcfa9d8cfa9, []int{1}
}

func (m *MessageInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageInfo.Unmarshal(m, b)
}
func (m *MessageInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageInfo.Marshal(b, m, deterministic)
}
func (m *MessageInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageInfo.Merge(m, src)
}
func (m *MessageInfo) XXX_Size() int {
	return xxx_messageInfo_MessageInfo.Size(m)
}
func (m *MessageInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MessageInfo proto.InternalMessageInfo

func (m *MessageInfo) GetType() uint32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *MessageInfo) GetVersion() []uint32 {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *MessageInfo) GetLength() uint32 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

func (m *MessageInfo) GetFieldInfos() []*FieldInfo {
	if m != nil {
		return m.FieldInfos
	}
	return nil
}

func (m *MessageInfo) GetObjectReferences() []uint64 {
	if m != nil {
		return m.ObjectReferences
	}
	return nil
}

func (m *MessageInfo) GetDataReferences() []uint64 {
	if m != nil {
		return m.DataReferences
	}
	return nil
}

type FieldInfo struct {
	Path                 *FieldPath      `protobuf:"bytes,1,req,name=path" json:"path,omitempty"`
	Type                 *FieldInfo_Type `protobuf:"varint,2,opt,name=type,enum=TSP.FieldInfo_Type,def=0" json:"type,omitempty"`
	Rule                 *FieldInfo_Rule `protobuf:"varint,3,opt,name=rule,enum=TSP.FieldInfo_Rule,def=0" json:"rule,omitempty"`
	ObjectReferences     []uint64        `protobuf:"varint,4,rep,packed,name=object_references,json=objectReferences" json:"object_references,omitempty"`
	DataReferences       []uint64        `protobuf:"varint,5,rep,packed,name=data_references,json=dataReferences" json:"data_references,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *FieldInfo) Reset()         { *m = FieldInfo{} }
func (m *FieldInfo) String() string { return proto.CompactTextString(m) }
func (*FieldInfo) ProtoMessage()    {}
func (*FieldInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_47802bcfa9d8cfa9, []int{2}
}

func (m *FieldInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldInfo.Unmarshal(m, b)
}
func (m *FieldInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldInfo.Marshal(b, m, deterministic)
}
func (m *FieldInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldInfo.Merge(m, src)
}
func (m *FieldInfo) XXX_Size() int {
	return xxx_messageInfo_FieldInfo.Size(m)
}
func (m *FieldInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FieldInfo proto.InternalMessageInfo

const Default_FieldInfo_Type FieldInfo_Type = FieldInfo_Value
const Default_FieldInfo_Rule FieldInfo_Rule = FieldInfo_IgnoreAndDrop

func (m *FieldInfo) GetPath() *FieldPath {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *FieldInfo) GetType() FieldInfo_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_FieldInfo_Type
}

func (m *FieldInfo) GetRule() FieldInfo_Rule {
	if m != nil && m.Rule != nil {
		return *m.Rule
	}
	return Default_FieldInfo_Rule
}

func (m *FieldInfo) GetObjectReferences() []uint64 {
	if m != nil {
		return m.ObjectReferences
	}
	return nil
}

func (m *FieldInfo) GetDataReferences() []uint64 {
	if m != nil {
		return m.DataReferences
	}
	return nil
}

type FieldPath struct {
	Path                 []uint32 `protobuf:"varint,1,rep,packed,name=path" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldPath) Reset()         { *m = FieldPath{} }
func (m *FieldPath) String() string { return proto.CompactTextString(m) }
func (*FieldPath) ProtoMessage()    {}
func (*FieldPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_47802bcfa9d8cfa9, []int{3}
}

func (m *FieldPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldPath.Unmarshal(m, b)
}
func (m *FieldPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldPath.Marshal(b, m, deterministic)
}
func (m *FieldPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldPath.Merge(m, src)
}
func (m *FieldPath) XXX_Size() int {
	return xxx_messageInfo_FieldPath.Size(m)
}
func (m *FieldPath) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldPath.DiscardUnknown(m)
}

var xxx_messageInfo_FieldPath proto.InternalMessageInfo

func (m *FieldPath) GetPath() []uint32 {
	if m != nil {
		return m.Path
	}
	return nil
}

type ComponentInfo struct {
	Identifier                               *uint64                       `protobuf:"varint,1,req,name=identifier" json:"identifier,omitempty"`
	PreferredLocator                         *string                       `protobuf:"bytes,2,req,name=preferred_locator,json=preferredLocator" json:"preferred_locator,omitempty"`
	Locator                                  *string                       `protobuf:"bytes,3,opt,name=locator" json:"locator,omitempty"`
	ReadVersion                              []uint32                      `protobuf:"varint,4,rep,packed,name=read_version,json=readVersion" json:"read_version,omitempty"`
	WriteVersion                             []uint32                      `protobuf:"varint,5,rep,packed,name=write_version,json=writeVersion" json:"write_version,omitempty"`
	ExternalReferences                       []*ComponentExternalReference `protobuf:"bytes,6,rep,name=external_references,json=externalReferences" json:"external_references,omitempty"`
	DataReferences                           []*ComponentDataReference     `protobuf:"bytes,7,rep,name=data_references,json=dataReferences" json:"data_references,omitempty"`
	AllowsDuplicatesOutsideOfDocumentPackage *bool                         `protobuf:"varint,8,opt,name=allows_duplicates_outside_of_document_package,json=allowsDuplicatesOutsideOfDocumentPackage,def=0" json:"allows_duplicates_outside_of_document_package,omitempty"`
	DirtiesDocumentPackage                   *bool                         `protobuf:"varint,9,opt,name=dirties_document_package,json=dirtiesDocumentPackage,def=1" json:"dirties_document_package,omitempty"`
	IsStoredOutsideObjectArchive             *bool                         `protobuf:"varint,10,opt,name=is_stored_outside_object_archive,json=isStoredOutsideObjectArchive,def=0" json:"is_stored_outside_object_archive,omitempty"`
	XXX_NoUnkeyedLiteral                     struct{}                      `json:"-"`
	XXX_unrecognized                         []byte                        `json:"-"`
	XXX_sizecache                            int32                         `json:"-"`
}

func (m *ComponentInfo) Reset()         { *m = ComponentInfo{} }
func (m *ComponentInfo) String() string { return proto.CompactTextString(m) }
func (*ComponentInfo) ProtoMessage()    {}
func (*ComponentInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_47802bcfa9d8cfa9, []int{4}
}

func (m *ComponentInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComponentInfo.Unmarshal(m, b)
}
func (m *ComponentInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComponentInfo.Marshal(b, m, deterministic)
}
func (m *ComponentInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComponentInfo.Merge(m, src)
}
func (m *ComponentInfo) XXX_Size() int {
	return xxx_messageInfo_ComponentInfo.Size(m)
}
func (m *ComponentInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ComponentInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ComponentInfo proto.InternalMessageInfo

const Default_ComponentInfo_AllowsDuplicatesOutsideOfDocumentPackage bool = false
const Default_ComponentInfo_DirtiesDocumentPackage bool = true
const Default_ComponentInfo_IsStoredOutsideObjectArchive bool = false

func (m *ComponentInfo) GetIdentifier() uint64 {
	if m != nil && m.Identifier != nil {
		return *m.Identifier
	}
	return 0
}

func (m *ComponentInfo) GetPreferredLocator() string {
	if m != nil && m.PreferredLocator != nil {
		return *m.PreferredLocator
	}
	return ""
}

func (m *ComponentInfo) GetLocator() string {
	if m != nil && m.Locator != nil {
		return *m.Locator
	}
	return ""
}

func (m *ComponentInfo) GetReadVersion() []uint32 {
	if m != nil {
		return m.ReadVersion
	}
	return nil
}

func (m *ComponentInfo) GetWriteVersion() []uint32 {
	if m != nil {
		return m.WriteVersion
	}
	return nil
}

func (m *ComponentInfo) GetExternalReferences() []*ComponentExternalReference {
	if m != nil {
		return m.ExternalReferences
	}
	return nil
}

func (m *ComponentInfo) GetDataReferences() []*ComponentDataReference {
	if m != nil {
		return m.DataReferences
	}
	return nil
}

func (m *ComponentInfo) GetAllowsDuplicatesOutsideOfDocumentPackage() bool {
	if m != nil && m.AllowsDuplicatesOutsideOfDocumentPackage != nil {
		return *m.AllowsDuplicatesOutsideOfDocumentPackage
	}
	return Default_ComponentInfo_AllowsDuplicatesOutsideOfDocumentPackage
}

func (m *ComponentInfo) GetDirtiesDocumentPackage() bool {
	if m != nil && m.DirtiesDocumentPackage != nil {
		return *m.DirtiesDocumentPackage
	}
	return Default_ComponentInfo_DirtiesDocumentPackage
}

func (m *ComponentInfo) GetIsStoredOutsideObjectArchive() bool {
	if m != nil && m.IsStoredOutsideObjectArchive != nil {
		return *m.IsStoredOutsideObjectArchive
	}
	return Default_ComponentInfo_IsStoredOutsideObjectArchive
}

type ComponentExternalReference struct {
	ComponentIdentifier  *uint64  `protobuf:"varint,1,req,name=component_identifier,json=componentIdentifier" json:"component_identifier,omitempty"`
	ObjectIdentifier     *uint64  `protobuf:"varint,2,opt,name=object_identifier,json=objectIdentifier" json:"object_identifier,omitempty"`
	IsWeak               *bool    `protobuf:"varint,3,opt,name=is_weak,json=isWeak" json:"is_weak,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComponentExternalReference) Reset()         { *m = ComponentExternalReference{} }
func (m *ComponentExternalReference) String() string { return proto.CompactTextString(m) }
func (*ComponentExternalReference) ProtoMessage()    {}
func (*ComponentExternalReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_47802bcfa9d8cfa9, []int{5}
}

func (m *ComponentExternalReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComponentExternalReference.Unmarshal(m, b)
}
func (m *ComponentExternalReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComponentExternalReference.Marshal(b, m, deterministic)
}
func (m *ComponentExternalReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComponentExternalReference.Merge(m, src)
}
func (m *ComponentExternalReference) XXX_Size() int {
	return xxx_messageInfo_ComponentExternalReference.Size(m)
}
func (m *ComponentExternalReference) XXX_DiscardUnknown() {
	xxx_messageInfo_ComponentExternalReference.DiscardUnknown(m)
}

var xxx_messageInfo_ComponentExternalReference proto.InternalMessageInfo

func (m *ComponentExternalReference) GetComponentIdentifier() uint64 {
	if m != nil && m.ComponentIdentifier != nil {
		return *m.ComponentIdentifier
	}
	return 0
}

func (m *ComponentExternalReference) GetObjectIdentifier() uint64 {
	if m != nil && m.ObjectIdentifier != nil {
		return *m.ObjectIdentifier
	}
	return 0
}

func (m *ComponentExternalReference) GetIsWeak() bool {
	if m != nil && m.IsWeak != nil {
		return *m.IsWeak
	}
	return false
}

type ComponentDataReference struct {
	DataIdentifier       *uint64  `protobuf:"varint,1,req,name=data_identifier,json=dataIdentifier" json:"data_identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComponentDataReference) Reset()         { *m = ComponentDataReference{} }
func (m *ComponentDataReference) String() string { return proto.CompactTextString(m) }
func (*ComponentDataReference) ProtoMessage()    {}
func (*ComponentDataReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_47802bcfa9d8cfa9, []int{6}
}

func (m *ComponentDataReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComponentDataReference.Unmarshal(m, b)
}
func (m *ComponentDataReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComponentDataReference.Marshal(b, m, deterministic)
}
func (m *ComponentDataReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComponentDataReference.Merge(m, src)
}
func (m *ComponentDataReference) XXX_Size() int {
	return xxx_messageInfo_ComponentDataReference.Size(m)
}
func (m *ComponentDataReference) XXX_DiscardUnknown() {
	xxx_messageInfo_ComponentDataReference.DiscardUnknown(m)
}

var xxx_messageInfo_ComponentDataReference proto.InternalMessageInfo

func (m *ComponentDataReference) GetDataIdentifier() uint64 {
	if m != nil && m.DataIdentifier != nil {
		return *m.DataIdentifier
	}
	return 0
}

type PackageMetadata struct {
	LastObjectIdentifier *uint64          `protobuf:"varint,1,req,name=last_object_identifier,json=lastObjectIdentifier" json:"last_object_identifier,omitempty"`
	Components           []*ComponentInfo `protobuf:"bytes,3,rep,name=components" json:"components,omitempty"`
	Datas                []*DataInfo      `protobuf:"bytes,4,rep,name=datas" json:"datas,omitempty"`
	ReadVersion          []uint32         `protobuf:"varint,5,rep,packed,name=read_version,json=readVersion" json:"read_version,omitempty"`
	WriteVersion         []uint32         `protobuf:"varint,6,rep,packed,name=write_version,json=writeVersion" json:"write_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PackageMetadata) Reset()         { *m = PackageMetadata{} }
func (m *PackageMetadata) String() string { return proto.CompactTextString(m) }
func (*PackageMetadata) ProtoMessage()    {}
func (*PackageMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_47802bcfa9d8cfa9, []int{7}
}

func (m *PackageMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PackageMetadata.Unmarshal(m, b)
}
func (m *PackageMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PackageMetadata.Marshal(b, m, deterministic)
}
func (m *PackageMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PackageMetadata.Merge(m, src)
}
func (m *PackageMetadata) XXX_Size() int {
	return xxx_messageInfo_PackageMetadata.Size(m)
}
func (m *PackageMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_PackageMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_PackageMetadata proto.InternalMessageInfo

func (m *PackageMetadata) GetLastObjectIdentifier() uint64 {
	if m != nil && m.LastObjectIdentifier != nil {
		return *m.LastObjectIdentifier
	}
	return 0
}

func (m *PackageMetadata) GetComponents() []*ComponentInfo {
	if m != nil {
		return m.Components
	}
	return nil
}

func (m *PackageMetadata) GetDatas() []*DataInfo {
	if m != nil {
		return m.Datas
	}
	return nil
}

func (m *PackageMetadata) GetReadVersion() []uint32 {
	if m != nil {
		return m.ReadVersion
	}
	return nil
}

func (m *PackageMetadata) GetWriteVersion() []uint32 {
	if m != nil {
		return m.WriteVersion
	}
	return nil
}

type PasteboardMetadata struct {
	Version              []uint32    `protobuf:"varint,1,rep,packed,name=version" json:"version,omitempty"`
	AppName              *string     `protobuf:"bytes,2,req,name=app_name,json=appName" json:"app_name,omitempty"`
	Datas                []*DataInfo `protobuf:"bytes,3,rep,name=datas" json:"datas,omitempty"`
	SourceDocumentUuid   *string     `protobuf:"bytes,4,opt,name=source_document_uuid,json=sourceDocumentUuid" json:"source_document_uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *PasteboardMetadata) Reset()         { *m = PasteboardMetadata{} }
func (m *PasteboardMetadata) String() string { return proto.CompactTextString(m) }
func (*PasteboardMetadata) ProtoMessage()    {}
func (*PasteboardMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_47802bcfa9d8cfa9, []int{8}
}

func (m *PasteboardMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PasteboardMetadata.Unmarshal(m, b)
}
func (m *PasteboardMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PasteboardMetadata.Marshal(b, m, deterministic)
}
func (m *PasteboardMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PasteboardMetadata.Merge(m, src)
}
func (m *PasteboardMetadata) XXX_Size() int {
	return xxx_messageInfo_PasteboardMetadata.Size(m)
}
func (m *PasteboardMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_PasteboardMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_PasteboardMetadata proto.InternalMessageInfo

func (m *PasteboardMetadata) GetVersion() []uint32 {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *PasteboardMetadata) GetAppName() string {
	if m != nil && m.AppName != nil {
		return *m.AppName
	}
	return ""
}

func (m *PasteboardMetadata) GetDatas() []*DataInfo {
	if m != nil {
		return m.Datas
	}
	return nil
}

func (m *PasteboardMetadata) GetSourceDocumentUuid() string {
	if m != nil && m.SourceDocumentUuid != nil {
		return *m.SourceDocumentUuid
	}
	return ""
}

type DataInfo struct {
	Identifier                 *uint64  `protobuf:"varint,1,req,name=identifier" json:"identifier,omitempty"`
	Digest                     []byte   `protobuf:"bytes,2,req,name=digest" json:"digest,omitempty"`
	PreferredFileName          *string  `protobuf:"bytes,3,req,name=preferred_file_name,json=preferredFileName" json:"preferred_file_name,omitempty"`
	FileName                   *string  `protobuf:"bytes,4,opt,name=file_name,json=fileName" json:"file_name,omitempty"`
	DocumentResourceLocator    *string  `protobuf:"bytes,5,opt,name=document_resource_locator,json=documentResourceLocator" json:"document_resource_locator,omitempty"`
	SourceBookmarkData         []byte   `protobuf:"bytes,6,opt,name=source_bookmark_data,json=sourceBookmarkData" json:"source_bookmark_data,omitempty"`
	PasteboardExternalFilePath *string  `protobuf:"bytes,99,opt,name=pasteboard_external_file_path,json=pasteboardExternalFilePath" json:"pasteboard_external_file_path,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *DataInfo) Reset()         { *m = DataInfo{} }
func (m *DataInfo) String() string { return proto.CompactTextString(m) }
func (*DataInfo) ProtoMessage()    {}
func (*DataInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_47802bcfa9d8cfa9, []int{9}
}

func (m *DataInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataInfo.Unmarshal(m, b)
}
func (m *DataInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataInfo.Marshal(b, m, deterministic)
}
func (m *DataInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataInfo.Merge(m, src)
}
func (m *DataInfo) XXX_Size() int {
	return xxx_messageInfo_DataInfo.Size(m)
}
func (m *DataInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DataInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DataInfo proto.InternalMessageInfo

func (m *DataInfo) GetIdentifier() uint64 {
	if m != nil && m.Identifier != nil {
		return *m.Identifier
	}
	return 0
}

func (m *DataInfo) GetDigest() []byte {
	if m != nil {
		return m.Digest
	}
	return nil
}

func (m *DataInfo) GetPreferredFileName() string {
	if m != nil && m.PreferredFileName != nil {
		return *m.PreferredFileName
	}
	return ""
}

func (m *DataInfo) GetFileName() string {
	if m != nil && m.FileName != nil {
		return *m.FileName
	}
	return ""
}

func (m *DataInfo) GetDocumentResourceLocator() string {
	if m != nil && m.DocumentResourceLocator != nil {
		return *m.DocumentResourceLocator
	}
	return ""
}

func (m *DataInfo) GetSourceBookmarkData() []byte {
	if m != nil {
		return m.SourceBookmarkData
	}
	return nil
}

func (m *DataInfo) GetPasteboardExternalFilePath() string {
	if m != nil && m.PasteboardExternalFilePath != nil {
		return *m.PasteboardExternalFilePath
	}
	return ""
}

type ViewStateMetadata struct {
	Version              []uint32       `protobuf:"varint,1,rep,packed,name=version" json:"version,omitempty"`
	DocumentVersionUuid  *string        `protobuf:"bytes,2,req,name=document_version_uuid,json=documentVersionUuid" json:"document_version_uuid,omitempty"`
	Component            *ComponentInfo `protobuf:"bytes,3,req,name=component" json:"component,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ViewStateMetadata) Reset()         { *m = ViewStateMetadata{} }
func (m *ViewStateMetadata) String() string { return proto.CompactTextString(m) }
func (*ViewStateMetadata) ProtoMessage()    {}
func (*ViewStateMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_47802bcfa9d8cfa9, []int{10}
}

func (m *ViewStateMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ViewStateMetadata.Unmarshal(m, b)
}
func (m *ViewStateMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ViewStateMetadata.Marshal(b, m, deterministic)
}
func (m *ViewStateMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ViewStateMetadata.Merge(m, src)
}
func (m *ViewStateMetadata) XXX_Size() int {
	return xxx_messageInfo_ViewStateMetadata.Size(m)
}
func (m *ViewStateMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ViewStateMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ViewStateMetadata proto.InternalMessageInfo

func (m *ViewStateMetadata) GetVersion() []uint32 {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *ViewStateMetadata) GetDocumentVersionUuid() string {
	if m != nil && m.DocumentVersionUuid != nil {
		return *m.DocumentVersionUuid
	}
	return ""
}

func (m *ViewStateMetadata) GetComponent() *ComponentInfo {
	if m != nil {
		return m.Component
	}
	return nil
}

func init() {
	proto.RegisterEnum("TSP.FieldInfo_Type", FieldInfo_Type_name, FieldInfo_Type_value)
	proto.RegisterEnum("TSP.FieldInfo_Rule", FieldInfo_Rule_name, FieldInfo_Rule_value)
	proto.RegisterType((*ArchiveInfo)(nil), "TSP.ArchiveInfo")
	proto.RegisterType((*MessageInfo)(nil), "TSP.MessageInfo")
	proto.RegisterType((*FieldInfo)(nil), "TSP.FieldInfo")
	proto.RegisterType((*FieldPath)(nil), "TSP.FieldPath")
	proto.RegisterType((*ComponentInfo)(nil), "TSP.ComponentInfo")
	proto.RegisterType((*ComponentExternalReference)(nil), "TSP.ComponentExternalReference")
	proto.RegisterType((*ComponentDataReference)(nil), "TSP.ComponentDataReference")
	proto.RegisterType((*PackageMetadata)(nil), "TSP.PackageMetadata")
	proto.RegisterType((*PasteboardMetadata)(nil), "TSP.PasteboardMetadata")
	proto.RegisterType((*DataInfo)(nil), "TSP.DataInfo")
	proto.RegisterType((*ViewStateMetadata)(nil), "TSP.ViewStateMetadata")
}

func init() { proto.RegisterFile("TSPArchiveMessages.proto", fileDescriptor_47802bcfa9d8cfa9) }

var fileDescriptor_47802bcfa9d8cfa9 = []byte{
	// 1064 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xcd, 0x6e, 0xdb, 0x46,
	0x17, 0x0d, 0x25, 0xea, 0xef, 0x4a, 0x72, 0xe8, 0x91, 0xa3, 0xd0, 0x49, 0xbe, 0xaf, 0x02, 0x83,
	0x22, 0x02, 0x82, 0x2a, 0xa9, 0xd1, 0x76, 0xe1, 0x45, 0x01, 0xa7, 0xae, 0x01, 0x03, 0x75, 0x2c,
	0x50, 0x8e, 0xbb, 0x24, 0xc6, 0xe2, 0xa5, 0xcd, 0x9a, 0xe2, 0x10, 0x33, 0x43, 0xbb, 0x79, 0x90,
	0xa2, 0xbb, 0xae, 0xfa, 0x20, 0x7d, 0x89, 0xbe, 0x42, 0xd7, 0x7d, 0x83, 0x16, 0x1c, 0x0e, 0x49,
	0xfd, 0xa5, 0x4d, 0xb5, 0xd2, 0xcc, 0x39, 0x73, 0xe7, 0x9c, 0xb9, 0x3f, 0x04, 0xfb, 0x62, 0x36,
	0x3d, 0xe2, 0xf3, 0x9b, 0xf0, 0x0e, 0xcf, 0x50, 0x08, 0x7a, 0x8d, 0x62, 0x92, 0x70, 0x26, 0x19,
	0xa9, 0x5f, 0xcc, 0xa6, 0x8e, 0x0f, 0x5d, 0x8d, 0x9e, 0xc6, 0x01, 0x23, 0xff, 0x07, 0x08, 0x7d,
	0x8c, 0x65, 0x18, 0x84, 0xc8, 0x6d, 0x63, 0x64, 0x8c, 0x4d, 0x77, 0x69, 0x87, 0x7c, 0x09, 0xfd,
	0x45, 0x1e, 0xc5, 0x0b, 0xe3, 0x80, 0x09, 0xbb, 0x36, 0xaa, 0x8f, 0xbb, 0x07, 0xd6, 0xe4, 0x62,
	0x36, 0x9d, 0xe8, 0xf8, 0x59, 0x20, 0xb7, 0xb7, 0xa8, 0x16, 0xc2, 0xf9, 0xc3, 0x80, 0xee, 0x12,
	0x4a, 0x08, 0x98, 0xf2, 0x7d, 0x82, 0xb6, 0x31, 0xaa, 0x8d, 0xfb, 0xae, 0xfa, 0x4f, 0x9e, 0x41,
	0xeb, 0x0e, 0xb9, 0x08, 0x59, 0xac, 0x82, 0xf6, 0xdf, 0xd4, 0x2c, 0xc3, 0x2d, 0xb6, 0xc8, 0x10,
	0x9a, 0x11, 0xc6, 0xd7, 0xf2, 0xc6, 0xae, 0xab, 0x33, 0x7a, 0x45, 0x5e, 0x41, 0x37, 0x08, 0x31,
	0xf2, 0xb5, 0x1c, 0x53, 0xc9, 0xd9, 0x51, 0x72, 0x4e, 0xb2, 0x7d, 0x25, 0x06, 0x82, 0xe2, 0xaf,
	0x20, 0xaf, 0x60, 0x97, 0x5d, 0xfd, 0x80, 0x73, 0xe9, 0x71, 0x0c, 0x90, 0x63, 0x3c, 0x47, 0x61,
	0x37, 0x46, 0xf5, 0xb1, 0xa9, 0x2e, 0xb4, 0x72, 0xd0, 0x2d, 0x31, 0xf2, 0x12, 0x1e, 0xfa, 0x54,
	0xd2, 0x65, 0x7a, 0xb3, 0xa4, 0xef, 0x64, 0x50, 0x45, 0x76, 0x7e, 0xa9, 0x43, 0xa7, 0xbc, 0x97,
	0x38, 0x60, 0x26, 0x54, 0xde, 0x28, 0x9b, 0x2b, 0xaa, 0xa6, 0x54, 0xde, 0xb8, 0x0a, 0x23, 0x13,
	0xfd, 0x14, 0xb5, 0x91, 0x31, 0xde, 0x39, 0x18, 0xac, 0x2a, 0x9f, 0x5c, 0xbc, 0x4f, 0xf0, 0xb0,
	0x71, 0x49, 0xa3, 0x14, 0xf5, 0x33, 0x7d, 0x05, 0x26, 0x4f, 0x23, 0xb4, 0xeb, 0x5b, 0xf9, 0x6e,
	0x1a, 0xe1, 0x61, 0xff, 0xf4, 0x3a, 0x66, 0x1c, 0x8f, 0x62, 0xff, 0x98, 0xb3, 0xc4, 0x55, 0xfc,
	0xed, 0xbe, 0xcd, 0xff, 0xe6, 0xbb, 0xf1, 0x41, 0xdf, 0x27, 0x60, 0x66, 0x52, 0x49, 0x07, 0x72,
	0xb1, 0xd6, 0x03, 0x32, 0x80, 0x87, 0xe7, 0xab, 0x31, 0x2d, 0x83, 0xec, 0x42, 0xff, 0x78, 0xf9,
	0xa4, 0x55, 0x23, 0x5d, 0x68, 0xe9, 0xd2, 0xb0, 0xea, 0x8e, 0x07, 0x66, 0x66, 0x21, 0xe3, 0xad,
	0x98, 0xb0, 0x1e, 0x90, 0x47, 0xb0, 0x5b, 0x6e, 0x4d, 0x39, 0x0a, 0xe4, 0x77, 0x59, 0x44, 0x02,
	0x3b, 0x67, 0xa9, 0x90, 0xef, 0x62, 0x1f, 0xb9, 0x90, 0x34, 0xf6, 0xad, 0x1a, 0xd9, 0x87, 0xde,
	0x5b, 0x26, 0x67, 0x69, 0x92, 0x30, 0x2e, 0xd1, 0xb7, 0xfe, 0x2a, 0x7e, 0x86, 0xf3, 0x5c, 0xe7,
	0x27, 0xcb, 0x00, 0x19, 0x96, 0xf9, 0x29, 0xea, 0x4d, 0xad, 0x9d, 0xdf, 0x4d, 0xe8, 0x7f, 0xc3,
	0x16, 0x09, 0x8b, 0x31, 0x96, 0x5b, 0xfb, 0xa2, 0xb6, 0xd6, 0x17, 0x2f, 0x61, 0x37, 0x51, 0x0f,
	0xc5, 0xd1, 0xf7, 0x22, 0x36, 0xa7, 0x92, 0x71, 0xbb, 0x36, 0xaa, 0x8d, 0x3b, 0xae, 0x55, 0x02,
	0xdf, 0xe5, 0xfb, 0xc4, 0x86, 0x56, 0x41, 0xc9, 0xb2, 0xd8, 0x71, 0x8b, 0x25, 0xf9, 0x14, 0x7a,
	0x1c, 0xa9, 0xef, 0x15, 0x8d, 0x60, 0x96, 0xc2, 0xba, 0xd9, 0xfe, 0xa5, 0x6e, 0x86, 0x17, 0xd0,
	0xbf, 0xe7, 0xa1, 0xc4, 0x92, 0xd7, 0x28, 0x79, 0x3d, 0x05, 0x14, 0xc4, 0x29, 0x0c, 0xf0, 0x47,
	0x89, 0x3c, 0xa6, 0xd1, 0x7a, 0xfd, 0x76, 0x0f, 0x3e, 0x51, 0xb5, 0x53, 0xfa, 0xfc, 0x56, 0x13,
	0xcb, 0xdc, 0xb8, 0x04, 0xd7, 0xb7, 0x04, 0x39, 0xde, 0xac, 0x8a, 0x96, 0x8a, 0xf6, 0x74, 0x35,
	0xda, 0x4a, 0x96, 0xd7, 0xcb, 0x85, 0x04, 0xf0, 0x19, 0x8d, 0x22, 0x76, 0x2f, 0x3c, 0x3f, 0x4d,
	0xa2, 0x70, 0x4e, 0x25, 0x0a, 0x8f, 0xa5, 0x52, 0x84, 0x3e, 0x7a, 0x2c, 0xf0, 0x7c, 0x36, 0x4f,
	0x17, 0x18, 0x4b, 0x2f, 0xa1, 0xf3, 0x5b, 0x7a, 0x8d, 0x76, 0x7b, 0x64, 0x8c, 0xdb, 0x87, 0x8d,
	0x80, 0x46, 0x02, 0xdd, 0x71, 0x7e, 0xf6, 0xb8, 0x3c, 0x7a, 0x9e, 0x9f, 0x3c, 0x0f, 0x8e, 0xf5,
	0xb9, 0x69, 0x7e, 0x8c, 0x7c, 0x0d, 0xb6, 0x1f, 0x72, 0x19, 0xa2, 0xd8, 0x0c, 0xd9, 0x51, 0x21,
	0x4d, 0xc9, 0x53, 0x74, 0x87, 0x9a, 0xb5, 0x7e, 0xfe, 0x0c, 0x46, 0xa1, 0xf0, 0x84, 0x64, 0x59,
	0x5a, 0x4b, 0x7d, 0x79, 0x1b, 0xd1, 0x7c, 0x6c, 0xda, 0xb0, 0x2c, 0xed, 0x59, 0x28, 0x66, 0x8a,
	0x5d, 0x48, 0x52, 0x5c, 0x3d, 0x61, 0x9d, 0x9f, 0x0c, 0x78, 0xf2, 0xe1, 0xf7, 0x26, 0x9f, 0xc3,
	0xde, 0xbc, 0x40, 0xbd, 0x8d, 0x72, 0x1b, 0x94, 0xd8, 0xe9, 0x4a, 0xdd, 0x69, 0x39, 0x4b, 0xfc,
	0x9a, 0x1a, 0xdb, 0xba, 0xa3, 0x97, 0xc8, 0x8f, 0xa1, 0x15, 0x0a, 0xef, 0x1e, 0xe9, 0xad, 0xaa,
	0xbb, 0xb6, 0xdb, 0x0c, 0xc5, 0xf7, 0x48, 0x6f, 0x9d, 0x23, 0x18, 0x6e, 0x4f, 0x1c, 0x79, 0xa1,
	0xd3, 0xbd, 0xa1, 0x46, 0x65, 0xb4, 0x8a, 0xed, 0xfc, 0x69, 0xc0, 0x43, 0xfd, 0x6a, 0x67, 0x28,
	0x69, 0x86, 0x92, 0x2f, 0x60, 0x18, 0x51, 0x21, 0xbd, 0x4d, 0x85, 0x79, 0x8c, 0xbd, 0x0c, 0x3d,
	0x5f, 0x57, 0x79, 0x00, 0x50, 0x3a, 0x15, 0x76, 0x5d, 0x15, 0x17, 0x59, 0x2d, 0xae, 0x7c, 0xa8,
	0x57, 0x2c, 0xf2, 0x1c, 0x1a, 0xd9, 0x8d, 0xc5, 0xfc, 0xef, 0x2b, 0x7a, 0xe6, 0x44, 0x31, 0x73,
	0x6c, 0xa3, 0xb9, 0x1a, 0x1f, 0xd9, 0x5c, 0xcd, 0xed, 0xcd, 0xe5, 0xfc, 0x6a, 0x00, 0x99, 0x52,
	0x21, 0xf1, 0x8a, 0x51, 0xee, 0x97, 0xae, 0x97, 0xbe, 0x63, 0xc6, 0xe6, 0x77, 0x6c, 0x1f, 0xda,
	0x34, 0x49, 0xbc, 0x98, 0x2e, 0x50, 0xcf, 0x87, 0x16, 0x4d, 0x92, 0xb7, 0x74, 0x81, 0x95, 0x89,
	0xfa, 0x3f, 0x98, 0x78, 0x0d, 0x7b, 0x82, 0xa5, 0x7c, 0x8e, 0x55, 0x41, 0xa7, 0x69, 0xe8, 0xdb,
	0xa6, 0x1a, 0x24, 0x24, 0xc7, 0x8a, 0x32, 0x7e, 0x97, 0x86, 0xbe, 0xf3, 0x5b, 0x0d, 0xda, 0x45,
	0x94, 0x7f, 0x9d, 0x63, 0x43, 0x68, 0xfa, 0xe1, 0x35, 0x0a, 0xa9, 0xc4, 0xf5, 0x5c, 0xbd, 0x22,
	0x13, 0x18, 0x54, 0xf3, 0x2d, 0x08, 0x23, 0xcc, 0x1d, 0xd4, 0x95, 0x83, 0x6a, 0xf4, 0x9d, 0x84,
	0x11, 0x2a, 0x2f, 0x4f, 0xa1, 0x53, 0xb1, 0x72, 0x6d, 0xed, 0xa0, 0x00, 0x0f, 0x61, 0xbf, 0x14,
	0xcf, 0x51, 0xdb, 0x29, 0x26, 0x62, 0x43, 0x91, 0x1f, 0x17, 0x04, 0x57, 0xe3, 0xc5, 0xec, 0xac,
	0xfc, 0x5f, 0x31, 0x76, 0xbb, 0xa0, 0xfc, 0xd6, 0xcb, 0x1e, 0xc6, 0x6e, 0x8e, 0x8c, 0x71, 0xaf,
	0xf0, 0xff, 0x46, 0x43, 0x99, 0x6d, 0x72, 0x04, 0xff, 0x4b, 0xca, 0x2c, 0x79, 0xe5, 0x38, 0x54,
	0xf2, 0xd4, 0xf4, 0x9f, 0xab, 0x1b, 0x9f, 0x54, 0xa4, 0xa2, 0x33, 0x33, 0x37, 0xd9, 0x77, 0xc2,
	0xf9, 0xd9, 0x80, 0xdd, 0xcb, 0x10, 0xef, 0x67, 0x92, 0x4a, 0xfc, 0xc8, 0x44, 0x1f, 0xc0, 0xa3,
	0xd2, 0xa4, 0xde, 0xcb, 0x33, 0x95, 0x67, 0x7d, 0x50, 0x80, 0xba, 0x9a, 0xb2, 0x54, 0x91, 0xd7,
	0xd0, 0x29, 0x8b, 0x5a, 0xbd, 0xed, 0xf6, 0xca, 0xaf, 0x48, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff,
	0xd5, 0x7b, 0xc3, 0x39, 0xde, 0x09, 0x00, 0x00,
}
