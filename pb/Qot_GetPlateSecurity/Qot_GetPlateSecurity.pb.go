// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Qot_GetPlateSecurity/Qot_GetPlateSecurity.proto

package Qot_GetPlateSecurity

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
	Plate                *Qot_Common.Security `protobuf:"bytes,1,req,name=plate" json:"plate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *C2S) Reset()         { *m = C2S{} }
func (m *C2S) String() string { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()    {}
func (*C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetPlateSecurity_bbd09fc8f21c192b, []int{0}
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

func (m *C2S) GetPlate() *Qot_Common.Security {
	if m != nil {
		return m.Plate
	}
	return nil
}

type S2C struct {
	StaticInfoList       []*Qot_Common.SecurityStaticInfo `protobuf:"bytes,1,rep,name=staticInfoList" json:"staticInfoList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetPlateSecurity_bbd09fc8f21c192b, []int{1}
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

func (m *S2C) GetStaticInfoList() []*Qot_Common.SecurityStaticInfo {
	if m != nil {
		return m.StaticInfoList
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
	return fileDescriptor_Qot_GetPlateSecurity_bbd09fc8f21c192b, []int{2}
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
	return fileDescriptor_Qot_GetPlateSecurity_bbd09fc8f21c192b, []int{3}
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
	proto.RegisterType((*C2S)(nil), "Qot_GetPlateSecurity.C2S")
	proto.RegisterType((*S2C)(nil), "Qot_GetPlateSecurity.S2C")
	proto.RegisterType((*Request)(nil), "Qot_GetPlateSecurity.Request")
	proto.RegisterType((*Response)(nil), "Qot_GetPlateSecurity.Response")
}

func init() {
	proto.RegisterFile("Qot_GetPlateSecurity/Qot_GetPlateSecurity.proto", fileDescriptor_Qot_GetPlateSecurity_bbd09fc8f21c192b)
}

var fileDescriptor_Qot_GetPlateSecurity_bbd09fc8f21c192b = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0x51, 0x4b, 0xfb, 0x30,
	0x14, 0xc5, 0xc9, 0xba, 0xfe, 0xfb, 0xf7, 0x0e, 0x7c, 0x88, 0x43, 0xe2, 0x84, 0x51, 0xfa, 0x54,
	0x14, 0xbb, 0x19, 0xc4, 0x07, 0x5f, 0x03, 0x8a, 0xe0, 0x40, 0x13, 0xdf, 0x65, 0xd4, 0xab, 0x14,
	0x5c, 0x53, 0x93, 0xbb, 0x87, 0x7d, 0x04, 0xbf, 0xb5, 0x94, 0xb4, 0x2a, 0x52, 0x9f, 0x92, 0x7b,
	0xcf, 0xef, 0x1e, 0xce, 0x81, 0xc5, 0x83, 0xa5, 0xa7, 0x1b, 0xa4, 0xfb, 0xb7, 0x35, 0xa1, 0xc1,
	0x72, 0xeb, 0x2a, 0xda, 0x0d, 0x2e, 0x8b, 0xc6, 0x59, 0xb2, 0x7c, 0x3a, 0xa4, 0xcd, 0x0e, 0x94,
	0xdd, 0x6c, 0x6c, 0xbd, 0x08, 0x4f, 0x40, 0x67, 0xc7, 0x2d, 0xda, 0x09, 0xdf, 0xdf, 0x20, 0x66,
	0xe7, 0x10, 0x29, 0x69, 0xf8, 0x09, 0xc4, 0x4d, 0xeb, 0x24, 0x58, 0x3a, 0xca, 0x27, 0x72, 0x5a,
	0xfc, 0x00, 0x7b, 0x77, 0x1d, 0x90, 0x6c, 0x05, 0x91, 0x91, 0x8a, 0x5f, 0xc3, 0xbe, 0xa7, 0x35,
	0x55, 0xe5, 0x6d, 0xfd, 0x62, 0xef, 0x2a, 0x4f, 0x82, 0xa5, 0x51, 0x3e, 0x91, 0xf3, 0xa1, 0x5b,
	0xf3, 0x45, 0xea, 0x5f, 0x57, 0xd9, 0x25, 0x24, 0x1a, 0xdf, 0xb7, 0xe8, 0x89, 0x9f, 0x42, 0x54,
	0x4a, 0xdf, 0x65, 0x38, 0x2a, 0x06, 0xeb, 0x2b, 0x69, 0x74, 0x4b, 0x65, 0x1f, 0x0c, 0xfe, 0x6b,
	0xf4, 0x8d, 0xad, 0x3d, 0xf2, 0x39, 0x24, 0x0e, 0xe9, 0x71, 0xd7, 0x84, 0x06, 0xf1, 0xd5, 0xf8,
	0xec, 0x62, 0xb9, 0xd4, 0xfd, 0x92, 0x1f, 0xc2, 0x3f, 0x87, 0xb4, 0xf2, 0xaf, 0x62, 0x94, 0xb2,
	0x7c, 0x4f, 0x77, 0x13, 0x17, 0x90, 0xa0, 0x73, 0xca, 0x3e, 0xa3, 0x88, 0x52, 0x96, 0xc7, 0xba,
	0x1f, 0xdb, 0x2c, 0x5e, 0x96, 0x62, 0x9c, 0xb2, 0xbf, 0xb3, 0x18, 0xa9, 0x74, 0x4b, 0x7d, 0x06,
	0x00, 0x00, 0xff, 0xff, 0xf0, 0x89, 0xf8, 0xeb, 0xbf, 0x01, 0x00, 0x00,
}
