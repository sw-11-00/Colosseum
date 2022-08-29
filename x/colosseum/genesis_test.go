package colosseum_test

import (
	"testing"

	keepertest "Colosseum/testutil/keeper"
	"Colosseum/testutil/nullify"
	"Colosseum/x/colosseum"
	"Colosseum/x/colosseum/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ColosseumKeeper(t)
	colosseum.InitGenesis(ctx, *k, genesisState)
	got := colosseum.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
