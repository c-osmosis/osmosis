package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	gogotypes "github.com/gogo/protobuf/types"

	"github.com/osmosis-labs/osmosis/x/gamm/types"
)

func (k Keeper) GetPool(ctx sdk.Context, poolId uint64) (types.PoolAccountI, error) {
	acc := k.accountKeeper.GetAccount(ctx, types.NewPoolAddress(poolId))
	if acc == nil {
		return nil, sdkerrors.Wrapf(types.ErrPoolNotFound, "pool %d does not exist", poolId)
	}

	poolAcc, ok := acc.(types.PoolAccountI)
	if !ok {
		return nil, sdkerrors.Wrapf(types.ErrPoolNotFound, "pool %d does not exist", poolId)
	}

	poolAcc.PokeTokenWeights(ctx.BlockTime())

	return poolAcc, nil
}

func (k Keeper) SetPool(ctx sdk.Context, poolAcc types.PoolAccountI) error {
	// Make sure that pool exists
	_, err := k.GetPool(ctx, poolAcc.GetId())
	if err != nil {
		return err
	}

	k.accountKeeper.SetAccount(ctx, poolAcc)
	return nil
}

// newPool is an internal function that creates a new Pool object with the provided
// parameters, initial assets, and future governor.
func (k Keeper) newPool(ctx sdk.Context, poolParams types.PoolParams, assets []types.PoolAsset, futureGovernor string) (types.PoolAccountI, error) {
	poolId := k.getNextPoolNumber(ctx)

	acc := k.accountKeeper.GetAccount(ctx, types.NewPoolAddress(poolId))
	if acc != nil {
		return nil, sdkerrors.Wrapf(types.ErrPoolAlreadyExist, "pool %d already exist", poolId)
	}

	poolAcc, err := types.NewPoolAccount(poolId, poolParams, assets, futureGovernor, ctx.BlockTime())
	if err != nil {
		return nil, err
	}
	poolAcc = k.accountKeeper.NewAccount(ctx, poolAcc).(types.PoolAccountI)

	k.accountKeeper.SetAccount(ctx, poolAcc)

	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetKeyPaginationPoolNumbers(poolId), sdk.Uint64ToBigEndian(poolId))

	return poolAcc, nil
}

// getNextPoolNumber returns the next pool number
func (k Keeper) getNextPoolNumber(ctx sdk.Context) uint64 {
	var poolNumber uint64
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.GlobalPoolNumber)
	if bz == nil {
		// initialize the account numbers
		poolNumber = 1
	} else {
		val := gogotypes.UInt64Value{}

		err := k.cdc.UnmarshalBinaryBare(bz, &val)
		if err != nil {
			panic(err)
		}

		poolNumber = val.GetValue()
	}

	bz = k.cdc.MustMarshalBinaryBare(&gogotypes.UInt64Value{Value: poolNumber + 1})
	store.Set(types.GlobalPoolNumber, bz)

	return poolNumber
}

func (k Keeper) getPoolAndInOutAssets(
	ctx sdk.Context, poolId uint64,
	tokenInDenom string,
	tokenOutDenom string) (
	pool types.PoolAccountI,
	inAsset types.PoolAsset,
	outAsset types.PoolAsset,
	err error,
) {
	pool, err = k.GetPool(ctx, poolId)
	if err != nil {
		return
	}

	inAsset, err = pool.GetPoolAsset(tokenInDenom)
	if err != nil {
		return
	}

	outAsset, err = pool.GetPoolAsset(tokenOutDenom)
	return
}
