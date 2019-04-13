package msgs

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/types"
)

type MsgOpenSpend struct {
	TrackerPublicKey  string
	PeerPublicKey     string
	From              types.AccAddress
	MaxAmount         types.Coins
	PriceAmount       types.Coins
	PriceQuantumPower uint8
	LifeTime          int64
}

func NewMsgOpenSpend(
	trackerPublicKey string,
	peerPublicKey string,
	from types.AccAddress,
	maxAmount types.Coins,
	priceAmount types.Coins,
	priceQuantumPower uint8,
	lifeTime int64,
) MsgOpenSpend {
	return MsgOpenSpend{
		TrackerPublicKey:  trackerPublicKey,
		PeerPublicKey:     peerPublicKey,
		From:              from,
		MaxAmount:         maxAmount,
		PriceAmount:       priceAmount,
		PriceQuantumPower: priceQuantumPower,
		LifeTime:          lifeTime}
}

func (msg MsgOpenSpend) Route() string {
	return "svcnodr"
}

// Returns a human-readable string for the message, intended for utilization
// within tags
func (msg MsgOpenSpend) Type() string {
	return "open_spend"
}

// ValidateBasic does a simple validation check that
// doesn't require access to any other information.
func (msg MsgOpenSpend) ValidateBasic() types.Error {
	if len(msg.PeerPublicKey) <= 0 {
		return types.ErrInvalidPubKey("Peer pub key in empty")
	}

	if msg.From.Empty() || msg.From == nil {
		return types.ErrInvalidAddress("peer from address is empty or nil")
	}

	return nil
}

// Get the canonical byte representation of the Msg.
func (msg MsgOpenSpend) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return types.MustSortJSON(b)
}

func (msg MsgOpenSpend) GetSigners() []types.AccAddress {
	return []types.AccAddress{}
}

func (msg MsgOpenSpend) String() string {
	return fmt.Sprintf("Msg type: %s", msg.Type())
}
