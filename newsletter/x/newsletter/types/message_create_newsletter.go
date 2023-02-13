package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateNewsletter = "create_newsletter"

var _ sdk.Msg = &MsgCreateNewsletter{}

func NewMsgCreateNewsletter(creator string, title string, description string, price uint64) *MsgCreateNewsletter {
	return &MsgCreateNewsletter{
		Creator:     creator,
		Title:       title,
		Description: description,
		Price:       price,
	}
}

func (msg *MsgCreateNewsletter) Route() string {
	return RouterKey
}

func (msg *MsgCreateNewsletter) Type() string {
	return TypeMsgCreateNewsletter
}

func (msg *MsgCreateNewsletter) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateNewsletter) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateNewsletter) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
