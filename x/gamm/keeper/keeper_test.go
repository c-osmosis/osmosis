package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/tendermint/tendermint/crypto/ed25519"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/c-osmosis/osmosis/app"

	"github.com/c-osmosis/osmosis/x/gamm/types"
)

type KeeperTestSuite struct {
	suite.Suite

	app         *app.OsmosisApp
	ctx         sdk.Context
	queryClient types.QueryClient
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.app = app.Setup(false)
	suite.ctx = suite.app.BaseApp.NewContext(false, tmproto.Header{})

	queryHelper := baseapp.NewQueryServerTestHelper(suite.ctx, suite.app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, suite.app.GAMMKeeper)
	suite.queryClient = types.NewQueryClient(queryHelper)
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

var (
	acc1 = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	acc2 = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	acc3 = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
)

func (suite *KeeperTestSuite) preparePoolWithPoolParams(poolParams types.PoolParams) uint64 {
	// Mint some assets to the accounts.
	for _, acc := range []sdk.AccAddress{acc1, acc2, acc3} {
		err := suite.app.BankKeeper.AddCoins(
			suite.ctx,
			acc,
			sdk.NewCoins(
				sdk.NewCoin("foo", sdk.NewInt(10000000)),
				sdk.NewCoin("bar", sdk.NewInt(10000000)),
				sdk.NewCoin("baz", sdk.NewInt(10000000)),
			),
		)
		if err != nil {
			panic(err)
		}
	}

	poolId, err := suite.app.GAMMKeeper.CreatePool(suite.ctx, acc1, poolParams, []types.PoolAsset{
		{
			Weight: sdk.NewInt(100),
			Token:  sdk.NewCoin("foo", sdk.NewInt(5000000)),
		},
		{
			Weight: sdk.NewInt(200),
			Token:  sdk.NewCoin("bar", sdk.NewInt(5000000)),
		},
		{
			Weight: sdk.NewInt(300),
			Token:  sdk.NewCoin("baz", sdk.NewInt(5000000)),
		},
	})
	suite.NoError(err)
	return poolId
}

func (suite *KeeperTestSuite) preparePool() uint64 {
	poolId := suite.preparePoolWithPoolParams(types.PoolParams{
		Lock:    false,
		SwapFee: sdk.NewDec(0),
		ExitFee: sdk.NewDec(0),
	})

	spotPrice, err := suite.app.GAMMKeeper.CalculateSpotPrice(suite.ctx, poolId, "foo", "bar")
	suite.NoError(err)
	suite.Equal(sdk.NewDec(2).String(), spotPrice.String())
	spotPrice, err = suite.app.GAMMKeeper.CalculateSpotPrice(suite.ctx, poolId, "bar", "baz")
	suite.NoError(err)
	suite.Equal(sdk.NewDecWithPrec(15, 1).String(), spotPrice.String())
	spotPrice, err = suite.app.GAMMKeeper.CalculateSpotPrice(suite.ctx, poolId, "baz", "foo")
	suite.NoError(err)
	suite.Equal(sdk.NewDec(1).Quo(sdk.NewDec(3)).String(), spotPrice.String())

	return poolId
}