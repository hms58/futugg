// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Qot_GetPlateSet/Qot_GetPlateSet.proto

package Qot_GetPlateSet

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "futugg/pb/Common"
import Qot_Common "futugg/pb/Qot_Common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type C2S struct {
	Market               *int32   `protobuf:"varint,1,req,name=market" json:"market,omitempty"`
	PlateSetType         *int32   `protobuf:"varint,2,req,name=plateSetType" json:"plateSetType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2S) Reset()         { *m = C2S{} }
func (m *C2S) String() string { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()    {}
func (*C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetPlateSet_309dce51f7aad9b2, []int{0}
}
func (m *C2S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2S.Unmarshal(m, b)
}
func (m *C2S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2S.Marshal(b, m, deterministic)
}
func (dst *C2S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2S.Merge(dst, src)
}
func (m *C2S) XXX_Size() int {
	return xxx_messageInfo_C2S.Size(m)
}
func (m *C2S) XXX_DiscardUnknown() {
	xxx_messageInfo_C2S.DiscardUnknown(m)
}

var xxx_messageInfo_C2S proto.InternalMessageInfo

func (m *C2S) GetMarket() int32 {
	if m != nil && m.Market != nil {
		return *m.Market
	}
	return 0
}

func (m *C2S) GetPlateSetType() int32 {
	if m != nil && m.PlateSetType != nil {
		return *m.PlateSetType
	}
	return 0
}

type PlateInfo struct {
	Plate                *Qot_Common.Security `protobuf:"bytes,1,req,name=plate" json:"plate,omitempty"`
	Name                 *string              `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PlateInfo) Reset()         { *m = PlateInfo{} }
