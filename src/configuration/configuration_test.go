package configuration_test

import (
	"credit-calc/configuration"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testFileContent = `ineligible_patterns:
  - TEST_PATTERN_1
  - TEST_PATTERN_2
  - TEST_PATTERN_3
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

	equalConfig := &configuration.Config{
		IneligiblePatterns: []string{"TEST_PATTERN_1", "TEST_PATTERN_2", "TEST_PATTERN_3"},
	}

	inequalConfig := &configuration.Config{
		IneligiblePatterns: []string{"TEST_PATTERN_4", "TEST_PATTERN_5", "TEST_PATTERN_6"},
	}

	assert.True(t, equalConfig.Equals(equalConfig))
	assert.False(t, equalConfig.Equals(inequalConfig))
}

func TestLoadConfig_ValidConfig(t *testing.T) {
	t.Parallel()

	filePath := createTestFile(t, []byte(testFileContent))

	expectedConfig := &configuration.Config{
		IneligiblePatterns: []string{"TEST_PATTERN_1", "TEST_PATTERN_2", "TEST_PATTERN_3"},
	}

	config, err := configuration.LoadConfig(filePath)

	require.NoError(t, err)
	require.NotNil(t, config)
	require.True(t, expectedConfig.Equals(config))
}

func TestLoadConfig_InvalidFile(t *testing.T) {
	t.Parallel()

	filename := "invalid-file.yaml"

	config, err := configuration.LoadConfig(filename)

	require.ErrorContains(t, err, configuration.ErrFailedToReadFileMsg)
	require.Nil(t, config)
}
