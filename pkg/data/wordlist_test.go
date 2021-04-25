package data

import (
	"testing"

	"github.com/chutified/smart-passwd/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestConnect(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		file    *utils.File
		wantErr bool
	}{
		{
			name:    "test file",
			file:    testDBFile,
			wantErr: false,
		},
		{
			name:    "nil file",
			file:    nil,
			wantErr: true,
		},
		{
			name:    "non found file",
			file:    utils.NewFile("test", "na", "db"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			wl, err := Connect(tt.file)
			if tt.wantErr {
				require.Error(t, err)
				require.Nil(t, wl)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, wl)

				err = wl.Close()
				require.NoError(t, err)
			}
		})
	}
}

func TestWordList_Word(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		len     int16
		wantErr bool
	}{
		{
			name:    "ok",
			len:     8,
			wantErr: false,
		},
		{
			name:    "zero",
			len:     0,
			wantErr: true,
		},
		{
			name:    "negative",
			len:     -20,
			wantErr: true,
		},
		{
			name:    "over max",
			len:     23,
			wantErr: true,
		},
	}

	wl := WordList{testDB}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			w, err := wl.Word(tt.len)
			if tt.wantErr {
				require.Error(t, err)
				require.Empty(t, w)
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, w)
				require.Len(t, w, int(tt.len))
			}
		})
	}
}
