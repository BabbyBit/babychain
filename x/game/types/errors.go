package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/game module sentinel errors
var (
	ErrInvalidSigner              = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrPlayerAlreadyJoinedToGame  = sdkerrors.Register(ModuleName, 1101, "player already joined to game")
	ErrPlayerNotFound             = sdkerrors.Register(ModuleName, 1102, "player not found")
	ErrNotEnoughBalanceForUpgrade = sdkerrors.Register(ModuleName, 1103, "not enough balance for upgrade")
)
