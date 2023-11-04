package keeper

import (
	"github.com/furyahub/furya/x/cron/types"
)

var _ types.QueryServer = Keeper{}
