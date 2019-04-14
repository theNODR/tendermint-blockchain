package types

import (
	"github.com/cosmos/cosmos-sdk/types"
)

const denom = "ndr"

type QuantumPowerType uint8

func (p QuantumPowerType) Volume() uint64 {
	return 1 << p
}

type IPrice interface {
	Cmp(IPrice) (int8, error)
	ToNewQuantumPower(QuantumPowerType) (IPrice, error)
	ToChannel() *ChannelPrice
}

type Price struct {
	Amount       types.Coins      `json:"amount"`
	QuantumPower QuantumPowerType `json:"quantumPower"`
}

type ChannelPrice struct {
	Amount       int64            `json:"amount"`
	QuantumPower QuantumPowerType `json:"quantumPower"`
}

func NewPrice(
	amount types.Coins,
	quantumPower QuantumPowerType,
) *Price {
	return &Price{amount, quantumPower}
}

// -1 if this less then another
// 0 if this equal another
// 1 if this greater than another
func (p *Price) Cmp(another Price) (int8, error) {
	//if another == nil {
	//	return 0, errors.New("another is nil")
	//}
	//
	//var first Price
	//var second Price
	//var err error = nil
	//
	//if another.QuantumPower > p.QuantumPower {
	//	first, err = p.ToNewQuantumPower(another.QuantumPower)
	//	second = another
	//} else if another.QuantumPower < p.QuantumPower {
	//	first = p
	//	second, err = another.ToNewQuantumPower(p.QuantumPower)
	//} else {
	//	first = p
	//	second = another
	//}
	//
	//if err != nil {
	//	return 0, err
	//}
	//
	//if first.Amount.IsAllGT(second.Amount) {
	//	return 1, nil
	//} else if first.Amount.IsAllLT(second.Amount) {
	//	return -1, nil
	//} else {
	//	return 0, nil
	//}
	return 0, nil
}

func (p *Price) ToNewQuantumPower(newQuantumPower QuantumPowerType) (*Price, error) {
	//if p.QuantumPower > newQuantumPower {
	//	return nil, errors.New("new quantum power should be greater or equal than current quantum power")
	//}
	//
	//diffSize := uint64(1 << (newQuantumPower - p.QuantumPower))
	//return NewPrice(
	//	p.Amount*TransactionAmountType(diffSize),
	//	newQuantumPower,
	//), nil

	return nil, nil
}

func (p *Price) ToChannel() *ChannelPrice {
	return &ChannelPrice{
		Amount:       p.Amount.AmountOf(denom).Int64(),
		QuantumPower: p.QuantumPower,
	}
}
