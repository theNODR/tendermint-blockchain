package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
)

type SpecialChannelState struct {
	planned map[string]*PlanChannelState
	Fact    *FactChannelState
	Plan    *PlanChannelState
}

func NewSpecialChannelState(
	amount sdk.Coins,
	quantumVolume uint64,
) *SpecialChannelState {
	return &SpecialChannelState{
		planned: make(map[string]*PlanChannelState),
		Fact:    NewFactPlanChannelState(amount, quantumVolume, quantumVolume),
		Plan:    NewPlanChannelState(amount, quantumVolume),
	}
}

func (a *SpecialChannelState) IncreaseFact(
	id string,
	amount sdk.Coins,
	quantumVolume uint64,
	volume uint64,
) error {
	state, ok := a.planned[id]

	if !ok {
		return errors.New("SpecialChannelState: IncreaseFact: inconsistent tran log")
	}

	plan := a.Plan
	fact := a.Fact

	fact.Amount.Add(amount)
	fact.QuantumVolume += quantumVolume
	fact.Volume += volume

	plan.Amount.Sub(state.Amount.Sub(amount))
	plan.QuantumVolume -= state.QuantumVolume - quantumVolume

	delete(a.planned, id)

	return nil
}

func (a *SpecialChannelState) DecreaseFact(
	id string,
	amount sdk.Coins,
	quantumVolume uint64,
	volume uint64,
) error {
	state, ok := a.planned[id]

	if !ok {
		return errors.New("SpecialChannelState: DecreaseFact: inconsistent tran log")
	}

	plan := a.Plan
	fact := a.Fact

	fact.Amount.Sub(amount)
	fact.QuantumVolume -= quantumVolume
	fact.Volume -= volume

	plan.Amount.Add(state.Amount.Sub(amount))
	plan.QuantumVolume += state.QuantumVolume - quantumVolume

	delete(a.planned, id)

	return nil
}

func (a *SpecialChannelState) IncreasePlan(
	id string,
	amount sdk.Coins,
	quantumVolume uint64,
) error {
	_, ok := a.planned[id]

	if ok {
		return errors.New("SpecialChannelState: IncreasePlan: inconsistent tran log")
	}

	a.planned[id] = NewPlanChannelState(amount, quantumVolume)

	a.Plan.Amount.Add(amount)
	a.Plan.QuantumVolume += quantumVolume

	return nil
}

func (a *SpecialChannelState) DecreasePlan(
	id string,
	amount sdk.Coins,
	quantumVolume uint64,
) error {
	_, ok := a.planned[id]

	if ok {
		return errors.New("SpecialChannelState: DecreasePlan: inconsistent tran log")
	}

	a.planned[id] = NewPlanChannelState(amount, quantumVolume)
	a.Plan.Amount.Sub(amount)
	a.Plan.QuantumVolume -= quantumVolume

	return nil
}
