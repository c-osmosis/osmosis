// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/claim/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/x/bank/types"
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

type Action int32

const (
	ActionAddLiquidity  Action = 0
	ActionSwap          Action = 1
	ActionVote          Action = 2
	ActionDelegateStake Action = 3
)

var Action_name = map[int32]string{
	0: "ActionAddLiquidity",
	1: "ActionSwap",
	2: "ActionVote",
	3: "ActionDelegateStake",
}

var Action_value = map[string]int32{
	"ActionAddLiquidity":  0,
	"ActionSwap":          1,
	"ActionVote":          2,
	"ActionDelegateStake": 3,
}

func (x Action) String() string {
	return proto.EnumName(Action_name, int32(x))
}

func (Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9236f2c69911ca0c, []int{0}
}

type UserActivities struct {
	User    string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Actions []Action `protobuf:"varint,2,rep,packed,name=actions,proto3,enum=osmosis.claim.v1beta1.Action" json:"actions,omitempty"`
}

func (m *UserActivities) Reset()         { *m = UserActivities{} }
func (m *UserActivities) String() string { return proto.CompactTextString(m) }
func (*UserActivities) ProtoMessage()    {}
func (*UserActivities) Descriptor() ([]byte, []int) {
	return fileDescriptor_9236f2c69911ca0c, []int{0}
}
func (m *UserActivities) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserActivities) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserActivities.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserActivities) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserActivities.Merge(m, src)
}
func (m *UserActivities) XXX_Size() int {
	return m.Size()
}
func (m *UserActivities) XXX_DiscardUnknown() {
	xxx_messageInfo_UserActivities.DiscardUnknown(m)
}

var xxx_messageInfo_UserActivities proto.InternalMessageInfo

func (m *UserActivities) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *UserActivities) GetActions() []Action {
	if m != nil {
		return m.Actions
	}
	return nil
}

type UserActivity struct {
	User   string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Action Action `protobuf:"varint,2,opt,name=action,proto3,enum=osmosis.claim.v1beta1.Action" json:"action,omitempty"`
}

func (m *UserActivity) Reset()         { *m = UserActivity{} }
func (m *UserActivity) String() string { return proto.CompactTextString(m) }
func (*UserActivity) ProtoMessage()    {}
func (*UserActivity) Descriptor() ([]byte, []int) {
	return fileDescriptor_9236f2c69911ca0c, []int{1}
}
func (m *UserActivity) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserActivity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserActivity.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserActivity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserActivity.Merge(m, src)
}
func (m *UserActivity) XXX_Size() int {
	return m.Size()
}
func (m *UserActivity) XXX_DiscardUnknown() {
	xxx_messageInfo_UserActivity.DiscardUnknown(m)
}

var xxx_messageInfo_UserActivity proto.InternalMessageInfo

func (m *UserActivity) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *UserActivity) GetAction() Action {
	if m != nil {
		return m.Action
	}
	return ActionAddLiquidity
}

