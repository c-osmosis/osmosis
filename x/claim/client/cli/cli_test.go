package cli_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/osmosis-labs/osmosis/app"
	"github.com/osmosis-labs/osmosis/app/params"
	"github.com/osmosis-labs/osmosis/x/claim/client/cli"
	"github.com/osmosis-labs/osmosis/x/claim/types"
	claimtypes "github.com/osmosis-labs/osmosis/x/claim/types"
	"github.com/stretchr/testify/suite"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	dbm "github.com/tendermint/tm-db"
)

var addr1 sdk.AccAddress
var addr2 sdk.AccAddress

func init() {
	params.SetAddressPrefixes()
	addr1 = ed25519.GenPrivKey().PubKey().Address().Bytes()
	addr2 = ed25519.GenPrivKey().PubKey().Address().Bytes()
}

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func (s *IntegrationTestSuite) SetupSuite() {

	s.T().Log("setting up integration test suite")
	encCfg := app.MakeEncodingConfig()

	genState := app.ModuleBasics.DefaultGenesis(encCfg.Marshaler)
	claimGenState := claimtypes.DefaultGenesis()
	claimGenState.ModuleAccountBalance = sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(30))
	claimGenState.InitialClaimables = []banktypes.Balance{
		{
			Address: addr1.String(),
			Coins:   sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 10)),
		},
		{
			Address: addr2.String(),
			Coins:   sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 20)),
		},
	}
	claimGenStateBz := encCfg.Marshaler.MustMarshalJSON(claimGenState)
	genState[claimtypes.ModuleName] = claimGenStateBz

	s.cfg = network.Config{
		Codec:             encCfg.Marshaler,
		TxConfig:          encCfg.TxConfig,
		LegacyAmino:       encCfg.Amino,
		InterfaceRegistry: encCfg.InterfaceRegistry,
		AccountRetriever:  authtypes.AccountRetriever{},
		AppConstructor: func(val network.Validator) servertypes.Application {
			return app.NewOsmosisApp(
				val.Ctx.Logger, dbm.NewMemDB(), nil, true, make(map[int64]bool), val.Ctx.Config.RootDir, 0,
				encCfg,
				simapp.EmptyAppOptions{},
				baseapp.SetMinGasPrices(val.AppConfig.MinGasPrices),
			)
		},
		GenesisState:    genState,
		TimeoutCommit:   2 * time.Second,
		ChainID:         "osmosis-1",
		NumValidators:   1,
		BondDenom:       sdk.DefaultBondDenom,
		MinGasPrices:    fmt.Sprintf("0.000006%s", sdk.DefaultBondDenom),
		AccountTokens:   sdk.TokensFromConsensusPower(1000),
		StakingTokens:   sdk.TokensFromConsensusPower(500),
		BondedTokens:    sdk.TokensFromConsensusPower(100),
		PruningStrategy: storetypes.PruningOptionNothing,
		CleanupDir:      true,
		SigningAlgo:     string(hd.Secp256k1Type),
		KeyringOptions:  []keyring.Option{},
	}

	s.network = network.New(s.T(), s.cfg)

	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

// TODO: Make sure CLI replaces with genesis time of chain.

// func (s *IntegrationTestSuite) TestGetCmdQueryParams() {
// 	val := s.network.Validators[0]

// 	testCases := []struct {
// 		name           string
// 		args           []string
// 		expectedOutput string
// 	}{
// 		{
// 			"json output",
// 			[]string{fmt.Sprintf("--%s=1", flags.FlagHeight), fmt.Sprintf("--%s=json", tmcli.OutputFlag)},
// 			`{"airdrop_start":"1970-01-01T00:00:00Z","duration_until_decay":"3600s","duration_of_decay":"18000s"}`,
// 		},
// 		{
// 			"text output",
// 			[]string{fmt.Sprintf("--%s=1", flags.FlagHeight), fmt.Sprintf("--%s=text", tmcli.OutputFlag)},
// 			`airdrop_start: "1970-01-01T00:00:00Z"
// duration_of_decay: 18000s
// duration_until_decay: 3600s`,
// 		},
// 	}

// 	for _, tc := range testCases {
// 		tc := tc

// 		s.Run(tc.name, func() {
// 			cmd := cli.GetCmdQueryParams()
// 			clientCtx := val.ClientCtx

// 			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
// 			s.Require().NoError(err)
// 			s.Require().Equal(tc.expectedOutput, strings.TrimSpace(out.String()))
// 		})
// 	}
// }

func (s *IntegrationTestSuite) TestCmdQueryClaimable() {
	val := s.network.Validators[0]

	testCases := []struct {
		name  string
		args  []string
		coins sdk.Coins
	}{
		{
			"query claimable amount",
			[]string{
				addr2.String(),
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			sdk.Coins{sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(20))},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryClaimable()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			s.Require().NoError(err)

			var result types.ClaimableResponse
			s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), &result))
			s.Require().Equal(tc.coins.String(), sdk.Coins(result.Coins).String())
		})
	}
}

func (s *IntegrationTestSuite) TestCmdQueryActivities() {
	val := s.network.Validators[0]

	testCases := []struct {
		name string
		args []string
	}{
		{
			"query activities amount",
			[]string{
				addr2.String(),
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryActivities()
			clientCtx := val.ClientCtx

			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			s.Require().NoError(err)

			var result types.ActivitiesResponse
			s.Require().NoError(clientCtx.JSONMarshaler.UnmarshalJSON(out.Bytes(), &result))
			s.Require().Equal([]string{"ActionAddLiquidity", "ActionSwap", "ActionVote", "ActionDelegateStake"}, result.All)
			s.Require().Equal([]string(nil), result.Completed)
		})
	}
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
