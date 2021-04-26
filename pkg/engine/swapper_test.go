package engine

import (
	"testing"

	"github.com/chutified/smart-passwd/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestNewSwapper(t *testing.T) {
	t.Parallel()

	require.NotEmpty(t, NewSwapper(utils.Alphabet(), utils.Specials(), utils.SwapList()))
}

func newSwapper(t *testing.T) *Swapper {
	t.Helper()

	return NewSwapper(utils.Alphabet(), utils.Specials(), utils.SwapList())
}

func TestSwapper_Alpha(t *testing.T) {
	t.Parallel()

	s := newSwapper(t)
	ok := false

	a := s.Alpha()
	for i := 0; i < 800; i++ {
		if s.Alpha() != a {
			ok = true

			break
		}
	}

	require.True(t, ok)
}

func TestSwapper_AlphaCap(t *testing.T) {
	t.Parallel()

	s := newSwapper(t)
	ok := false

	a := s.AlphaCap()
	for i := 0; i < 800; i++ {
		if s.AlphaCap() != a {
			ok = true

			break
		}
	}

	require.True(t, ok)
}

func TestSwapper_Special(t *testing.T) {
	t.Parallel()

	s := newSwapper(t)
	ok := false

	a := s.Special()
	for i := 0; i < 800; i++ {
		if s.AlphaCap() != a {
			ok = true

			break
		}
	}

	require.True(t, ok)
}

func TestSwapper_Num(t *testing.T) {
	t.Parallel()

	s := newSwapper(t)
	ok := false

	a := s.Num()
	for i := 0; i < 800; i++ {
		if s.Num() != a {
			ok = true

			break
		}
	}

	require.True(t, ok)
}

func TestSwapper_Swap(t *testing.T) {
	t.Parallel()

	s := newSwapper(t)
	ok := false

	a := s.Swap('a')
	for i := 0; i < 800; i++ {
		if s.Swap('a') != a {
			ok = true

			break
		}
	}

	require.True(t, ok)
}
