package msgs

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/types"
)

type MsgCloseSpend struct {
	PeerPublicKey string
	From          types.AccAddress
	To            types.AccAddress
}

func NewMsgCloseSpend(peerPk string, from types.AccAddress, to types.AccAddress) MsgCloseSpend {
	return MsgCloseSpend{
		PeerPublicKey: peerPk,
		From:          from,
		To:            to,
	}
}

func (msg MsgCloseSpend) Route() string {
	return "svcnodr"
}

// Returns a human-readable string for the message, intended for utilization
// within tags
func (msg MsgCloseSpend) Type() string {
	return "close_spend"
}

// ValidateBasic does a simple validation check that
// doesn't require access to any other information.
func (msg MsgCloseSpend) ValidateBasic() types.Error {
	if len(msg.PeerPublicKey) < 0 {
		return types.ErrInvalidPubKey("Peer public key can`t be empty")
	}

	if msg.From.Empty() {
		return types.ErrInvalidAddress("Address From can`t be empty")
	}

	if msg.To.Empty() {
		return types.ErrInvalidAddress("Address To can`t be empty")
	}

	return nil
}

// Get the canonical byte representation of the Msg.
func (msg MsgCloseSpend) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return types.MustSortJSON(b)
}

func (msg MsgCloseSpend) GetSigners() []types.AccAddress {
	return []types.AccAddress{}
}

func (msg MsgCloseSpend) String() string {
	return fmt.Sprintf("Msg type: %s", msg.Type())
}
