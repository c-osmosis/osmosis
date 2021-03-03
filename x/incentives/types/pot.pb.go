// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/incentives/pot.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type LockType int32

const (
	by_duration LockType = 0
	by_time     LockType = 1
)

var LockType_name = map[int32]string{
	0: "by_duration",
	1: "by_time",
}

var LockType_value = map[string]int32{
	"by_duration": 0,
	"by_time":     1,
}

func (x LockType) String() string {
	return proto.EnumName(LockType_name, int32(x))
}

func (LockType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2bf3751e058d24ca, []int{0}
}

type DistrCondition struct {
	LockType  LockType             `protobuf:"varint,1,opt,name=lock_type,json=lockType,proto3,enum=osmosis.incentives.LockType" json:"lock_type,omitempty"`
	Denom     string               `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	Duration  *duration.Duration   `protobuf:"bytes,3,opt,name=duration,proto3" json:"duration,omitempty"`
	Timestamp *timestamp.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (m *DistrCondition) Reset()         { *m = DistrCondition{} }
func (m *DistrCondition) String() string { return proto.CompactTextString(m) }
func (*DistrCondition) ProtoMessage()    {}
func (*DistrCondition) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bf3751e058d24ca, []int{0}
}
func (m *DistrCondition) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DistrCondition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DistrCondition.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DistrCondition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistrCondition.Merge(m, src)
}
func (m *DistrCondition) XXX_Size() int {
	return m.Size()
}
func (m *DistrCondition) XXX_DiscardUnknown() {
	xxx_messageInfo_DistrCondition.DiscardUnknown(m)
}

var xxx_messageInfo_DistrCondition proto.InternalMessageInfo

func (m *DistrCondition) GetLockType() LockType {
	if m != nil {
		return m.LockType
	}
	return by_duration
}

func (m *DistrCondition) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *DistrCondition) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *DistrCondition) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type Pot struct {
	Id           uint64                                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DistributeTo []*DistrCondition                        `protobuf:"bytes,2,rep,name=distribute_to,json=distributeTo,proto3" json:"distribute_to,omitempty"`
	Coins        github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
	StartTime    *timestamp.Timestamp                     `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	NumEpochs    uint64                                   `protobuf:"varint,5,opt,name=num_epochs,json=numEpochs,proto3" json:"num_epochs,omitempty"`
}

func (m *Pot) Reset()         { *m = Pot{} }
func (m *Pot) String() string { return proto.CompactTextString(m) }
func (*Pot) ProtoMessage()    {}
func (*Pot) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bf3751e058d24ca, []int{1}
}
func (m *Pot) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pot.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pot.Merge(m, src)
}
func (m *Pot) XXX_Size() int {
	return m.Size()
}
func (m *Pot) XXX_DiscardUnknown() {
	xxx_messageInfo_Pot.DiscardUnknown(m)
}

var xxx_messageInfo_Pot proto.InternalMessageInfo

func (m *Pot) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Pot) GetDistributeTo() []*DistrCondition {
	if m != nil {
		return m.DistributeTo
	}
	return nil
}

func (m *Pot) GetCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Coins
	}
	return nil
}

func (m *Pot) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Pot) GetNumEpochs() uint64 {
	if m != nil {
		return m.NumEpochs
	}
	return 0
}

func init() {
	proto.RegisterEnum("osmosis.incentives.LockType", LockType_name, LockType_value)
	proto.RegisterType((*DistrCondition)(nil), "osmosis.incentives.DistrCondition")
	proto.RegisterType((*Pot)(nil), "osmosis.incentives.Pot")
}

func init() { proto.RegisterFile("osmosis/incentives/pot.proto", fileDescriptor_2bf3751e058d24ca) }

