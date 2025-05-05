package configuration

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	IneligiblePatterns []string `yaml:"ineligible_patterns"`
}

func (c *Config) Equals(other *Config) bool {
	if c == nil || other == nil {
		return c == other
	}
	if len(c.IneligiblePatterns) != len(other.IneligiblePatterns) {
		return false
	}
	for i, v := range c.IneligiblePatterns {
		if v != other.IneligiblePatterns[i] {
			return false
		}
	}
	return true
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, ErrFailedToReadFile(err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, ErrFailedToUnmarshal(err)
	}

	return &config, nil
}
