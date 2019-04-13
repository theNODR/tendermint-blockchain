package svcnodr

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"svcnodr/store"
	ntypes "svcnodr/types"
)

const (
	QueryIncomeAddressState   = "income"
	QuerySpendAddressState    = "spend"
	QueryTransferAddressState = "transfer"
)

func NewQuerier(keeper store.Keeper) types.Querier {
	return func(ctx types.Context, path []string, req abci.RequestQuery) ([]byte, types.Error) {
		switch path[0] {
		case QueryIncomeAddressState:
			return queryIncomeAddressState(ctx, path[1:], req, keeper)
		case QuerySpendAddressState:
			return querySpendAddressState(ctx, path[1:], req, keeper)
		case QueryTransferAddressState:
			return queryTransferChannelState(ctx, path[1:], req, keeper)
		default:
			return nil, types.ErrUnknownRequest("Unknown srvnodr query endpoint")
		}
	}
}

func queryIncomeAddressState(ctx types.Context, path []string, req abci.RequestQuery, keeper store.Keeper) (res []byte, err types.Error) {
	request := &ntypes.GetIncomeStateRequest{}
	err1 := keeper.Cdc.UnmarshalJSON(req.Data, request)
	if err1 != nil {
		return nil, types.ErrInternal("Can`t parse reqeust income state")
	}

	state, err := keeper.GetIncome(ctx, types.AccAddress(request.Address))
	if err != nil {
		return nil, err
	}

	response := &ntypes.IncomeChannelStateResponse{
		Address:  string(state.Address),
		Current:  state.State.Fact.ToChannel(),
		Limit:    state.State.Plan.ToChannel(),
		LifeTime: state.LifeTime,
		TimeLock: state.TimeLock,
		Price:    state.Price.ToChannel(),
	}

	bytesResponse := keeper.Cdc.MustMarshalJSON(response)
	return bytesResponse, nil
}

func querySpendAddressState(ctx types.Context, path []string, req abci.RequestQuery, keeper store.Keeper) (res []byte, err types.Error) {
	request := &ntypes.GetSpendStateRequest{}
	keeper.Cdc.MustUnmarshalJSON(req.Data, request)

	state, err := keeper.GetSpend(ctx, types.AccAddress(request.Address))
	if err != nil {
		return nil, err
	}

	response := &ntypes.SpendChannelStateResponse{
		Address:  string(state.Address),
		Current:  state.State.Fact.ToChannel(),
		Limit:    state.State.Plan.ToChannel(),
		LifeTime: state.LifeTime,
		TimeLock: state.TimeLock,
		Price:    state.Price.ToChannel(),
	}

	bytesResponse := keeper.Cdc.MustMarshalJSON(response)
	return bytesResponse, nil
}

func queryTransferChannelState(ctx types.Context, path []string, req abci.RequestQuery, keeper store.Keeper) ([]byte, types.Error) {
	request := &ntypes.GetTransferStateRequest{}
	err := keeper.Cdc.UnmarshalJSON(req.Data, request)
	if err != nil {
		return nil, types.ErrUnknownRequest(fmt.Sprintf("Can`t parse request data: %s", req.String()))
	}

	_, err = keeper.GetTransfer(ctx, request.ChannelId)

	response := &ntypes.TransferChannelStateResponse{}
	if err != nil {
		response.Status = ntypes.NotExistTransferChannelStatus
	} else {
		response.Status = ntypes.OpenTransferChannelStatus
	}

	responseBytes, err := keeper.Cdc.MarshalJSON(response)
	if err != nil {
		return nil, types.ErrInternal(fmt.Sprintf("Can`t create response: %v", response))
	}

	return responseBytes, nil
}
