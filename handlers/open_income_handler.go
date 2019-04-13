package handlers

import (
	"common"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"
	"svcnodr/helpers"
	"svcnodr/logger"
	"svcnodr/msgs"
	"svcnodr/store"
	ntypes "svcnodr/types"
	states "svcnodr/types"
)

func HandleOpenIncomeMsg(ctx types.Context, keeper store.Keeper, cdc *codec.Codec, msg msgs.MsgOpenIncome) types.Result {
	common.Log.Event(logger.EventOpenIncomeMsgHandle, common.Printf("Handle: %", msg.Type()))

	err := msg.ValidateBasic()
	if err != nil {
		return err.Result()
	}

	address := createIncomeAddress(msg)
	timestamp := common.GetNowUnixMs()
	timelock := common.GetNowUnixMs() + msg.LifeTime
	state := states.NewIncomeChannelState(address, timestamp, msg.PriceAmount, msg.PriceQuantumPower, timelock)
	err = keeper.SetIncome(ctx, state)
	if err != nil {
		return err.Result()
	}

	response := &ntypes.IncomeChannelStateResponse{
		Address:  string(state.Address.Bytes()),
		Current:  state.State.Fact.ToChannel(),
		Limit:    state.State.Plan.ToChannel(),
		Price:    state.Price.ToChannel(),
		TimeLock: state.TimeLock,
		LifeTime: state.LifeTime,
	}

	stateResponse := cdc.MustMarshalJSON(response)
	return types.Result{Data: stateResponse}
}

func createIncomeAddress(msg msgs.MsgOpenIncome) types.AccAddress {
	timestamp := common.GetNowUnixMs()
	addressStr := helpers.CreateIncomeAddress(msg.TrackerPublicKey, msg.PeerPublicKey, timestamp)
	address := types.AccAddress(addressStr)
	return address
}
