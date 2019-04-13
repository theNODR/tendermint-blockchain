package msgs

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/types"
	types2 "svcnodr/types"
)

type MsgOpenIncome struct {
	TrackerPublicKey  string
	PeerPublicKey     string
	From              types.AccAddress
	PriceAmount       types.Coins
	PriceQuantumPower types2.QuantumPowerType
	LifeTime          int64
}

func NewMsgOpenIncome(
	trackerPublicKey string,
	peerPublicKey string,
	from types.AccAddress,
	priceAmount types.Coins,
	priceQuantumPower types2.QuantumPowerType,
	lifeTime int64) MsgOpenIncome {
	return MsgOpenIncome{
		trackerPublicKey,
		peerPublicKey,
		from,
		priceAmount,
		priceQuantumPower,
		lifeTime}
}

func (msg MsgOpenIncome) Route() string {
	return "svcnodr"
}

// Returns a human-readable string for the message, intended for utilization
// within tags
func (msg MsgOpenIncome) Type() string {
	return "open_income"
}

func (msg MsgOpenIncome) ValidateBasic() types.Error {
	if len(msg.PeerPublicKey) <= 0 {
		return types.ErrInvalidPubKey("Peer pub key in empty")
	}

	if msg.From.Empty() || msg.From == nil {
		return types.ErrInvalidAddress("peer from address is empty or nil")
	}

	if len(msg.TrackerPublicKey) <= 0 {
		return types.ErrInvalidPubKey("Tracker public key can`t be empty")
	}

	return nil
}

// Get the canonical byte representation of the Msg.
func (msg MsgOpenIncome) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return types.MustSortJSON(b)
}

func (msg MsgOpenIncome) GetSigners() []types.AccAddress {
	return []types.AccAddress{}
}

func (msg MsgOpenIncome) String() string {
	return fmt.Sprintf("Msg type: %s", msg.Type())
}
