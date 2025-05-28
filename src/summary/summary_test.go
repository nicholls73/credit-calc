package summary_test

import (
	"credit-calc/config"
	"credit-calc/summary"
	"credit-calc/transactions"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	testConfig = config.Config{
		IneligiblePatterns:     []string{"PAYMENT", "TRANSFER"},
		BonusSpendPeriodInDays: 30,
		CardStartDate:          time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
		SpendAmountForBonus:    4000,
	}

	testTransactions = []*transactions.Transaction{
		{
			Date:        time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			Amount:      -1000,
			Description: "Test Transaction 1",
		},
		{
			Date:        time.Date(2025, 1, 2, 0, 0, 0, 0, time.UTC),
			Amount:      -2000,
			Description: "Test Transaction 2",
		},
		{
			Date:        time.Date(2025, 1, 3, 0, 0, 0, 0, time.UTC),
			Amount:      -2000,
			Description: "PAYMENT",
		},
	}
)

func TestGenerateSummary_ValidInputs(t *testing.T) {
	t.Parallel()

	now := time.Date(2025, 1, 15, 0, 0, 0, 0, time.UTC)
	sum, err := summary.GenerateSummary(testTransactions, &testConfig, now)

	require.NoError(t, err)
	assert.Equal(t, sum.TotalAmountSpent, float64(3000))
	assert.Equal(t, sum.TotalPointsEarned, 3000)
	assert.Equal(t, sum.AmountLeft, float64(1000))
	assert.Equal(t, sum.DaysLeft, 16)
}

func TestGenerateSummary_NoTransactions(t *testing.T) {
	t.Parallel()

	now := time.Date(2025, 1, 15, 0, 0, 0, 0, time.UTC)
	sum, err := summary.GenerateSummary([]*transactions.Transaction{}, &testConfig, now)

	require.NoError(t, err)
	assert.Equal(t, sum.TotalAmountSpent, float64(0))
	assert.Equal(t, sum.TotalPointsEarned, 0)
	assert.Equal(t, sum.AmountLeft, float64(4000))
	assert.Equal(t, sum.DaysLeft, 16)
}

func TestGenerateSummary_NoEligibleTransactions(t *testing.T) {
	t.Parallel()

	testTransactions[0].Description = "PAYMENT TO CREDIT CARD"
	testTransactions[1].Description = "BANK TRANSFER"

	now := time.Date(2025, 1, 15, 0, 0, 0, 0, time.UTC)
	sum, err := summary.GenerateSummary(testTransactions, &testConfig, now)

	require.NoError(t, err)
	assert.Equal(t, sum.TotalAmountSpent, float64(0))
	assert.Equal(t, sum.TotalPointsEarned, 0)
	assert.Equal(t, sum.AmountLeft, float64(4000))
	assert.Equal(t, sum.DaysLeft, 16)
}

func TestGenerateSummary_MissingConfig(t *testing.T) {
	t.Parallel()

	now := time.Date(2025, 1, 15, 0, 0, 0, 0, time.UTC)
	sum, err := summary.GenerateSummary(testTransactions, nil, now)

	require.ErrorContains(t, err, summary.ErrEmptyConfigMsg)
	assert.Equal(t, sum.TotalAmountSpent, float64(0))
	assert.Equal(t, sum.TotalPointsEarned, 0)
	assert.Equal(t, sum.AmountLeft, float64(0))
	assert.Equal(t, sum.DaysLeft, 0)
}
