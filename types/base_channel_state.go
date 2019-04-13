package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type BaseChannelState struct {
	Address  sdk.AccAddress
	Price    *Price
	State    *SpecialChannelState
	TimeLock int64
	LifeTime int64
}

func (state *BaseChannelState) IsTimeExpired(timestamp int64) bool {
	return state.TimeLock > timestamp
}
