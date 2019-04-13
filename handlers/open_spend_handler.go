package handlers

import (
	"common"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"svcnodr/helpers"
	"svcnodr/logger"
	"svcnodr/msgs"
	"svcnodr/store"
	ntypes "svcnodr/types"
	states "svcnodr/types"
)

func HandleOpenSpendMsg(ctx sdk.Context, keeper store.Keeper, cdc *codec.Codec, msg msgs.MsgOpenSpend) sdk.Result {
	common.Log.Event(logger.EventOpenSpendMsgHandle, common.Printf("Handle: %", msg.Type()))

	err := msg.ValidateBasic()
	if err != nil {
		return err.Result()
	}

	address := createAddress(msg)
	timelock := common.GetNowUnixMs() + msg.LifeTime
	state := states.NewSpendChannelState(address, msg.PriceAmount, msg.MaxAmount, states.QuantumPowerType(msg.PriceQuantumPower), timelock)
	if state == nil {
		return sdk.ErrInternal("Can`t create open spend channel").Result()
	}

	err = keeper.SetSpend(ctx, state)
	if err != nil {
		return err.Result()
	}

	_, err = keeper.CoinKeeper.SendCoins(ctx, msg.From, state.Address, msg.MaxAmount)
	if err != nil {
		return err.Result()
	}
	response := &ntypes.SpendChannelStateResponse{
		Address:  string(state.Address.Bytes()),
		Current:  state.State.Fact.ToChannel(),
		Limit:    state.State.Plan.ToChannel(),
		Price:    state.Price.ToChannel(),
		TimeLock: state.TimeLock,
		LifeTime: state.LifeTime,
	}

	responseBytes := cdc.MustMarshalJSON(response)
	return sdk.Result{Data: responseBytes}
}

func createAddress(msg msgs.MsgOpenSpend) sdk.AccAddress {
	addressString := helpers.CreateSpendAddress(msg.TrackerPublicKey, msg.PeerPublicKey, common.GetNowUnixMs())
	return sdk.AccAddress(addressString)
}
