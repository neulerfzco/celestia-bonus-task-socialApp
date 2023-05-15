package keeper

import (
	"context"
	"fmt"

	"socialapp/x/socialapp/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_, found := k.GetProfile(ctx, msg.Creator)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrProfileInexistant, fmt.Sprintf("Cannot publish a post, the address %s doesn't have an account.", msg.Creator))
	}

	var post = types.Post{
		Owner: msg.Creator,
		Title: msg.Title,
		Body:  msg.Body,
	}
	id := k.AppendPost(
		ctx,
		post,
	)
	return &types.MsgCreatePostResponse{
		Id: id,
	}, nil
}