var fileDescriptor_2bf3751e058d24ca = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xf6, 0xe5, 0x07, 0x24, 0x2f, 0x10, 0xaa, 0x53, 0x07, 0x37, 0x2a, 0x4e, 0x94, 0xc9, 0x42,
	0xea, 0x1d, 0x0d, 0x42, 0xa2, 0x6b, 0x5b, 0x84, 0x90, 0x18, 0x90, 0x95, 0x89, 0xc5, 0xf2, 0x8f,
	0xc3, 0x3d, 0x25, 0xf6, 0xb3, 0x72, 0xe7, 0x8a, 0xfc, 0x07, 0x8c, 0xec, 0x8c, 0x6c, 0xfc, 0x25,
	0x1d, 0x2b, 0x26, 0x26, 0x40, 0xc9, 0x3f, 0x82, 0x7c, 0xb6, 0x9b, 0x42, 0x19, 0x98, 0xee, 0xde,
	0xbd, 0xef, 0x7d, 0xf7, 0x7d, 0xdf, 0x1d, 0x1c, 0xa2, 0x4a, 0x51, 0x49, 0xc5, 0x65, 0x16, 0x89,
	0x4c, 0xcb, 0x4b, 0xa1, 0x78, 0x8e, 0x9a, 0xe5, 0x2b, 0xd4, 0x48, 0x69, 0xdd, 0x65, 0xbb, 0xee,
	0x68, 0x3f, 0xc1, 0x04, 0x4d, 0x9b, 0x97, 0xbb, 0x0a, 0x39, 0x72, 0x12, 0xc4, 0x64, 0x29, 0xb8,
	0xa9, 0xc2, 0xe2, 0x3d, 0x8f, 0x8b, 0x55, 0xa0, 0x25, 0x66, 0x75, 0x7f, 0xfc, 0x77, 0x5f, 0xcb,
	0x54, 0x28, 0x1d, 0xa4, 0x79, 0x43, 0x10, 0x99, 0xbb, 0x78, 0x18, 0x28, 0xc1, 0x2f, 0x8f, 0x43,
	0xa1, 0x83, 0x63, 0x1e, 0xa1, 0xac, 0x09, 0xa6, 0xdf, 0x08, 0x0c, 0xcf, 0xa5, 0xd2, 0xab, 0x33,
	0xcc, 0x62, 0x59, 0x32, 0xd3, 0x13, 0xe8, 0x2f, 0x31, 0x5a, 0xf8, 0x7a, 0x9d, 0x0b, 0x9b, 0x4c,
	0x88, 0x3b, 0x9c, 0x1d, 0xb2, 0xbb, 0x8a, 0xd9, 0x1b, 0x8c, 0x16, 0xf3, 0x75, 0x2e, 0xbc, 0xde,
	0xb2, 0xde, 0xd1, 0x7d, 0xe8, 0xc6, 0x22, 0xc3, 0xd4, 0x6e, 0x4d, 0x88, 0xdb, 0xf7, 0xaa, 0x82,
	0x3e, 0x87, 0x5e, 0x23, 0xdb, 0x6e, 0x4f, 0x88, 0x3b, 0x98, 0x1d, 0xb0, 0x4a, 0x37, 0x6b, 0x74,
	0xb3, 0xf3, 0x1a, 0xe0, 0xdd, 0x40, 0xe9, 0x0b, 0xe8, 0xdf, 0xb8, 0xb1, 0x3b, 0x66, 0x6e, 0x74,
	0x67, 0x6e, 0xde, 0x20, 0xbc, 0x1d, 0x78, 0xfa, 0xb9, 0x05, 0xed, 0xb7, 0xa8, 0xe9, 0x10, 0x5a,
	0x32, 0x36, 0x16, 0x3a, 0x5e, 0x4b, 0xc6, 0xf4, 0x15, 0x3c, 0x8c, 0x4b, 0xaf, 0x32, 0x2c, 0xb4,
	0xf0, 0x35, 0xda, 0xad, 0x49, 0xdb, 0x1d, 0xcc, 0xa6, 0xff, 0x72, 0xf7, 0x67, 0x28, 0xde, 0x83,
	0xdd, 0xe0, 0x1c, 0x69, 0x00, 0xdd, 0x32, 0x43, 0x65, 0xb7, 0x0d, 0xc1, 0x01, 0xab, 0x52, 0x66,
	0x65, 0xca, 0xac, 0x4e, 0x99, 0x9d, 0xa1, 0xcc, 0x4e, 0x9f, 0x5e, 0xfd, 0x18, 0x5b, 0x5f, 0x7f,
	0x8e, 0xdd, 0x44, 0xea, 0x8b, 0x22, 0x64, 0x11, 0xa6, 0xbc, 0x7e, 0x92, 0x6a, 0x39, 0x52, 0xf1,
	0x82, 0x97, 0x51, 0x2b, 0x33, 0xa0, 0xbc, 0x8a, 0x99, 0x9e, 0x00, 0x28, 0x1d, 0xac, 0xb4, 0x5f,
	0xda, 0xfa, 0x1f, 0xfb, 0x06, 0x5d, 0xd6, 0xf4, 0x31, 0x40, 0x56, 0xa4, 0xbe, 0xc8, 0x31, 0xba,
	0x50, 0x76, 0xd7, 0xd8, 0xef, 0x67, 0x45, 0xfa, 0xd2, 0x1c, 0x3c, 0x61, 0xd0, 0x6b, 0x9e, 0x8e,
	0x3e, 0x82, 0x41, 0xb8, 0xf6, 0x9b, 0xc8, 0xf7, 0x2c, 0x3a, 0x80, 0xfb, 0xe1, 0xda, 0xdc, 0xb9,
	0x47, 0x46, 0x9d, 0x8f, 0x5f, 0x1c, 0xeb, 0xf4, 0xf5, 0xd5, 0xc6, 0x21, 0xd7, 0x1b, 0x87, 0xfc,
	0xda, 0x38, 0xe4, 0xd3, 0xd6, 0xb1, 0xae, 0xb7, 0x8e, 0xf5, 0x7d, 0xeb, 0x58, 0xef, 0xf8, 0x6d,
	0x53, 0x47, 0xcd, 0x97, 0x6f, 0xd6, 0x0f, 0xb7, 0x3f, 0xbf, 0x71, 0x18, 0xde, 0x33, 0xc2, 0x9f,
	0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x37, 0x86, 0x2e, 0x1f, 0x03, 0x00, 0x00,
}

