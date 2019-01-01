// Code generated by protoc-gen-go. DO NOT EDIT.
// source: GetGlobalState/GetGlobalState.proto

package GetGlobalState

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "futugg/pb/Common"
import _ "futugg/pb/Qot_Common"

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
	UserID               *uint64  `protobuf:"varint,1,req,name=userID" json:"userID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2S) Reset()         { *m = C2S{} }
func (m *C2S) String() string { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()    {}
func (*C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_GetGlobalState_67e8c0cdea9d35bf, []int{0}
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

func (m *C2S) GetUserID() uint64 {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return 0
}

type S2C struct {
	MarketHK             *int32   `protobuf:"varint,1,req,name=marketHK" json:"marketHK,omitempty"`
	MarketUS             *int32   `protobuf:"varint,2,req,name=marketUS" json:"marketUS,omitempty"`
	MarketSH             *int32   `protobuf:"varint,3,req,name=marketSH" json:"marketSH,omitempty"`
	MarketSZ             *int32   `protobuf:"varint,4,req,name=marketSZ" json:"marketSZ,omitempty"`
	MarketHKFuture       *int32   `protobuf:"varint,5,req,name=marketHKFuture" json:"marketHKFuture,omitempty"`
	QotLogined           *bool    `protobuf:"varint,6,req,name=qotLogined" json:"qotLogined,omitempty"`
	TrdLogined           *bool    `protobuf:"varint,7,req,name=trdLogined" json:"trdLogined,omitempty"`
	ServerVer            *int32   `protobuf:"varint,8,req,name=serverVer" json:"serverVer,omitempty"`
	ServerBuildNo        *int32   `protobuf:"varint,9,req,name=serverBuildNo" json:"serverBuildNo,omitempty"`
	Time                 *int64   `protobuf:"varint,10,req,name=time" json:"time,omitempty"`
	LocalTime            *float64 `protobuf:"fixed64,11,opt,name=localTime" json:"localTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_GetGlobalState_67e8c0cdea9d35bf, []int{1}
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

func (m *S2C) GetMarketHK() int32 {
	if m != nil && m.MarketHK != nil {
		return *m.MarketHK
	}
	return 0
}

func (m *S2C) GetMarketUS() int32 {
	if m != nil && m.MarketUS != nil {
		return *m.MarketUS
	}
	return 0
}

func (m *S2C) GetMarketSH() int32 {
	if m != nil && m.MarketSH != nil {
		return *m.MarketSH
	}
	return 0
}

func (m *S2C) GetMarketSZ() int32 {
	if m != nil && m.MarketSZ != nil {
		return *m.MarketSZ
	}
	return 0
}

func (m *S2C) GetMarketHKFuture() int32 {
	if m != nil && m.MarketHKFuture != nil {
		return *m.MarketHKFuture
	}
	return 0
}

func (m *S2C) GetQotLogined() bool {
	if m != nil && m.QotLogined != nil {
		return *m.QotLogined
	}
	return false
}

func (m *S2C) GetTrdLogined() bool {
	if m != nil && m.TrdLogined != nil {
		return *m.TrdLogined
	}
	return false
}

func (m *S2C) GetServerVer() int32 {
	if m != nil && m.ServerVer != nil {
		return *m.ServerVer
	}
	return 0
}

func (m *S2C) GetServerBuildNo() int32 {
	if m != nil && m.ServerBuildNo != nil {
		return *m.ServerBuildNo
	}
	return 0
}

