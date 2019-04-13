package msgs

import (
	"encoding/json"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgCloseTransfer struct {
	ChannelId         string
	FromAddress       sdk.AccAddress
	FromPk            string
	ToAddress         sdk.AccAddress
	ToPk              string
	PriceAmount       uint64
	PriceQuantumPower uint8
	Amount            sdk.Coins
	QuantumCount      uint64
	Volume            uint64
}

func NewMsgCloseTransfer(
	channelId string,
	fromAddress sdk.AccAddress,
	fromPk string,
	toAddress sdk.AccAddress,
	toPk string,
	priceAmount uint64,
	priceQuantumPower uint8,
	amount sdk.Coins,
	quantumCount uint64,
	volume uint64,
) MsgCloseTransfer {
	return MsgCloseTransfer{
		channelId,
		fromAddress,
		fromPk,
		toAddress,
		toPk,
		priceAmount,
		priceQuantumPower,
		amount,
		quantumCount,
		volume,
	}
}

func (msg MsgCloseTransfer) Route() string {
	return "svcnodr"
}

func (msg MsgCloseTransfer) Type() string {
	return "close_transfer"
}

func (msg MsgCloseTransfer) ValidateBasic() sdk.Error {
	if len(msg.ChannelId) <= 0 {
		return sdk.ErrInternal("ChannelsIds can`t be empty")
	}

	if len(msg.ToPk) <= 0 {
		return sdk.ErrInvalidPubKey("Public key To empty")
	}

	if len(msg.FromPk) <= 0 {
		return sdk.ErrInvalidPubKey("Public key From empty")
	}

	if msg.ToAddress.Empty() {
		return sdk.ErrInvalidAddress("Address To empty")
	}

	if msg.FromAddress.Empty() {
		return sdk.ErrInvalidAddress("Address From empty")
	}

	return nil
}

func (msg MsgCloseTransfer) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return sdk.MustSortJSON(b)
}

func (msg MsgCloseTransfer) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}

func (msg MsgCloseTransfer) String() string {
	return fmt.Sprintf("Msg type: %s", msg.Type())
}
