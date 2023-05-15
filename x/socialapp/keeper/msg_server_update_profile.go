package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"socialapp/x/socialapp/types"
)

func (k msgServer) UpdateProfile(goCtx context.Context, msg *types.MsgUpdateProfile) (*types.MsgUpdateProfileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUpdateProfileResponse{}, nil
}