// GenesisState defines the claim module's genesis state.
type GenesisState struct {
	AirdropAmount      github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=airdrop_amount,json=airdropAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"airdrop_amount"`
	AirdropStart       time.Time                              `protobuf:"bytes,2,opt,name=airdrop_start,json=airdropStart,proto3,stdtime" json:"airdrop_start" yaml:"airdrop_start"`
	DurationUntilDecay time.Duration                          `protobuf:"bytes,3,opt,name=duration_until_decay,json=durationUntilDecay,proto3,stdduration" json:"duration_until_decay,omitempty" yaml:"duration_until_decay"`
	DurationOfDecay    time.Duration                          `protobuf:"bytes,4,opt,name=duration_of_decay,json=durationOfDecay,proto3,stdduration" json:"duration_of_decay,omitempty" yaml:"duration_of_decay"`
	Claimables         []types.Balance                        `protobuf:"bytes,5,rep,name=claimables,proto3" json:"claimables"`
	Activities         []UserActivities                       `protobuf:"bytes,6,rep,name=activities,proto3" json:"activities"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_9236f2c69911ca0c, []int{2}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetAirdropStart() time.Time {
	if m != nil {
		return m.AirdropStart
	}
	return time.Time{}
}

func (m *GenesisState) GetDurationUntilDecay() time.Duration {
	if m != nil {
		return m.DurationUntilDecay
	}
	return 0
}

func (m *GenesisState) GetDurationOfDecay() time.Duration {
	if m != nil {
		return m.DurationOfDecay
	}
	return 0
}

func (m *GenesisState) GetClaimables() []types.Balance {
	if m != nil {
		return m.Claimables
	}
	return nil
}

func (m *GenesisState) GetActivities() []UserActivities {
	if m != nil {
		return m.Activities
	}
	return nil
}

func init() {
	proto.RegisterEnum("osmosis.claim.v1beta1.Action", Action_name, Action_value)
	proto.RegisterType((*UserActivities)(nil), "osmosis.claim.v1beta1.UserActivities")
	proto.RegisterType((*UserActivity)(nil), "osmosis.claim.v1beta1.UserActivity")
	proto.RegisterType((*GenesisState)(nil), "osmosis.claim.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("osmosis/claim/v1beta1/genesis.proto", fileDescriptor_9236f2c69911ca0c)
}

var fileDescriptor_9236f2c69911ca0c = []byte{
	// 618 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0x6e, 0xd6, 0x52, 0x84, 0x37, 0x4a, 0x31, 0x03, 0x42, 0x07, 0x49, 0x29, 0x02, 0x4d, 0x13,
	0x73, 0xb4, 0x21, 0x2e, 0x48, 0x1c, 0x1a, 0x2a, 0xa1, 0x09, 0x24, 0xa4, 0x96, 0x21, 0xc1, 0xa5,
	0x72, 0x12, 0x2f, 0x58, 0x4b, 0xe2, 0x10, 0x3b, 0x83, 0xfc, 0x07, 0x88, 0xd3, 0x4e, 0x88, 0x3b,
	0xff, 0xcc, 0x8e, 0x3b, 0x22, 0x0e, 0x05, 0x6d, 0x37, 0x8e, 0xfb, 0x07, 0x40, 0xb1, 0x9d, 0xfd,
	0x2c, 0xda, 0x29, 0xf6, 0x7b, 0xdf, 0x7b, 0xdf, 0x97, 0xe7, 0xcf, 0x06, 0xf7, 0x18, 0x8f, 0x19,
	0xa7, 0xdc, 0xf1, 0x23, 0x4c, 0x63, 0x67, 0x6b, 0xc5, 0x23, 0x02, 0xaf, 0x38, 0x21, 0x49, 0x08,
	0xa7, 0x1c, 0xa5, 0x19, 0x13, 0x0c, 0x5e, 0xd7, 0x20, 0x24, 0x41, 0x48, 0x83, 0x3a, 0xf3, 0x21,
	0x0b, 0x99, 0x44, 0x38, 0xe5, 0x4a, 0x81, 0x3b, 0x77, 0x7d, 0x89, 0x76, 0x3c, 0x9c, 0x6c, 0x4e,
	0xef, 0xd7, 0xb1, 0x42, 0xc6, 0xc2, 0x88, 0x38, 0x72, 0xe7, 0xe5, 0x1b, 0x4e, 0x90, 0x67, 0x58,
	0x50, 0x96, 0xe8, 0xbc, 0x7d, 0x3a, 0x2f, 0x68, 0x4c, 0xb8, 0xc0, 0x71, 0xaa, 0x00, 0x3d, 0x1f,
	0xb4, 0xd6, 0x39, 0xc9, 0xfa, 0xbe, 0xa0, 0x5b, 0x54, 0x50, 0xc2, 0x21, 0x04, 0x8d, 0x9c, 0x93,
	0xcc, 0x34, 0xba, 0xc6, 0xe2, 0xa5, 0xa1, 0x5c, 0xc3, 0xa7, 0xe0, 0x22, 0xf6, 0xcb, 0xb6, 0xdc,
	0x9c, 0xe9, 0xd6, 0x17, 0x5b, 0xab, 0x77, 0xd0, 0xd4, 0x1f, 0x41, 0x7d, 0x89, 0x72, 0x1b, 0x3b,
	0x13, 0xbb, 0x36, 0xac, 0x6a, 0x7a, 0x6f, 0xc1, 0xdc, 0x31, 0x92, 0x62, 0x2a, 0xc5, 0x63, 0xd0,
	0x54, 0x70, 0x73, 0xa6, 0x6b, 0x9c, 0xcb, 0x30, 0xd4, 0xe0, 0xde, 0xdf, 0x06, 0x98, 0x7b, 0xae,
	0x46, 0x32, 0x12, 0x58, 0x10, 0xb8, 0x0e, 0x5a, 0x98, 0x66, 0x41, 0xc6, 0xd2, 0x31, 0x8e, 0x59,
	0x9e, 0x08, 0xc5, 0xe2, 0xa2, 0x52, 0xd2, 0xcf, 0x89, 0xfd, 0x20, 0xa4, 0xe2, 0x7d, 0xee, 0x21,
	0x9f, 0xc5, 0x8e, 0x9e, 0xaf, 0xfa, 0x2c, 0xf3, 0x60, 0xd3, 0x11, 0x45, 0x4a, 0x38, 0x5a, 0x4b,
	0xc4, 0xf0, 0xb2, 0xee, 0xd2, 0x97, 0x4d, 0x20, 0x06, 0x55, 0x60, 0xcc, 0x05, 0xce, 0x84, 0x54,
	0x39, 0xbb, 0xda, 0x41, 0x6a, 0xc0, 0xa8, 0x1a, 0x30, 0x7a, 0x5d, 0x0d, 0xd8, 0xed, 0x96, 0x8c,
	0x07, 0x13, 0x7b, 0xbe, 0xc0, 0x71, 0xf4, 0xa4, 0x77, 0xa2, 0xbc, 0xb7, 0xfd, 0xcb, 0x36, 0x86,
	0x73, 0x3a, 0x36, 0x2a, 0x43, 0xf0, 0xab, 0x01, 0xe6, 0xab, 0xe3, 0x1b, 0xe7, 0x89, 0xa0, 0xd1,
	0x38, 0x20, 0x3e, 0x2e, 0xcc, 0xba, 0xa4, 0xba, 0x75, 0x86, 0x6a, 0xa0, 0xc1, 0xee, 0x5a, 0xc9,
	0xf4, 0x67, 0x62, 0x5b, 0xd3, 0xca, 0x1f, 0xb2, 0x98, 0x0a, 0x12, 0xa7, 0xa2, 0x38, 0x98, 0xd8,
	0x0b, 0x4a, 0xcb, 0x34, 0x5c, 0xef, 0x5b, 0x29, 0x09, 0x56, 0xa9, 0xf5, 0x32, 0x33, 0x28, 0x13,
	0xf0, 0x8b, 0x01, 0xae, 0x1e, 0x56, 0xb0, 0x0d, 0xad, 0xaa, 0x71, 0x9e, 0xaa, 0x67, 0x5a, 0xd5,
	0xc2, 0x99, 0xda, 0x13, 0x92, 0xcc, 0x53, 0x92, 0x2a, 0x90, 0xd2, 0x73, 0xa5, 0x8a, 0xbf, 0xda,
	0x50, 0x62, 0x5c, 0x00, 0xa4, 0x21, 0xb0, 0x17, 0x11, 0x6e, 0x5e, 0xe8, 0xd6, 0x17, 0x67, 0x57,
	0x6f, 0x23, 0x75, 0x84, 0xa8, 0xbc, 0x29, 0x87, 0x4e, 0x71, 0x71, 0x84, 0x13, 0x9f, 0x68, 0x33,
	0x1e, 0xab, 0x82, 0x2f, 0x00, 0xc0, 0x87, 0x86, 0x37, 0x9b, 0xb2, 0xc7, 0xfd, 0xff, 0xf8, 0xed,
	0xe4, 0xed, 0xa8, 0x9a, 0x1d, 0x95, 0x2f, 0x8d, 0x41, 0x53, 0x79, 0x12, 0xde, 0x00, 0x50, 0xad,
	0xfa, 0x41, 0xf0, 0x92, 0x7e, 0xc8, 0x69, 0x40, 0x45, 0xd1, 0xae, 0xc1, 0x16, 0x00, 0x2a, 0x3e,
	0xfa, 0x88, 0xd3, 0xb6, 0x71, 0xb4, 0x7f, 0xc3, 0x04, 0x69, 0xcf, 0xc0, 0x9b, 0xe0, 0x9a, 0xda,
	0x0f, 0x48, 0x44, 0x42, 0x2c, 0xc8, 0x48, 0xe0, 0x4d, 0xd2, 0xae, 0x77, 0x1a, 0x9f, 0xbf, 0x5b,
	0x35, 0x77, 0xb0, 0xb3, 0x67, 0x19, 0xbb, 0x7b, 0x96, 0xf1, 0x7b, 0xcf, 0x32, 0xb6, 0xf7, 0xad,
	0xda, 0xee, 0xbe, 0x55, 0xfb, 0xb1, 0x6f, 0xd5, 0xde, 0x2d, 0x1d, 0xf7, 0xf2, 0x72, 0xf5, 0xfe,
	0x54, 0xdf, 0x4f, 0xfa, 0x25, 0x92, 0x9e, 0xf6, 0x9a, 0xf2, 0x80, 0x1e, 0xfd, 0x0b, 0x00, 0x00,
	0xff, 0xff, 0xae, 0x77, 0x4f, 0x5f, 0xa7, 0x04, 0x00, 0x00,
}

func (m *UserActivities) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserActivities) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserActivities) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Actions) > 0 {
		dAtA2 := make([]byte, len(m.Actions)*10)
		var j1 int
		for _, num := range m.Actions {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintGenesis(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x12
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UserActivity) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserActivity) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserActivity) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Action != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Action))
		i--
		dAtA[i] = 0x10
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Activities) > 0 {
		for iNdEx := len(m.Activities) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Activities[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Claimables) > 0 {
		for iNdEx := len(m.Claimables) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Claimables[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	n3, err3 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.DurationOfDecay, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.DurationOfDecay):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintGenesis(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x22
	n4, err4 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.DurationUntilDecay, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.DurationUntilDecay):])
	if err4 != nil {
		return 0, err4
	}
	i -= n4
	i = encodeVarintGenesis(dAtA, i, uint64(n4))
	i--
	dAtA[i] = 0x1a
	n5, err5 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.AirdropStart, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.AirdropStart):])
	if err5 != nil {
		return 0, err5
	}
	i -= n5
	i = encodeVarintGenesis(dAtA, i, uint64(n5))
	i--
	dAtA[i] = 0x12
	{
		size := m.AirdropAmount.Size()
		i -= size
		if _, err := m.AirdropAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UserActivities) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.Actions) > 0 {
		l = 0
		for _, e := range m.Actions {
			l += sovGenesis(uint64(e))
		}
		n += 1 + sovGenesis(uint64(l)) + l
	}
	return n
}

func (m *UserActivity) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Action != 0 {
		n += 1 + sovGenesis(uint64(m.Action))
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.AirdropAmount.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.AirdropStart)
	n += 1 + l + sovGenesis(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.DurationUntilDecay)
	n += 1 + l + sovGenesis(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.DurationOfDecay)
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Claimables) > 0 {
		for _, e := range m.Claimables {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Activities) > 0 {
		for _, e := range m.Activities {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UserActivities) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: UserActivities: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserActivities: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v Action
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenesis
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= Action(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Actions = append(m.Actions, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenesis
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGenesis
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthGenesis
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				if elementCount != 0 && len(m.Actions) == 0 {
					m.Actions = make([]Action, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v Action
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= Action(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Actions = append(m.Actions, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Actions", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *UserActivity) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: UserActivity: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserActivity: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			m.Action = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Action |= Action(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AirdropAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AirdropAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AirdropStart", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.AirdropStart, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DurationUntilDecay", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.DurationUntilDecay, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DurationOfDecay", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.DurationOfDecay, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Claimables", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Claimables = append(m.Claimables, types.Balance{})
			if err := m.Claimables[len(m.Claimables)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Activities", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Activities = append(m.Activities, UserActivities{})
			if err := m.Activities[len(m.Activities)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)