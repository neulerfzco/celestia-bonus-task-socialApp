package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/socialapp module sentinel errors
var (
	ErrProfileAlreadyExists = sdkerrors.Register(ModuleName, 1000, "This address already owns a profile, try updating it instead")
	ErrProfileInexistant    = sdkerrors.Register(ModuleName, 1001, "This address does not have an account")
	ErrHandleNameTaken      = sdkerrors.Register(ModuleName, 1010, "This handlename is already taken by someone else, try another one") // NOT IMPLEMENTED YET
	ErrPostInexistant       = sdkerrors.Register(ModuleName, 1100, "The post does not exist")
)
