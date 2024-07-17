package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/BabbyBit/babychain/x/game/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetPlayer set a specific player in the store from its index
func (k Keeper) SetPlayer(ctx context.Context, player types.Player) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PlayerKeyPrefix))
	b := k.cdc.MustMarshal(&player)
	store.Set(types.PlayerKey(
		player.Address,
	), b)
}

// GetPlayer returns a player from its index
func (k Keeper) GetPlayer(
	ctx context.Context,
	address string,

) (val types.Player, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PlayerKeyPrefix))

	b := store.Get(types.PlayerKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePlayer removes a player from the store
func (k Keeper) RemovePlayer(
	ctx context.Context,
	address string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PlayerKeyPrefix))
	store.Delete(types.PlayerKey(
		address,
	))
}

// GetAllPlayer returns all player
func (k Keeper) GetAllPlayer(ctx context.Context) (list []types.Player) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PlayerKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Player
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
