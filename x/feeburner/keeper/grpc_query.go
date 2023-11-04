package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/furyahub/furya/x/feeburner/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) TotalBurnedFuryasAmount(goCtx context.Context, _ *types.QueryTotalBurnedFuryasAmountRequest) (*types.QueryTotalBurnedFuryasAmountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	totalBurnedFuryasAmount := k.GetTotalBurnedFuryasAmount(ctx)

	return &types.QueryTotalBurnedFuryasAmountResponse{TotalBurnedFuryasAmount: totalBurnedFuryasAmount}, nil
}
