package config

import (
	"io/ioutil"
	"path/filepath"
	"runtime"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// ErrFileNotFound is returned if the file was not found
var ErrFileNotFound = errors.New("file config.yml was not found")

// ErrInvalidYamlFile is return when the file does not satisfy the yaml file type.
var ErrInvalidYamlFile = errors.New("file config.yaml has invalid content")

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
		return nil, ErrFileNotFound
	}

	// validate and apply the settings
	var cfg Config
	err = yaml.Unmarshal(content, &cfg)
	if err != nil {
		return nil, ErrInvalidYamlFile
	}

	return &cfg, nil
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Dir(b)
}
