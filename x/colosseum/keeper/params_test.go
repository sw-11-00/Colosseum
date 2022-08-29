package keeper_test

import (
	"testing"

	testkeeper "Colosseum/testutil/keeper"
	"Colosseum/x/colosseum/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ColosseumKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
