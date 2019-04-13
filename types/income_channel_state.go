package types

import (
	"common"
	"github.com/cosmos/cosmos-sdk/types"
)

type IncomeChannelState struct {
	*BaseChannelState
	ChannelsIds
}

func NewIncomeChannelState(address types.AccAddress, timestamp int64, priceAmount types.Coins, priceQuantumPower QuantumPowerType, timelock int64) *IncomeChannelState {
	return &IncomeChannelState{
		&BaseChannelState{
			Address:  address,
			Price:    NewPrice(priceAmount, priceQuantumPower),
			State:    NewSpecialChannelState(priceAmount, 0),
			TimeLock: timelock,
			LifeTime: timelock + common.GetNowUnixMs(),
		},
		NewChannelsIds(),
	}
}

//
//func (s *IncomeChannelState) MinPrice() IPrice {
//	return s.IPrice
//}
//
//func (s *IncomeChannelState) AddPlan(tran *LedgerTransaction) error {
//	return s.State.IncreasePlan(tran.Id, tran.PlannedAmount(), tran.PlannedVolume())
//}
//
//func (s *IncomeChannelState) AddFact(tran *LedgerTransaction) error {
//	Amount := tran.Amount
//	quantumVolume := tran.QuantumCount * tran.PriceQuantumPower.Volume()
//	volume := tran.Volume
//	return s.State.IncreaseFact(tran.ParentId, Amount, quantumVolume, volume)
//}
