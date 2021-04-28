package engine

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAlphabet(t *testing.T) {
	t.Parallel()
	require.NotEmpty(t, Alphabet())
}

func TestSpecials(t *testing.T) {
	t.Parallel()
	require.NotEmpty(t, Specials())
}

func TestSwapList(t *testing.T) {
	t.Parallel()
	require.NotEmpty(t, SwapList())
}
