package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	legacyerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRedelegate = "redelegate"

var _ sdk.Msg = &MsgRedelegate{}

func NewMsgRedelegate(delegator string, from_provider string, from_chainID string, to_provider string, to_chainID string, amount sdk.Coin) *MsgRedelegate {
	return &MsgRedelegate{
		Delegator:    delegator,
		FromProvider: from_provider,
		FromChainID:  from_chainID,
		ToProvider:   to_provider,
		ToChainID:    to_chainID,
		Amount:       amount,
	}
}

func (msg *MsgRedelegate) Route() string {
	return RouterKey
}

func (msg *MsgRedelegate) Type() string {
	return TypeMsgRedelegate
}

func (msg *MsgRedelegate) GetSigners() []sdk.AccAddress {
	delegator, err := sdk.AccAddressFromBech32(msg.Delegator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{delegator}
}

func (msg *MsgRedelegate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRedelegate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Delegator)
	if err != nil {
		return sdkerrors.Wrapf(legacyerrors.ErrInvalidAddress, "invalid delegator address (%s)", err)
	}

	if !msg.Amount.IsValid() {
		return legacyerrors.ErrInvalidCoins
	}

	return nil
}
