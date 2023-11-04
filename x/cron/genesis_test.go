package cron_test

import (
	"testing"

	"github.com/furyahub/furya/testutil/cron/keeper"
	"github.com/furyahub/furya/testutil/cron/nullify"
	"github.com/furyahub/furya/x/cron"
	"github.com/furyahub/furya/x/cron/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	k, ctx := keeper.CronKeeper(t, nil, nil)

	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		ScheduleList: []types.Schedule{
			{
				Name:              "a",
				Period:            5,
				Msgs:              nil,
				LastExecuteHeight: uint64(ctx.BlockHeight()),
			},
		},
	}

	cron.InitGenesis(ctx, *k, genesisState)
	got := cron.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.Params, got.Params)
	require.ElementsMatch(t, genesisState.ScheduleList, got.ScheduleList)
}
