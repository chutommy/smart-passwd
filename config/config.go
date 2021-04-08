package config

import (
	"errors"
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

// Keys represents the key values of the configuration.
const (
	KeyHTTPPort = "HTTPPort"
	KeyDBFile   = "DBFile"
	KeyDebug    = "Debug"
)

// File represents metadata of the configuration file.
type File struct {
	Name string
	Type string
	Path string
}

// Config represents an application configuration.
type Config struct {
	HTTPPort int64  `yaml:"HTTPPort"`
	DBFile   string `yaml:"DBFile"`
	Debug    bool   `yaml:"Debug"`
}

// GetConfig sets defaults, replaces them with the values from the configuration file
// and finally overrides them with flags.
func GetConfig(defCfg *Config, file *File, args []string) (*Config, error) {
	vi := viper.New()

	setDefault(vi, defCfg)

	if err := setFromFile(vi, file); err != nil {
		return nil, fmt.Errorf("set config from file: %w", err)
	}

	if err := setFromFlags(vi, args); err != nil {
		return nil, fmt.Errorf("set config from flags: %w", err)
	}

	var cfg Config
	if err := vi.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal viper config: %w", err)
	}

	return &cfg, nil
}

func setDefault(vi *viper.Viper, cfg *Config) {
	vi.SetDefault(KeyHTTPPort, cfg.HTTPPort)
	vi.SetDefault(KeyDBFile, cfg.DBFile)
	vi.SetDefault(KeyDebug, cfg.Debug)
}

func setFromFile(vi *viper.Viper, f *File) error {
	fileData(vi, f)

	if err := loadFile(vi); err != nil {
		return fmt.Errorf("")
	}

	return nil
}

func loadFile(vi *viper.Viper) error {
	err := vi.ReadInConfig()
	if err != nil {
		if !errors.Is(err, viper.ConfigFileNotFoundError{}) {
			return fmt.Errorf("read config file: %w", err)
		}

		err = viper.SafeWriteConfig()
		if err != nil {
			return fmt.Errorf("generate config file: %w", err)
		}
	}
	return nil
}

func fileData(vi *viper.Viper, f *File) {
	vi.SetConfigName(f.Name)
	vi.SetConfigType(f.Type)
	vi.AddConfigPath(f.Path)
}

func setFromFlags(vi *viper.Viper, args []string) error {
	fs := flag.NewFlagSet("Smart Pass-WD", flag.ExitOnError)

	vi.Set(KeyHTTPPort, *fs.Int64("http", vi.GetInt64(KeyHTTPPort), "port of the application to serve"))
	vi.Set(KeyDBFile, *fs.String("db", vi.GetString(KeyDBFile), "path to SQLite3 database file"))
	vi.Set(KeyDebug, *fs.Bool("debug", vi.GetBool(KeyDebug), "debug mode"))

	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("parse flags: %w", err)
	}

	return nil
}
