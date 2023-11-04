package contractmanager_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/furyahub/furya/testutil/contractmanager/keeper"
	"github.com/furyahub/furya/testutil/contractmanager/nullify"
	"github.com/furyahub/furya/x/contractmanager"
	"github.com/furyahub/furya/x/contractmanager/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		FailuresList: []types.Failure{
			{
				Address: "address1",
				Id:      1,
			},
			{
				Address: "address1",
				Id:      2,
			},
		},
	}

	k, ctx := keepertest.ContractManagerKeeper(t, nil)
	contractmanager.InitGenesis(ctx, *k, genesisState)
	got := contractmanager.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.FailuresList, got.FailuresList)
}
