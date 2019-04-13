package handlers

import (
	"common"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"svcnodr/msgs"
)

func HandleDebugMsg(ctx sdk.Context, msg msgs.MsgDebug) sdk.Result {
	common.Log.PrintFull(common.Printf("Handle: %", msg.Type()))
	return sdk.Result{}
}
