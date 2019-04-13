package handlers

import (
	"common"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"svcnodr/logger"
	"svcnodr/msgs"
	"svcnodr/store"
	states "svcnodr/types"
)

type OpenTransferResponse struct {
	TimeLock int64 `json:"timelock"`
	LifeTime int64 `json:"lifetime"`
}

func HandleOpenTransferMsg(ctx sdk.Context, keeper store.Keeper, cdc *codec.Codec, msg msgs.MsgOpenTransfer) sdk.Result {
	common.Log.Event(logger.EventOpenTransferMsghandle, common.Printf("Handle: %", msg.Type()))
	fmt.Printf("\nOpen transfer %s\n", msg.ChannelId)

	if !keeper.HasIncome(ctx, msg.ToAddress) {
		return sdk.ErrInvalidAddress(fmt.Sprintf("No such income channel: %s", msg.ToAddress.String())).Result()
	}

	if !keeper.HasSpend(ctx, msg.FromAddress) {
		return sdk.ErrInvalidAddress(fmt.Sprintf("No such spend channel:%s", msg.FromAddress.String())).Result()
	}

	state := states.NewTransferChannelState(msg.ChannelId,
		msg.FromAddress,
		msg.FromPk,
		msg.ToAddress,
		msg.ToPk,
		msg.PriceAmount,
		msg.PriceQuantumPower,
		msg.PlannedQuantumCount,
		msg.Lifetime)

	err := keeper.AddTransfer(ctx, state)
	if err != nil {
		err.Result()
	}

	err = keeper.AddChannelIdToSpend(ctx, msg.FromAddress, state.ChannelId)
	if err != nil {
		return err.Result()
	}

	err = keeper.AddChannelIdToIncome(ctx, msg.ToAddress, state.ChannelId)
	if err != nil {
		return err.Result()
	}

	iter := keeper.GetTransferIterator(ctx)
	for ; iter.Valid(); iter.Next() {
		fmt.Printf("\n%s", iter.Key())
	}
	fmt.Printf("\n")

	response := &OpenTransferResponse{state.TimeLock, msg.Lifetime}
	bytesResponse := cdc.MustMarshalJSON(response)
	return sdk.Result{Data: bytesResponse}
}
