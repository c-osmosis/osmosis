// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/gamm/v1beta1/pool.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/x/auth/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type PoolAsset struct {
	// Denormalized weight.
	Weight github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=weight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"weight" yaml:"denormalized_weight"`
	Token  types.Coin                             `protobuf:"bytes,2,opt,name=token,proto3" json:"token" yaml:"token"`
}

func (m *PoolAsset) Reset()         { *m = PoolAsset{} }
func (m *PoolAsset) String() string { return proto.CompactTextString(m) }
func (*PoolAsset) ProtoMessage()    {}
func (*PoolAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5ab9bc6d45f98ce, []int{0}
}
func (m *PoolAsset) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolAsset.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolAsset.Merge(m, src)
}
func (m *PoolAsset) XXX_Size() int {
	return m.Size()
}
func (m *PoolAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolAsset.DiscardUnknown(m)
}

var xxx_messageInfo_PoolAsset proto.InternalMessageInfo

func (m *PoolAsset) GetToken() types.Coin {
	if m != nil {
		return m.Token
	}
	return types.Coin{}
}

// PoolParams defined the parameters that will be managed by the pool governance
// in the future. This params are not managed by the chain governanace.
type PoolParams struct {
	// If the pool is locked, users can't join, exit or swap the pool.
	Lock    bool                                   `protobuf:"varint,1,opt,name=lock,proto3" json:"lock,omitempty"`
	SwapFee github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=swapFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"swapFee" yaml:"swap_fee"`
	ExitFee github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=exitFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"exitFee" yaml:"exit_fee"`
}

func (m *PoolParams) Reset()         { *m = PoolParams{} }
func (m *PoolParams) String() string { return proto.CompactTextString(m) }
func (*PoolParams) ProtoMessage()    {}
func (*PoolParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5ab9bc6d45f98ce, []int{1}
}
func (m *PoolParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolParams.Merge(m, src)
}
func (m *PoolParams) XXX_Size() int {
	return m.Size()
}
func (m *PoolParams) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolParams.DiscardUnknown(m)
}

var xxx_messageInfo_PoolParams proto.InternalMessageInfo

func (m *PoolParams) GetLock() bool {
	if m != nil {
		return m.Lock
	}
	return false
}

type PoolAccount struct {
	*types1.BaseAccount `protobuf:"bytes,1,opt,name=baseAccount,proto3,embedded=baseAccount" json:"baseAccount,omitempty" yaml:"base_account"`
	Id                  uint64                                 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	PoolParams          PoolParams                             `protobuf:"bytes,3,opt,name=poolParams,proto3" json:"poolParams" yaml:"pool_params"`
	TotalWeight         github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=totalWeight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"totalWeight" yaml:"total_weight"`
	TotalShare          types.Coin                             `protobuf:"bytes,5,opt,name=totalShare,proto3" json:"totalShare" yaml:"total_share"`
	PoolAssets          []PoolAsset                            `protobuf:"bytes,6,rep,name=poolAssets,proto3" json:"poolAssets" yaml:"pool_assets"`
	FuturePoolGoverner  string                                 `protobuf:"bytes,7,opt,name=future_pool_governer,json=futurePoolGoverner,proto3" json:"future_pool_governer,omitempty" yaml:"future_pool_governer"`
}

func (m *PoolAccount) Reset()      { *m = PoolAccount{} }
func (*PoolAccount) ProtoMessage() {}
func (*PoolAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5ab9bc6d45f98ce, []int{2}
}
func (m *PoolAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PoolAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PoolAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PoolAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolAccount.Merge(m, src)
}
func (m *PoolAccount) XXX_Size() int {
	return m.Size()
}
func (m *PoolAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolAccount.DiscardUnknown(m)
}

var xxx_messageInfo_PoolAccount proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PoolAsset)(nil), "osmosis.gamm.v1beta1.PoolAsset")
	proto.RegisterType((*PoolParams)(nil), "osmosis.gamm.v1beta1.PoolParams")
	proto.RegisterType((*PoolAccount)(nil), "osmosis.gamm.v1beta1.PoolAccount")
}

