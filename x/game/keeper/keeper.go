package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/BabbyBit/babychain/x/game/types"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority string

		bankKeeper types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,

	bankKeeper types.BankKeeper,
) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	return Keeper{
		cdc:          cdc,
		storeService: storeService,
		authority:    authority,
		logger:       logger,

		bankKeeper: bankKeeper,
	}
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) MintRewards(goCtx context.Context) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	currentBlock := uint64(ctx.BlockHeight())

	for _, player := range k.GetAllPlayer(ctx) {
		if currentBlock-player.LastIncomeBlock >= 6 {
			playerRewardAmount := calculateIncome(int64(player.CurrentLevel), types.RewardExponent)
			playerReward := sdk.NewCoin(types.RewardDenom, playerRewardAmount)
			playerRewards := sdk.NewCoins(playerReward)
			playerAddress, err := sdk.AccAddressFromBech32(player.Address)

			if err != nil {
				return err
			}

			if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, playerRewards); err != nil {
				return err
			}

			if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, playerAddress, playerRewards); err != nil {
				return err
			}

			player.LastIncomeBlock = currentBlock
			k.SetPlayer(ctx, player)
		}
	}

	return nil
}
