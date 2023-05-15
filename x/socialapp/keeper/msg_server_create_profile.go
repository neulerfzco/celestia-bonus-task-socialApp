package keeper

import (
	"context"
	"fmt"

	"socialapp/x/socialapp/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateProfile(goCtx context.Context, msg *types.MsgCreateProfile) (*types.MsgCreateProfileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var profile = types.Profile{
		Id:         msg.Creator,
		Handlename: msg.Handlename,
	}

	_, found := k.GetProfile(ctx, msg.Creator)
	if found {
		return nil, sdkerrors.Wrap(types.ErrProfileAlreadyExists, fmt.Sprintf("address %s already owns a profile, try updating it instead", msg.Creator))
	}
	id := k.AppendProfile(
		ctx,
		profile,
	)
	return &types.MsgCreateProfileResponse{
		Id: id,
	}, nil
}
