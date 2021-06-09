package simulation

// DONTCOVER

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/osmosis-labs/osmosis/x/incentives/types"
)

// Simulation parameter constants
const (
	ParamsDistrEpochIdentifier       = "distr_epoch_identifier"
)

// RandomizedGenState generates a random GenesisState for gov
func RandomizedGenState(simState *module.SimulationState) {
	var distrEpochIdentifier string
	simState.AppParams.GetOrGenerate(
		simState.Cdc, ParamsDistrEpochIdentifier, &distrEpochIdentifier, simState.Rand,
		func(r *rand.Rand) { distrEpochIdentifier = GenParamsDistrEpochIdentifier(r) },
	)

	incentivesGenesis := &types.GenesisState{
		Params: types.Params {
			DistrEpochIdentifier: distrEpochIdentifier,
		},
		// Gauges: gauges,
	}

	bz, err := json.MarshalIndent(&incentivesGenesis, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated governance parameters:\n%s\n", bz)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(incentivesGenesis)
}
