// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/mint/v1beta1/mint.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Minter represents the minting state.
type Minter struct {
	// current epoch provisions
	EpochProvisions github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=epoch_provisions,json=epochProvisions,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"epoch_provisions" yaml:"epoch_provisions"`
}

func (m *Minter) Reset()         { *m = Minter{} }
func (m *Minter) String() string { return proto.CompactTextString(m) }
func (*Minter) ProtoMessage()    {}
func (*Minter) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccb38f8335e0f45b, []int{0}
}
func (m *Minter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Minter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Minter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Minter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Minter.Merge(m, src)
}
func (m *Minter) XXX_Size() int {
	return m.Size()
}
func (m *Minter) XXX_DiscardUnknown() {
	xxx_messageInfo_Minter.DiscardUnknown(m)
}

var xxx_messageInfo_Minter proto.InternalMessageInfo

// Params holds parameters for the mint module.
type Params struct {
	// type of coin to mint
	MintDenom string `protobuf:"bytes,1,opt,name=mint_denom,json=mintDenom,proto3" json:"mint_denom,omitempty"`
	// epoch provisions from the first epoch
	GenesisEpochProvisions github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=genesis_epoch_provisions,json=genesisEpochProvisions,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"genesis_epoch_provisions" yaml:"genesis_epoch_provisions"`
	// duration of an epoch
	EpochDuration time.Duration `protobuf:"bytes,3,opt,name=epoch_duration,json=epochDuration,proto3,stdduration" json:"epoch_duration" yaml:"epoch_duration"`
	// number of epochs take to reduce rewards
	ReductionPeriodInEpochs int64 `protobuf:"varint,4,opt,name=reduction_period_in_epochs,json=reductionPeriodInEpochs,proto3" json:"reduction_period_in_epochs,omitempty" yaml:"reduction_period_in_epochs"`
	// reduction multiplier to execute on each period
	ReductionFactorForEvent github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=reduction_factor_for_event,json=reductionFactorForEvent,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reduction_factor_for_event" yaml:"reduction_factor_for_event"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccb38f8335e0f45b, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMintDenom() string {
	if m != nil {
		return m.MintDenom
	}
	return ""
}

func (m *Params) GetEpochDuration() time.Duration {
	if m != nil {
		return m.EpochDuration
	}
	return 0
}

func (m *Params) GetReductionPeriodInEpochs() int64 {
	if m != nil {
		return m.ReductionPeriodInEpochs
	}
	return 0
}

func init() {
	proto.RegisterType((*Minter)(nil), "osmosis.mint.v1beta1.Minter")
	proto.RegisterType((*Params)(nil), "osmosis.mint.v1beta1.Params")
}

func init() { proto.RegisterFile("osmosis/mint/v1beta1/mint.proto", fileDescriptor_ccb38f8335e0f45b) }

