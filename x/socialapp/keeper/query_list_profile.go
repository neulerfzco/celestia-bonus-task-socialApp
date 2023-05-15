package keeper

import (
	"context"
	"socialapp/x/socialapp/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListProfile(goCtx context.Context, req *types.QueryListProfileRequest) (*types.QueryListProfileResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var profiles []types.Profile
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	profileStore := prefix.NewStore(store, types.KeyPrefix(types.ProfileKey))

	pageRes, err := query.Paginate(profileStore, req.Pagination, func(key []byte, value []byte) error {
		var profile types.Profile
		if err := k.cdc.Unmarshal(value, &profile); err != nil {
			return err
		}

		profiles = append(profiles, profile)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListProfileResponse{Profile: profiles, Pagination: pageRes}, nil
}
