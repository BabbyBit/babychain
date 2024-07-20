package keeper

import (
	"context"

	"github.com/BabbyBit/babychain/x/game/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Upgrade(goCtx context.Context, msg *types.MsgUpgrade) (*types.MsgUpgradeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	player, playerFound := k.GetPlayer(ctx, msg.Creator)

	if !playerFound {
		return nil, types.ErrPlayerNotFound
	}

	playerAddress, _ := sdk.AccAddressFromBech32(msg.Creator)
	playerBalance := k.bankKeeper.GetBalance(ctx, playerAddress, types.RewardDenom)

	upgradeAvailable := playerBalance.Amount.GTE(player.NextLevelPrice.Amount)

	if !upgradeAvailable {
		return nil, types.ErrNotEnoughBalanceForUpgrade
	}

	player.CurrentLevel++
	player.CurrentIncome.Amount = player.NextIncome.Amount
	player.NextIncome.Amount = calculateIncome(int64(player.CurrentLevel+1), types.RewardExponent)
	player.NextLevelPrice.Amount = calculateLevelPrice(int64(player.CurrentLevel+1), types.RewardExponent)

	k.SetPlayer(ctx, player)

	return &types.MsgUpgradeResponse{
		Player: &player,
	}, nil
}
