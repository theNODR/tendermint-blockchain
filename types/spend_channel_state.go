package types

import (
	"common"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type SpendChannelState struct {
	*BaseChannelState
	ChannelsIds
	TestCh ChannelsIds
}

func NewSpendChannelState(addres sdk.AccAddress, priceAmount sdk.Coins, maxAmount sdk.Coins, priceQuantumPower QuantumPowerType, timeLock int64) *SpendChannelState {
	return &SpendChannelState{
		&BaseChannelState{
			Address:  addres,
			Price:    NewPrice(priceAmount, priceQuantumPower),
			State:    NewSpecialChannelState(maxAmount, 0),
			TimeLock: timeLock,
			LifeTime: timeLock - common.GetNowUnixMs(),
		},
		NewChannelsIds(),
		NewChannelsIds(),
	}
}

//
//func (s *SpendChannelState) MaxPrice() IPrice {
//	return s.IPrice
//}
//
//func (s *SpendChannelState) AddPlan(tran *LedgerTransaction) error {
//	return s.State.DecreasePlan(tran.Id, tran.PlannedAmount(), tran.PlannedVolume())
//}
//
//func (s *SpendChannelState) AddFact(tran *LedgerTransaction) error {
//	Amount := tran.Amount
//	quantumVolume := tran.QuantumCount * tran.PriceQuantumPower.Volume()
//	volume := tran.Volume
//
//	return s.State.DecreaseFact(tran.ParentId, Amount, quantumVolume, volume)
//}
