package keeper

import (
	"github.com/furyahub/furya/x/contractmanager/types"
)

var _ types.QueryServer = Keeper{}