func (m *PlateInfo) String() string { return proto.CompactTextString(m) }
func (*PlateInfo) ProtoMessage()    {}
func (*PlateInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetPlateSet_309dce51f7aad9b2, []int{1}
}
func (m *PlateInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlateInfo.Unmarshal(m, b)
}
func (m *PlateInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlateInfo.Marshal(b, m, deterministic)
}
func (dst *PlateInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlateInfo.Merge(dst, src)
}
func (m *PlateInfo) XXX_Size() int {
	return xxx_messageInfo_PlateInfo.Size(m)
}
func (m *PlateInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PlateInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PlateInfo proto.InternalMessageInfo

func (m *PlateInfo) GetPlate() *Qot_Common.Security {
	if m != nil {
		return m.Plate
	}
	return nil
}

func (m *PlateInfo) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type S2C struct {
	PlateInfoList        []*PlateInfo `protobuf:"bytes,1,rep,name=plateInfoList" json:"plateInfoList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetPlateSet_309dce51f7aad9b2, []int{2}
}
func (m *S2C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2C.Unmarshal(m, b)
}
func (m *S2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2C.Marshal(b, m, deterministic)
}
func (dst *S2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C.Merge(dst, src)
}
func (m *S2C) XXX_Size() int {
	return xxx_messageInfo_S2C.Size(m)
}
func (m *S2C) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C.DiscardUnknown(m)
}

var xxx_messageInfo_S2C proto.InternalMessageInfo

func (m *S2C) GetPlateInfoList() []*PlateInfo {
	if m != nil {
		return m.PlateInfoList
	}
	return nil
}

type Request struct {
	C2S                  *C2S     `protobuf:"bytes,1,req,name=c2s" json:"c2s,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetPlateSet_309dce51f7aad9b2, []int{3}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetC2S() *C2S {
	if m != nil {
		return m.C2S
	}
	return nil
}

type Response struct {
	RetType              *int32   `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"`
	RetMsg               *string  `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg,omitempty"`
	ErrCode              *int32   `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	S2C                  *S2C     `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetPlateSet_309dce51f7aad9b2, []int{4}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

const Default_Response_RetType int32 = -400

func (m *Response) GetRetType() int32 {
	if m != nil && m.RetType != nil {
		return *m.RetType
	}
	return Default_Response_RetType
}

func (m *Response) GetRetMsg() string {
	if m != nil && m.RetMsg != nil {
		return *m.RetMsg
	}
	return ""
}

func (m *Response) GetErrCode() int32 {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *Response) GetS2C() *S2C {
	if m != nil {
		return m.S2C
	}
	return nil
}

func init() {
	proto.RegisterType((*C2S)(nil), "Qot_GetPlateSet.C2S")
	proto.RegisterType((*PlateInfo)(nil), "Qot_GetPlateSet.PlateInfo")
	proto.RegisterType((*S2C)(nil), "Qot_GetPlateSet.S2C")
	proto.RegisterType((*Request)(nil), "Qot_GetPlateSet.Request")
	proto.RegisterType((*Response)(nil), "Qot_GetPlateSet.Response")
}

func init() {
	proto.RegisterFile("Qot_GetPlateSet/Qot_GetPlateSet.proto", fileDescriptor_Qot_GetPlateSet_309dce51f7aad9b2)
}

var fileDescriptor_Qot_GetPlateSet_309dce51f7aad9b2 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0xa5, 0xeb, 0xe6, 0xdc, 0x37, 0x45, 0x88, 0x22, 0x61, 0x82, 0x94, 0x80, 0x52, 0x04, 0xb7,
	0x19, 0x3c, 0x79, 0x52, 0x72, 0x18, 0xa2, 0x82, 0x26, 0xde, 0x65, 0xcc, 0x4f, 0x19, 0xda, 0xa6,
	0x26, 0xd9, 0x61, 0x77, 0x7f, 0xb8, 0x24, 0x6d, 0x75, 0x2d, 0x9e, 0xfa, 0xbd, 0xef, 0xbd, 0xbe,
	0xef, 0xe5, 0xc1, 0xc9, 0x93, 0x76, 0x2f, 0x33, 0x74, 0x8f, 0x9f, 0x73, 0x87, 0x0a, 0xdd, 0xa4,
	0x85, 0xc7, 0x85, 0xd1, 0x4e, 0x93, 0xbd, 0xd6, 0x7a, 0xb4, 0x2f, 0x74, 0x96, 0xe9, 0x7c, 0x52,
	0x7e, 0x4a, 0xd5, 0xe8, 0xc8, 0xab, 0x2a, 0xe2, 0x6f, 0x2c, 0x49, 0x76, 0x03, 0xb1, 0xe0, 0x8a,
	0x1c, 0xc2, 0x56, 0x36, 0x37, 0x1f, 0xe8, 0x68, 0x94, 0x74, 0xd2, 0x9e, 0xac, 0x10, 0x61, 0xb0,
	0x53, 0x54, 0xe6, 0xcf, 0xeb, 0x02, 0x69, 0x27, 0xb0, 0x8d, 0x1d, 0xbb, 0x83, 0x41, 0x08, 0x70,
	0x9b, 0xbf, 0x69, 0x72, 0x06, 0xbd, 0x40, 0x06, 0x9f, 0x21, 0x3f, 0x18, 0x6f, 0x5c, 0x54, 0xb8,
	0x58, 0x99, 0xa5, 0x5b, 0xcb, 0x52, 0x42, 0x08, 0x74, 0xf3, 0x79, 0x56, 0x9a, 0x0e, 0x64, 0x98,
	0xd9, 0x0c, 0x62, 0xc5, 0x05, 0xb9, 0x86, 0xdd, 0xa2, 0xf6, 0xbc, 0x5f, 0x5a, 0x1f, 0x2b, 0x4e,
	0x87, 0x7c, 0x34, 0x6e, 0x17, 0xf1, 0x7b, 0x59, 0x36, 0x7f, 0x60, 0x17, 0xd0, 0x97, 0xf8, 0xb5,
	0x42, 0xeb, 0xc8, 0x29, 0xc4, 0x0b, 0x6e, 0x1b, 0x89, 0x36, 0x2d, 0x04, 0x57, 0xd2, 0x0b, 0xd8,
	0x77, 0x04, 0xdb, 0x12, 0x6d, 0xa1, 0x73, 0x8b, 0xe4, 0x18, 0xfa, 0xa6, 0x7a, 0x74, 0xa8, 0xe4,
	0xaa, 0x7b, 0x7e, 0x39, 0x9d, 0xca, 0x7a, 0xe9, 0x1b, 0x33, 0xe8, 0x1e, 0xec, 0x3b, 0xed, 0x24,
	0x51, 0x3a, 0x90, 0x15, 0x22, 0x14, 0xfa, 0x68, 0x8c, 0xd0, 0xaf, 0x48, 0xe3, 0x24, 0x4a, 0x7b,
	0xb2, 0x86, 0x3e, 0x86, 0xe5, 0x0b, 0xda, 0x4d, 0xa2, 0x7f, 0x63, 0x28, 0x2e, 0xa4, 0x17, 0xfc,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x5e, 0xb0, 0xcd, 0x36, 0xfd, 0x01, 0x00, 0x00,
}
