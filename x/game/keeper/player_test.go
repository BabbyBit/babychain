package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "github.com/BabbyBit/babychain/testutil/keeper"
	"github.com/BabbyBit/babychain/testutil/nullify"
	"github.com/BabbyBit/babychain/x/game/keeper"
	"github.com/BabbyBit/babychain/x/game/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNPlayer(keeper keeper.Keeper, ctx context.Context, n int) []types.Player {
	items := make([]types.Player, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetPlayer(ctx, items[i])
	}
	return items
}

func TestPlayerGet(t *testing.T) {
	keeper, ctx := keepertest.GameKeeper(t)
	items := createNPlayer(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetPlayer(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestPlayerRemove(t *testing.T) {
	keeper, ctx := keepertest.GameKeeper(t)
	items := createNPlayer(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePlayer(ctx,
			item.Address,
		)
		_, found := keeper.GetPlayer(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestPlayerGetAll(t *testing.T) {
	keeper, ctx := keepertest.GameKeeper(t)
	items := createNPlayer(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPlayer(ctx)),
	)
}
