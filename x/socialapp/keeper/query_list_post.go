package keeper

import (
	"context"
	"fmt"

	"socialapp/x/socialapp/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TODO use the pagination from the cosmos-sdk instead of this one
type MyPaginationRequest struct {
	Limit int    // Number of results per page
	Key   []byte // Starting key for pagination (optional)
}

func (k Keeper) ListPost(goCtx context.Context, req *types.QueryListPostRequest) (*types.QueryListPostResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	_, found := k.GetProfile(ctx, req.ProfileId)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrProfileInexistant, fmt.Sprintf("The address %s doesn't have an account", req.ProfileId))
	}

	var posts []types.Post
	posts, found = k.GetPosts(ctx, req.ProfileId)
	pageResponse := query.PageResponse{
		Total:   0,
		NextKey: nil,
	}
	emptyPost := types.Post{}
	emptyPosts := []types.Post{emptyPost}

	if !found {
		return &types.QueryListPostResponse{Post: emptyPosts, Pagination: &pageResponse}, nil
	}

	return &types.QueryListPostResponse{Post: posts, Pagination: &pageResponse}, nil
}
