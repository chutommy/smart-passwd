package engine

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func sumLen(a []string) int16 {
	var sum int
	for _, v := range a {
		sum += len(v)
	}

	return int16(sum)
}

func TestEngine_helperGen(t *testing.T) {
	t.Parallel()

	e := testEngine
	ei := testInvalidEngine

	tests := []struct {
		name    string
		length  int16
		wantErr bool
	}{
		{
			name:    "basic 10",
			length:  10,
			wantErr: false,
		},
		{
			name:    "basic 100",
			length:  100,
			wantErr: false,
		},
		{
			name:    "basic 1000",
			length:  1000,
			wantErr: false,
		},
		{
			name:    "negative length",
			length:  -1,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ll, err := e.helperGen(context.Background(), tt.length)
			if tt.wantErr {
				require.Error(t, err)
				require.Nil(t, ll)
			} else {
				require.NoError(t, err)
				require.NotNil(t, ll)

				require.Equal(t, tt.length, sumLen(ll))
			}
		})
	}

	// test invalid wordlist
	t.Run("invalid wordlist", func(t *testing.T) {
		ll, err := ei.helperGen(context.Background(), 10)
		require.Error(t, err)
		require.Nil(t, ll)
	})
}

func TestEngine_helper(t *testing.T) {
	t.Parallel()

	e := testEngine

	tests := []struct {
		name    string
		helper  string
		length  int16
		wantErr bool
	}{
		{
			name:    "ok short",
			helper:  "abc abcde",
			wantErr: false,
		},
		{
			name:    "ok long",
			helper:  "abcde abcde abcde abcde abcde",
			wantErr: false,
		},
		{
			name:    "len 10",
			length:  10,
			wantErr: false,
		},
		{
			name:    "len 100",
			length:  100,
			wantErr: false,
		},
		{
			name:    "empty",
			wantErr: true,
		},
		{
			name:    "both provided",
			helper:  "abcde abcdefg",
			length:  5,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			h, err := e.helper(context.Background(), NewRequest(tt.length, 0, tt.helper))
			if tt.wantErr {
				require.Error(t, err)
				require.Nil(t, h)
			} else {
				require.NoError(t, err)
				require.NotNil(t, h)

				if tt.helper != "" {
					require.Equal(t, sumLen(strings.Split(tt.helper, " ")), sumLen(h))
				} else {
					require.Equal(t, tt.length, sumLen(h))
				}
			}
		})
	}
}

func TestEngine_swap(t *testing.T) {
	t.Parallel()

	e := testEngine

	tests := []struct {
		name   string
		helper string
	}{
		{
			name:   "one word",
			helper: "abcdef",
		},
		{
			name:   "two words",
			helper: "abcd abcdef",
		},
		{
			name:   "three words",
			helper: "abc abcde abcdef",
		},
		{
			name:   "empty",
			helper: "",
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			h := e.swap(strings.Split(tt.helper, " "))
			require.Len(t, h, int(sumLen(strings.Split(tt.helper, " "))))
		})
	}
}

func TestEngine_Generate(t *testing.T) {
	t.Parallel()

	e := testEngine

	tests := []struct {
		name    string
		req     *Request
		wantErr bool
	}{
		{
			name: "invalid length 0",
			req: &Request{
				length: 0,
			},
			wantErr: true,
		},
		{
			name: "invalid length 2", // too short
			req: &Request{
				length: 2,
			},
			wantErr: true,
		},
		{
			name: "length 10",
			req: &Request{
				length: 10,
			},
			wantErr: false,
		},
		{
			name: "length 100",
			req: &Request{
				length: 100,
			},
			wantErr: false,
		},
		{
			name: "invalid length helper",
			req: &Request{
				length: 10,
				helper: "abcde",
			},
			wantErr: true,
		},
		{
			name: "short custom helper ab",
			req: &Request{
				helper: "ab",
			},
			wantErr: false,
		},
		{
			name: "helper abcde",
			req: &Request{
				helper: "abcde",
			},
			wantErr: false,
		},
		{
			name: "helper abc abc",
			req: &Request{
				helper: "abc abc",
			},
			wantErr: false,
		},
		{
			name: "extra sec -1",
			req: &Request{
				length:   10,
				extraSec: -1,
			},
			wantErr: true,
		},
		{
			name: "extra sec 10",
			req: &Request{
				length:   10,
				extraSec: 10,
			},
			wantErr: false,
		},
		{
			name: "extra sec 100",
			req: &Request{
				length:   10,
				extraSec: 100,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			resp, err := e.Generate(context.Background(), tt.req)
			if tt.wantErr {
				require.Error(t, err)
				require.Nil(t, resp)
			} else {
				require.NoError(t, err)
				require.NotNil(t, resp)

				if tt.req.helper == "" {
					pl := int(tt.req.length + tt.req.extraSec)
					hl := sumLen(strings.Split(resp.helper, " "))
					require.Len(t, resp.password, pl)
					require.Equal(t, int(hl+tt.req.extraSec), pl)
				} else {
					pl := int(sumLen(strings.Split(tt.req.helper, " ")) + tt.req.extraSec)
					require.Len(t, resp.password, pl)
					require.Equal(t, tt.req.helper, resp.helper)
				}
			}
		})
	}
}
