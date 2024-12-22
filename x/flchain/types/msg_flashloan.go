package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const RouterKey = "flchain"

// MsgFlashLoan defines a message for requesting a flash loan.
type MsgFlashLoan struct {
	Borrower sdk.AccAddress `json:"borrower"`
	Amount   sdk.Coin       `json:"amount"`
	Callback string         `json:"callback"`
}

// Route implements the sdk.Msg interface, returning the module name.
func (msg MsgFlashLoan) Route() string {
	return RouterKey
}

// Type implements the sdk.Msg interface, returning the action type.
func (msg MsgFlashLoan) Type() string {
	return "FlashLoan"
}

// ValidateBasic performs stateless validation of the message fields.
func (msg MsgFlashLoan) ValidateBasic() error {
	if msg.Borrower.Empty() {
		return sdkerrors.ErrInvalidAddress.Wrap("borrower address cannot be empty")
	}
	if msg.Amount.IsZero() || msg.Amount.IsNegative() {
		return sdkerrors.ErrInvalidRequest.Wrap("amount must be positive")
	}
	if len(msg.Callback) == 0 {
		return sdkerrors.ErrInvalidRequest.Wrap("callback function required")
	}
	return nil
}

// GetSigners defines whose signature is required for the message.
func (msg MsgFlashLoan) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Borrower}
}
