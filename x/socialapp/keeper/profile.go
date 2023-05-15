package keeper

import (
	"socialapp/x/socialapp/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AppendProfile(ctx sdk.Context, profile types.Profile) string {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProfileKey))
	appendedValue := k.cdc.MustMarshal(&profile)
	store.Set([]byte(profile.Id), appendedValue)
	return profile.Id
}

func (k Keeper) GetProfile(ctx sdk.Context, id string) (val types.Profile, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProfileKey))
	b := store.Get([]byte(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetProfile(ctx sdk.Context, profile types.Profile) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProfileKey))
	b := k.cdc.MustMarshal(&profile)
	store.Set([]byte(profile.Id), b)
}

func (k Keeper) RemoveProfile(ctx sdk.Context, id string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProfileKey))
	store.Delete([]byte(id))
}
