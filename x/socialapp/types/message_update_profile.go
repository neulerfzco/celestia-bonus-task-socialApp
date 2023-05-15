package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateProfile = "update_profile"

var _ sdk.Msg = &MsgUpdateProfile{}

func NewMsgUpdateProfile(creator string, handlename string, id string) *MsgUpdateProfile {
	return &MsgUpdateProfile{
		Creator:    creator,
		Handlename: handlename,
		Id:         id,
	}
}

func (msg *MsgUpdateProfile) Route() string {
	return RouterKey
}

func (msg *MsgUpdateProfile) Type() string {
	return TypeMsgUpdateProfile
}

func (msg *MsgUpdateProfile) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateProfile) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateProfile) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