func init() { proto.RegisterFile("osmosis/gamm/v1beta1/pool.proto", fileDescriptor_e5ab9bc6d45f98ce) }

var fileDescriptor_e5ab9bc6d45f98ce = []byte{
	// 619 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xcd, 0x4a, 0xdd, 0x40,
	0x14, 0xc7, 0x6f, 0xf4, 0xaa, 0x75, 0x22, 0x2d, 0x8c, 0x77, 0x11, 0xaf, 0x90, 0xb9, 0x64, 0x51,
	0xec, 0xc2, 0x04, 0xed, 0xce, 0x9d, 0xb1, 0xb6, 0x08, 0x5d, 0xd8, 0x94, 0x52, 0xa8, 0x85, 0x30,
	0x37, 0x19, 0x73, 0x83, 0x49, 0x26, 0x64, 0xe6, 0xfa, 0xd1, 0x27, 0xe8, 0xb2, 0xcb, 0x2e, 0x7d,
	0x88, 0x2e, 0xfa, 0x08, 0x2e, 0xa5, 0xab, 0xe2, 0x22, 0x14, 0x85, 0x3e, 0xc0, 0x7d, 0x82, 0x32,
	0x1f, 0xb9, 0x06, 0xb1, 0x14, 0xe9, 0x2a, 0x33, 0xe7, 0x9c, 0xf9, 0x9d, 0x33, 0xff, 0x73, 0x32,
	0x00, 0x51, 0x96, 0x53, 0x96, 0x32, 0x2f, 0xc1, 0x79, 0xee, 0x1d, 0x6f, 0x0c, 0x09, 0xc7, 0x1b,
	0x5e, 0x49, 0x69, 0xe6, 0x96, 0x15, 0xe5, 0x14, 0xf6, 0x74, 0x80, 0x2b, 0x02, 0x5c, 0x1d, 0xd0,
	0x5f, 0x89, 0xa4, 0x39, 0x94, 0x31, 0x9e, 0xda, 0xa8, 0x03, 0xfd, 0x5e, 0x42, 0x13, 0xaa, 0xec,
	0x62, 0xa5, 0xad, 0xb6, 0x8a, 0xf1, 0xf0, 0x98, 0x8f, 0xa6, 0x69, 0xc4, 0xe6, 0x8e, 0x7f, 0x88,
	0x19, 0x99, 0xfa, 0x23, 0x9a, 0x16, 0xca, 0xef, 0x7c, 0x37, 0xc0, 0xe2, 0x3e, 0xa5, 0xd9, 0x36,
	0x63, 0x84, 0xc3, 0x18, 0xcc, 0x9f, 0x90, 0x34, 0x19, 0x71, 0xcb, 0x18, 0x18, 0x6b, 0x8b, 0xfe,
	0xeb, 0x8b, 0x1a, 0x75, 0xae, 0x6a, 0xf4, 0x34, 0x49, 0xf9, 0x68, 0x3c, 0x74, 0x23, 0x9a, 0xeb,
	0xa2, 0xf4, 0x67, 0x9d, 0xc5, 0x47, 0x1e, 0x3f, 0x2b, 0x09, 0x73, 0xf7, 0x0a, 0x3e, 0xa9, 0x51,
	0xff, 0x0c, 0xe7, 0xd9, 0x96, 0x13, 0x93, 0x82, 0x56, 0x39, 0xce, 0xd2, 0x4f, 0x24, 0x0e, 0x15,
	0xd2, 0x09, 0x34, 0x1b, 0xee, 0x82, 0x39, 0x4e, 0x8f, 0x48, 0x61, 0xcd, 0x0c, 0x8c, 0x35, 0x73,
	0x73, 0xc5, 0xd5, 0xf7, 0x14, 0x35, 0x36, 0x4a, 0xb8, 0x3b, 0x34, 0x2d, 0xfc, 0x9e, 0xc8, 0x3f,
	0xa9, 0xd1, 0x92, 0xa2, 0xca, 0x53, 0x4e, 0xa0, 0x4e, 0x3b, 0x57, 0x06, 0x00, 0xa2, 0xf4, 0x7d,
	0x5c, 0xe1, 0x9c, 0x41, 0x08, 0xba, 0x19, 0x8d, 0x8e, 0x64, 0xe5, 0x8f, 0x02, 0xb9, 0x86, 0x07,
	0x60, 0x81, 0x9d, 0xe0, 0xf2, 0x25, 0x21, 0x32, 0xd7, 0xa2, 0xbf, 0xfd, 0x80, 0x0b, 0xbd, 0x20,
	0xd1, 0xa4, 0x46, 0x4f, 0x54, 0x6a, 0x81, 0x09, 0x0f, 0x09, 0x71, 0x82, 0x86, 0x28, 0xe0, 0xe4,
	0x34, 0xe5, 0x02, 0x3e, 0xfb, 0x7f, 0x70, 0x81, 0xd1, 0x70, 0x4d, 0x74, 0x7e, 0x77, 0x81, 0x29,
	0xfb, 0x12, 0x45, 0x74, 0x5c, 0x70, 0x18, 0x02, 0x53, 0xc8, 0xa3, 0xb7, 0xf2, 0x92, 0xe6, 0xe6,
	0xa0, 0x51, 0x4e, 0x36, 0xbc, 0x51, 0xce, 0xbf, 0x8d, 0xf3, 0x57, 0x2f, 0x6b, 0x64, 0x4c, 0x6a,
	0xb4, 0xac, 0x12, 0x09, 0x44, 0x88, 0x95, 0xcf, 0x09, 0xda, 0x44, 0xf8, 0x18, 0xcc, 0xa4, 0xb1,
	0x54, 0xa9, 0x1b, 0xcc, 0xa4, 0x31, 0xfc, 0x08, 0x40, 0x39, 0x15, 0x57, 0x5e, 0x50, 0xe4, 0xbb,
	0x6f, 0x68, 0xdd, 0xdb, 0x26, 0xf8, 0x7d, 0xdd, 0x30, 0xa8, 0xf2, 0x09, 0x42, 0x58, 0x4a, 0x97,
	0x13, 0xb4, 0x78, 0x30, 0x01, 0x26, 0xa7, 0x1c, 0x67, 0xef, 0xd5, 0xb4, 0x75, 0xa5, 0x7e, 0xbb,
	0x0f, 0x9e, 0xb6, 0xe5, 0x66, 0x2e, 0x38, 0xce, 0xa6, 0x63, 0xd6, 0x26, 0xc3, 0x77, 0x00, 0xc8,
	0xed, 0xdb, 0x11, 0xae, 0x88, 0x35, 0xf7, 0xaf, 0x81, 0xbb, 0x53, 0xbf, 0x02, 0x33, 0x71, 0xd6,
	0x09, 0x5a, 0x20, 0x78, 0xa0, 0xd4, 0x91, 0x7f, 0x0d, 0xb3, 0xe6, 0x07, 0xb3, 0x6b, 0xe6, 0x26,
	0xfa, 0xbb, 0x3a, 0x32, 0xee, 0x5e, 0x71, 0xb0, 0x24, 0x68, 0x71, 0x14, 0x0e, 0xbe, 0x01, 0xbd,
	0xc3, 0x31, 0x1f, 0x57, 0x24, 0x94, 0x21, 0x09, 0x3d, 0x26, 0x55, 0x41, 0x2a, 0x6b, 0x41, 0xaa,
	0x84, 0x26, 0x35, 0x5a, 0x55, 0x84, 0xfb, 0xa2, 0x9c, 0x00, 0x2a, 0xb3, 0xc8, 0xfb, 0x4a, 0x1b,
	0xb7, 0xac, 0xcf, 0xe7, 0xa8, 0xf3, 0xf5, 0x1c, 0x75, 0x7e, 0x7c, 0x5b, 0x5f, 0x6a, 0xcd, 0xd5,
	0x9e, 0xbf, 0x73, 0x71, 0x6d, 0x1b, 0x97, 0xd7, 0xb6, 0xf1, 0xeb, 0xda, 0x36, 0xbe, 0xdc, 0xd8,
	0x9d, 0xcb, 0x1b, 0xbb, 0xf3, 0xf3, 0xc6, 0xee, 0x7c, 0x78, 0xd6, 0x6e, 0xc3, 0x7a, 0xf3, 0x9e,
	0x35, 0xdf, 0x53, 0xf5, 0xb2, 0xc9, 0x6e, 0x0c, 0xe7, 0xe5, 0x63, 0xf2, 0xfc, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x66, 0x33, 0x26, 0xe4, 0xf6, 0x04, 0x00, 0x00,
}