var fileDescriptor_ccb38f8335e0f45b = []byte{
	// 474 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4d, 0x6f, 0xd3, 0x30,
	0x1c, 0xc6, 0x63, 0x56, 0x2a, 0xcd, 0x88, 0x17, 0x45, 0x83, 0x65, 0x95, 0x88, 0xbb, 0x48, 0xa0,
	0x72, 0x58, 0xa2, 0xc1, 0x6d, 0xc7, 0xd2, 0x4d, 0xda, 0x01, 0xa9, 0x84, 0x1b, 0x97, 0x28, 0x2f,
	0x6e, 0x66, 0xb1, 0xf8, 0x1f, 0xd9, 0x6e, 0x45, 0x2f, 0x7c, 0x01, 0x2e, 0xbb, 0xb1, 0x23, 0x1f,
	0x67, 0xc7, 0x1d, 0x11, 0x87, 0x80, 0xda, 0x6f, 0xd0, 0x4f, 0x80, 0x62, 0x27, 0x50, 0x3a, 0xed,
	0xb0, 0x93, 0xeb, 0xe7, 0x79, 0xfc, 0xf7, 0xaf, 0x4f, 0x64, 0x4c, 0x40, 0x16, 0x20, 0x99, 0x0c,
	0x0a, 0xc6, 0x55, 0x30, 0x3b, 0x4c, 0xa8, 0x8a, 0x0f, 0xf5, 0xc6, 0x2f, 0x05, 0x28, 0xb0, 0x77,
	0x9a, 0x80, 0xaf, 0xb5, 0x26, 0xd0, 0xdb, 0xc9, 0x21, 0x07, 0x1d, 0x08, 0xea, 0x5f, 0x26, 0xdb,
	0x23, 0x39, 0x40, 0x7e, 0x4e, 0x03, 0xbd, 0x4b, 0xa6, 0x93, 0x40, 0xb1, 0x82, 0x4a, 0x15, 0x17,
	0x65, 0x13, 0xd8, 0xdb, 0x0c, 0xc4, 0x7c, 0xde, 0x58, 0xee, 0xa6, 0x95, 0x4d, 0x45, 0xac, 0x18,
	0x70, 0xe3, 0x7b, 0x5f, 0x70, 0xf7, 0x1d, 0xe3, 0x8a, 0x0a, 0x5b, 0xe1, 0x27, 0xb4, 0x84, 0xf4,
	0x2c, 0x2a, 0x05, 0xcc, 0x98, 0x64, 0xc0, 0xa5, 0x83, 0xfa, 0x68, 0xb0, 0x3d, 0x3c, 0xbd, 0xaa,
	0x88, 0xf5, 0xb3, 0x22, 0x2f, 0x73, 0xa6, 0xce, 0xa6, 0x89, 0x9f, 0x42, 0x11, 0xa4, 0x9a, 0xbf,
	0x59, 0x0e, 0x64, 0xf6, 0x29, 0x50, 0xf3, 0x92, 0x4a, 0x7f, 0x44, 0xd3, 0x55, 0x45, 0x76, 0xe7,
	0x71, 0x71, 0x7e, 0xe4, 0x6d, 0xce, 0xf3, 0xc2, 0xc7, 0x5a, 0x1a, 0xff, 0x53, 0xbe, 0x75, 0x70,
	0x77, 0x1c, 0x8b, 0xb8, 0x90, 0xf6, 0x73, 0x8c, 0xeb, 0x32, 0xa2, 0x8c, 0x72, 0x28, 0xcc, 0xd5,
	0xe1, 0x76, 0xad, 0x8c, 0x6a, 0xc1, 0xfe, 0x8a, 0xb0, 0x93, 0x53, 0x4e, 0x25, 0x93, 0xd1, 0x0d,
	0xd0, 0x7b, 0x1a, 0xf4, 0xfd, 0x9d, 0x41, 0x89, 0x01, 0xbd, 0x6d, 0xae, 0x17, 0x3e, 0x6b, 0xac,
	0xe3, 0xff, 0xb9, 0xed, 0x14, 0x3f, 0x32, 0xe1, 0xb6, 0x4f, 0x67, 0xab, 0x8f, 0x06, 0x0f, 0x5e,
	0xef, 0xf9, 0xa6, 0x70, 0xbf, 0x2d, 0xdc, 0x1f, 0x35, 0x81, 0xe1, 0x7e, 0x4d, 0xb7, 0xaa, 0xc8,
	0xd3, 0xf5, 0x72, 0xda, 0xe3, 0xde, 0xe5, 0x2f, 0x82, 0xc2, 0x87, 0x5a, 0x6c, 0x4f, 0xd8, 0x09,
	0xee, 0x09, 0x9a, 0x4d, 0xd3, 0x7a, 0x13, 0x95, 0x54, 0x30, 0xc8, 0x22, 0xc6, 0x0d, 0xa5, 0x74,
	0x3a, 0x7d, 0x34, 0xd8, 0x1a, 0xbe, 0x58, 0x55, 0x64, 0xdf, 0x4c, 0xbc, 0x3d, 0xeb, 0x85, 0xbb,
	0x7f, 0xcd, 0xb1, 0xf6, 0x4e, 0xb9, 0xfe, 0x47, 0xd2, 0xbe, 0x40, 0xeb, 0x97, 0x4c, 0xe2, 0x54,
	0x81, 0x88, 0x26, 0x20, 0x22, 0x3a, 0xa3, 0x5c, 0x39, 0xf7, 0x75, 0xb1, 0x1f, 0xee, 0x5c, 0xec,
	0x0d, 0xa4, 0xcd, 0xc9, 0xeb, 0x48, 0x27, 0xda, 0x3b, 0x01, 0x71, 0x5c, 0x3b, 0x47, 0x9d, 0xcb,
	0xef, 0xc4, 0x1a, 0xbe, 0xbd, 0x5a, 0xb8, 0xe8, 0x7a, 0xe1, 0xa2, 0xdf, 0x0b, 0x17, 0x5d, 0x2c,
	0x5d, 0xeb, 0x7a, 0xe9, 0x5a, 0x3f, 0x96, 0xae, 0xf5, 0xf1, 0xd5, 0x3a, 0xc5, 0x41, 0xfb, 0xd2,
	0xda, 0xf5, 0xb3, 0x79, 0x73, 0x1a, 0x26, 0xe9, 0xea, 0xcf, 0xf0, 0xe6, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xca, 0xab, 0xe8, 0xc9, 0x90, 0x03, 0x00, 0x00,
}

func (m *Minter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Minter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Minter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.EpochProvisions.Size()
		i -= size
		if _, err := m.EpochProvisions.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.ReductionFactorForEvent.Size()
		i -= size
		if _, err := m.ReductionFactorForEvent.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.ReductionPeriodInEpochs != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.ReductionPeriodInEpochs))
		i--
		dAtA[i] = 0x20
	}
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.EpochDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.EpochDuration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintMint(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1a
	{
		size := m.GenesisEpochProvisions.Size()
		i -= size
		if _, err := m.GenesisEpochProvisions.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.MintDenom) > 0 {
		i -= len(m.MintDenom)
		copy(dAtA[i:], m.MintDenom)
		i = encodeVarintMint(dAtA, i, uint64(len(m.MintDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMint(dAtA []byte, offset int, v uint64) int {
	offset -= sovMint(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Minter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.EpochProvisions.Size()
	n += 1 + l + sovMint(uint64(l))
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MintDenom)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	l = m.GenesisEpochProvisions.Size()
	n += 1 + l + sovMint(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.EpochDuration)
	n += 1 + l + sovMint(uint64(l))
	if m.ReductionPeriodInEpochs != 0 {
		n += 1 + sovMint(uint64(m.ReductionPeriodInEpochs))
	}
	l = m.ReductionFactorForEvent.Size()
	n += 1 + l + sovMint(uint64(l))
	return n
}

func sovMint(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMint(x uint64) (n int) {
	return sovMint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Minter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: Minter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Minter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochProvisions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EpochProvisions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMint
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMint
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisEpochProvisions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GenesisEpochProvisions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.EpochDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReductionPeriodInEpochs", wireType)
			}
			m.ReductionPeriodInEpochs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReductionPeriodInEpochs |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReductionFactorForEvent", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReductionFactorForEvent.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMint
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMint
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
func skipMint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
				return 0, ErrInvalidLengthMint
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMint
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMint
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMint        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMint          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMint = fmt.Errorf("proto: unexpected end of group")
)
