package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/functionx/fx-core/v5/x/evm/types"
)

func TestNewNoOpTracer(t *testing.T) {
	require.Equal(t, &types.NoOpTracer{}, types.NewNoOpTracer())
}
