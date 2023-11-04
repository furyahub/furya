package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/furyahub/furya/x/feerefunder/types"
)

func (k Keeper) CheckFees(ctx sdk.Context, fees types.Fee) error {
	return k.checkFees(ctx, fees)
}