func (m *PoolAsset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolAsset) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolAsset) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Token.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Weight.Size()
		i -= size
		if _, err := m.Weight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *PoolParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.ExitFee.Size()
		i -= size
		if _, err := m.ExitFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.SwapFee.Size()
		i -= size
		if _, err := m.SwapFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Lock {
		i--
		if m.Lock {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *PoolAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PoolAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PoolAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FuturePoolGoverner) > 0 {
		i -= len(m.FuturePoolGoverner)
		copy(dAtA[i:], m.FuturePoolGoverner)
		i = encodeVarintPool(dAtA, i, uint64(len(m.FuturePoolGoverner)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.PoolAssets) > 0 {
		for iNdEx := len(m.PoolAssets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolAssets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPool(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	{
		size, err := m.TotalShare.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.TotalWeight.Size()
		i -= size
		if _, err := m.TotalWeight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.PoolParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.Id != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if m.BaseAccount != nil {
		{
			size, err := m.BaseAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPool(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PoolAsset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Weight.Size()
	n += 1 + l + sovPool(uint64(l))
	l = m.Token.Size()
	n += 1 + l + sovPool(uint64(l))
	return n
}

func (m *PoolParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Lock {
		n += 2
	}
	l = m.SwapFee.Size()
	n += 1 + l + sovPool(uint64(l))
	l = m.ExitFee.Size()
	n += 1 + l + sovPool(uint64(l))
	return n
}

func (m *PoolAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseAccount != nil {
		l = m.BaseAccount.Size()
		n += 1 + l + sovPool(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovPool(uint64(m.Id))
	}
	l = m.PoolParams.Size()
	n += 1 + l + sovPool(uint64(l))
	l = m.TotalWeight.Size()
	n += 1 + l + sovPool(uint64(l))
	l = m.TotalShare.Size()
	n += 1 + l + sovPool(uint64(l))
	if len(m.PoolAssets) > 0 {
		for _, e := range m.PoolAssets {
			l = e.Size()
			n += 1 + l + sovPool(uint64(l))
		}
	}
	l = len(m.FuturePoolGoverner)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	return n
}

func sovPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPool(x uint64) (n int) {
	return sovPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PoolAsset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PoolAsset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolAsset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Weight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Token.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPool
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPool
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PoolParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PoolParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lock", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Lock = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SwapFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExitFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ExitFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPool
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPool
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PoolAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PoolAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PoolAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseAccount == nil {
				m.BaseAccount = &types1.BaseAccount{}
			}
			if err := m.BaseAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PoolParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalWeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalWeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalShare", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalShare.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolAssets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolAssets = append(m.PoolAssets, PoolAsset{})
			if err := m.PoolAssets[len(m.PoolAssets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FuturePoolGoverner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FuturePoolGoverner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPool
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPool
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPool
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPool
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowPool
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPool = fmt.Errorf("proto: unexpected end of group")
)
