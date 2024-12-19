package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/kaushik-i-b/fl-chain/testutil/keeper"
	"github.com/kaushik-i-b/fl-chain/x/flchain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.FlchainKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
