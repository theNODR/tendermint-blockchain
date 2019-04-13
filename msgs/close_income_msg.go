package msgs

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/types"
)

type MsgCloseIncome struct {
	PeerPublicKey string
	From          types.AccAddress
	To            types.AccAddress
}

func NewMsgCloseIncome(peerPk string, from types.AccAddress, to types.AccAddress) MsgCloseIncome {
	return MsgCloseIncome{
		PeerPublicKey: peerPk,
		From:          from,
		To:            to,
	}
}

func (msg MsgCloseIncome) Route() string {
	return "svcnodr"
}

func (msg MsgCloseIncome) Type() string {
	return "close_income"
}

func (msg MsgCloseIncome) ValidateBasic() types.Error {
	if msg.From.Empty() {
		return types.ErrInvalidAddress("Address From can`t be empty")
	}

	if msg.To.Empty() {
		return types.ErrInvalidAddress("Address To can`t be empty")
	}

	return nil
}

func (msg MsgCloseIncome) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return types.MustSortJSON(b)
}

func (msg MsgCloseIncome) GetSigners() []types.AccAddress {
	return []types.AccAddress{}
}

func (msg MsgCloseIncome) String() string {
	return fmt.Sprintf("Msg type: %s", msg.Type())
}