func (m *DistrCondition) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DistrCondition) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DistrCondition) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Timestamp != nil {
		{
			size, err := m.Timestamp.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPot(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Duration != nil {
		{
			size, err := m.Duration.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPot(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintPot(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if m.LockType != 0 {
		i = encodeVarintPot(dAtA, i, uint64(m.LockType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Pot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pot) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pot) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NumEpochs != 0 {
		i = encodeVarintPot(dAtA, i, uint64(m.NumEpochs))
		i--
		dAtA[i] = 0x28
	}
	if m.StartTime != nil {
		{
			size, err := m.StartTime.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPot(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPot(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.DistributeTo) > 0 {
		for iNdEx := len(m.DistributeTo) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DistributeTo[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPot(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Id != 0 {
		i = encodeVarintPot(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPot(dAtA []byte, offset int, v uint64) int {
	offset -= sovPot(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DistrCondition) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LockType != 0 {
		n += 1 + sovPot(uint64(m.LockType))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovPot(uint64(l))
	}
	if m.Duration != nil {
		l = m.Duration.Size()
		n += 1 + l + sovPot(uint64(l))
	}
	if m.Timestamp != nil {
		l = m.Timestamp.Size()
		n += 1 + l + sovPot(uint64(l))
	}
	return n
}

func (m *Pot) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovPot(uint64(m.Id))
	}
	if len(m.DistributeTo) > 0 {
		for _, e := range m.DistributeTo {
			l = e.Size()
			n += 1 + l + sovPot(uint64(l))
		}
	}
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovPot(uint64(l))
		}
	}
	if m.StartTime != nil {
		l = m.StartTime.Size()
		n += 1 + l + sovPot(uint64(l))
	}
	if m.NumEpochs != 0 {
		n += 1 + sovPot(uint64(m.NumEpochs))
	}
	return n
}

func sovPot(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPot(x uint64) (n int) {
	return sovPot(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DistrCondition) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPot
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
			return fmt.Errorf("proto: DistrCondition: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DistrCondition: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockType", wireType)
			}
			m.LockType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LockType |= LockType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPot
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
				return ErrInvalidLengthPot
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPot
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
				return ErrInvalidLengthPot
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Duration == nil {
				m.Duration = &duration.Duration{}
			}
			if err := m.Duration.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPot
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
				return ErrInvalidLengthPot
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timestamp == nil {
				m.Timestamp = &timestamp.Timestamp{}
			}
			if err := m.Timestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPot
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPot
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
func (m *Pot) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPot
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
			return fmt.Errorf("proto: Pot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPot
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributeTo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPot
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
				return ErrInvalidLengthPot
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DistributeTo = append(m.DistributeTo, &DistrCondition{})
			if err := m.DistributeTo[len(m.DistributeTo)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPot
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
				return ErrInvalidLengthPot
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPot
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
				return ErrInvalidLengthPot
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StartTime == nil {
				m.StartTime = &timestamp.Timestamp{}
			}
			if err := m.StartTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumEpochs", wireType)
			}
			m.NumEpochs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumEpochs |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPot
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthPot
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
func skipPot(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPot
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
					return 0, ErrIntOverflowPot
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
					return 0, ErrIntOverflowPot
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
				return 0, ErrInvalidLengthPot
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPot
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPot
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPot        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPot          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPot = fmt.Errorf("proto: unexpected end of group")
)
