package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
<<<<<<< HEAD
	"github.com/strangelove-ventures/noble/v4/x/stabletokenfactory/types"
=======
	"github.com/noble-assets/noble/v5/x/stabletokenfactory/types"
>>>>>>> a4ad980 (chore: rename module path (#283))
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) MasterMinter(c context.Context, req *types.QueryGetMasterMinterRequest) (*types.QueryGetMasterMinterResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetMasterMinter(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetMasterMinterResponse{MasterMinter: val}, nil
}
