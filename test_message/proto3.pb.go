// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto3.proto

package test_message

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

// 枚举类型
type TestMessage_Status int32

const (
	TestMessage_OK   TestMessage_Status = 0
	TestMessage_FAIL TestMessage_Status = 1
)

var TestMessage_Status_name = map[int32]string{
	0: "OK",
	1: "FAIL",
}

var TestMessage_Status_value = map[string]int32{
	"OK":   0,
	"FAIL": 1,
}

func (x TestMessage_Status) String() string {
	return proto.EnumName(TestMessage_Status_name, int32(x))
}

func (TestMessage_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4fee6d65e34a64b6, []int{0, 0}
}

// 测试message 采用为驼峰的命令方式
type TestMessage struct {
	TeName  string               `protobuf:"bytes,1,opt,name=te_name,json=teName,proto3" json:"te_name,omitempty"`
	TeAge   int32                `protobuf:"varint,2,opt,name=te_age,json=teAge,proto3" json:"te_age,omitempty"`
	TeCount int32                `protobuf:"varint,3,opt,name=te_count,json=teCount,proto3" json:"te_count,omitempty"`
	TeMoney float64              `protobuf:"fixed64,4,opt,name=te_money,json=teMoney,proto3" json:"te_money,omitempty"`
	TeScore float32              `protobuf:"fixed32,5,opt,name=te_score,json=teScore,proto3" json:"te_score,omitempty"`
	TeFat   bool                 `protobuf:"varint,6,opt,name=te_fat,json=teFat,proto3" json:"te_fat,omitempty"`
	TeChar  []byte               `protobuf:"bytes,7,opt,name=te_char,json=teChar,proto3" json:"te_char,omitempty"`
	Childs  *TestMessage_TeChild `protobuf:"bytes,9,opt,name=childs,proto3" json:"childs,omitempty"`
	// map类型
	TeMap                map[string]int32 `protobuf:"bytes,10,rep,name=te_map,json=teMap,proto3" json:"te_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TestMessage) Reset()         { *m = TestMessage{} }
func (m *TestMessage) String() string { return proto.CompactTextString(m) }
func (*TestMessage) ProtoMessage()    {}
func (*TestMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fee6d65e34a64b6, []int{0}
}

func (m *TestMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestMessage.Unmarshal(m, b)
}
func (m *TestMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestMessage.Marshal(b, m, deterministic)
}
func (m *TestMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestMessage.Merge(m, src)
}
func (m *TestMessage) XXX_Size() int {
	return xxx_messageInfo_TestMessage.Size(m)
}
func (m *TestMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_TestMessage.DiscardUnknown(m)
}

var xxx_messageInfo_TestMessage proto.InternalMessageInfo

func (m *TestMessage) GetTeName() string {
	if m != nil {
		return m.TeName
	}
	return ""
}

func (m *TestMessage) GetTeAge() int32 {
	if m != nil {
		return m.TeAge
	}
	return 0
}

func (m *TestMessage) GetTeCount() int32 {
	if m != nil {
		return m.TeCount
	}
	return 0
}

func (m *TestMessage) GetTeMoney() float64 {
	if m != nil {
		return m.TeMoney
	}
	return 0
}

func (m *TestMessage) GetTeScore() float32 {
	if m != nil {
		return m.TeScore
	}
	return 0
}

func (m *TestMessage) GetTeFat() bool {
	if m != nil {
		return m.TeFat
	}
	return false
}

func (m *TestMessage) GetTeChar() []byte {
	if m != nil {
		return m.TeChar
	}
	return nil
}

func (m *TestMessage) GetChilds() *TestMessage_TeChild {
	if m != nil {
		return m.Childs
	}
	return nil
}

func (m *TestMessage) GetTeMap() map[string]int32 {
	if m != nil {
		return m.TeMap
	}
	return nil
}

type TestMessage_TeChild struct {
	ChName               string   `protobuf:"bytes,1,opt,name=ch_name,json=chName,proto3" json:"ch_name,omitempty"`
	ChSex                string   `protobuf:"bytes,2,opt,name=ch_sex,json=chSex,proto3" json:"ch_sex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestMessage_TeChild) Reset()         { *m = TestMessage_TeChild{} }
func (m *TestMessage_TeChild) String() string { return proto.CompactTextString(m) }
func (*TestMessage_TeChild) ProtoMessage()    {}
func (*TestMessage_TeChild) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fee6d65e34a64b6, []int{0, 0}
}

func (m *TestMessage_TeChild) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestMessage_TeChild.Unmarshal(m, b)
}
func (m *TestMessage_TeChild) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestMessage_TeChild.Marshal(b, m, deterministic)
}
func (m *TestMessage_TeChild) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestMessage_TeChild.Merge(m, src)
}
func (m *TestMessage_TeChild) XXX_Size() int {
	return xxx_messageInfo_TestMessage_TeChild.Size(m)
}
func (m *TestMessage_TeChild) XXX_DiscardUnknown() {
	xxx_messageInfo_TestMessage_TeChild.DiscardUnknown(m)
}

var xxx_messageInfo_TestMessage_TeChild proto.InternalMessageInfo

func (m *TestMessage_TeChild) GetChName() string {
	if m != nil {
		return m.ChName
	}
	return ""
}

