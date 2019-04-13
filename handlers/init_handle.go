package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"svcnodr/msgs"
	"svcnodr/store"
)

func HandleInitMsg(ctx sdk.Context, keeper store.Keeper, msg msgs.MsgInit) sdk.Result {

	if msg.To.Empty() {
		return sdk.ErrInvalidAddress("Init To address invalid").Result()
	}

	_, _, err := keeper.CoinKeeper.AddCoins(ctx, msg.To, msg.Amount)

	if err != nil {
		return err.Result()
	} else {
		return sdk.Result{}
	}
}
