package keeper

import (
	"time"

	"github.com/c-osmosis/osmosis/x/incentives/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) iteratorAfterTime(ctx sdk.Context, prefix []byte, time time.Time) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	timeKey := getTimeKey(time)
	key := combineKeys(prefix, timeKey)
	return store.Iterator(storetypes.InclusiveEndBytes(key), storetypes.PrefixEndBytes(prefix))
}

func (k Keeper) iteratorBeforeTime(ctx sdk.Context, prefix []byte, time time.Time) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	timeKey := getTimeKey(time)
	key := combineKeys(prefix, timeKey)
	return store.Iterator(prefix, storetypes.InclusiveEndBytes(key))
}

func (k Keeper) iterator(ctx sdk.Context, prefix []byte) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, prefix)
}

// UpcomingPotsIteratorAfterTime returns the iterator to get upcoming pots that start distribution after specific time
func (k Keeper) UpcomingPotsIteratorAfterTime(ctx sdk.Context, time time.Time) sdk.Iterator {
	return k.iteratorAfterTime(ctx, types.KeyPrefixUpcomingPots, time)
}

// UpcomingPotsIteratorBeforeTime returns the iterator to get upcoming pots that already started distribution before specific time
func (k Keeper) UpcomingPotsIteratorBeforeTime(ctx sdk.Context, time time.Time) sdk.Iterator {
	return k.iteratorBeforeTime(ctx, types.KeyPrefixUpcomingPots, time)
}

// UpcomingPotsIterator returns iterator for upcoming pots
func (k Keeper) UpcomingPotsIterator(ctx sdk.Context) sdk.Iterator {
	return k.iterator(ctx, types.KeyPrefixUpcomingPots)
}

// ActivePotsIterator returns iterator for active pots
func (k Keeper) ActivePotsIterator(ctx sdk.Context) sdk.Iterator {
	return k.iterator(ctx, types.KeyPrefixActivePots)
}

// FinishedPotsIterator returns iterator for finished pots
func (k Keeper) FinishedPotsIterator(ctx sdk.Context) sdk.Iterator {
	return k.iterator(ctx, types.KeyPrefixFinishedPots)
}
