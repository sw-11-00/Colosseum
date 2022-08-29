package keeper_test

import (
	"context"
	"testing"

	keepertest "Colosseum/testutil/keeper"
	"Colosseum/x/colosseum/keeper"
	"Colosseum/x/colosseum/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ColosseumKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
