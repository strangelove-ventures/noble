package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
<<<<<<< HEAD
	"github.com/strangelove-ventures/noble/v4/x/tokenfactory/types"
=======
	"github.com/noble-assets/noble/v5/x/tokenfactory/types"
>>>>>>> a4ad980 (chore: rename module path (#283))
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Pauser(c context.Context, req *types.QueryGetPauserRequest) (*types.QueryGetPauserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetPauser(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetPauserResponse{Pauser: val}, nil
}
