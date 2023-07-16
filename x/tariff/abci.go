package tariff

import (
	"github.com/strangelove-ventures/noble/x/tariff/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlocker sets the proposer for determining distribution during endblock
// and distribute rewards for the previous block
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	k.AllocateTokens(ctx)
}
