package keeper

import (
	"context"

	"github.com/BabbyBit/babychain/x/game/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Join(goCtx context.Context, msg *types.MsgJoin) (*types.MsgJoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if _, playerFound := k.GetPlayer(ctx, msg.Creator); playerFound {
		return nil, types.ErrPlayerAlreadyJoinedToGame
	}

	currentIncome := calculateIncome(1, types.RewardExponent)
	nextIncome := calculateIncome(2, types.RewardExponent)
	nextLevelPrice := calculateLevelPrice(2, types.RewardExponent)

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
