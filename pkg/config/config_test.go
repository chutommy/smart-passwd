package config

import (
	"testing"

	"github.com/chutified/smart-passwd/pkg/utils"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func testDecode(t *testing.T, vi *viper.Viper) *Config {
	t.Helper()

	cfg, err := decodeViper(vi)
	require.NoError(t, err)
	require.NotNil(t, cfg)

	return cfg
}

func TestGetConfig(t *testing.T) {
	t.Parallel()

	type input struct {
		defCfg *Config
		file   *utils.File
		args   []string
	}

	type output struct {
		cfg     *Config
		wantErr bool
	}

	tests := []struct {
		name string
		inp  input
		out  output
	}{
		{
			name: "complete",
			inp: input{
				defCfg: NewConfig(8080, "data/words-test.db", true),
				file:   utils.NewFile("test", "config4", "yaml"),
				args: []string{
					"--" + KeyHTTPPort, "10500",
					"--" + KeyDBFile, "data/words-prod.db",
					"--" + KeyDebug,
				},
			},
			out: output{
				cfg:     NewConfig(10500, "data/words-prod.db", true),
				wantErr: false,
			},
		},
		{
			name: "default and file values",
			inp: input{
				defCfg: NewConfig(8080, "data/words-test.db", true),
				file:   utils.NewFile("test", "config4", "yaml"),
				args:   nil,
			},
			out: output{
				cfg:     NewConfig(80, "data/words.db", false),
				wantErr: false,
			},
		},
		{
			name: "default and empty config",
			inp: input{
				defCfg: NewConfig(8080, "data/words-test.db", true),
				file:   utils.NewFile("test", "config5", "yaml"),
				args:   nil,
			},
			out: output{
				cfg:     NewConfig(8080, "data/words-test.db", true),
				wantErr: false,
			},
		},
		{
			name: "nil file",
			inp: input{
				defCfg: NewConfig(8080, "data/words-test.db", true),
				file:   nil,
				args:   nil,
			},
			out: output{
				cfg:     nil,
				wantErr: true,
			},
		},
		{
			name: "nil arguments",
			inp: input{
				defCfg: nil,
				file:   nil,
				args:   nil,
			},
			out: output{
				cfg:     nil,
				wantErr: true,
			},
		},
		{
			name: "invalid flags",
			inp: input{
				defCfg: NewConfig(8080, "data/words-test.db", true),
				file:   utils.NewFile("test", "config4", "yaml"),
				args:   []string{"-invalid"},
			},
			out: output{
				cfg:     nil,
				wantErr: true,
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			cfg, err := GetConfig(tt.inp.defCfg, tt.inp.file, tt.inp.args)
			if tt.out.wantErr {
				require.Error(t, err)
				require.Nil(t, cfg)
			} else {
				require.NoError(t, err)
				require.NotNil(t, cfg)

				require.Equal(t, *tt.out.cfg, *cfg)
			}
		})
	}
}

func TestSetDefault(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		cfg     *Config
		wantErr bool
	}{
		{
			name:    "empty values",
			cfg:     NewConfig(0, "", false),
			wantErr: false,
		},
		{
			name:    "default values",
			cfg:     NewConfig(80, "words.db", false),
			wantErr: false,
		},
		{
			name:    "complete",
			cfg:     NewConfig(10503, "data/words-test.db", true),
			wantErr: false,
		},
		{
			name:    "nil config",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			vi := viper.New()
			err := setDefault(vi, tt.cfg)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				cfg := testDecode(t, vi)

				require.Equal(t, tt.cfg.HTTPPort, cfg.HTTPPort)
				require.Equal(t, tt.cfg.DBFile, cfg.DBFile)
				require.Equal(t, tt.cfg.Debug, cfg.Debug)
			}
		})
	}
}

func TestSetFromFile(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		file    *utils.File
		cfg     *Config
		wantErr bool
	}{
		{
			name:    "default",
			file:    utils.NewFile("test", "config1", "yaml"),
			cfg:     NewConfig(80, "data/words.db", false),
			wantErr: false,
		},
		{
			name:    "empty",
			file:    utils.NewFile("test", "config2", "yaml"),
			cfg:     NewConfig(0, "", false),
			wantErr: false,
		},
		{
			name:    "debug",
			file:    utils.NewFile("test", "config3", "yaml"),
			cfg:     NewConfig(8080, "data/words-test.db", true),
			wantErr: false,
		},
		{
			name:    "na file",
			file:    utils.NewFile("test", "na", "yaml"),
			wantErr: true,
		},
		{
			name:    "nil file",
			file:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			vi := viper.New()
			err := setFromFile(vi, tt.file)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				cfg := testDecode(t, vi)

				require.Equal(t, tt.cfg.HTTPPort, cfg.HTTPPort)
				require.Equal(t, tt.cfg.DBFile, cfg.DBFile)
				require.Equal(t, tt.cfg.Debug, cfg.Debug)
			}
		})
	}
}

