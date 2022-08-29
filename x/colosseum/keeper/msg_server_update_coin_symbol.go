package keeper

import (
	"context"
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strings"
	"time"

	"Colosseum/x/colosseum/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateCoinSymbol(goCtx context.Context, msg *types.MsgUpdateCoinSymbol) (*types.MsgUpdateCoinSymbolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetCoinSymbol(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	var coinSymbol = types.CoinSymbol{
		Creator:   msg.Creator,
		Id:        msg.Id,
		Symbol:    strings.ToUpper(msg.Symbol),
		CreatedAt: val.CreatedAt,
		UpdatedAt: int32(time.Now().Unix()),
		IsDeleted: false,
	}

	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetCoinSymbol(ctx, coinSymbol)

	return &types.MsgUpdateCoinSymbolResponse{}, nil
}
