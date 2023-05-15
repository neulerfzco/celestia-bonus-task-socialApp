package keeper

import (
	"encoding/binary"
	"socialapp/x/socialapp/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) NewPostStore(ctx sdk.Context, profile_key string) prefix.Store {
	PrefixStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Posts))
	PostStore := prefix.NewStore(PrefixStore, types.KeyPrefix(profile_key))
	return PostStore
}

func (k Keeper) GetPostCount(ctx sdk.Context, owner string) uint64 {
	store := k.NewPostStore(ctx, owner)
	byteKey := types.KeyPrefix(types.PostCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetPostCount(ctx sdk.Context, owner string, count uint64) {
	postStore := k.NewPostStore(ctx, owner)
	byteKey := types.KeyPrefix(types.PostCountKey)
	bz := GetPostIDBytes(count)
	postStore.Set(byteKey, bz)
}

func GetPostIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func (k Keeper) AppendPost(ctx sdk.Context, post types.Post) uint64 {
	count := k.GetPostCount(ctx, post.Owner)
	post.Id = count
	postStore := k.NewPostStore(ctx, post.Owner)
	store := prefix.NewStore(postStore, types.KeyPrefix(types.PostKey))
	appendedValue := k.cdc.MustMarshal(&post)
	store.Set(GetPostIDBytes(post.Id), appendedValue)
	k.SetPostCount(ctx, post.Owner, count+1)
	return count
}

func (k Keeper) GetPost(ctx sdk.Context, owner string, id uint64) (val types.Post, found bool) {
	postStore := k.NewPostStore(ctx, owner)
	store := prefix.NewStore(postStore, types.KeyPrefix(types.PostKey))
	b := store.Get(GetPostIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) GetPosts(ctx sdk.Context, owner string) (posts []types.Post, found bool) {
	postStore := k.NewPostStore(ctx, owner)
	store := prefix.NewStore(postStore, types.KeyPrefix(types.PostKey))
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var post types.Post
		k.cdc.MustUnmarshal(iterator.Value(), &post)
		posts = append(posts, post)
	}
	return posts, true
}

func (k Keeper) RemovePost(ctx sdk.Context, owner string, id uint64) {
	postStore := k.NewPostStore(ctx, owner)
	store := prefix.NewStore(postStore, types.KeyPrefix(types.PostKey))
	store.Delete(GetPostIDBytes(id))
}

func (k Keeper) RemovePosts(ctx sdk.Context, owner string) {
	postStore := k.NewPostStore(ctx, owner)
	store := prefix.NewStore(postStore, types.KeyPrefix(types.PostKey))
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		store.Delete(iterator.Key())
	}
}
