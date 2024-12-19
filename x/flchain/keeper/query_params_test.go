package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/kaushik-i-b/fl-chain/testutil/keeper"
	"github.com/kaushik-i-b/fl-chain/x/flchain/types"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := keepertest.FlchainKeeper(t)
	params := types.DefaultParams()
	require.NoError(t, keeper.SetParams(ctx, params))

	response, err := keeper.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
