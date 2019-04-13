package msgs

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgDebug struct {
	Data string
}

func NewMsgDebug(data string) MsgDebug {
	return MsgDebug{
		data,
	}
}

// Return the message type.
// Must be alphanumeric or empty.
func (msg MsgDebug) Route() string {
	return "svcnodr"
}

// Returns a human-readable string for the message, intended for utilization
// within tags
func (msg MsgDebug) Type() string {
	return "debug_msg"
}

// ValidateBasic does a simple validation check that
// doesn't require access to any other information.
func (msg MsgDebug) ValidateBasic() sdk.Error {
	return nil
}

// Get the canonical byte representation of the Msg.
func (msg MsgDebug) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return sdk.MustSortJSON(b)
}

// Signers returns the addrs of signers that must sign.
// CONTRACT: All signatures must be present to be valid.
// CONTRACT: Returns addrs in some deterministic order.
func (msg MsgDebug) GetSigners() []sdk.AccAddress {
	return make([]sdk.AccAddress, 0)
}
