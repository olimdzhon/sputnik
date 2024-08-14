package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/olimdzhon/sputnik/testutil/keeper"
	"github.com/olimdzhon/sputnik/x/sputnik/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.SputnikKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
