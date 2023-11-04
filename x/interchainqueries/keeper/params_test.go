package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/furyahub/furya/testutil/interchainqueries/keeper"
	"github.com/furyahub/furya/x/interchainqueries/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.InterchainQueriesKeeper(t, nil, nil, nil, nil)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
