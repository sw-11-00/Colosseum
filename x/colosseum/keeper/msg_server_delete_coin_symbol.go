package keeper

import (
	"context"
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"

	"Colosseum/x/colosseum/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) DeleteCoinSymbol(goCtx context.Context, msg *types.MsgDeleteCoinSymbol) (*types.MsgDeleteCoinSymbolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetCoinSymbol(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	var coinSymbol = types.CoinSymbol{
		Creator:   msg.Creator,
		Id:        msg.Id,
		CreatedAt: val.CreatedAt,
		UpdatedAt: int32(time.Now().Unix()),
		IsDeleted: !val.IsDeleted,
	}

	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetCoinSymbol(ctx, coinSymbol)

	return &types.MsgDeleteCoinSymbolResponse{}, nil
}
