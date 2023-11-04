package keeper_test

import (
	"testing"

	"github.com/furyahub/furya/app"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/furyahub/furya/testutil/feeburner/keeper"
	"github.com/furyahub/furya/x/feeburner/types"
)

func TestGetParams(t *testing.T) {
	_ = app.GetDefaultConfig()

	k, ctx := testkeeper.FeeburnerKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
