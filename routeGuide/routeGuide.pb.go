// Code generated by protoc-gen-go. DO NOT EDIT.
// source: routeGuide.proto

package routeGuide

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Tuid struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tuid) Reset()         { *m = Tuid{} }
func (m *Tuid) String() string { return proto.CompactTextString(m) }
func (*Tuid) ProtoMessage()    {}
func (*Tuid) Descriptor() ([]byte, []int) {
	return fileDescriptor_routeGuide_b8e450cdd0b1490a, []int{0}
}
func (m *Tuid) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tuid.Unmarshal(m, b)
}
func (m *Tuid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tuid.Marshal(b, m, deterministic)
}
func (dst *Tuid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tuid.Merge(dst, src)
}
func (m *Tuid) XXX_Size() int {
	return xxx_messageInfo_Tuid.Size(m)
}
func (m *Tuid) XXX_DiscardUnknown() {
	xxx_messageInfo_Tuid.DiscardUnknown(m)
}

var xxx_messageInfo_Tuid proto.InternalMessageInfo

func (m *Tuid) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type TaInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Passwd               string   `protobuf:"bytes,2,opt,name=Passwd,proto3" json:"Passwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaInfo) Reset()         { *m = TaInfo{} }
func (m *TaInfo) String() string { return proto.CompactTextString(m) }
func (*TaInfo) ProtoMessage()    {}
func (*TaInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_routeGuide_b8e450cdd0b1490a, []int{1}
}
func (m *TaInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaInfo.Unmarshal(m, b)
}
func (m *TaInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaInfo.Marshal(b, m, deterministic)
}
func (dst *TaInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaInfo.Merge(dst, src)
}
func (m *TaInfo) XXX_Size() int {
	return xxx_messageInfo_TaInfo.Size(m)
}
func (m *TaInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TaInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TaInfo proto.InternalMessageInfo

func (m *TaInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TaInfo) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

type ItemInfo struct {
	Ta                   *TaInfo  `protobuf:"bytes,1,opt,name=Ta,proto3" json:"Ta,omitempty"`
	Duration             int64    `protobuf:"varint,2,opt,name=Duration,proto3" json:"Duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemInfo) Reset()         { *m = ItemInfo{} }
func (m *ItemInfo) String() string { return proto.CompactTextString(m) }
func (*ItemInfo) ProtoMessage()    {}
func (*ItemInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_routeGuide_b8e450cdd0b1490a, []int{2}
}
func (m *ItemInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemInfo.Unmarshal(m, b)
}
func (m *ItemInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemInfo.Marshal(b, m, deterministic)
}
func (dst *ItemInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemInfo.Merge(dst, src)
}
func (m *ItemInfo) XXX_Size() int {
	return xxx_messageInfo_ItemInfo.Size(m)
}
func (m *ItemInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ItemInfo proto.InternalMessageInfo

func (m *ItemInfo) GetTa() *TaInfo {
	if m != nil {
		return m.Ta
	}
	return nil
}

func (m *ItemInfo) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func init() {
	proto.RegisterType((*Tuid)(nil), "routeGuide.Tuid")
	proto.RegisterType((*TaInfo)(nil), "routeGuide.TaInfo")
	proto.RegisterType((*ItemInfo)(nil), "routeGuide.ItemInfo")
}

func init() { proto.RegisterFile("routeGuide.proto", fileDescriptor_routeGuide_b8e450cdd0b1490a) }

var fileDescriptor_routeGuide_b8e450cdd0b1490a = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0xca, 0x2f, 0x2d,
	0x49, 0x75, 0x2f, 0xcd, 0x4c, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x28, 0x49, 0x70, 0xb1, 0x84, 0x94, 0x66, 0xa6, 0x08, 0x09, 0x70, 0x31, 0x87, 0x66, 0xa6, 0x48,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0x98, 0x4a, 0x26, 0x5c, 0x6c, 0x21, 0x89, 0x9e, 0x79,
	0x69, 0xf9, 0x42, 0x42, 0x5c, 0x2c, 0x7e, 0x89, 0xb9, 0xa9, 0x50, 0x49, 0x30, 0x5b, 0x48, 0x8c,
	0x8b, 0x2d, 0x20, 0xb1, 0xb8, 0xb8, 0x3c, 0x45, 0x82, 0x09, 0x2c, 0x0a, 0xe5, 0x29, 0x79, 0x71,
	0x71, 0x78, 0x96, 0xa4, 0xe6, 0x82, 0xf5, 0x29, 0x71, 0x31, 0x85, 0x24, 0x82, 0x75, 0x71, 0x1b,
	0x09, 0xe9, 0x21, 0x39, 0x03, 0x62, 0x6e, 0x10, 0x53, 0x48, 0xa2, 0x90, 0x14, 0x17, 0x87, 0x4b,
	0x69, 0x51, 0x62, 0x49, 0x66, 0x7e, 0x1e, 0xd8, 0x24, 0xe6, 0x20, 0x38, 0xdf, 0xa8, 0x94, 0x8b,
	0x3d, 0x28, 0x35, 0x25, 0xb3, 0xd8, 0xbf, 0x40, 0xc8, 0x90, 0x8b, 0xdd, 0x3d, 0xb5, 0xc4, 0xb1,
	0xb4, 0x24, 0x43, 0x48, 0x00, 0xc5, 0xa4, 0xd2, 0xcc, 0x14, 0x29, 0x2c, 0x66, 0x2b, 0x31, 0x08,
	0x99, 0x73, 0x71, 0x07, 0x43, 0xb4, 0x80, 0x1d, 0x23, 0x82, 0xac, 0x08, 0xe6, 0x44, 0x29, 0x0c,
	0xc3, 0x94, 0x18, 0x92, 0xd8, 0xc0, 0xa1, 0x64, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x02, 0x1f,
	0x32, 0xed, 0x39, 0x01, 0x00, 0x00,
}