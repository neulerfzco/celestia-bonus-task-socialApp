package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"socialapp/x/socialapp/types"
)

func (k msgServer) DeleteProfile(goCtx context.Context, msg *types.MsgDeleteProfile) (*types.MsgDeleteProfileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgDeleteProfileResponse{}, nil
}
