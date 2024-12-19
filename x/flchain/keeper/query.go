package keeper

import (
	"github.com/kaushik-i-b/fl-chain/x/flchain/types"
)

var _ types.QueryServer = Keeper{}
