package flchain_test

import (
	"testing"

	keepertest "github.com/kaushik-i-b/fl-chain/testutil/keeper"
	"github.com/kaushik-i-b/fl-chain/testutil/nullify"
	flchain "github.com/kaushik-i-b/fl-chain/x/flchain/module"
	"github.com/kaushik-i-b/fl-chain/x/flchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FlchainKeeper(t)
	flchain.InitGenesis(ctx, k, genesisState)
	got := flchain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
