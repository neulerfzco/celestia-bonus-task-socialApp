package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"socialapp/x/socialapp/types"
)

func (k msgServer) CreateProfile(goCtx context.Context, msg *types.MsgCreateProfile) (*types.MsgCreateProfileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateProfileResponse{}, nil
}
