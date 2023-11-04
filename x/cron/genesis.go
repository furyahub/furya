package cron

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furyahub/furya/x/cron/keeper"
	"github.com/furyahub/furya/x/cron/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the schedules
	for _, elem := range genState.ScheduleList {
		err := k.AddSchedule(ctx, elem.Name, elem.Period, elem.Msgs)
		if err != nil {
			panic(err)
		}
	}
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.ScheduleList = k.GetAllSchedules(ctx)

	return genesis
}