// TestLoadFile test whether loading file without setting
// its data returns an error.
func TestLoadFile(t *testing.T) {
	t.Parallel()
	require.Error(t, loadFile(viper.New()))
}

func TestSetFromFlags(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		arg     []string
		cfg     *Config
		wantErr bool
	}{
		{
			name:    "empty",
			arg:     []string{},
			cfg:     NewConfig(0, "", false),
			wantErr: false,
		},
		{
			name:    "empty with nil",
			arg:     nil,
			cfg:     NewConfig(0, "", false),
			wantErr: false,
		},
		{
			name:    "boolean short",
			arg:     []string{"-d"},
			cfg:     NewConfig(0, "", true),
			wantErr: false,
		},
		{
			name: "shorthands with space",
			arg: []string{
				"-p", "1313",
				"-f", "data/words.db",
				"-d", "true",
			},
			cfg:     NewConfig(1313, "data/words.db", true),
			wantErr: false,
		},
		{
			name: "full name with spaces",
			arg: []string{
				"--" + KeyHTTPPort, "1313",
				"--" + KeyDBFile, "data/words.db",
				"--" + KeyDebug, "true",
			},
			cfg:     NewConfig(1313, "data/words.db", true),
			wantErr: false,
		},
		{
			name: "shorthands with equals",
			arg: []string{
				"-p=1313",
				"-f=data/words.db",
				"-d=true",
			},
			cfg:     NewConfig(1313, "data/words.db", true),
			wantErr: false,
		},
		{
			name: "full name with equals",
			arg: []string{
				"--" + KeyHTTPPort + "=1313",
				"--" + KeyDBFile + "=data/words.db",
				"--" + KeyDebug + "=true",
			},
			cfg:     NewConfig(1313, "data/words.db", true),
			wantErr: false,
		},
		{
			name:    "invalid flag",
			arg:     []string{"--invalid", "value"},
			wantErr: true,
		},
		{
			name:    "invalid value",
			arg:     []string{"-p"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			vi := viper.New()
			err := setFromFlags(vi, tt.arg)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				cfg := testDecode(t, vi)

				require.Equal(t, tt.cfg.HTTPPort, cfg.HTTPPort)
				require.Equal(t, tt.cfg.DBFile, cfg.DBFile)
				require.Equal(t, tt.cfg.Debug, cfg.Debug)
			}
		})
	}
}

func TestDecodeViper(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		buildViper func(vi *viper.Viper)
		cfg        *Config
		wantErr    bool
	}{
		{
			name: "default",
			buildViper: func(vi *viper.Viper) {
				vi.Set(KeyHTTPPort, 80)
				vi.Set(KeyDBFile, "data/words-1.db")
				vi.Set(KeyDebug, false)
			},
			cfg:     NewConfig(80, "data/words-1.db", false),
			wantErr: false,
		},
		{
			name: "debug",
			buildViper: func(vi *viper.Viper) {
				vi.Set(KeyHTTPPort, 8080)
				vi.Set(KeyDBFile, "data/words-test.db")
				vi.Set(KeyDebug, true)
			},
			cfg:     NewConfig(8080, "data/words-test.db", true),
			wantErr: false,
		},
		{
			name: "invalid type",
			buildViper: func(vi *viper.Viper) {
				vi.Set(KeyHTTPPort, "invalid int")
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			vi := viper.New()
			tt.buildViper(vi)
			cfg, err := decodeViper(vi)
			if tt.wantErr {
				require.Error(t, err)
				require.Nil(t, cfg)
			} else {
				require.NoError(t, err)
				require.NotNil(t, cfg)

				cfg := testDecode(t, vi)

				require.Equal(t, tt.cfg.HTTPPort, cfg.HTTPPort)
				require.Equal(t, tt.cfg.DBFile, cfg.DBFile)
				require.Equal(t, tt.cfg.Debug, cfg.Debug)
			}
		})
	}
}
