// Code generated by protoc-gen-go. DO NOT EDIT.
// source: KeepAlive/KeepAlive.proto

package KeepAlive

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "futugg/pb/Common"

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
	Time                 *int64   `protobuf:"varint,1,req,name=time" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2S) Reset()         { *m = C2S{} }
func (m *C2S) String() string { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()    {}
func (*C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_KeepAlive_8ce838656139b308, []int{0}
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

func (m *C2S) GetTime() int64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

type S2C struct {
	Time                 *int64   `protobuf:"varint,1,req,name=time" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_KeepAlive_8ce838656139b308, []int{1}
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

func (m *S2C) GetTime() int64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
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
	return fileDescriptor_KeepAlive_8ce838656139b308, []int{2}
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
	return fileDescriptor_KeepAlive_8ce838656139b308, []int{3}
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
	proto.RegisterType((*C2S)(nil), "KeepAlive.C2S")
	proto.RegisterType((*S2C)(nil), "KeepAlive.S2C")
	proto.RegisterType((*Request)(nil), "KeepAlive.Request")
	proto.RegisterType((*Response)(nil), "KeepAlive.Response")
}

func init() {
	proto.RegisterFile("KeepAlive/KeepAlive.proto", fileDescriptor_KeepAlive_8ce838656139b308)
}

var fileDescriptor_KeepAlive_8ce838656139b308 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8d, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0xc9, 0xa6, 0x6b, 0xdd, 0x11, 0x3c, 0x44, 0x90, 0xac, 0x07, 0x09, 0x7b, 0x2a, 0x88,
	0x6d, 0x09, 0x9e, 0xbc, 0x49, 0x8e, 0xe2, 0x65, 0xe2, 0x3f, 0xa8, 0x83, 0x14, 0x6c, 0x13, 0x93,
	0x28, 0x78, 0xf1, 0xb7, 0x4b, 0x5a, 0x5b, 0x41, 0xf6, 0x34, 0xef, 0xbd, 0x0f, 0xe6, 0x83, 0xfd,
	0x23, 0x91, 0x7f, 0x78, 0xeb, 0x3f, 0xa9, 0x59, 0x53, 0xed, 0x83, 0x4b, 0x4e, 0xec, 0xd6, 0xe1,
	0xea, 0xc2, 0xb8, 0x61, 0x70, 0x63, 0x33, 0x9f, 0x99, 0x1f, 0xf6, 0xc0, 0x8d, 0xb6, 0x42, 0x40,
	0x91, 0xfa, 0x81, 0x24, 0x53, 0x9b, 0x8a, 0xe3, 0x94, 0x33, 0xb2, 0xda, 0x1c, 0x45, 0x37, 0x50,
	0x22, 0xbd, 0x7f, 0x50, 0x4c, 0x42, 0x01, 0xef, 0x74, 0x9c, 0xe8, 0x99, 0x3e, 0xaf, 0xff, 0xfc,
	0x46, 0x5b, 0xcc, 0xe8, 0xf0, 0x0d, 0xa7, 0x48, 0xd1, 0xbb, 0x31, 0x92, 0xb8, 0x86, 0x32, 0x50,
	0x7a, 0xfe, 0xf2, 0xf3, 0xbf, 0xed, 0x7d, 0x71, 0x7b, 0xd7, 0xb6, 0xb8, 0x8c, 0xe2, 0x12, 0x4e,
	0x02, 0xa5, 0xa7, 0xf8, 0x2a, 0x37, 0x8a, 0x55, 0x3b, 0xfc, 0x6d, 0x42, 0x42, 0x49, 0x21, 0x18,
	0xf7, 0x42, 0x92, 0x2b, 0x56, 0x6d, 0x71, 0xa9, 0xd9, 0x1f, 0x75, 0x27, 0x0b, 0xc5, 0xfe, 0xf9,
	0xad, 0x36, 0x98, 0xd1, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa7, 0xc7, 0x1c, 0x98, 0x1e, 0x01,
	0x00, 0x00,
}
