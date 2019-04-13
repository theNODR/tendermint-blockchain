package msgs

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/types"
)

type MsgAutoclose struct {
	Timestamp    int64
	LedgerAddres types.AccAddress
}

func NewMsgAutoclose(timestamp int64, address types.AccAddress) MsgAutoclose {
	return MsgAutoclose{
		Timestamp:    timestamp,
		LedgerAddres: address,
	}
}

func (msg MsgAutoclose) Route() string {
	return "svcnodr"
}

func (msg MsgAutoclose) Type() string {
	return "auto_close"
}

func (msg MsgAutoclose) ValidateBasic() types.Error {
	if msg.LedgerAddres.Empty() {
		return types.ErrInvalidAddress("Tracker address empty")
	}

	return nil
}

func (msg MsgAutoclose) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return types.MustSortJSON(b)
}

func (msg MsgAutoclose) GetSigners() []types.AccAddress {
	return []types.AccAddress{}
}

func (msg MsgAutoclose) String() string {
	return fmt.Sprintf("Msg type: %s", msg.Type())
}
