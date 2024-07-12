package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/BabbyBit/babychain/testutil/keeper"
	"github.com/BabbyBit/babychain/x/game/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.GameKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
