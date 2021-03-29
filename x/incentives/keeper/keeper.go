package keeper

import (
	"fmt"

	"github.com/c-osmosis/osmosis/x/incentives/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"
)

type Keeper struct {
	cdc        codec.Marshaler
	storeKey   sdk.StoreKey
	paramSpace paramtypes.Subspace
	ak         authkeeper.AccountKeeper
	bk         types.BankKeeper
	lk         types.LockupKeeper
}

func NewKeeper(cdc codec.Marshaler, storeKey sdk.StoreKey, paramSpace paramtypes.Subspace, ak authkeeper.AccountKeeper, bk types.BankKeeper, lk types.LockupKeeper) *Keeper {
	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		paramSpace: paramSpace,
		ak:         ak,
		bk:         bk,
		lk:         lk,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
