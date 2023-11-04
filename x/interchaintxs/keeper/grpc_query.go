package keeper

import (
	"github.com/furyahub/furya/x/interchaintxs/types"
)

var _ types.QueryServer = Keeper{}
