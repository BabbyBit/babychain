package keeper

import (
	"context"

	"github.com/BabbyBit/babychain/x/game/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func (k msgServer) Join(goCtx context.Context, msg *types.MsgJoin) (*types.MsgJoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	denom, found := k.bankKeeper.GetDenomMetaData(ctx, "uexperience")
	if !found {
		panic("uexperience denom not found")
	}

	denomUnit := FindDenomUnit(&denom, "experience")

	currentIncome := calculateIncome(1, denomUnit.Exponent)
	nextIncome := calculateIncome(2, denomUnit.Exponent)
	nextLevelPrice := calculateLevelPrice(2, denomUnit.Exponent)

	player := types.Player{
		Address:        msg.Creator,
		CurrentLevel:   1,
		CurrentIncome:  sdk.NewCoin("uexperience", currentIncome),
		NextIncome:     sdk.NewCoin("uexperience", nextIncome),
		NextLevelPrice: sdk.NewCoin("uexperience", nextLevelPrice),
	}

	k.SetPlayer(ctx, player)

	return &types.MsgJoinResponse{
		Player: &player,
	}, nil
}

func FindDenomUnit(metadata *banktypes.Metadata, unitName string) *banktypes.DenomUnit {
	for _, unit := range metadata.DenomUnits {
		if unit.Denom == unitName {
			return unit
		}
	}

	panic("denom not found")
}
