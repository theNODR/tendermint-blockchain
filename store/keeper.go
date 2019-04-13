package store

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"svcnodr/types"
)

type Keeper struct {
	CoinKeeper bank.Keeper

	spendKey    sdk.StoreKey
	incomeKey   sdk.StoreKey
	transferKey sdk.StoreKey

	addressesIncome sdk.StoreKey
	addressesSpend  sdk.StoreKey
	Cdc             *codec.Codec
}

func NewKeeper(
	coinKeeper bank.Keeper,
	cdc *codec.Codec,
	spend sdk.StoreKey,
	income sdk.StoreKey,
	transfer sdk.StoreKey,
	addressesIncome sdk.StoreKey,
	addressesSpend sdk.StoreKey) Keeper {
	return Keeper{
		CoinKeeper:      coinKeeper,
		Cdc:             cdc,
		spendKey:        spend,
		incomeKey:       income,
		transferKey:     transfer,
		addressesIncome: addressesIncome,
		addressesSpend:  addressesSpend,
	}
}

func (k Keeper) SetIncome(ctx sdk.Context, state *types.IncomeChannelState) sdk.Error {
	store := ctx.KVStore(k.incomeKey)
	bytes := k.Cdc.MustMarshalBinaryBare(state)
	store.Set(state.Address, bytes)
	return nil
}

func (k Keeper) GetIncome(ctx sdk.Context, address sdk.AccAddress) (*types.IncomeChannelState, sdk.Error) {
	store := ctx.KVStore(k.incomeKey)
	if store.Has(address) {
		state := &types.IncomeChannelState{}
		k.Cdc.MustUnmarshalBinaryBare(store.Get(address), state)
		return state, nil
	} else {
		return nil, sdk.ErrInvalidAddress(fmt.Sprintf("No such state with address %s", address.String()))
	}
}

func (k Keeper) GetTransfer(ctx sdk.Context, channelId string) (*types.TransferChannelState, sdk.Error) {
	store := ctx.KVStore(k.transferKey)
	if store.Has([]byte(channelId)) {
		bytes := store.Get([]byte(channelId))
		state := &types.TransferChannelState{}
		k.Cdc.MustUnmarshalBinaryBare(bytes, &state)
		return state, nil
	} else {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("No transfer channel with id: %s", channelId))
	}
}

func (k Keeper) GetSpend(ctx sdk.Context, address sdk.AccAddress) (*types.SpendChannelState, sdk.Error) {
	store := ctx.KVStore(k.spendKey)
	if store.Has(address) {
		state := &types.SpendChannelState{}
		k.Cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(address), state)
		return state, nil
	}
	return nil, sdk.ErrInvalidAddress(fmt.Sprintf("No Spend channel with adrres: %s", address.String()))
}

func (k Keeper) DeleteIncome(ctx sdk.Context, address sdk.AccAddress) {
	store := ctx.KVStore(k.incomeKey)
	store.Delete(address)
}

func (k Keeper) DeleteSpend(ctx sdk.Context, address sdk.AccAddress) {
	store := ctx.KVStore(k.spendKey)
	store.Delete(address)
}

func (k Keeper) DeleteTransfer(ctx sdk.Context, channelId string) {
	store := ctx.KVStore(k.transferKey)
	store.Delete([]byte(channelId))
}

func (k Keeper) SetSpend(ctx sdk.Context, state *types.SpendChannelState) sdk.Error {
	store := ctx.KVStore(k.spendKey)
	bytes := k.Cdc.MustMarshalBinaryLengthPrefixed(state)
	store.Set(state.Address, bytes)
	return nil
}

func (k Keeper) AddTransfer(ctx sdk.Context, state *types.TransferChannelState) sdk.Error {
	store := ctx.KVStore(k.transferKey)
	stateEncoded := k.Cdc.MustMarshalBinaryBare(state)
	store.Set([]byte(state.ChannelId), stateEncoded)
	return nil
}

func (k Keeper) HasSpend(ctx sdk.Context, address sdk.AccAddress) bool {
	store := ctx.KVStore(k.spendKey)
	return store.Has(address)
}

func (k Keeper) HasIncome(ctx sdk.Context, address sdk.AccAddress) bool {
	store := ctx.KVStore(k.incomeKey)
	return store.Has(address)
}

func (k Keeper) AddChannelIdToIncome(ctx sdk.Context, addrIncome sdk.AccAddress, channelId string) sdk.Error {
	income, err := k.GetIncome(ctx, addrIncome)
	if err != nil {
		return err
	}

	if income.ChannelsIds == nil {
		income.ChannelsIds = types.NewChannelsIds()
	}

	income.ChannelsIds.AddChannelId(channelId)
	err = k.SetIncome(ctx, income)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) AddChannelIdToSpend(ctx sdk.Context, addrSpend sdk.AccAddress, channelId string) sdk.Error {
	spend, err := k.GetSpend(ctx, addrSpend)
	if err != nil {
		return err
	}

	spend.ChannelsIds.AddChannelId(channelId)
	err = k.SetSpend(ctx, spend)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) DeleteChannelIdFromIncome(ctx sdk.Context, addrIncome sdk.AccAddress, channelId string) sdk.Error {
	income, err := k.GetIncome(ctx, addrIncome)
	if err != nil {
		return err
	}

	income.ChannelsIds.DeleteChannel(channelId)
	err = k.SetIncome(ctx, income)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) DeleteChannelIdFromSpend(ctx sdk.Context, addrSpend sdk.AccAddress, channelId string) sdk.Error {
	spend, err := k.GetSpend(ctx, addrSpend)
	if err != nil {
		return err
	}

	spend.ChannelsIds.DeleteChannel(channelId)
	err = k.SetSpend(ctx, spend)
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) GetSpendIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.spendKey)
	return sdk.KVStorePrefixIterator(store, nil)
}

func (k Keeper) GetIncomeIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.incomeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}

func (k Keeper) GetTransferIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.transferKey)
	return sdk.KVStorePrefixIterator(store, nil)
}
