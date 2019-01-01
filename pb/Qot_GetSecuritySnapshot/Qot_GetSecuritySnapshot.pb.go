// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Qot_GetSecuritySnapshot/Qot_GetSecuritySnapshot.proto

package Qot_GetSecuritySnapshot

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
	SecurityList         []*Qot_Common.Security `protobuf:"bytes,1,rep,name=securityList" json:"securityList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *C2S) Reset()         { *m = C2S{} }
func (m *C2S) String() string { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()    {}
func (*C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetSecuritySnapshot_c01df103305b8087, []int{0}
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

func (m *C2S) GetSecurityList() []*Qot_Common.Security {
	if m != nil {
		return m.SecurityList
	}
	return nil
}

// 正股类型额外数据
type EquitySnapshotExData struct {
	IssuedShares         *int64   `protobuf:"varint,1,req,name=issuedShares" json:"issuedShares,omitempty"`
	IssuedMarketVal      *float64 `protobuf:"fixed64,2,req,name=issuedMarketVal" json:"issuedMarketVal,omitempty"`
	NetAsset             *float64 `protobuf:"fixed64,3,req,name=netAsset" json:"netAsset,omitempty"`
	NetProfit            *float64 `protobuf:"fixed64,4,req,name=netProfit" json:"netProfit,omitempty"`
	EarningsPershare     *float64 `protobuf:"fixed64,5,req,name=earningsPershare" json:"earningsPershare,omitempty"`
	OutstandingShares    *int64   `protobuf:"varint,6,req,name=outstandingShares" json:"outstandingShares,omitempty"`
	OutstandingMarketVal *float64 `protobuf:"fixed64,7,req,name=outstandingMarketVal" json:"outstandingMarketVal,omitempty"`
	NetAssetPershare     *float64 `protobuf:"fixed64,8,req,name=netAssetPershare" json:"netAssetPershare,omitempty"`
	EyRate               *float64 `protobuf:"fixed64,9,req,name=eyRate" json:"eyRate,omitempty"`
	PeRate               *float64 `protobuf:"fixed64,10,req,name=peRate" json:"peRate,omitempty"`
	PbRate               *float64 `protobuf:"fixed64,11,req,name=pbRate" json:"pbRate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EquitySnapshotExData) Reset()         { *m = EquitySnapshotExData{} }
func (m *EquitySnapshotExData) String() string { return proto.CompactTextString(m) }
func (*EquitySnapshotExData) ProtoMessage()    {}
func (*EquitySnapshotExData) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetSecuritySnapshot_c01df103305b8087, []int{1}
}
func (m *EquitySnapshotExData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EquitySnapshotExData.Unmarshal(m, b)
}
func (m *EquitySnapshotExData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EquitySnapshotExData.Marshal(b, m, deterministic)
}
func (dst *EquitySnapshotExData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EquitySnapshotExData.Merge(dst, src)
}
func (m *EquitySnapshotExData) XXX_Size() int {
	return xxx_messageInfo_EquitySnapshotExData.Size(m)
}
func (m *EquitySnapshotExData) XXX_DiscardUnknown() {
	xxx_messageInfo_EquitySnapshotExData.DiscardUnknown(m)
}

var xxx_messageInfo_EquitySnapshotExData proto.InternalMessageInfo

func (m *EquitySnapshotExData) GetIssuedShares() int64 {
	if m != nil && m.IssuedShares != nil {
		return *m.IssuedShares
	}
	return 0
}

func (m *EquitySnapshotExData) GetIssuedMarketVal() float64 {
	if m != nil && m.IssuedMarketVal != nil {
		return *m.IssuedMarketVal
	}
	return 0
}

func (m *EquitySnapshotExData) GetNetAsset() float64 {
	if m != nil && m.NetAsset != nil {
		return *m.NetAsset
	}
	return 0
}

func (m *EquitySnapshotExData) GetNetProfit() float64 {
	if m != nil && m.NetProfit != nil {
		return *m.NetProfit
	}
	return 0
}

func (m *EquitySnapshotExData) GetEarningsPershare() float64 {
	if m != nil && m.EarningsPershare != nil {
		return *m.EarningsPershare
	}
	return 0
}

func (m *EquitySnapshotExData) GetOutstandingShares() int64 {
	if m != nil && m.OutstandingShares != nil {
		return *m.OutstandingShares
	}
	return 0
}

