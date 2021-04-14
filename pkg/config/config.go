package config

import (
	"errors"
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// ErrNilValue is returned whenever variable with nil value is not expected.
var ErrNilValue = errors.New("invalid nil value")

// Keys represents the key values of the configuration.
const (
	KeyHTTPPort = "HTTPPort"
	KeyDBFile   = "DBFile"
	KeyDebug    = "Debug"
)

// Config represents an application configuration.
type Config struct {
	HTTPPort int64  `yaml:"HTTPPort"`
	DBFile   string `yaml:"DBFile"`
	Debug    bool   `yaml:"Debug"`
}

// File represents metadata of the configuration file.
type File struct {
	Name string
	Type string
	Path string
}

// GetConfig sets defaults, replaces them with the values from the configuration file
// and finally overrides them with flags.
func GetConfig(defCfg *Config, file *File, args []string) (*Config, error) {
	vi := viper.New()

	if err := setDefault(vi, defCfg); err != nil {
		return nil, fmt.Errorf("set value: %w", err)
	}

	if err := setFromFile(vi, file); err != nil {
		return nil, fmt.Errorf("set config from file: %w", err)
	}

	if err := setFromFlags(vi, args); err != nil {
		return nil, fmt.Errorf("set config from flags: %w", err)
	}

	cfg, err := decodeViper(vi)
	if err != nil {
		return nil, fmt.Errorf("decode viper to config: %w", err)
	}

	return cfg, nil
}

func setDefault(vi *viper.Viper, cfg *Config) error {
	if cfg == nil {
		return fmt.Errorf("cfg: %w", ErrNilValue)
	}

	vi.SetDefault(KeyHTTPPort, cfg.HTTPPort)
	vi.SetDefault(KeyDBFile, cfg.DBFile)
	vi.SetDefault(KeyDebug, cfg.Debug)

	return nil
}

func setFromFile(vi *viper.Viper, f *File) error {
	if err := fileData(vi, f); err != nil {
		return fmt.Errorf("set file data: %w", err)
	}

	if err := loadFile(vi); err != nil {
		return fmt.Errorf("load file: %w", err)
	}

	return nil
}

func loadFile(vi *viper.Viper) error {
	if err := vi.ReadInConfig(); err != nil {
		if !errors.Is(err, viper.ConfigFileNotFoundError{}) {
			return fmt.Errorf("read config file: %w", err)
		}

		if err = vi.SafeWriteConfig(); err != nil {
			return fmt.Errorf("generate config file: %w", err)
		}
	}

	return nil
}

func fileData(vi *viper.Viper, f *File) error {
	if f == nil {
		return fmt.Errorf("f: %w", ErrNilValue)
	}

	vi.SetConfigName(f.Name)
	vi.SetConfigType(f.Type)
	vi.AddConfigPath(f.Path)

	return nil
}

func setFromFlags(vi *viper.Viper, args []string) error {
	fs := pflag.NewFlagSet("smart-passwd", pflag.ContinueOnError)

	fs.Int64P(KeyHTTPPort, "p", vi.GetInt64(KeyHTTPPort), "port of the application to serve")
	fs.StringP(KeyDBFile, "f", vi.GetString(KeyDBFile), "path to SQLite3 database file")
	fs.BoolP(KeyDebug, "d", vi.GetBool(KeyDebug), "debug mode")

	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("parse flags: %w", err)
	}

	if err := vi.BindPFlags(fs); err != nil {
		return fmt.Errorf("bind pflags: %w", err)
	}

	return nil
}

func decodeViper(vi *viper.Viper) (*Config, error) {
	var cfg Config
	if err := vi.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal viper config: %w", err)
	}

	return &cfg, nil
}
