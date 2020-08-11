package config

import (
	"io/ioutil"
	"path/filepath"
	"runtime"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// ErrFileNotFound is returned if the file was not found
type ErrFileNotFound error

// ErrInvalidYamlFile is return when the file does not satisfy the yaml file type.
type ErrInvalidYamlFile error

// Config defines a configuration for the web API.
type Config struct {
	Port     int       `yaml:"Port"`
	DBConfig *DBConfig `yaml:"DB"`
}

// GetConfig tries to load and handle the configuration file.
func GetConfig(path string) (*Config, error) {

	// load configuration
	content, err := ioutil.ReadFile(filepath.Join(rootDir(), "..", path))
	if err != nil {
		return nil, ErrFileNotFound(errors.Wrap(err, "could not read config file"))
	}

	// validate and apply the settings
	var cfg Config
	err = yaml.Unmarshal(content, &cfg)
	if err != nil {
		return nil, ErrInvalidYamlFile(errors.Wrap(err, "invalid config file"))
	}

	return &cfg, nil
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Dir(b)
}
