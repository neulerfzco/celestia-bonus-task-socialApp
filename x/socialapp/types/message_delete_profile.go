package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDeleteProfile = "delete_profile"

var _ sdk.Msg = &MsgDeleteProfile{}

func NewMsgDeleteProfile(creator string, id string) *MsgDeleteProfile {
	return &MsgDeleteProfile{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgDeleteProfile) Route() string {
	return RouterKey
}

func (msg *MsgDeleteProfile) Type() string {
	return TypeMsgDeleteProfile
}

func (msg *MsgDeleteProfile) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteProfile) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteProfile) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
