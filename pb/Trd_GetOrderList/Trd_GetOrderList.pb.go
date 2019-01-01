// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Trd_GetOrderList/Trd_GetOrderList.proto

package Trd_GetOrderList

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "futugg/pb/Common"
import Trd_Common "futugg/pb/Trd_Common"

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
	Header               *Trd_Common.TrdHeader           `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	FilterConditions     *Trd_Common.TrdFilterConditions `protobuf:"bytes,2,opt,name=filterConditions" json:"filterConditions,omitempty"`
	FilterStatusList     []int32                         `protobuf:"varint,3,rep,name=filterStatusList" json:"filterStatusList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *C2S) Reset()         { *m = C2S{} }
func (m *C2S) String() string { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()    {}
func (*C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_Trd_GetOrderList_3375e4451a8bc156, []int{0}
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

func (m *C2S) GetHeader() *Trd_Common.TrdHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *C2S) GetFilterConditions() *Trd_Common.TrdFilterConditions {
	if m != nil {
		return m.FilterConditions
	}
	return nil
}

func (m *C2S) GetFilterStatusList() []int32 {
	if m != nil {
		return m.FilterStatusList
	}
	return nil
}

type S2C struct {
	Header               *Trd_Common.TrdHeader `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	OrderList            []*Trd_Common.Order   `protobuf:"bytes,2,rep,name=orderList" json:"orderList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_Trd_GetOrderList_3375e4451a8bc156, []int{1}
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

func (m *S2C) GetHeader() *Trd_Common.TrdHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *S2C) GetOrderList() []*Trd_Common.Order {
	if m != nil {
		return m.OrderList
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
	return fileDescriptor_Trd_GetOrderList_3375e4451a8bc156, []int{2}
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
	// 以下3个字段每条协议都有，注释说明在InitConnect.proto中
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
	return fileDescriptor_Trd_GetOrderList_3375e4451a8bc156, []int{3}
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
	proto.RegisterType((*C2S)(nil), "Trd_GetOrderList.C2S")
	proto.RegisterType((*S2C)(nil), "Trd_GetOrderList.S2C")
	proto.RegisterType((*Request)(nil), "Trd_GetOrderList.Request")
	proto.RegisterType((*Response)(nil), "Trd_GetOrderList.Response")
}

func init() {
	proto.RegisterFile("Trd_GetOrderList/Trd_GetOrderList.proto", fileDescriptor_Trd_GetOrderList_3375e4451a8bc156)
}

var fileDescriptor_Trd_GetOrderList_3375e4451a8bc156 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x85, 0x99, 0xa6, 0x3f, 0xf6, 0x76, 0x53, 0x23, 0x4a, 0xa8, 0xa0, 0xc3, 0x6c, 0x3a, 0x08,
	0xed, 0x94, 0xe0, 0xca, 0x6d, 0x40, 0x05, 0x15, 0x21, 0x33, 0x7b, 0x29, 0xcd, 0x55, 0x07, 0xec,
	0x64, 0x4c, 0xd2, 0x85, 0x0f, 0xe0, 0xe3, 0xf8, 0x8e, 0x32, 0x7f, 0xb4, 0x9d, 0xba, 0x71, 0x95,
	0xdc, 0x73, 0xbf, 0x93, 0xdc, 0x9c, 0xc0, 0x34, 0x31, 0xea, 0xe5, 0x0e, 0xdd, 0xb3, 0x51, 0x68,
	0x1e, 0x53, 0xeb, 0xa2, 0xb6, 0x30, 0xcf, 0x8d, 0x76, 0x9a, 0x8e, 0xdb, 0xfa, 0xe4, 0x44, 0xe8,
	0xf5, 0x5a, 0x67, 0x51, 0xb5, 0x54, 0xd8, 0xe4, 0xbc, 0xc0, 0xea, 0xc6, 0x76, 0x5b, 0x35, 0x83,
	0x1f, 0x0f, 0x88, 0xe0, 0x31, 0x9d, 0x41, 0xff, 0x1d, 0x97, 0x0a, 0x0d, 0xf3, 0xfc, 0x4e, 0x38,
	0xe2, 0xa7, 0xf3, 0x1d, 0x34, 0x31, 0xea, 0xbe, 0x6c, 0xca, 0x1a, 0xa2, 0x0f, 0x30, 0x7e, 0x4d,
	0x3f, 0x1c, 0x1a, 0xa1, 0x33, 0x95, 0xba, 0x54, 0x67, 0x96, 0x75, 0x7c, 0x2f, 0x1c, 0xf1, 0xcb,
	0x96, 0xf1, 0xb6, 0x85, 0xc9, 0x03, 0x23, 0xbd, 0x6a, 0x0e, 0x8b, 0xdd, 0xd2, 0x6d, 0x6c, 0xf1,
	0x12, 0x46, 0x7c, 0x12, 0xf6, 0xe4, 0x81, 0x1e, 0x20, 0x90, 0x98, 0x8b, 0xff, 0x8e, 0x1b, 0xc1,
	0x50, 0x37, 0x21, 0xb1, 0x8e, 0x4f, 0xc2, 0x11, 0x3f, 0xde, 0x75, 0x94, 0x09, 0xca, 0x2d, 0x13,
	0x70, 0x18, 0x48, 0xfc, 0xdc, 0xa0, 0x75, 0x74, 0x0a, 0x64, 0xc5, 0xed, 0xde, 0x3d, 0x7b, 0x7f,
	0x21, 0x78, 0x2c, 0x0b, 0x22, 0xf8, 0xf6, 0xe0, 0x48, 0xa2, 0xcd, 0x75, 0x66, 0x91, 0x5e, 0xc0,
	0xc0, 0xa0, 0x4b, 0xbe, 0x72, 0x2c, 0x9d, 0xbd, 0x9b, 0xee, 0xec, 0x7a, 0xb1, 0x90, 0x8d, 0x48,
	0xcf, 0xa0, 0x6f, 0xd0, 0x3d, 0xd9, 0xb7, 0x32, 0xb6, 0xa1, 0xac, 0x2b, 0xca, 0x60, 0x80, 0xc6,
	0x08, 0xad, 0x90, 0x11, 0xdf, 0x0b, 0x7b, 0xb2, 0x29, 0x8b, 0x39, 0x2c, 0x5f, 0xb1, 0x6e, 0x99,
	0xf2, 0x1f, 0x73, 0xc4, 0x5c, 0xc8, 0x82, 0xf8, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xa6, 0xff, 0x7b,
	0x5d, 0x40, 0x02, 0x00, 0x00,
}
