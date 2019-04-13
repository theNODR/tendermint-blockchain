package types

type TransferChannelStatus = uint

const (
	NotExistTransferChannelStatus TransferChannelStatus = 1 << iota
	OpenTransferChannelStatus     TransferChannelStatus = 1 << iota
	CloseTransferChannelStatus    TransferChannelStatus = 1 << iota
	ErrorTransferChannelStatus    TransferChannelStatus = 0
)

type ChannelsIds []string

func NewChannelsIds() ChannelsIds {
	return make([]string, 1)
}

func (c ChannelsIds) AddChannelId(chanId string) {
	if !c.HasChan(chanId) {
		c = append(c, chanId)
	}
}

func (c ChannelsIds) HasChannels() bool {
	return c != nil && len(c) > 1
}

func (c ChannelsIds) DeleteChannel(chanId string) bool {
	for i := 0; i < len(c); i++ {
		if c[i] == chanId {
			c[len(c)-1], c[i] = c[i], c[len(c)-1]
			c = c[:len(c)-1]
			return true
		}
	}

	return false
}

func (c ChannelsIds) HasChan(chanId string) bool {
	for i := 0; i < len(c); i++ {
		if c[i] == chanId {
			return true
		}
	}
	return false
}

type GetIncomeStateRequest struct {
	Address string `json:"address"`
}

type GetSpendStateRequest struct {
	Address string `json:"address"`
}

type GetTransferStateRequest struct {
	ChannelId string `json:"address"`
}

type SpendChannelStateResponse struct {
	Address  string        `json:"address"`
	Current  *ChannelFact  `json:"current"`
	Limit    *ChannelPlan  `json:"limit"`
	Price    *ChannelPrice `json:"price"`
	TimeLock int64         `json:"timelock"`
	LifeTime int64         `json:"lifetime"`
}

type IncomeChannelStateResponse struct {
	Address  string        `json:"address"`
	Current  *ChannelFact  `json:"current"`
	Limit    *ChannelPlan  `json:"limit"`
	Price    *ChannelPrice `json:"price"`
	TimeLock int64         `json:"timelock"`
	LifeTime int64         `json:"lifetime"`
}

type TransferChannelStateResponse struct {
	Status TransferChannelStatus `json:"status"`
}

type TrackerBalanceStateResponse struct {
}
