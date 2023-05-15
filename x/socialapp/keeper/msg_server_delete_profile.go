package keeper

import (
	"context"
	"fmt"
	"socialapp/x/socialapp/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) DeleteProfile(goCtx context.Context, msg *types.MsgDeleteProfile) (*types.MsgDeleteProfileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if msg.Creator != msg.Id {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	_, found := k.GetProfile(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrProfileInexistant, fmt.Sprintf("The address %s doesn't have an account", msg.Id))
	}
	k.RemovePosts(ctx, msg.Id)
	k.RemoveProfile(ctx, msg.Id)
	return &types.MsgDeleteProfileResponse{}, nil
}
