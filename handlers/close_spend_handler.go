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

type CloseSpendResponse struct {
	State *states.ChannelFact `json:"state"`
}

func HandleCloseSpendMsg(ctx types.Context, keeper store.Keeper, cdc *codec.Codec, msg msgs.MsgCloseSpend) types.Result {
	common.Log.Event(logger.EventCloseSpendMsgHandle, common.Printf("Handle %s", msg.Type()))

	state, err := keeper.GetSpend(ctx, msg.From)
	if err != nil {
		return err.Result()
	}

	if state.HasChannels() {
		return types.ErrInternal(fmt.Sprintf("Spend channel %s has opened transfer channels:", state.Address.String())).Result()
	}

	coins := keeper.CoinKeeper.GetCoins(ctx, state.Address)
	_, err = keeper.CoinKeeper.SendCoins(ctx, state.Address, msg.To, coins)
	if err != nil {
		return err.Result()
	}

	closeResponse := &CloseSpendResponse{State: state.State.Fact.ToChannel()}
	bytes := cdc.MustMarshalJSON(closeResponse)
	return types.Result{Data: bytes}
}
