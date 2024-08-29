package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/noble-assets/noble/v7/x/tokenfactory/keeper"
	"github.com/noble-assets/noble/v7/x/tokenfactory/types"
)

func SimulateMsgBurn(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k *keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgBurn{
			From: simAccount.Address.String(),
		}

		// TODO: Handling the Burn simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Burn simulation not implemented"), nil, nil
	}
}
