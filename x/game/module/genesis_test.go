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

		PlayerList: []types.Player{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GameKeeper(t)
	game.InitGenesis(ctx, k, genesisState)
	got := game.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.PlayerList, got.PlayerList)
	// this line is used by starport scaffolding # genesis/test/assert
}
