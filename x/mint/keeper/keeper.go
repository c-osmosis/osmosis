package keeper

import (
	"fmt"
	"time"

	"github.com/c-osmosis/osmosis/x/mint/types"
	poolincentivestypes "github.com/c-osmosis/osmosis/x/pool-incentives/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"
)

// Keeper of the mint store
type Keeper struct {
	cdc              codec.BinaryMarshaler
	storeKey         sdk.StoreKey
	paramSpace       paramtypes.Subspace
	accountKeeper    types.AccountKeeper
	bankKeeper       types.BankKeeper
	hooks            types.MintHooks
	feeCollectorName string
}

// NewKeeper creates a new mint Keeper instance
func NewKeeper(
	cdc codec.BinaryMarshaler, key sdk.StoreKey, paramSpace paramtypes.Subspace,
	ak types.AccountKeeper, bk types.BankKeeper,
	feeCollectorName string,
) Keeper {
	// ensure mint module account is set
	if addr := ak.GetModuleAddress(types.ModuleName); addr == nil {
		panic("the mint module account has not been set")
	}

	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:              cdc,
		storeKey:         key,
		paramSpace:       paramSpace,
		accountKeeper:    ak,
		bankKeeper:       bk,
		feeCollectorName: feeCollectorName,
	}
}

//______________________________________________________________________

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

// Set the mint hooks
func (k *Keeper) SetHooks(h types.MintHooks) *Keeper {
	if k.hooks != nil {
		panic("cannot set mint hooks twice")
	}

	k.hooks = h

	return k
}

// GetEpochNum returns epoch number
func (k Keeper) GetEpochNum(ctx sdk.Context) int64 {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.EpochKey)
	if b == nil {
		return 0
	}

	return int64(sdk.BigEndianToUint64(b))
}

// SetEpochNum set epoch number
func (k Keeper) SetEpochNum(ctx sdk.Context, epochNum int64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.EpochKey, sdk.Uint64ToBigEndian(uint64(epochNum)))
}

// GetLastEpochTime returns last epoch time
func (k Keeper) GetLastEpochTime(ctx sdk.Context) time.Time {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.EpochTimeKey)
	if b == nil {
		return time.Time{}
	}
	epochTime, err := sdk.ParseTimeBytes(b)
	if err != nil {
		return time.Time{}
	}
	return epochTime
}

// SetLastEpochTime set last epoch time
func (k Keeper) SetLastEpochTime(ctx sdk.Context, epochTime time.Time) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.EpochTimeKey, sdk.FormatTimeBytes(epochTime))
}

// GetLastHalvenEpochNum returns last halven epoch number
func (k Keeper) GetLastHalvenEpochNum(ctx sdk.Context) int64 {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.LastHalvenEpochKey)
	if b == nil {
		return 0
	}

	return int64(sdk.BigEndianToUint64(b))
}

// SetLastHalvenEpochNum set last halven epoch number
func (k Keeper) SetLastHalvenEpochNum(ctx sdk.Context, epochNum int64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.LastHalvenEpochKey, sdk.Uint64ToBigEndian(uint64(epochNum)))
}

// get the minter
func (k Keeper) GetMinter(ctx sdk.Context) (minter types.Minter) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.MinterKey)
	if b == nil {
		panic("stored minter should not have been nil")
	}

	k.cdc.MustUnmarshalBinaryBare(b, &minter)
	return
}

// set the minter
func (k Keeper) SetMinter(ctx sdk.Context, minter types.Minter) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalBinaryBare(&minter)
	store.Set(types.MinterKey, b)
}

//______________________________________________________________________

// GetParams returns the total set of minting parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the total set of minting parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

//______________________________________________________________________

// MintCoins implements an alias call to the underlying supply keeper's
// MintCoins to be used in BeginBlocker.
func (k Keeper) MintCoins(ctx sdk.Context, newCoins sdk.Coins) error {
	if newCoins.Empty() {
		// skip as no coins need to be minted
		return nil
	}

	return k.bankKeeper.MintCoins(ctx, types.ModuleName, newCoins)
}

// GetPoolAllocatableAsset gets the balance of the `MintedDenom` from fees and returns coins according to the `AllocationRatio`
func (k Keeper) GetProportions(ctx sdk.Context, fees sdk.Coins, ratio sdk.Dec) sdk.Coin {
	params := k.GetParams(ctx)
	amount := fees.AmountOf(params.MintDenom)
	return sdk.NewCoin(params.MintDenom, amount.ToDec().Mul(ratio).TruncateInt())
}

// DistributeMintedCoins implements distribution of minted coins from mint to external modules.
func (k Keeper) DistributeMintedCoins(ctx sdk.Context, mintedCoins sdk.Coins) error {
	proportions := k.GetParams(ctx).DistributionProportions

	// allocate staking incentives into fee collector account to be moved to on next begin blocker by staking module
	stakingIncentivesCoins := sdk.NewCoins(k.GetProportions(ctx, mintedCoins, proportions.Staking))
	err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, k.feeCollectorName, stakingIncentivesCoins)
	if err != nil {
		return err
	}

	// allocate pool allocation ratio to pool-incentives module account account
	poolIncentivesCoins := sdk.NewCoins(k.GetProportions(ctx, mintedCoins, proportions.PoolIncentives))
	err = k.bankKeeper.SendCoinsFromModuleToModule(ctx, k.feeCollectorName, poolincentivestypes.ModuleName, poolIncentivesCoins)
	if err != nil {
		return err
	}

	// allocate developer rewards to an address, for now empty address, TODO: update it
	devRewardCoins := sdk.NewCoins(k.GetProportions(ctx, mintedCoins, proportions.DeveloperRewards))
	devRewardsAddr := sdk.AccAddress{}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, k.feeCollectorName, devRewardsAddr, devRewardCoins)
	if err != nil {
		return err
	}

	fmt.Println("FeeCollectorName", k.feeCollectorName)
	fmt.Println("mintedCoins", mintedCoins)
	fmt.Println("stakingIncentivesCoins", stakingIncentivesCoins)
	fmt.Println("poolIncentivesCoins", poolIncentivesCoins)
	fmt.Println("devRewardCoins", devRewardCoins)

	// call an hook after deposition of coin into fee pool
	k.hooks.AfterDistributeMintedCoins(ctx, mintedCoins)

	return err
}
