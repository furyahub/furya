package nextupgrade

import (
	store "github.com/cosmos/cosmos-sdk/store/types"

	"github.com/cosmos/gaia/v8/x/globalfee"
	"github.com/furyahub/furya/app/upgrades"
)

const (
	// UpgradeName defines the on-chain upgrade name.
	UpgradeName = "Next-Upgrade"
)

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{
			globalfee.ModuleName,
		},
	},
}
