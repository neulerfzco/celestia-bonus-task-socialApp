package keeper

import (
	"context"
	"fmt"

	"socialapp/x/socialapp/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateProfile(goCtx context.Context, msg *types.MsgUpdateProfile) (*types.MsgUpdateProfileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var profile = types.Profile{
		Id:         msg.Id,
		Handlename: msg.Handlename,
	}
	val, found := k.GetProfile(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Id != val.Id {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	k.SetProfile(ctx, profile)
	return &types.MsgUpdateProfileResponse{}, nil
}
