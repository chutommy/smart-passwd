package engine

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func sumInt16(t *testing.T, a []int16) int16 {
	t.Helper()

	var sum int16
	for _, i := range a {
		sum += i
	}

	return sum
}

func TestConstructor(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		min     int16
		max     int16
		l       int16
		wantErr bool
	}{
		{
			name:    "invalid",
			min:     2,
			max:     8,
			l:       0,
			wantErr: true,
		},
		{
			name:    "small",
			min:     2,
			max:     8,
			l:       100,
			wantErr: false,
		},
		{
			name:    "large",
			min:     8,
			max:     21,
			l:       1000,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := NewConstructor(tt.min, tt.max)

			for i := 0; i < 10; i++ {
				ll, err := c.Distribute(tt.l)
				if tt.wantErr {
					require.Error(t, err)
					require.Nil(t, ll)
				} else {
					require.NoError(t, err)
					require.Equal(t, tt.l, sumInt16(t, ll))
				}
			}
		})
	}
}
