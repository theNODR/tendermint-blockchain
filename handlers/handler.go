package handlers

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types"
	"svcnodr/msgs"
	"svcnodr/store"
)

const eventHandleMsg = "handle msg"

func NewHandler(keeper store.Keeper, cdc *codec.Codec) types.Handler {
	return func(ctx types.Context, msg types.Msg) types.Result {
		switch msg := msg.(type) {
		case msgs.MsgInit:
			return HandleInitMsg(ctx, keeper, msg)
		case msgs.MsgAutoclose:
			return HandleAutocloseMsg(ctx, keeper, msg)
		case msgs.MsgCloseIncome:
			return HandleCloseIncomeMsg(ctx, keeper, cdc, msg)
		case msgs.MsgCloseSpend:
			return HandleCloseSpendMsg(ctx, keeper, cdc, msg)
		case msgs.MsgCloseTransfer:
			return HandleCloseTransferMsg(ctx, keeper, cdc, msg)
		case msgs.MsgOpenSpend:
			return HandleOpenSpendMsg(ctx, keeper, cdc, msg)
		case msgs.MsgOpenIncome:
			return HandleOpenIncomeMsg(ctx, keeper, cdc, msg)
		case msgs.MsgOpenTransfer:
			return HandleOpenTransferMsg(ctx, keeper, cdc, msg)
		case msgs.MsgDebug:
			return HandleDebugMsg(ctx, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized nodk Msg type: %v", msg.Type())
			return types.ErrUnknownRequest(errMsg).Result()
		}
	}
}
