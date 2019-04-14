package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type ChannelState struct {
	Amount        sdk.Coins
	QuantumVolume uint64
}

type PlanChannelState struct {
	*ChannelState
}

type ChannelPlan struct {
	Amount        int64  `json:"amount"`
	QuantumVolume uint64 `json:"quantumVolume"`
}

func NewPlanChannelState(amount sdk.Coins, quantumVolume uint64) *PlanChannelState {
	return &PlanChannelState{
		&ChannelState{
			Amount:        amount,
			QuantumVolume: quantumVolume,
		},
	}
}

func (s *PlanChannelState) ToChannel() *ChannelPlan {
	return &ChannelPlan{
		Amount:        s.Amount.AmountOf(denom).Int64(),
		QuantumVolume: s.QuantumVolume,
	}
}

type FactChannelState struct {
	*ChannelState
	Volume uint64
}

type ChannelFact struct {
	Amount        int64  `json:"amount"`
	QuantumVolume uint64 `json:"quantumVolume"`
	Volume        uint64 `json:"volume"`
}

func NewFactPlanChannelState(amount sdk.Coins, quantumVolume uint64, volume uint64) *FactChannelState {
	return &FactChannelState{
		&ChannelState{
			Amount:        amount,
			QuantumVolume: quantumVolume,
		},
		volume,
	}
}

func (s *FactChannelState) ToChannel() *ChannelFact {
	return &ChannelFact{
		Amount:        s.Amount.AmountOf(denom).Int64(),
		QuantumVolume: s.QuantumVolume,
		Volume:        s.Volume,
	}
}
