package keeper

import (
	"context"

	"github.com/BabbyBit/babychain/x/game/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Join(goCtx context.Context, msg *types.MsgJoin) (*types.MsgJoinResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgJoinResponse{}, nil
}
