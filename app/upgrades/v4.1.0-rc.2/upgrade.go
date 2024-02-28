package v4m1p0rc2

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	cfg module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		if ctx.ChainID() != TestnetChainID {
			return vm, errors.New(fmt.Sprintf("%s upgrade not allowed to execute on %s chain", UpgradeName, ctx.ChainID()))
		}

		return mm.RunMigrations(ctx, cfg, vm)
	}
}
