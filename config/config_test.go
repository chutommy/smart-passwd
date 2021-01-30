package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	tests := []struct {
		name string
		file string
		err  error
	}{
		{
			name: "success",
			file: "config/tests/config_ok.yml",
			err:  nil,
		},
		{
			name: "invalid",
			file: "config/tests/config_invalid.yml",
			err:  ErrInvalidYamlFile,
		},
		{
			name: "not found",
			file: "config/tests/config.yml",
			err:  ErrFileNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t1 *testing.T) {
			_, err := GetConfig(test.file)
			assert.Equal(t1, err, test.err)
		})
	}
}
