package keeper

import (
	"Colosseum/x/colosseum/types"
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strings"
	"time"
)

func (k msgServer) CreateCoinSymbol(goCtx context.Context, msg *types.MsgCreateCoinSymbol) (*types.MsgCreateCoinSymbolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var coinSymbol = types.CoinSymbol{
		Creator:   msg.Creator,
		Symbol:    strings.ToUpper(msg.Symbol),
		CreatedAt: int32(time.Now().Unix()),
		UpdatedAt: int32(time.Now().Unix()),
		IsDeleted: false,
	}

	list := k.GetAllCoinSymbol(ctx)
	for _, v := range list {
		if v.Symbol == coinSymbol.Symbol {
			return &types.MsgCreateCoinSymbolResponse{
				Id: 999999999999,
			}, sdkerrors.Wrapf(types.ErrCoinSymbolExists, "coin symbol already exists")
		}
	}

	id := k.AppendCoinSymbol(
		ctx,
		coinSymbol,
	)

	return &types.MsgCreateCoinSymbolResponse{
		Id: id,
	}, nil
}
