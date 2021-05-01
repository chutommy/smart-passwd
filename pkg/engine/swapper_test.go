package engine

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewSwapper(t *testing.T) {
	t.Parallel()

	require.NotEmpty(t, NewSwapper(Alphabet(), Specials(), SwapList()))
}

func newSwapper(t *testing.T) *Swapper {
	t.Helper()

	return NewSwapper(Alphabet(), Specials(), SwapList())
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

	// non-listed character
	b := s.Swap('.')
	require.Equal(t, '.', b)
}

func TestExtras(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		lvl      int16
		expSpecs int16
		expNums  int16
	}{
		{
			name:     "zero",
			lvl:      0,
			expSpecs: 0,
			expNums:  0,
		},
		{
			name:     "one",
			lvl:      1,
			expSpecs: 0,
			expNums:  1,
		},
		{
			name:     "teo",
			lvl:      2,
			expSpecs: 0,
			expNums:  2,
		},
		{
			name:     "three",
			lvl:      3,
			expSpecs: 1,
			expNums:  2,
		},
		{
			name:     "four",
			lvl:      4,
			expSpecs: 1,
			expNums:  3,
		},
		{
			name:     "five",
			lvl:      5,
			expSpecs: 1,
			expNums:  4,
		},
		{
			name:     "six",
			lvl:      6,
			expSpecs: 2,
			expNums:  4,
		},
		{
			name:     "seven",
			lvl:      7,
			expSpecs: 2,
			expNums:  5,
		},
		{
			name:     "eight",
			lvl:      8,
			expSpecs: 2,
			expNums:  6,
		},
		{
			name:     "nine",
			lvl:      9,
			expSpecs: 3,
			expNums:  6,
		},
		{
			name:     "one hundred",
			lvl:      100,
			expSpecs: 33,
			expNums:  67,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			a, b := extras(tt.lvl)
			require.Equal(t, tt.lvl, a+b)
			require.Equal(t, tt.expSpecs, a)
			require.Equal(t, tt.expNums, b)
		})
	}
}

func TestSwapper_ExtraSec(t *testing.T) {
	t.Parallel()

	s := newSwapper(t)
	tests := []struct {
		name string
		text string
		lvl  int16
	}{
		{
			name: "foo 10",
			text: "foo",
			lvl:  10,
		},
		{
			name: "empty 10",
			text: "",
			lvl:  10,
		},
		{
			name: "foobar 0",
			text: "foobar",
			lvl:  0,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			out := s.ExtraSec(tt.text, tt.lvl)
			require.Len(t, out, len(tt.text)+int(tt.lvl))
		})
	}
}
