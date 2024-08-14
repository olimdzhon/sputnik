package sputnik_test

import (
	"testing"

	keepertest "github.com/olimdzhon/sputnik/testutil/keeper"
	"github.com/olimdzhon/sputnik/testutil/nullify"
	sputnik "github.com/olimdzhon/sputnik/x/sputnik/module"
	"github.com/olimdzhon/sputnik/x/sputnik/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SputnikKeeper(t)
	sputnik.InitGenesis(ctx, k, genesisState)
	got := sputnik.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
