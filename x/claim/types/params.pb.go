// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/claim/v1beta1/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// Params defines the claim module's parameters.
type Params struct {
	AirdropStart       time.Time     `protobuf:"bytes,1,opt,name=airdrop_start,json=airdropStart,proto3,stdtime" json:"airdrop_start" yaml:"airdrop_start"`
	DurationUntilDecay time.Duration `protobuf:"bytes,2,opt,name=duration_until_decay,json=durationUntilDecay,proto3,stdduration" json:"duration_until_decay,omitempty" yaml:"duration_until_decay"`
	DurationOfDecay    time.Duration `protobuf:"bytes,3,opt,name=duration_of_decay,json=durationOfDecay,proto3,stdduration" json:"duration_of_decay,omitempty" yaml:"duration_of_decay"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1687b9ddfb80c0a, []int{0}
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

func (m *Params) GetAirdropStart() time.Time {
	if m != nil {
		return m.AirdropStart
	}
	return time.Time{}
}

func (m *Params) GetDurationUntilDecay() time.Duration {
	if m != nil {
		return m.DurationUntilDecay
	}
	return 0
}

func (m *Params) GetDurationOfDecay() time.Duration {
	if m != nil {
		return m.DurationOfDecay
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "osmosis.claim.v1beta1.Params")
}

func init() {
	proto.RegisterFile("osmosis/claim/v1beta1/params.proto", fileDescriptor_a1687b9ddfb80c0a)
}

var fileDescriptor_a1687b9ddfb80c0a = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xbd, 0x4e, 0xeb, 0x30,
	0x14, 0xc7, 0xe3, 0x5b, 0xa9, 0x43, 0xee, 0xbd, 0x42, 0x44, 0x45, 0x2a, 0xad, 0xe4, 0x54, 0x99,
	0x10, 0x02, 0x5b, 0x85, 0x8d, 0xb1, 0x74, 0x61, 0x02, 0xf1, 0xb1, 0xb0, 0x54, 0x4e, 0x9a, 0x06,
	0x4b, 0x71, 0x1d, 0xc5, 0x0e, 0x22, 0xaf, 0xc0, 0x54, 0x16, 0xc4, 0x23, 0x75, 0xec, 0xc8, 0x14,
	0x50, 0xbb, 0x31, 0xf6, 0x09, 0x50, 0x1c, 0x1b, 0xa9, 0x1f, 0x12, 0x53, 0xe2, 0xf3, 0xff, 0x9d,
	0x73, 0x7e, 0xc3, 0xb1, 0x3d, 0x2e, 0x18, 0x17, 0x54, 0xe0, 0x20, 0x26, 0x94, 0xe1, 0xc7, 0xae,
	0x1f, 0x4a, 0xd2, 0xc5, 0x09, 0x49, 0x09, 0x13, 0x28, 0x49, 0xb9, 0xe4, 0xce, 0x9e, 0x66, 0x90,
	0x62, 0x90, 0x66, 0x5a, 0x8d, 0x88, 0x47, 0x5c, 0x11, 0xb8, 0xfc, 0xab, 0xe0, 0x16, 0x8c, 0x38,
	0x8f, 0xe2, 0x10, 0xab, 0x97, 0x9f, 0x8d, 0xf0, 0x30, 0x4b, 0x89, 0xa4, 0x7c, 0xac, 0x73, 0x77,
	0x3d, 0x97, 0x94, 0x85, 0x42, 0x12, 0x96, 0x54, 0x80, 0xf7, 0x52, 0xb3, 0xeb, 0x57, 0x6a, 0xbd,
	0x43, 0xec, 0xff, 0x84, 0xa6, 0xc3, 0x94, 0x27, 0x03, 0x21, 0x49, 0x2a, 0x9b, 0xa0, 0x03, 0x0e,
	0xfe, 0x9e, 0xb4, 0x50, 0x35, 0x03, 0x99, 0x19, 0xe8, 0xd6, 0xcc, 0xe8, 0x75, 0xa6, 0x85, 0x6b,
	0x2d, 0x0b, 0xb7, 0x91, 0x13, 0x16, 0x9f, 0x79, 0x2b, 0xed, 0xde, 0xe4, 0xc3, 0x05, 0xd7, 0xff,
	0x74, 0xed, 0xa6, 0x2c, 0x39, 0xaf, 0xc0, 0x6e, 0x18, 0xc3, 0x41, 0x36, 0x96, 0x34, 0x1e, 0x0c,
	0xc3, 0x80, 0xe4, 0xcd, 0x3f, 0x6a, 0xd5, 0xfe, 0xc6, 0xaa, 0xbe, 0x86, 0x7b, 0x17, 0xe5, 0xa6,
	0xaf, 0xc2, 0x85, 0xdb, 0xda, 0x8f, 0x38, 0xa3, 0x32, 0x64, 0x89, 0xcc, 0x97, 0x85, 0xdb, 0xae,
	0x5c, 0xb6, 0x71, 0xde, 0x5b, 0xa9, 0xe4, 0x98, 0xe8, 0xae, 0x4c, 0xfa, 0x65, 0xe0, 0x3c, 0x03,
	0x7b, 0xf7, 0xa7, 0x83, 0x8f, 0xb4, 0x55, 0xed, 0x37, 0xab, 0x73, 0x6d, 0xd5, 0xde, 0xe8, 0x5d,
	0x51, 0x6a, 0xae, 0x29, 0x19, 0xa8, 0xf2, 0xd9, 0x31, 0xf5, 0xcb, 0x91, 0x92, 0xe9, 0xf5, 0xa7,
	0x73, 0x08, 0x66, 0x73, 0x08, 0x3e, 0xe7, 0x10, 0x4c, 0x16, 0xd0, 0x9a, 0x2d, 0xa0, 0xf5, 0xbe,
	0x80, 0xd6, 0xfd, 0x61, 0x44, 0xe5, 0x43, 0xe6, 0xa3, 0x80, 0x33, 0x1c, 0x1c, 0x9b, 0x63, 0x32,
	0xdf, 0x27, 0x7d, 0x56, 0x32, 0x4f, 0x42, 0xe1, 0xd7, 0x95, 0xee, 0xe9, 0x77, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xae, 0xb5, 0x31, 0x1f, 0x74, 0x02, 0x00, 0x00,
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
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.DurationOfDecay, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.DurationOfDecay):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintParams(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1a
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.DurationUntilDecay, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.DurationUntilDecay):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintParams(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x12
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.AirdropStart, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.AirdropStart):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintParams(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.AirdropStart)
	n += 1 + l + sovParams(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.DurationUntilDecay)
	n += 1 + l + sovParams(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.DurationOfDecay)
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
				return fmt.Errorf("proto: wrong wireType = %d for field AirdropStart", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.AirdropStart, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DurationUntilDecay", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.DurationUntilDecay, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DurationOfDecay", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.DurationOfDecay, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthParams
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)