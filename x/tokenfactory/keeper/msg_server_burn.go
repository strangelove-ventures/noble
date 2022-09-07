package keeper

import (
	"context"

	"noble/x/tokenfactory/types"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Burn(goCtx context.Context, msg *types.MsgBurn) (*types.MsgBurnResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetMinters(ctx, msg.From)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "you are not a minter")
	}

	_, found = k.GetBlacklisted(ctx, msg.From)
	if found {
		return nil, sdkerrors.Wrapf(types.ErrBurn, "minter address is blacklisted")
	}

	paused, found := k.GetPaused(ctx)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrBurn, "paused value is not found")
	}

	if paused.Paused {
		return nil, sdkerrors.Wrapf(types.ErrBurn, "burning is paused")
	}

	minterAddress, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrParseAddress, err.Error())
	}

	amount := sdk.NewCoins(msg.Amount)

	k.bankKeeper.SendCoinsFromAccountToModule(ctx, minterAddress, types.ModuleName, amount)

	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, amount); err != nil {
		return nil, sdkerrors.Wrap(types.ErrBurn, err.Error())
	}

	return &types.MsgBurnResponse{}, nil
}