package keeper

import (
	"context"
	"fmt"

	"socialapp/x/socialapp/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) ShowPost(goCtx context.Context, req *types.QueryShowPostRequest) (*types.QueryShowPostResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	_, found := k.GetProfile(ctx, req.ProfileId)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrProfileInexistant, fmt.Sprintf("The address %s doesn't have an account", req.ProfileId))
	}
	post, found := k.GetPost(ctx, req.ProfileId, req.Id)
	if !found {
		return nil, types.ErrPostInexistant
	}

	return &types.QueryShowPostResponse{Post: post}, nil
}
