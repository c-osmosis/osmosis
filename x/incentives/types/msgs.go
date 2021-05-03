package types

import (
	"errors"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	lockuptypes "github.com/osmosis-labs/osmosis/x/lockup/types"
)

// constants
const (
	TypeMsgCreatePot = "create_pot"
	TypeMsgAddToPot  = "add_to_pot"
)

var _ sdk.Msg = &MsgCreatePot{}

func (m MsgCreatePot) Route() string { return RouterKey }
func (m MsgCreatePot) Type() string  { return TypeMsgCreatePot }
func (m MsgCreatePot) ValidateBasic() error {
	if m.Owner.Empty() {
		return errors.New("owner should be set")
	}
	if sdk.ValidateDenom(m.DistributeTo.Denom) != nil {
		return errors.New("denom should be valid for the condition")
	}
	if lockuptypes.LockQueryType_name[int32(m.DistributeTo.LockQueryType)] == "" {
		return errors.New("lock query type is invalid")
	}
	if m.StartTime.Equal(time.Time{}) {
		return errors.New("distribution start time should be set")
	}
	if m.NumEpochsPaidOver == 0 {
		return errors.New("distribution period should be at least 1 epoch")
	}
	if m.IsPerpetual && m.NumEpochsPaidOver != 1 {
		return errors.New("distribution period should be 1 epoch for perpetual pot")
	}

	return nil
}
func (m MsgCreatePot) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}
func (m MsgCreatePot) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Owner}
}

var _ sdk.Msg = &MsgAddToPot{}

func (m MsgAddToPot) Route() string { return RouterKey }
func (m MsgAddToPot) Type() string  { return TypeMsgAddToPot }
func (m MsgAddToPot) ValidateBasic() error {
	if m.Owner.Empty() {
		return errors.New("owner should be set")
	}
	if m.Rewards.Empty() {
		return errors.New("additional rewards should not be empty")
	}

	return nil
}
func (m MsgAddToPot) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}
func (m MsgAddToPot) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Owner}
}
