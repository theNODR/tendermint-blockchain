package svcnodr

import (
	"common"
	"common/flags"
	"fmt"
	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/staking"
	abci "github.com/tendermint/tendermint/abci/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	dbm "github.com/tendermint/tendermint/libs/db"
	"svcnodr/handlers"
	"svcnodr/store"
	"svcnodr/ws"

	"github.com/tendermint/tendermint/libs/log"
)

const (
	appName = "nodrservice"
)

const (
	denom = "ndr"

	eventStartBlocker = "start blocker"
	eventEndBlocker   = "end blocker"
	eventAntHandle    = "ant handle"
)

type nodrServiceApp struct {
	*bam.BaseApp
	cdc *codec.Codec

	blockSender *ws.WsBlockSender

	nodrKeeper    store.Keeper
	bank          bank.Keeper
	paramsKeeper  params.Keeper
	accountKeeper auth.AccountKeeper

	keyMain             *sdk.KVStoreKey
	keyParams           *sdk.KVStoreKey
	keyAccount          *sdk.KVStoreKey
	tkeyParams          *sdk.TransientStoreKey
	keyFeeCollection    *sdk.KVStoreKey
	feeCollectionKeeper auth.FeeCollectionKeeper

	keySpend         *sdk.KVStoreKey
	keyIncome        *sdk.KVStoreKey
	keyTransfer      *sdk.KVStoreKey
	keyAddressIncome *sdk.KVStoreKey
	keyAddressSpend  *sdk.KVStoreKey
}

var blockSendEndpoint = flags.String("block-serve-endpoint", ":8081", "Send block info")

func NewNodrServiceApp(logger log.Logger, db dbm.DB) *nodrServiceApp {
	cdc := MakeCodec()

	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc))
	app := &nodrServiceApp{
		BaseApp:          bApp,
		cdc:              cdc,
		keyMain:          sdk.NewKVStoreKey("main"),
		tkeyParams:       sdk.NewTransientStoreKey("transient_params"),
		keyParams:        sdk.NewKVStoreKey("params"),
		keyAccount:       sdk.NewKVStoreKey("acc"),
		keyFeeCollection: sdk.NewKVStoreKey("fee_collection"),

		keySpend:         sdk.NewKVStoreKey("spend"),
		keyIncome:        sdk.NewKVStoreKey("income"),
		keyTransfer:      sdk.NewKVStoreKey("transfer"),
		keyAddressIncome: sdk.NewKVStoreKey("addr_income"),
		keyAddressSpend:  sdk.NewKVStoreKey("addr_transfer"),
		blockSender:      ws.NewWsBlockSender(*blockSendEndpoint),
	}

	//Parameter storage for the application
	app.feeCollectionKeeper = auth.NewFeeCollectionKeeper(app.cdc, app.keyFeeCollection)
	app.paramsKeeper = params.NewKeeper(app.cdc, app.keyParams, app.tkeyParams)
	app.accountKeeper = auth.NewAccountKeeper(app.cdc, app.keyAccount, app.paramsKeeper.Subspace(auth.DefaultParamspace), auth.ProtoBaseAccount)
	app.bank = bank.NewBaseKeeper(app.accountKeeper, app.paramsKeeper.Subspace(bank.DefaultParamspace), bank.DefaultParamspace)
	app.nodrKeeper = store.NewKeeper(app.bank, app.cdc, app.keySpend, app.keyIncome, app.keyTransfer, app.keyAddressIncome, app.keyAddressSpend)

	app.QueryRouter().
		AddRoute("svcnodr", NewQuerier(app.nodrKeeper))

	app.Router().
		AddRoute("svcnodr", handlers.NewHandler(app.nodrKeeper, app.cdc)).
		AddRoute("bank", bank.NewHandler(app.bank))

	app.MountStores(
		app.keyMain,
		app.keyParams,
		app.tkeyParams,
		app.keyAccount,
		app.keyFeeCollection,
		app.keyIncome,
		app.keyTransfer,
		app.keySpend,
		app.keyAddressSpend,
		app.keyAddressIncome,
	)

	app.SetBeginBlocker(app.beginBlocker)
	app.SetEndBlocker(app.endBlocker)
	app.SetAnteHandler(app.anteHandler)

	err := app.startWebsocket()
	if err != nil {
		fmt.Print("can`t start web socket")
	}

	err = app.LoadLatestVersion(app.keyMain)

	if err != nil {
		cmn.Exit(err.Error())
	}
	return app
}

func (svc nodrServiceApp) startWebsocket() error {
	return svc.blockSender.Start()
}

func MakeCodec() *codec.Codec {
	cdc := codec.New()
	RegisterCodec(cdc)
	auth.RegisterCodec(cdc)
	bank.RegisterCodec(cdc)
	staking.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	return cdc
}

func (srv *nodrServiceApp) beginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	common.Log.Event(eventStartBlocker, common.Printf("BeginBlocker. Request: %s", string(req.Hash)))
	_ = srv.blockSender.Send(ws.NewBlockInfoFromReq(req))
	return abci.ResponseBeginBlock{}
}

func (srv *nodrServiceApp) endBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	common.Log.Event(eventEndBlocker, common.Printf("EndBlocker. Request: %s", req.String()))
	return abci.ResponseEndBlock{}
}

func (srv *nodrServiceApp) anteHandler(ctx sdk.Context, tx sdk.Tx, simulate bool) (newCtx sdk.Context, result sdk.Result, abort bool) {
	common.Log.Event(eventAntHandle, common.Printf("AnteHandler. Tx: %v", tx.GetMsgs()))
	return ctx, sdk.Result{Code: sdk.CodeOK}, false
}
