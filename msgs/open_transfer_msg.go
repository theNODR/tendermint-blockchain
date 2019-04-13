package msgs

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/types"
)

type MsgOpenTransfer struct {
	ChannelId           string
	FromAddress         types.AccAddress
	FromPk              string
	ToAddress           types.AccAddress
	ToPk                string
	PriceAmount         types.Coins
	PriceQuantumPower   uint8
	PlannedQuantumCount uint64
	Lifetime            int64
}

func NewMsgOpenTransfer(
	channelId string,
	fromAddr types.AccAddress,
	fromPk string,
	toAddress types.AccAddress,
	toPk string,
	priceAmount types.Coins,
	priceQuantumPower uint8,
	plannedQuanimCount uint64,
	lifetime int64) MsgOpenTransfer {
	return MsgOpenTransfer{
		ChannelId:           channelId,
		FromAddress:         fromAddr,
		FromPk:              fromPk,
		ToAddress:           toAddress,
		ToPk:                toPk,
		PriceAmount:         priceAmount,
		PriceQuantumPower:   priceQuantumPower,
		PlannedQuantumCount: plannedQuanimCount,
		Lifetime:            lifetime,
	}
}

func (msg MsgOpenTransfer) Route() string {
	return "svcnodr"
}

func (msg MsgOpenTransfer) Type() string {
	return "open_transfer"
}

func (msg MsgOpenTransfer) ValidateBasic() types.Error {

	if len(msg.FromPk) <= 0 {
		return types.ErrInvalidPubKey("Peer pub key in empty")
	}

	if msg.FromAddress.Empty() || msg.FromAddress == nil {
		return types.ErrInvalidAddress("peer from address is empty or nil")
	}

	if len(msg.ToPk) <= 0 {
		return types.ErrInvalidPubKey("Peer to pub key is empty")
	}

	if msg.ToAddress.Empty() || msg.ToAddress == nil {
		return types.ErrInvalidAddress("peer to address is empty or nil")
	}

	return nil
}

// Get the canonical byte representation of the Msg.
func (msg MsgOpenTransfer) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return types.MustSortJSON(b)
}

func (msg MsgOpenTransfer) GetSigners() []types.AccAddress {
	return []types.AccAddress{}
}

func (msg MsgOpenTransfer) String() string {
	return fmt.Sprintf("Msg type: %s", msg.Type())
}
