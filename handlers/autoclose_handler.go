package handlers

import (
	"common"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"svcnodr/logger"
	"svcnodr/msgs"
	"svcnodr/store"
	"svcnodr/types"
)

func HandleAutocloseMsg(ctx sdk.Context, keeper store.Keeper, msg msgs.MsgAutoclose) sdk.Result {
	common.Log.Event(logger.EventAutoCloseHandle, common.Printf("Handle $s", msg.Type()))

	//Close transfer channels
	transfer := &types.TransferChannelState{}
	iterator := keeper.GetTransferIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		err := keeper.Cdc.UnmarshalBinaryBare(iterator.Value(), transfer)
		if err != nil {
			continue
		}

		if transfer.IsNeedClose(msg.Timestamp) {
			CloseTransfer(ctx, keeper, transfer, msg)
		}
	}

	//Close income channels
	income := &types.IncomeChannelState{}
	iterator = keeper.GetIncomeIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		err := keeper.Cdc.UnmarshalBinaryBare(iterator.Value(), income)
		if err != nil {
			continue
		}

		if income.IsTimeExpired(msg.Timestamp) && !income.HasChannels() {
			CloseIncome(ctx, keeper, income, msg)
		}
	}

	//Close spend channels
	spend := &types.SpendChannelState{}
	iterator = keeper.GetSpendIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		err := keeper.Cdc.UnmarshalBinaryBare(iterator.Value(), spend)
		if err != nil {
			continue
		}

		if spend.IsTimeExpired(msg.Timestamp) && !spend.HasChannels() {
			CloseSpend(ctx, keeper, spend, msg)
		}
	}

	return sdk.Result{}
}

func CloseIncome(ctx sdk.Context, keeper store.Keeper, state *types.IncomeChannelState, msg msgs.MsgAutoclose) {
	coins := keeper.CoinKeeper.GetCoins(ctx, state.Address)
	_, err := keeper.CoinKeeper.SendCoins(ctx, state.Address, msg.LedgerAddres, coins)
	if err != nil {
		_, _, _ = keeper.CoinKeeper.AddCoins(ctx, msg.LedgerAddres, coins)
	}
	keeper.DeleteIncome(ctx, state.Address)
}

func CloseSpend(ctx sdk.Context, keeper store.Keeper, state *types.SpendChannelState, msg msgs.MsgAutoclose) {
	coins := keeper.CoinKeeper.GetCoins(ctx, state.Address)
	_, err := keeper.CoinKeeper.SendCoins(ctx, state.Address, msg.LedgerAddres, coins)
	if err != nil {
		_, _, _ = keeper.CoinKeeper.AddCoins(ctx, msg.LedgerAddres, coins)
	}
	keeper.DeleteSpend(ctx, state.Address)
}

func CloseTransfer(ctx sdk.Context, keeper store.Keeper, state *types.TransferChannelState, msg msgs.MsgAutoclose) {
	keeper.DeleteTransfer(ctx, state.ChannelId)
}
