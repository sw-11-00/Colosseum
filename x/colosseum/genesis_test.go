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

		CoinSymbolList: []types.CoinSymbol{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		CoinSymbolCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ColosseumKeeper(t)
	colosseum.InitGenesis(ctx, *k, genesisState)
	got := colosseum.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.CoinSymbolList, got.CoinSymbolList)
	require.Equal(t, genesisState.CoinSymbolCount, got.CoinSymbolCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