func (m *S2C) GetTime() int64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *S2C) GetLocalTime() float64 {
	if m != nil && m.LocalTime != nil {
		return *m.LocalTime
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
	return fileDescriptor_GetGlobalState_67e8c0cdea9d35bf, []int{2}
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
	return fileDescriptor_GetGlobalState_67e8c0cdea9d35bf, []int{3}
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
	proto.RegisterType((*C2S)(nil), "GetGlobalState.C2S")
	proto.RegisterType((*S2C)(nil), "GetGlobalState.S2C")
	proto.RegisterType((*Request)(nil), "GetGlobalState.Request")
	proto.RegisterType((*Response)(nil), "GetGlobalState.Response")
}

func init() {
	proto.RegisterFile("GetGlobalState/GetGlobalState.proto", fileDescriptor_GetGlobalState_67e8c0cdea9d35bf)
}

var fileDescriptor_GetGlobalState_67e8c0cdea9d35bf = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4b, 0xcb, 0xda, 0x40,
	0x14, 0x86, 0xc9, 0x45, 0x13, 0x8f, 0xd4, 0xc5, 0x08, 0x65, 0xb0, 0x17, 0x42, 0x7a, 0x21, 0x9b,
	0xaa, 0x84, 0xae, 0xba, 0x6c, 0x4a, 0xb5, 0xf4, 0x02, 0x9d, 0xd1, 0x2e, 0xdc, 0x94, 0xd4, 0x1c,
	0x44, 0x9a, 0x38, 0x3a, 0x33, 0x29, 0x74, 0xdd, 0x9f, 0xd8, 0x3f, 0xf4, 0x31, 0x89, 0xf9, 0x92,
	0xc8, 0xb7, 0xca, 0xbc, 0xcf, 0xf3, 0x32, 0x39, 0xc9, 0x81, 0x17, 0x2b, 0xd4, 0xab, 0x5c, 0xfc,
	0x4a, 0x73, 0xae, 0x53, 0x8d, 0x8b, 0x7e, 0x9c, 0x9f, 0xa5, 0xd0, 0x82, 0x4c, 0xfa, 0x74, 0x36,
	0x4d, 0x44, 0x51, 0x88, 0xd3, 0xa2, 0x7e, 0xd4, 0xa5, 0xd9, 0x93, 0xef, 0x42, 0xff, 0xbc, 0x8a,
	0xf6, 0x58, 0xcb, 0xf0, 0x19, 0x38, 0x49, 0xcc, 0xc9, 0x63, 0x18, 0x96, 0x0a, 0xe5, 0xa7, 0x0f,
	0xd4, 0x0a, 0xec, 0xc8, 0x65, 0xd7, 0x14, 0xfe, 0xb7, 0xc1, 0xe1, 0x71, 0x42, 0x66, 0xe0, 0x17,
	0xa9, 0xfc, 0x8d, 0x7a, 0xfd, 0xb9, 0x6a, 0x0c, 0xd8, 0x7d, 0x6e, 0xdd, 0x96, 0x53, 0xbb, 0xeb,
	0xb6, 0xbc, 0x75, 0x7c, 0x4d, 0x9d, 0xae, 0xe3, 0xeb, 0x8e, 0xdb, 0x51, 0xb7, 0xe7, 0x76, 0xe4,
	0x35, 0x4c, 0x9a, 0xfb, 0x3f, 0x96, 0xba, 0x94, 0x48, 0x07, 0x55, 0xe3, 0x86, 0x92, 0xe7, 0x00,
	0x17, 0xa1, 0xbf, 0x88, 0xc3, 0xf1, 0x84, 0x19, 0x1d, 0x06, 0x76, 0xe4, 0xb3, 0x0e, 0x31, 0x5e,
	0xcb, 0xac, 0xf1, 0x5e, 0xed, 0x5b, 0x42, 0x9e, 0xc2, 0x48, 0xa1, 0xfc, 0x83, 0xf2, 0x07, 0x4a,
	0xea, 0x57, 0xaf, 0x68, 0x01, 0x79, 0x09, 0x8f, 0xea, 0xf0, 0xbe, 0x3c, 0xe6, 0xd9, 0x37, 0x41,
	0x47, 0x55, 0xa3, 0x0f, 0x09, 0x01, 0x57, 0x1f, 0x0b, 0xa4, 0x10, 0xd8, 0x91, 0xc3, 0xaa, 0xb3,
	0xb9, 0x37, 0x17, 0xfb, 0x34, 0xdf, 0x18, 0x31, 0x0e, 0xac, 0xc8, 0x62, 0x2d, 0x08, 0x97, 0xe0,
	0x31, 0xbc, 0x94, 0xa8, 0x34, 0x79, 0x05, 0xce, 0x3e, 0x56, 0xd5, 0x3f, 0x1d, 0xc7, 0xd3, 0xf9,
	0xcd, 0x96, 0x93, 0x98, 0x33, 0xe3, 0xc3, 0x7f, 0x16, 0xf8, 0x0c, 0xd5, 0x59, 0x9c, 0x94, 0xf9,
	0x68, 0x4f, 0xa2, 0xde, 0xfc, 0x3d, 0x63, 0xbd, 0x8b, 0x77, 0xee, 0x9b, 0xb7, 0xcb, 0x25, 0x6b,
	0xa0, 0x59, 0xa6, 0x44, 0xfd, 0x55, 0x1d, 0xa8, 0x1d, 0x58, 0xd1, 0x88, 0x5d, 0x13, 0xa1, 0xe0,
	0xa1, 0x94, 0x89, 0xc8, 0x90, 0x3a, 0x81, 0x15, 0x0d, 0x58, 0x13, 0xcd, 0x14, 0x2a, 0xde, 0x53,
	0x37, 0xb0, 0x1e, 0x9a, 0x82, 0xc7, 0x09, 0x33, 0xfe, 0x2e, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x58,
	0x5b, 0xcb, 0x94, 0x02, 0x00, 0x00,
}
