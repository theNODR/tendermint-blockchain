package handlers

import (
	"common"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"
	"svcnodr/logger"
	"svcnodr/msgs"
	"svcnodr/store"
)

type CloseTransferResponse struct{}

func HandleCloseTransferMsg(ctx types.Context, keeper store.Keeper, cdc *codec.Codec, msg msgs.MsgCloseTransfer) types.Result {
	common.Log.Event(logger.EventCloseTransferMsgHandle, common.Printf("Handle %s", msg.Type()))
	fmt.Printf("\nClose transfer %s\n", msg.ChannelId)

	state, err := keeper.GetTransfer(ctx, msg.ChannelId)
	if err != nil {
		fmt.Printf("Close transfer. No such transfer channel %s ", msg.ChannelId)
		return err.Result()
	}

	hasCoins := keeper.CoinKeeper.HasCoins(ctx, state.FromAddress, msg.Amount)
	if !hasCoins {
		return types.ErrInsufficientCoins(fmt.Sprintf("Insufficient coins on %s", state.FromAddress.String())).Result()
	}

	_, err = keeper.CoinKeeper.SendCoins(ctx, state.FromAddress, state.ToAddress, msg.Amount)
	if err != nil {
		return err.Result()
	}

	err = keeper.DeleteChannelIdFromIncome(ctx, msg.ToAddress, state.ChannelId)
	if err != nil {
		return err.Result()
	}

	err = keeper.DeleteChannelIdFromSpend(ctx, msg.FromAddress, state.ChannelId)
	if err != nil {
		return err.Result()
	}
	return types.Result{}
}
