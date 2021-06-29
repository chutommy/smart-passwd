package config

import (
	"errors"
	"fmt"

	"github.com/chutommy/smart-passwd/pkg/utils"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Keys represents the key values of the configuration.
const (
	KeyHTTPPort = "HTTPPort"
	KeyDBFile   = "DBFile"
	KeyDebug    = "Debug"
	KeyRootPath = "RootPath"
)

// Config represents an application configuration.
type Config struct {
	HTTPPort int64  `yaml:"HTTPPort"`
	DBFile   string `yaml:"DBFile"`
	Debug    bool   `yaml:"Debug"`
	RootPath string `yaml:"RootPath"`
}

// NewConfig is a constructor of the Config struct.
func NewConfig(httpPort int64, dbFile string, debug bool, rootPath string) *Config {
	return &Config{
		HTTPPort: httpPort,
		DBFile:   dbFile,
		Debug:    debug,
		RootPath: rootPath,
	}
}

// GetConfig sets defaults, replaces them with the values from the configuration file
// and finally overrides them with flags.
func GetConfig(defaultCfg *Config, file *utils.File, args []string) (*Config, error) {
	vi := viper.New()

	if err := setDefault(vi, defaultCfg); err != nil {
		return nil, fmt.Errorf("set default values: %w", err)
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
		return fmt.Errorf("cfg: %w", utils.ErrNilValue)
	}

	vi.SetDefault(KeyHTTPPort, cfg.HTTPPort)
	vi.SetDefault(KeyDBFile, cfg.DBFile)
	vi.SetDefault(KeyDebug, cfg.Debug)
	vi.SetDefault(KeyRootPath, cfg.RootPath)

	return nil
}

func setFromFile(vi *viper.Viper, f *utils.File) error {
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
	}

	return nil
}

func fileData(vi *viper.Viper, f *utils.File) error {
	if f == nil {
		return fmt.Errorf("f: %w", utils.ErrNilValue)
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
	fs.StringP(KeyRootPath, "r", vi.GetString(KeyRootPath), "project root directory")

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
