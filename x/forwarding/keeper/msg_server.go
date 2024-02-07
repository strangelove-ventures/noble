package keeper

import (
	"context"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/noble-assets/noble/v5/x/forwarding/types"
)

var _ types.MsgServer = &Keeper{}

func (k *Keeper) RegisterAccount(goCtx context.Context, msg *types.MsgRegisterAccount) (*types.MsgRegisterAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	address := types.GenerateAddress(msg.Channel, msg.Recipient)

	if k.authKeeper.HasAccount(ctx, address) {
		rawAccount := k.authKeeper.GetAccount(ctx, address)

		switch account := rawAccount.(type) {
		case *authtypes.BaseAccount:
			k.authKeeper.SetAccount(ctx, &types.ForwardingAccount{
				BaseAccount: account,
				Channel:     msg.Channel,
				Recipient:   msg.Recipient,
				CreatedAt:   ctx.BlockHeight(),
			})

			k.IncrementNumOfAccounts(ctx, msg.Channel)
		case *types.ForwardingAccount:
			return nil, errors.New("account has already been registered")
		default:
			break
		}

		if !k.bankKeeper.GetAllBalances(ctx, address).IsZero() {
			account, ok := rawAccount.(*types.ForwardingAccount)

			if ok {
				k.SetPendingForward(ctx, account)
			}
		}

		return &types.MsgRegisterAccountResponse{Address: address.String()}, nil
	}

	account := types.ForwardingAccount{
		BaseAccount: authtypes.NewBaseAccountWithAddress(address),
		Channel:     msg.Channel,
		Recipient:   msg.Recipient,
		CreatedAt:   ctx.BlockHeight(),
	}

	k.authKeeper.SetAccount(ctx, &account)
	k.IncrementNumOfAccounts(ctx, msg.Channel)

	return &types.MsgRegisterAccountResponse{Address: address.String()}, nil
}
