package config_test

import (
	c "credit-calc/config"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const testFileContent = `
ineligible_patterns:
  - TEST_PATTERN_1
  - TEST_PATTERN_2
  - TEST_PATTERN_3

spend_amount_for_bonus: 1000
bonus_spend_period_in_days: 30
card_start_date: 2025-01-01
`

func createTestFile(t *testing.T, content []byte) string {
	t.Helper()

	testFile, err := os.CreateTemp("", "test-*.yaml")
	if err != nil {
		t.Fatal(err)
	}

	defer testFile.Close()

	if _, err := testFile.Write(content); err != nil {
		t.Fatal(err)
	}

	filename := testFile.Name()
	t.Cleanup(func() {
		os.Remove(filename)
	})

	return filename
}

func TestLoadConfig_Equals(t *testing.T) {
	t.Parallel()

	equalConfig := &c.Config{
		IneligiblePatterns:     []string{"TEST_PATTERN_1", "TEST_PATTERN_2", "TEST_PATTERN_3"},
		SpendAmountForBonus:    1000,
		BonusSpendPeriodInDays: 30,
		CardStartDate:          time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	inequalConfig := &c.Config{
		IneligiblePatterns:     []string{"TEST_PATTERN_4", "TEST_PATTERN_5", "TEST_PATTERN_6"},
		SpendAmountForBonus:    2000,
		BonusSpendPeriodInDays: 60,
		CardStartDate:          time.Date(2026, 2, 2, 0, 0, 0, 0, time.UTC),
	}

	require.True(t, equalConfig.Equals(equalConfig))
	require.False(t, equalConfig.Equals(inequalConfig))
}

func TestLoadConfig_ValidConfig(t *testing.T) {
	t.Parallel()

	filePath := createTestFile(t, []byte(testFileContent))

	expectedConfig := &c.Config{
		IneligiblePatterns:     []string{"TEST_PATTERN_1", "TEST_PATTERN_2", "TEST_PATTERN_3"},
		SpendAmountForBonus:    1000,
		BonusSpendPeriodInDays: 30,
		CardStartDate:          time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	config, err := c.LoadConfig(filePath)

	require.NoError(t, err)
	require.NotNil(t, config)
	require.True(t, expectedConfig.Equals(config))
}

func TestLoadConfig_InvalidFile(t *testing.T) {
	t.Parallel()

	filename := "invalid-file.yaml"

	config, err := c.LoadConfig(filename)

	require.ErrorContains(t, err, c.ErrFailedToReadFileMsg)
	require.Nil(t, config)
}

func TestLoadConfig_InvalidSpendAmountForBonus(t *testing.T) {
	t.Parallel()

	const invalidFileContent = `
ineligible_patterns:
	- TEST_PATTERN_1
	- TEST_PATTERN_2
	- TEST_PATTERN_3

spend_amount_for_bonus: invalid
bonus_spend_period_in_days: 30
card_start_date: 2025-01-01
`

	filePath := createTestFile(t, []byte(invalidFileContent))

	config, err := c.LoadConfig(filePath)

	require.ErrorContains(t, err, c.ErrFailedToUnmarshalMsg)
	require.Nil(t, config)
}

func TestLoadConfig_InvalidBonusSpendPeriodInDays(t *testing.T) {
	t.Parallel()

	const invalidFileContent = `
ineligible_patterns:
	- TEST_PATTERN_1
	- TEST_PATTERN_2
	- TEST_PATTERN_3

spend_amount_for_bonus: 1000
bonus_spend_period_in_days: invalid
card_start_date: 2025-01-01
`

	filePath := createTestFile(t, []byte(invalidFileContent))

	config, err := c.LoadConfig(filePath)

	require.ErrorContains(t, err, c.ErrFailedToUnmarshalMsg)
	require.Nil(t, config)
}

func TestLoadConfig_InvalidCardStartDate(t *testing.T) {
	t.Parallel()

	const invalidFileContent = `
ineligible_patterns:
  - TEST_PATTERN_1
  - TEST_PATTERN_2
  - TEST_PATTERN_3

spend_amount_for_bonus: 1000
bonus_spend_period_in_days: 30
card_start_date: invalid
`

	filePath := createTestFile(t, []byte(invalidFileContent))

	config, err := c.LoadConfig(filePath)

	require.ErrorContains(t, err, c.ErrFailedToUnmarshalMsg)
	require.Nil(t, config)
}
