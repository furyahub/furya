package feeburner_test

import (
	"testing"

	"github.com/furyahub/furya/app"

	"github.com/furyahub/furya/testutil/feeburner/keeper"
	"github.com/furyahub/furya/testutil/feeburner/nullify"
	"github.com/furyahub/furya/x/feeburner"
	"github.com/furyahub/furya/x/feeburner/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	_ = app.GetDefaultConfig()

	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	k, ctx := keeper.FeeburnerKeeper(t)
	feeburner.InitGenesis(ctx, *k, genesisState)
	got := feeburner.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)
}
