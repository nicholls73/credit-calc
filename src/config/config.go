package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	IneligiblePatterns     []string  `yaml:"ineligible_patterns"`
	SpendAmountForBonus    float64   `yaml:"spend_amount_for_bonus"`
	BonusSpendPeriodInDays int       `yaml:"bonus_spend_period_in_days"`
	CardStartDate          time.Time `yaml:"card_start_date"`
}

func (c *Config) Equals(other *Config) bool {
	if c == nil || other == nil {
		return c == other
	}
	if c.SpendAmountForBonus != other.SpendAmountForBonus {
		return false
	}
	if c.BonusSpendPeriodInDays != other.BonusSpendPeriodInDays {
		return false
	}
	if !c.CardStartDate.Equal(other.CardStartDate) {
		return false
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
