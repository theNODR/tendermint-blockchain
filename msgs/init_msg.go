package msgs

import (
	"encoding/json"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgInit struct {
	To     sdk.AccAddress
	Amount sdk.Coins
}

func NewMsgInit(to sdk.AccAddress, amount sdk.Coins) MsgInit {
	return MsgInit{
		To:     to,
		Amount: amount,
	}
}

// Return the message type.
// Must be alphanumeric or empty.
func (msg MsgInit) Route() string {
	return "svcnodr"
}

// Returns a human-readable string for the message, intended for utilization
// within tags
func (msg MsgInit) Type() string {
	return "init_msg"
}

// ValidateBasic does a simple validation check that
// doesn't require access to any other information.
func (msg MsgInit) ValidateBasic() sdk.Error {
	if len(msg.To) == 0 {
		return sdk.ErrInvalidAddress("Init addres To can`t be empty")
	}

	return nil
}

// Get the canonical byte representation of the Msg.
func (msg MsgInit) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return sdk.MustSortJSON(b)
}

// Signers returns the addrs of signers that must sign.
// CONTRACT: All signatures must be present to be valid.
// CONTRACT: Returns addrs in some deterministic order.
func (msg MsgInit) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.To}
}

func (msg MsgInit) String() string {
	return fmt.Sprintf("MsgInit: To %s, \n Amount: %d", msg.To.String(), msg.Amount)
}