func (m *TestMessage_TeChild) GetChSex() string {
	if m != nil {
		return m.ChSex
	}
	return ""
}

func init() {
	proto.RegisterEnum("main.TestMessage_Status", TestMessage_Status_name, TestMessage_Status_value)
	proto.RegisterType((*TestMessage)(nil), "main.TestMessage")
	proto.RegisterMapType((map[string]int32)(nil), "main.TestMessage.TeMapEntry")
	proto.RegisterType((*TestMessage_TeChild)(nil), "main.TestMessage.TeChild")
}

func init() { proto.RegisterFile("proto3.proto", fileDescriptor_4fee6d65e34a64b6) }

var fileDescriptor_4fee6d65e34a64b6 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x4d, 0x6f, 0xea, 0x30,
	0x10, 0x7c, 0x4e, 0x48, 0x02, 0x0b, 0x07, 0x64, 0xbd, 0xa7, 0x67, 0x50, 0x0f, 0x16, 0x27, 0x9f,
	0x22, 0x15, 0x2e, 0xb4, 0x37, 0x84, 0x8a, 0x54, 0xb5, 0x69, 0xa5, 0xc0, 0x3d, 0x72, 0xdd, 0x2d,
	0x41, 0x25, 0x1f, 0x4a, 0x96, 0x0a, 0x7e, 0x59, 0xff, 0x5e, 0xe5, 0x34, 0xfd, 0x92, 0x7a, 0xb2,
	0x47, 0xb3, 0xbb, 0x33, 0x3b, 0x0b, 0x83, 0xb2, 0x2a, 0xa8, 0x98, 0x85, 0xcd, 0xc3, 0x3b, 0x99,
	0xde, 0xe5, 0x93, 0x57, 0x17, 0xfa, 0x1b, 0xac, 0x29, 0xc2, 0xba, 0xd6, 0x5b, 0xe4, 0xff, 0x21,
	0x20, 0x4c, 0x72, 0x9d, 0xa1, 0x60, 0x92, 0xa9, 0x5e, 0xec, 0x13, 0xde, 0xe9, 0x0c, 0xf9, 0x3f,
	0xf0, 0x09, 0x13, 0xbd, 0x45, 0xe1, 0x48, 0xa6, 0xbc, 0xd8, 0x23, 0x5c, 0x6c, 0x91, 0x8f, 0xa0,
	0x4b, 0x98, 0x98, 0xe2, 0x90, 0x93, 0x70, 0x1b, 0x22, 0x20, 0x5c, 0x5a, 0xd8, 0x52, 0x59, 0x91,
	0xe3, 0x49, 0x74, 0x24, 0x53, 0xcc, 0x52, 0x91, 0x85, 0x2d, 0x55, 0x9b, 0xa2, 0x42, 0xe1, 0x49,
	0xa6, 0x1c, 0x4b, 0xad, 0x2d, 0x6c, 0x75, 0x9e, 0x34, 0x09, 0x5f, 0x32, 0xd5, 0xb5, 0x3a, 0x2b,
	0x4d, 0xad, 0x2f, 0x93, 0xea, 0x4a, 0x04, 0x92, 0xa9, 0x81, 0xf5, 0xb5, 0x4c, 0x75, 0xc5, 0xcf,
	0xc1, 0x37, 0xe9, 0x6e, 0xff, 0x58, 0x8b, 0x9e, 0x64, 0xaa, 0x3f, 0x1d, 0x85, 0x76, 0xaf, 0xf0,
	0xdb, 0x4e, 0xe1, 0x06, 0x97, 0xb6, 0x22, 0x6e, 0x0b, 0xf9, 0xac, 0x91, 0xc8, 0x74, 0x29, 0x40,
	0xba, 0xaa, 0x3f, 0x3d, 0xfb, 0xad, 0x25, 0xd2, 0xe5, 0x55, 0x4e, 0xd5, 0xc9, 0x1a, 0x88, 0x74,
	0x39, 0xbe, 0x80, 0xa0, 0x9d, 0x63, 0xbd, 0x98, 0xf4, 0x47, 0x46, 0x26, 0xfd, 0xc8, 0xc8, 0xa4,
	0x49, 0x8d, 0xc7, 0x26, 0xa3, 0x5e, 0xec, 0x99, 0x74, 0x8d, 0xc7, 0xf1, 0x1c, 0xe0, 0x6b, 0x1e,
	0x1f, 0x82, 0xfb, 0x8c, 0xa7, 0xb6, 0xd3, 0x7e, 0xf9, 0x5f, 0xf0, 0x5e, 0xf4, 0xfe, 0xf0, 0x99,
	0x6c, 0x03, 0x2e, 0x9d, 0x39, 0x9b, 0x8c, 0xc1, 0x5f, 0x93, 0xa6, 0x43, 0xcd, 0x7d, 0x70, 0xee,
	0x6f, 0x86, 0x7f, 0x78, 0x17, 0x3a, 0xab, 0xc5, 0xf5, 0xed, 0x90, 0x3d, 0xf8, 0xef, 0xd7, 0x7c,
	0x0b, 0x00, 0x00, 0xff, 0xff, 0x30, 0x6b, 0xfd, 0xd9, 0xd6, 0x01, 0x00, 0x00,
}