package keeper

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/kaushik-i-b/fl-chain/x/flchain/types"
)

func (k Keeper) GetLoanFeeRate(ctx sdk.Context) sdk.DecCoin {
	// Define 5% fee rate using LegacyNewDec
	rate := math.LegacyNewDec(5).QuoInt64(100) // 0.05 or 5%
	return sdk.NewDecCoinFromDec("token", rate)
}

func (k Keeper) GetMaxLoanAmount(ctx sdk.Context) sdk.Coin {
	// Define the maximum loan amount
	intVal := math.NewInt(1000000)
	return sdk.NewCoin("token", intVal)
}

func (k Keeper) CallCallback(ctx sdk.Context, callback string) error {
	// Implement the callback logic here
	return nil
}

func (k Keeper) FlashLoan(ctx sdk.Context, msg types.MsgFlashLoan) error {
	// Ensure the loan is within allowed limits
	if msg.Amount.Amount.GT(k.GetMaxLoanAmount(ctx).Amount) {
		return sdkerrors.ErrInvalidRequest.Wrap("loan amount exceeds limit")
	}

	// Deduct loan fee
	fee := sdk.NewCoin(
		msg.Amount.Denom,
		math.LegacyNewDecFromInt(msg.Amount.Amount).Mul(k.GetLoanFeeRate(ctx).Amount).TruncateInt(),
	)
	// Transfer the loan amount
	err := k.bankKeeper.SendCoins(ctx, k.moduleAddress, msg.Borrower, sdk.NewCoins(msg.Amount))
	if err != nil {
		return err
	}

	// Call the callback function
	if err := k.CallCallback(ctx, msg.Callback); err != nil {
		return err
	}

	// Check repayment
	balance := k.bankKeeper.GetBalance(ctx, k.moduleAddress, msg.Amount.Denom)
	if balance.Amount.LT(msg.Amount.Amount.Add(fee.Amount)) {
		return sdkerrors.ErrInsufficientFunds
	}

	return nil
}
