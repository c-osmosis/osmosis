package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	yaml "gopkg.in/yaml.v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestPoolAccountMarshalYAML(t *testing.T) {
	var poolId uint64 = 10

	ymlAssetTest := []PoolAsset{
		{
			Weight: sdk.NewInt(200),
			Token:  sdk.NewCoin("test2", sdk.NewInt(50000)),
		},
		{
			Weight: sdk.NewInt(100),
			Token:  sdk.NewCoin("test1", sdk.NewInt(10000)),
		},
	}
	pacc, err := NewPoolAccount(poolId, PoolParams{
		SwapFee: defaultSwapFee,
		ExitFee: defaultExitFee,
	}, ymlAssetTest, defaultFutureGovernor, defaultCurBlockTime)
	require.NoError(t, err)

	bs, err := yaml.Marshal(pacc)
	require.NoError(t, err)

	want := `|
  address: cosmos1m48tfmd0e6yqgfhraxl9ddt7lygpsnsrhtwpas
  public_key: ""
  account_number: 0
  sequence: 0
  id: 10
  pool_params:
    swap_fee: "0.025000000000000000"
    exit_fee: "0.025000000000000000"
    smooth_weight_change_params: null
  future_pool_governor: ""
  total_weight: "300.000000000000000000"
  total_share:
    denom: gamm/pool/10
    amount: "0"
  pool_assets:
  - |
    token:
      denom: test1
      amount: "10000"
    weight: "100.000000000000000000"
  - |
    token:
      denom: test2
      amount: "50000"
    weight: "200.000000000000000000"
`
	require.Equal(t, want, string(bs))
}

func TestPoolAccountJson(t *testing.T) {
	var poolId uint64 = 10

	jsonAssetTest := []PoolAsset{
		{
			Weight: sdk.NewInt(200),
			Token:  sdk.NewCoin("test2", sdk.NewInt(50000)),
		},
		{
			Weight: sdk.NewInt(100),
			Token:  sdk.NewCoin("test1", sdk.NewInt(10000)),
		},
	}
	pacc, err := NewPoolAccount(poolId, PoolParams{
		SwapFee: defaultSwapFee,
		ExitFee: defaultExitFee,
	}, jsonAssetTest, defaultFutureGovernor, defaultCurBlockTime)
	require.NoError(t, err)

	paccInternal := pacc.(*PoolAccount)

	bz, err := json.Marshal(pacc)
	require.NoError(t, err)

	bz1, err := paccInternal.MarshalJSON()
	require.NoError(t, err)
	require.Equal(t, string(bz1), string(bz))

	var a PoolAccount
	require.NoError(t, json.Unmarshal(bz, &a))
	require.Equal(t, pacc.String(), a.String())
}