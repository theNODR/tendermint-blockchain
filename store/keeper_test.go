package store

import (
	"github.com/cosmos/cosmos-sdk/cmd/gaia/app"
	"svcnodr/types"
	"testing"
)

type Tmp struct {
	Addres string
}

func TestKeeper_Addresss(t *testing.T) {
	cdc := app.MakeCodec()
	channeslId := types.NewChannelsIds()
	bytes := cdc.MustMarshalJSON(channeslId)
	t.Log(bytes)

}
