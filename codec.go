package svcnodr

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"svcnodr/msgs"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(msgs.MsgInit{}, "svcnodr/Init", nil)
	cdc.RegisterConcrete(msgs.MsgDebug{}, "svcnodr/Debug", nil)
	cdc.RegisterConcrete(msgs.MsgOpenTransfer{}, "svcnodr/OpenTransfer", nil)
	cdc.RegisterConcrete(msgs.MsgOpenSpend{}, "svcnodr/OpenSpend", nil)
	cdc.RegisterConcrete(msgs.MsgOpenIncome{}, "svcnodr/OpenIncome", nil)
	cdc.RegisterConcrete(msgs.MsgCloseTransfer{}, "svcnodr/CloseTransfer", nil)
	cdc.RegisterConcrete(msgs.MsgCloseIncome{}, "svcnodr/CloseIncome", nil)
	cdc.RegisterConcrete(msgs.MsgCloseSpend{}, "svcnodr/CloseSpend", nil)
	cdc.RegisterConcrete(msgs.MsgAutoclose{}, "svcnodr/Autoclose", nil)
}
