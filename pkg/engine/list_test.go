package engine

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAlphabet(t *testing.T) {
	t.Parallel()
	require.NotEmpty(t, alphabet())
}

func TestSpecials(t *testing.T) {
	t.Parallel()
	require.NotEmpty(t, specials())
}

func TestSwapList(t *testing.T) {
	t.Parallel()
	require.NotEmpty(t, swapList())
}
