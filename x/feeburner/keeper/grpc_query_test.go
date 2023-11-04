package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	feekeeperutil "github.com/furyahub/furya/testutil/feeburner/keeper"
	"github.com/furyahub/furya/x/feeburner/types"
)

func TestGrpcQuery_TotalBurnedFuryasAmount(t *testing.T) {
	feeKeeper, ctx := feekeeperutil.FeeburnerKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)

	feeKeeper.RecordBurnedFees(ctx, sdk.NewCoin(types.DefaultFuryaDenom, sdk.NewInt(100)))

	request := types.QueryTotalBurnedFuryasAmountRequest{}
	response, err := feeKeeper.TotalBurnedFuryasAmount(wctx, &request)
	require.NoError(t, err)
	require.Equal(t, &types.QueryTotalBurnedFuryasAmountResponse{TotalBurnedFuryasAmount: types.TotalBurnedFuryasAmount{Coin: sdk.Coin{Denom: types.DefaultFuryaDenom, Amount: sdk.NewInt(100)}}}, response)
}
