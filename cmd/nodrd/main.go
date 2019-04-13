package main

import (
	"common"
	"common/flags"
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/cmd/gaia/app"
	gaiaInit "github.com/cosmos/cosmos-sdk/cmd/gaia/init"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	abci "github.com/tendermint/tendermint/abci/types"
	cfg "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/cli"
	tcommon "github.com/tendermint/tendermint/libs/common"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/types"
	"io"
	"os"
	"path/filepath"
	"svcnodr"
)

var DefaultNodeHome = os.ExpandEnv("$HOME/.nodrd")

const (
	flagOverwrite = "overwrite"
)

var (
	chain_id = flags.String("chain_id", "teschain", "chain id")
)

func main() {
	common.InitializeDefault()
	common.Log.Start()
	flags.Parse()
	viper.Set(client.FlagChainID, chain_id)

	cobra.EnableCommandSorting = false

	cdc := app.MakeCodec()
	ctx := server.NewDefaultContext()

	rootCmd := &cobra.Command{
		Use:               "nodrd",
		Short:             "Nodr node server",
		PersistentPreRunE: server.PersistentPreRunEFn(ctx),
	}

	rootCmd.AddCommand(InitCmd(ctx, cdc))

	//todo поменять appExporter с  nil на нормальную реализацию
	server.AddCommands(ctx, cdc, rootCmd, newApp, nil)

	executor := cli.PrepareBaseCmd(rootCmd, "NODR", DefaultNodeHome)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

//
//func appExporter() server.AppExporter {
//	return func(logger log.Logger, db dbm.DB, _ io.Writer, _ int64, _ bool, _ []string) (
//		json.RawMessage, []tmtypes.GenesisValidator, error) {
//		dapp := srvnodr.NewNodrServiceApp(logger, db)
//		return dapp.()
//	}
//}

func InitCmd(ctx *server.Context, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize genesis config, priv-validator file, and p2p-node file",
		Args:  cobra.NoArgs,
		RunE: func(_ *cobra.Command, _ []string) error {
			config := ctx.Config
			config.Consensus.CreateEmptyBlocks = false
			config.SetRoot(viper.GetString(cli.HomeFlag))

			config.P2P.Seeds = "192.168.100.9,192.168.100.6"

			chainID := viper.GetString(client.FlagChainID)
			if chainID == "" {
				chainID = fmt.Sprintf("test-chain-%v", tcommon.RandStr(6))
			}

			_, pk, err := gaiaInit.InitializeNodeValidatorFiles(config)
			if err != nil {
				return err
			}

			var appState json.RawMessage
			genFile := config.GenesisFile()

			if !viper.GetBool(flagOverwrite) && tcommon.FileExists(genFile) {
				return fmt.Errorf("genesis.json file already exists: %v", genFile)
			}

			genesis := app.GenesisState{
				AuthData: auth.DefaultGenesisState(),
				BankData: bank.DefaultGenesisState(),
			}

			appState, err = codec.MarshalJSONIndent(cdc, genesis)
			if err != nil {
				return err
			}

			_, _, validator, err := SimpleAppGenTx(cdc, pk)
			if err != nil {
				return err
			}

			if err = gaiaInit.ExportGenesisFile(genFile, chainID, []tmtypes.GenesisValidator{validator}, appState); err != nil {
				return err
			}

			cfg.WriteConfigFile(filepath.Join(config.RootDir, "config", "config.toml"), config)

			fmt.Printf("Initialized nsd configuration and bootstrapping files in %s...\n", viper.GetString(cli.HomeFlag))
			return nil
		},
	}

	cmd.Flags().String(cli.HomeFlag, DefaultNodeHome, "node's home directory")
	cmd.Flags().String(client.FlagChainID, "", "genesis file chain-id, if left blank will be randomly created")
	cmd.Flags().BoolP(flagOverwrite, "o", false, "overwrite the genesis.json file")

	return cmd
}

// SimpleAppGenTx returns a simple GenTx command that makes the node a valdiator from the start
func SimpleAppGenTx(cdc *codec.Codec, pk crypto.PubKey) (
	appGenTx, cliPrint json.RawMessage, validator tmtypes.GenesisValidator, err error) {

	addr, secret, err := server.GenerateCoinKey()
	if err != nil {
		return
	}

	bz, err := cdc.MarshalJSON(struct {
		Addr sdk.AccAddress `json:"addr"`
	}{addr})
	if err != nil {
		return
	}

	appGenTx = json.RawMessage(bz)

	bz, err = cdc.MarshalJSON(map[string]string{"secret": secret})
	if err != nil {
		return
	}

	cliPrint = json.RawMessage(bz)

	validator = tmtypes.GenesisValidator{
		PubKey: pk,
		Power:  10,
	}

	return
}

func newApp(logger log.Logger, db dbm.DB, traceStore io.Writer) abci.Application {
	return svcnodr.NewNodrServiceApp(logger, db)
}
