package keeper

import (
	"github.com/BabbyBit/babychain/x/game/types"
)

var _ types.QueryServer = Keeper{}
