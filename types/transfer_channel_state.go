package types

import (
	"common"
	"github.com/cosmos/cosmos-sdk/types"
)

type TransferChannelState struct {
	ChannelId           string
	FromAddress         types.AccAddress
	FromPk              string
	ToAddress           types.AccAddress
	ToPk                string
	PriceAmount         types.Coins
	PriceQuantumPower   uint8
	PlannedQuantumCount uint64
	TimeLock            int64
}

func NewTransferChannelState(channelId string,
	fromAddres types.AccAddress,
	fromPk string,
	toAddres types.AccAddress,
	toPk string,
	amount types.Coins,
	priceQuantumPower uint8,
	plannedQuantumPower uint64,
	lifetime int64) *TransferChannelState {
	return &TransferChannelState{
		ChannelId:           channelId,
		FromAddress:         fromAddres,
		FromPk:              fromPk,
		ToAddress:           toAddres,
		ToPk:                toPk,
		PriceAmount:         amount,
		PriceQuantumPower:   priceQuantumPower,
		PlannedQuantumCount: plannedQuantumPower,
		TimeLock:            lifetime + common.GetNowUnixMs(),
	}
}

func (ts *TransferChannelState) IsNeedClose(timestamp int64) bool {
	return ts.TimeLock < timestamp
}