func (m *EquitySnapshotExData) GetOutstandingMarketVal() float64 {
	if m != nil && m.OutstandingMarketVal != nil {
		return *m.OutstandingMarketVal
	}
	return 0
}

func (m *EquitySnapshotExData) GetNetAssetPershare() float64 {
	if m != nil && m.NetAssetPershare != nil {
		return *m.NetAssetPershare
	}
	return 0
}

func (m *EquitySnapshotExData) GetEyRate() float64 {
	if m != nil && m.EyRate != nil {
		return *m.EyRate
	}
	return 0
}

func (m *EquitySnapshotExData) GetPeRate() float64 {
	if m != nil && m.PeRate != nil {
		return *m.PeRate
	}
	return 0
}

func (m *EquitySnapshotExData) GetPbRate() float64 {
	if m != nil && m.PbRate != nil {
		return *m.PbRate
	}
	return 0
}

// 涡轮类型额外数据
type WarrantSnapshotExData struct {
	ConversionRate       *float64             `protobuf:"fixed64,1,req,name=conversionRate" json:"conversionRate,omitempty"`
	WarrantType          *int32               `protobuf:"varint,2,req,name=warrantType" json:"warrantType,omitempty"`
	StrikePrice          *float64             `protobuf:"fixed64,3,req,name=strikePrice" json:"strikePrice,omitempty"`
	MaturityTime         *string              `protobuf:"bytes,4,req,name=maturityTime" json:"maturityTime,omitempty"`
	EndTradeTime         *string              `protobuf:"bytes,5,req,name=endTradeTime" json:"endTradeTime,omitempty"`
	Owner                *Qot_Common.Security `protobuf:"bytes,6,req,name=owner" json:"owner,omitempty"`
	RecoveryPrice        *float64             `protobuf:"fixed64,7,req,name=recoveryPrice" json:"recoveryPrice,omitempty"`
	StreetVolumn         *int64               `protobuf:"varint,8,req,name=streetVolumn" json:"streetVolumn,omitempty"`
	IssueVolumn          *int64               `protobuf:"varint,9,req,name=issueVolumn" json:"issueVolumn,omitempty"`
	StreetRate           *float64             `protobuf:"fixed64,10,req,name=streetRate" json:"streetRate,omitempty"`
	Delta                *float64             `protobuf:"fixed64,11,req,name=delta" json:"delta,omitempty"`
	ImpliedVolatility    *float64             `protobuf:"fixed64,12,req,name=impliedVolatility" json:"impliedVolatility,omitempty"`
	Premium              *float64             `protobuf:"fixed64,13,req,name=premium" json:"premium,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *WarrantSnapshotExData) Reset()         { *m = WarrantSnapshotExData{} }
func (m *WarrantSnapshotExData) String() string { return proto.CompactTextString(m) }
func (*WarrantSnapshotExData) ProtoMessage()    {}
func (*WarrantSnapshotExData) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetSecuritySnapshot_c01df103305b8087, []int{2}
}
func (m *WarrantSnapshotExData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WarrantSnapshotExData.Unmarshal(m, b)
}
func (m *WarrantSnapshotExData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WarrantSnapshotExData.Marshal(b, m, deterministic)
}
func (dst *WarrantSnapshotExData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WarrantSnapshotExData.Merge(dst, src)
}
func (m *WarrantSnapshotExData) XXX_Size() int {
	return xxx_messageInfo_WarrantSnapshotExData.Size(m)
}
func (m *WarrantSnapshotExData) XXX_DiscardUnknown() {
	xxx_messageInfo_WarrantSnapshotExData.DiscardUnknown(m)
}

var xxx_messageInfo_WarrantSnapshotExData proto.InternalMessageInfo

func (m *WarrantSnapshotExData) GetConversionRate() float64 {
	if m != nil && m.ConversionRate != nil {
		return *m.ConversionRate
	}
	return 0
}

func (m *WarrantSnapshotExData) GetWarrantType() int32 {
	if m != nil && m.WarrantType != nil {
		return *m.WarrantType
	}
	return 0
}

func (m *WarrantSnapshotExData) GetStrikePrice() float64 {
	if m != nil && m.StrikePrice != nil {
		return *m.StrikePrice
	}
	return 0
}

func (m *WarrantSnapshotExData) GetMaturityTime() string {
	if m != nil && m.MaturityTime != nil {
		return *m.MaturityTime
	}
	return ""
}

func (m *WarrantSnapshotExData) GetEndTradeTime() string {
	if m != nil && m.EndTradeTime != nil {
		return *m.EndTradeTime
	}
	return ""
}

func (m *WarrantSnapshotExData) GetOwner() *Qot_Common.Security {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *WarrantSnapshotExData) GetRecoveryPrice() float64 {
	if m != nil && m.RecoveryPrice != nil {
		return *m.RecoveryPrice
	}
	return 0
}

func (m *WarrantSnapshotExData) GetStreetVolumn() int64 {
	if m != nil && m.StreetVolumn != nil {
		return *m.StreetVolumn
	}
	return 0
}

func (m *WarrantSnapshotExData) GetIssueVolumn() int64 {
	if m != nil && m.IssueVolumn != nil {
		return *m.IssueVolumn
	}
	return 0
}

func (m *WarrantSnapshotExData) GetStreetRate() float64 {
	if m != nil && m.StreetRate != nil {
		return *m.StreetRate
	}
	return 0
}

func (m *WarrantSnapshotExData) GetDelta() float64 {
	if m != nil && m.Delta != nil {
		return *m.Delta
	}
	return 0
}

func (m *WarrantSnapshotExData) GetImpliedVolatility() float64 {
	if m != nil && m.ImpliedVolatility != nil {
		return *m.ImpliedVolatility
	}
	return 0
}

func (m *WarrantSnapshotExData) GetPremium() float64 {
	if m != nil && m.Premium != nil {
		return *m.Premium
	}
	return 0
}

// 基本快照数据
type SnapshotBasicData struct {
	Security             *Qot_Common.Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`
	Type                 *int32               `protobuf:"varint,2,req,name=type" json:"type,omitempty"`
	IsSuspend            *bool                `protobuf:"varint,3,req,name=isSuspend" json:"isSuspend,omitempty"`
	ListTime             *string              `protobuf:"bytes,4,req,name=listTime" json:"listTime,omitempty"`
	LotSize              *int32               `protobuf:"varint,5,req,name=lotSize" json:"lotSize,omitempty"`
	PriceSpread          *float64             `protobuf:"fixed64,6,req,name=priceSpread" json:"priceSpread,omitempty"`
	UpdateTime           *string              `protobuf:"bytes,7,req,name=updateTime" json:"updateTime,omitempty"`
	HighPrice            *float64             `protobuf:"fixed64,8,req,name=highPrice" json:"highPrice,omitempty"`
	OpenPrice            *float64             `protobuf:"fixed64,9,req,name=openPrice" json:"openPrice,omitempty"`
	LowPrice             *float64             `protobuf:"fixed64,10,req,name=lowPrice" json:"lowPrice,omitempty"`
	LastClosePrice       *float64             `protobuf:"fixed64,11,req,name=lastClosePrice" json:"lastClosePrice,omitempty"`
	CurPrice             *float64             `protobuf:"fixed64,12,req,name=curPrice" json:"curPrice,omitempty"`
	Volume               *int64               `protobuf:"varint,13,req,name=volume" json:"volume,omitempty"`
	Turnover             *float64             `protobuf:"fixed64,14,req,name=turnover" json:"turnover,omitempty"`
	TurnoverRate         *float64             `protobuf:"fixed64,15,req,name=turnoverRate" json:"turnoverRate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SnapshotBasicData) Reset()         { *m = SnapshotBasicData{} }
func (m *SnapshotBasicData) String() string { return proto.CompactTextString(m) }
func (*SnapshotBasicData) ProtoMessage()    {}
func (*SnapshotBasicData) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetSecuritySnapshot_c01df103305b8087, []int{3}
}
func (m *SnapshotBasicData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnapshotBasicData.Unmarshal(m, b)
}
func (m *SnapshotBasicData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnapshotBasicData.Marshal(b, m, deterministic)
}
func (dst *SnapshotBasicData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotBasicData.Merge(dst, src)
}
func (m *SnapshotBasicData) XXX_Size() int {
	return xxx_messageInfo_SnapshotBasicData.Size(m)
}
func (m *SnapshotBasicData) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotBasicData.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotBasicData proto.InternalMessageInfo

func (m *SnapshotBasicData) GetSecurity() *Qot_Common.Security {
	if m != nil {
		return m.Security
	}
	return nil
}

func (m *SnapshotBasicData) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func (m *SnapshotBasicData) GetIsSuspend() bool {
	if m != nil && m.IsSuspend != nil {
		return *m.IsSuspend
	}
	return false
}

func (m *SnapshotBasicData) GetListTime() string {
	if m != nil && m.ListTime != nil {
		return *m.ListTime
	}
	return ""
}

func (m *SnapshotBasicData) GetLotSize() int32 {
	if m != nil && m.LotSize != nil {
		return *m.LotSize
	}
	return 0
}

func (m *SnapshotBasicData) GetPriceSpread() float64 {
	if m != nil && m.PriceSpread != nil {
		return *m.PriceSpread
	}
	return 0
}

func (m *SnapshotBasicData) GetUpdateTime() string {
	if m != nil && m.UpdateTime != nil {
		return *m.UpdateTime
	}
	return ""
}

func (m *SnapshotBasicData) GetHighPrice() float64 {
	if m != nil && m.HighPrice != nil {
		return *m.HighPrice
	}
	return 0
}

func (m *SnapshotBasicData) GetOpenPrice() float64 {
	if m != nil && m.OpenPrice != nil {
		return *m.OpenPrice
	}
	return 0
}

func (m *SnapshotBasicData) GetLowPrice() float64 {
	if m != nil && m.LowPrice != nil {
		return *m.LowPrice
	}
	return 0
}

func (m *SnapshotBasicData) GetLastClosePrice() float64 {
	if m != nil && m.LastClosePrice != nil {
		return *m.LastClosePrice
	}
	return 0
}

func (m *SnapshotBasicData) GetCurPrice() float64 {
	if m != nil && m.CurPrice != nil {
		return *m.CurPrice
	}
	return 0
}

func (m *SnapshotBasicData) GetVolume() int64 {
	if m != nil && m.Volume != nil {
		return *m.Volume
	}
	return 0
}

func (m *SnapshotBasicData) GetTurnover() float64 {
	if m != nil && m.Turnover != nil {
		return *m.Turnover
	}
	return 0
}

func (m *SnapshotBasicData) GetTurnoverRate() float64 {
	if m != nil && m.TurnoverRate != nil {
		return *m.TurnoverRate
	}
	return 0
}

type Snapshot struct {
	Basic                *SnapshotBasicData     `protobuf:"bytes,1,req,name=basic" json:"basic,omitempty"`
	EquityExData         *EquitySnapshotExData  `protobuf:"bytes,2,opt,name=equityExData" json:"equityExData,omitempty"`
	WarrantExData        *WarrantSnapshotExData `protobuf:"bytes,3,opt,name=warrantExData" json:"warrantExData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Snapshot) Reset()         { *m = Snapshot{} }
func (m *Snapshot) String() string { return proto.CompactTextString(m) }
func (*Snapshot) ProtoMessage()    {}
func (*Snapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetSecuritySnapshot_c01df103305b8087, []int{4}
}
func (m *Snapshot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Snapshot.Unmarshal(m, b)
}
func (m *Snapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Snapshot.Marshal(b, m, deterministic)
}
func (dst *Snapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Snapshot.Merge(dst, src)
}
func (m *Snapshot) XXX_Size() int {
	return xxx_messageInfo_Snapshot.Size(m)
}
func (m *Snapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_Snapshot.DiscardUnknown(m)
}

var xxx_messageInfo_Snapshot proto.InternalMessageInfo

func (m *Snapshot) GetBasic() *SnapshotBasicData {
	if m != nil {
		return m.Basic
	}
	return nil
}

func (m *Snapshot) GetEquityExData() *EquitySnapshotExData {
	if m != nil {
		return m.EquityExData
	}
	return nil
}

func (m *Snapshot) GetWarrantExData() *WarrantSnapshotExData {
	if m != nil {
		return m.WarrantExData
	}
	return nil
}

type S2C struct {
	SnapshotList         []*Snapshot `protobuf:"bytes,1,rep,name=snapshotList" json:"snapshotList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_Qot_GetSecuritySnapshot_c01df103305b8087, []int{5}
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

func (m *S2C) GetSnapshotList() []*Snapshot {
	if m != nil {
		return m.SnapshotList
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
	return fileDescriptor_Qot_GetSecuritySnapshot_c01df103305b8087, []int{6}
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
	return fileDescriptor_Qot_GetSecuritySnapshot_c01df103305b8087, []int{7}
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
	proto.RegisterType((*C2S)(nil), "Qot_GetSecuritySnapshot.C2S")
	proto.RegisterType((*EquitySnapshotExData)(nil), "Qot_GetSecuritySnapshot.EquitySnapshotExData")
	proto.RegisterType((*WarrantSnapshotExData)(nil), "Qot_GetSecuritySnapshot.WarrantSnapshotExData")
	proto.RegisterType((*SnapshotBasicData)(nil), "Qot_GetSecuritySnapshot.SnapshotBasicData")
	proto.RegisterType((*Snapshot)(nil), "Qot_GetSecuritySnapshot.Snapshot")
	proto.RegisterType((*S2C)(nil), "Qot_GetSecuritySnapshot.S2C")
	proto.RegisterType((*Request)(nil), "Qot_GetSecuritySnapshot.Request")
	proto.RegisterType((*Response)(nil), "Qot_GetSecuritySnapshot.Response")
}

func init() {
	proto.RegisterFile("Qot_GetSecuritySnapshot/Qot_GetSecuritySnapshot.proto", fileDescriptor_Qot_GetSecuritySnapshot_c01df103305b8087)
}

var fileDescriptor_Qot_GetSecuritySnapshot_c01df103305b8087 = []byte{
	// 868 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0x5f, 0x6f, 0xe3, 0x44,
	0x10, 0x57, 0xea, 0xe6, 0x92, 0x4c, 0xd3, 0x3b, 0x6e, 0x29, 0x60, 0x95, 0xd3, 0x29, 0x44, 0x08,
	0x45, 0x15, 0x97, 0xab, 0x2c, 0x90, 0x80, 0x17, 0xfe, 0x84, 0x8a, 0x97, 0x3b, 0xa9, 0xb7, 0xae,
	0xca, 0x23, 0xda, 0x8b, 0x87, 0x76, 0x75, 0xce, 0xae, 0xbb, 0xbb, 0x6e, 0x09, 0x9f, 0x81, 0x0f,
	0xc1, 0x57, 0x43, 0x7c, 0x06, 0xde, 0xd1, 0xce, 0xda, 0x89, 0x9d, 0xd6, 0xba, 0xa7, 0xf8, 0xf7,
	0x9b, 0xdf, 0x8e, 0xc7, 0xf3, 0x9b, 0xd9, 0xc0, 0xd7, 0x6f, 0xb4, 0xfb, 0xed, 0x17, 0x74, 0x29,
	0x2e, 0x4b, 0x23, 0xdd, 0x3a, 0x55, 0xa2, 0xb0, 0xd7, 0xda, 0xbd, 0xec, 0xe0, 0xe7, 0x85, 0xd1,
	0x4e, 0xb3, 0x4f, 0x3a, 0xc2, 0xc7, 0x1f, 0x2e, 0xf4, 0x6a, 0xa5, 0xd5, 0xcb, 0xf0, 0x13, 0xd4,
	0xc7, 0x9f, 0x7a, 0x75, 0x15, 0xd8, 0x3e, 0x86, 0xe0, 0xf4, 0x7b, 0x88, 0x16, 0x49, 0xca, 0xbe,
	0x81, 0xb1, 0xad, 0x92, 0xbd, 0x92, 0xd6, 0xc5, 0xbd, 0x49, 0x34, 0x3b, 0x48, 0x8e, 0xe6, 0x0d,
	0x7d, 0xfd, 0x32, 0xde, 0x52, 0x4e, 0xff, 0x8e, 0xe0, 0xe8, 0xec, 0xa6, 0x6c, 0x54, 0x71, 0xf6,
	0xc7, 0xcf, 0xc2, 0x09, 0x36, 0x85, 0xb1, 0xb4, 0xb6, 0xc4, 0x2c, 0xbd, 0x16, 0x06, 0x6d, 0xdc,
	0x9b, 0xec, 0xcd, 0x22, 0xde, 0xe2, 0xd8, 0x0c, 0x9e, 0x04, 0xfc, 0x5a, 0x98, 0x77, 0xe8, 0x2e,
	0x45, 0x1e, 0xef, 0x4d, 0xf6, 0x66, 0x3d, 0xbe, 0x4b, 0xb3, 0x63, 0x18, 0x2a, 0x74, 0x3f, 0x5a,
	0x8b, 0x2e, 0x8e, 0x48, 0xb2, 0xc1, 0xec, 0x19, 0x8c, 0x14, 0xba, 0x73, 0xa3, 0x7f, 0x97, 0x2e,
	0xde, 0xa7, 0xe0, 0x96, 0x60, 0x27, 0xf0, 0x01, 0x0a, 0xa3, 0xa4, 0xba, 0xb2, 0xe7, 0x68, 0xac,
	0x7f, 0x71, 0xdc, 0x27, 0xd1, 0x3d, 0x9e, 0x7d, 0x09, 0x4f, 0x75, 0xe9, 0xac, 0x13, 0x2a, 0x93,
	0xea, 0xaa, 0x2a, 0xfc, 0x11, 0x15, 0x7e, 0x3f, 0xc0, 0x12, 0x38, 0x6a, 0x90, 0xdb, 0x4f, 0x18,
	0x50, 0xf6, 0x07, 0x63, 0xbe, 0x9a, 0xba, 0xee, 0x4d, 0x35, 0xc3, 0x50, 0xcd, 0x2e, 0xcf, 0x3e,
	0x86, 0x47, 0xb8, 0xe6, 0xc2, 0x61, 0x3c, 0x22, 0x45, 0x85, 0x3c, 0x5f, 0x20, 0xf1, 0x10, 0xf8,
	0x80, 0x88, 0x7f, 0x4b, 0xfc, 0x41, 0xc5, 0x13, 0x9a, 0xfe, 0x13, 0xc1, 0x47, 0xbf, 0x0a, 0x63,
	0x84, 0x72, 0x3b, 0x1e, 0x7d, 0x01, 0x8f, 0x97, 0x5a, 0xdd, 0xa2, 0xb1, 0x52, 0x2b, 0x3a, 0xd9,
	0xa3, 0x93, 0x3b, 0x2c, 0x9b, 0xc0, 0xc1, 0x5d, 0x48, 0x70, 0xb1, 0x2e, 0x90, 0x3c, 0xea, 0xf3,
	0x26, 0xe5, 0x15, 0xd6, 0x19, 0xf9, 0x0e, 0xcf, 0x8d, 0x5c, 0x62, 0x65, 0x51, 0x93, 0xf2, 0xf3,
	0xb0, 0x12, 0x8e, 0x06, 0xe7, 0x42, 0xae, 0x90, 0x8c, 0x1a, 0xf1, 0x16, 0xe7, 0x35, 0xa8, 0xb2,
	0x0b, 0x23, 0x32, 0x24, 0x4d, 0x3f, 0x68, 0x9a, 0x1c, 0x3b, 0x81, 0xbe, 0xbe, 0x53, 0x68, 0xc8,
	0x97, 0xae, 0x19, 0x0d, 0x12, 0xf6, 0x39, 0x1c, 0x1a, 0x5c, 0xea, 0x5b, 0x34, 0xeb, 0x50, 0x57,
	0xb0, 0xa6, 0x4d, 0xfa, 0xb7, 0x5a, 0x67, 0x10, 0xdd, 0xa5, 0xce, 0xcb, 0x95, 0x22, 0x3f, 0x22,
	0xde, 0xe2, 0xfc, 0xf7, 0xd1, 0x48, 0x56, 0x92, 0x11, 0x49, 0x9a, 0x14, 0x7b, 0x0e, 0x10, 0x4e,
	0x34, 0x9c, 0x69, 0x30, 0xec, 0x08, 0xfa, 0x19, 0xe6, 0x4e, 0x54, 0xe6, 0x04, 0xe0, 0x27, 0x4e,
	0xae, 0x8a, 0x5c, 0x62, 0x76, 0xa9, 0x73, 0xe1, 0x64, 0x2e, 0xdd, 0x3a, 0x1e, 0x93, 0xe2, 0x7e,
	0x80, 0xc5, 0x30, 0x28, 0x0c, 0xae, 0x64, 0xb9, 0x8a, 0x0f, 0x49, 0x53, 0xc3, 0xe9, 0xbf, 0x11,
	0x3c, 0xad, 0xcd, 0xfd, 0x49, 0x58, 0xb9, 0x24, 0x7f, 0x4f, 0x61, 0x58, 0x2f, 0x2b, 0x39, 0xdb,
	0xd5, 0xae, 0x8d, 0x8a, 0x31, 0xd8, 0x77, 0x5b, 0x8b, 0xe9, 0xd9, 0xef, 0x97, 0xb4, 0x69, 0x69,
	0x0b, 0x54, 0x19, 0x39, 0x3b, 0xe4, 0x5b, 0xc2, 0x6f, 0x66, 0x2e, 0xad, 0x6b, 0x78, 0xba, 0xc1,
	0xbe, 0xde, 0x5c, 0xbb, 0x54, 0xfe, 0x19, 0xac, 0xec, 0xf3, 0x1a, 0xfa, 0x7e, 0x16, 0xbe, 0xf9,
	0x69, 0x61, 0x50, 0x64, 0xe4, 0x65, 0x8f, 0x37, 0x29, 0xdf, 0xcf, 0xb2, 0xc8, 0x84, 0x0b, 0x93,
	0x30, 0xa0, 0xcc, 0x0d, 0xc6, 0x57, 0x75, 0x2d, 0xaf, 0xae, 0x83, 0xaf, 0x61, 0x85, 0xb6, 0x84,
	0x8f, 0xea, 0x02, 0x55, 0x88, 0x86, 0xf5, 0xd9, 0x12, 0x54, 0xb3, 0xbe, 0x0b, 0xc1, 0xe0, 0xd4,
	0x06, 0xfb, 0x9d, 0xc8, 0x85, 0x75, 0x8b, 0x5c, 0xdb, 0x6a, 0x98, 0x83, 0x61, 0x3b, 0xac, 0xcf,
	0xb1, 0x2c, 0x4d, 0x50, 0x04, 0xc3, 0x36, 0xd8, 0x6f, 0xe2, 0xad, 0x9f, 0x0a, 0x24, 0x9b, 0x22,
	0x5e, 0x21, 0x7f, 0xc6, 0x95, 0x46, 0xf9, 0xd9, 0x8b, 0x1f, 0x87, 0x33, 0x35, 0xf6, 0x53, 0x58,
	0x3f, 0xd3, 0x04, 0x3d, 0xa1, 0x78, 0x8b, 0x9b, 0xfe, 0xd7, 0x83, 0x61, 0xed, 0x32, 0xfb, 0x01,
	0xfa, 0x6f, 0xbd, 0xd3, 0x95, 0xb3, 0x27, 0xf3, 0xae, 0x3f, 0x8d, 0x7b, 0x73, 0xc1, 0xc3, 0x41,
	0xf6, 0x06, 0xc6, 0x48, 0x57, 0x77, 0xb8, 0x0e, 0xe2, 0xbd, 0x49, 0x6f, 0x76, 0x90, 0xbc, 0xe8,
	0x4c, 0xf4, 0xd0, 0x3d, 0xcf, 0x5b, 0x29, 0xd8, 0x05, 0x1c, 0x56, 0xd7, 0x42, 0x95, 0x33, 0xa2,
	0x9c, 0xf3, 0xce, 0x9c, 0x0f, 0x5e, 0x4c, 0xbc, 0x9d, 0x64, 0xfa, 0x0a, 0xa2, 0x34, 0x59, 0xb0,
	0x33, 0x18, 0xdb, 0x4a, 0xd7, 0xf8, 0x97, 0xfa, 0xec, 0xbd, 0x1f, 0xce, 0x5b, 0xc7, 0xa6, 0xdf,
	0xc2, 0x80, 0xe3, 0x4d, 0x89, 0xd6, 0xb1, 0x39, 0x44, 0xcb, 0xc4, 0x56, 0x1d, 0x7c, 0xd6, 0x99,
	0x68, 0x91, 0xa4, 0xdc, 0x0b, 0xa7, 0x7f, 0xf5, 0x60, 0xc8, 0xd1, 0x16, 0x5a, 0x59, 0x64, 0xcf,
	0x61, 0x60, 0x30, 0xdc, 0x88, 0x3e, 0x41, 0xff, 0xbb, 0xfd, 0x17, 0x5f, 0x9d, 0x9e, 0xf2, 0x9a,
	0xf4, 0x53, 0x60, 0xd0, 0xbd, 0xb6, 0x57, 0xd4, 0xd8, 0x11, 0xaf, 0x90, 0xdf, 0x0a, 0x34, 0x66,
	0xa1, 0x33, 0xa4, 0xee, 0xf4, 0x79, 0x0d, 0x7d, 0x39, 0x36, 0x59, 0xc6, 0xfb, 0xd4, 0xb3, 0xee,
	0x72, 0xd2, 0x64, 0xc1, 0xbd, 0xf0, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x56, 0x3e, 0x27, 0xed,
	0x40, 0x08, 0x00, 0x00,
}
