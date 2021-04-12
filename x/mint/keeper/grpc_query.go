package keeper

import (
	"context"

	"github.com/c-osmosis/osmosis/x/mint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.QueryServer = Keeper{}

// Params returns params of the mint module.
func (k Keeper) Params(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}

// EpochProvisions returns minter.EpochProvisions of the mint module.
func (k Keeper) EpochProvisions(c context.Context, _ *types.QueryEpochProvisionsRequest) (*types.QueryEpochProvisionsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	minter := k.GetMinter(ctx)

	return &types.QueryEpochProvisionsResponse{EpochProvisions: minter.EpochProvisions}, nil
}
