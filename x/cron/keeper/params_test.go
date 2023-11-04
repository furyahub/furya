package keeper_test

import (
	"testing"

	"github.com/furyahub/furya/testutil"

	"github.com/furyahub/furya/app"

	testkeeper "github.com/furyahub/furya/testutil/cron/keeper"

	"github.com/furyahub/furya/x/cron/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	_ = app.GetDefaultConfig()

	k, ctx := testkeeper.CronKeeper(t, nil, nil)
	params := types.Params{
		SecurityAddress: testutil.TestOwnerAddress,
		Limit:           5,
	}

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
