package interchaintxs_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/furyahub/furya/testutil/interchaintxs/keeper"
	"github.com/furyahub/furya/testutil/interchaintxs/nullify"
	"github.com/furyahub/furya/x/interchaintxs"
	"github.com/furyahub/furya/x/interchaintxs/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
	}

	k, ctx, _ := keepertest.InterchainTxsKeeper(t, nil, nil, nil, nil, nil)
	interchaintxs.InitGenesis(ctx, *k, genesisState)
	got := interchaintxs.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)
}
