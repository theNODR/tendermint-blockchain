package handlers

import (
	"common"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"
	"svcnodr/logger"
	"svcnodr/msgs"
	"svcnodr/store"
	states "svcnodr/types"
)

type CloseIncomeResponse struct {
	State *states.ChannelFact `json:"state"`
}

func HandleCloseIncomeMsg(ctx types.Context, keeper store.Keeper, cdc *codec.Codec, msg msgs.MsgCloseIncome) types.Result {
	common.Log.Event(logger.EventCloseIncomeMsgHandle, common.Printf("Handle CloseIncomeMsg"))

	state, err := keeper.GetIncome(ctx, msg.From)
	if err != nil {
		return err.Result()
	}

	if state.HasChannels() {
		return types.ErrInternal(fmt.Sprintf("Income %s has open transfer channesl:", state.Address)).Result()
	}

	coins := keeper.CoinKeeper.GetCoins(ctx, state.Address)
	_, err = keeper.CoinKeeper.SendCoins(ctx, state.Address, msg.To, coins)
	if err != nil {
		return err.Result()
	}

	closeResponse := &CloseIncomeResponse{State: state.State.Fact.ToChannel()}
	bytes := cdc.MustMarshalJSON(closeResponse)
	return types.Result{Data: bytes}
}
