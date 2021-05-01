package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRand(t *testing.T) {
	t.Parallel()

	r := Rand()
	i1 := r.Int()
	i2 := r.Int()
	require.NotEqual(t, i1, i2)

	i3 := Rand().Int()
	i4 := Rand().Int()
	require.NotEqual(t, i3, i4)
}
