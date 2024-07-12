package game_test

import (
	"testing"

	keepertest "github.com/BabbyBit/babychain/testutil/keeper"
	"github.com/BabbyBit/babychain/testutil/nullify"
	game "github.com/BabbyBit/babychain/x/game/module"
	"github.com/BabbyBit/babychain/x/game/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GameKeeper(t)
	game.InitGenesis(ctx, k, genesisState)
	got := game.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
