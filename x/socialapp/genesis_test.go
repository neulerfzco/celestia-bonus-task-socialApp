package socialapp_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "socialapp/testutil/keeper"
	"socialapp/testutil/nullify"
	"socialapp/x/socialapp"
	"socialapp/x/socialapp/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SocialappKeeper(t)
	socialapp.InitGenesis(ctx, *k, genesisState)
	got := socialapp.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
