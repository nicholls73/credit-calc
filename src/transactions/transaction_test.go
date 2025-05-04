package transactions_test

import (
	"credit-calc/transactions"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var validRow = []string{"20/03/2025", "500.00", "VENDOR ONE"}

var validTransaction = &transactions.Transaction{
	Date:   time.Date(2025, 3, 20, 0, 0, 0, 0, time.Local),
	Amount: 500.00,
	Vendor: "VENDOR ONE",
}

func TestTransaction_Equals(t *testing.T) {
	t.Parallel()

	equalTransaction := &transactions.Transaction{
		Date:   time.Date(2025, 3, 20, 0, 0, 0, 0, time.Local),
		Amount: 500.00,
		Vendor: "VENDOR ONE",
	}

	inequalTransaction := &transactions.Transaction{
		Date:   time.Date(2025, 3, 21, 0, 0, 0, 0, time.Local),
		Amount: 600.00,
		Vendor: "VENDOR TWO",
	}

	assert.True(t, validTransaction.Equals(equalTransaction))
	assert.False(t, validTransaction.Equals(inequalTransaction))
}

func TestFromCSVRow_ValidInputs(t *testing.T) {
	t.Parallel()

	transaction, err := transactions.FromCSVRow(validRow)

	assert.NoError(t, err)
	assert.True(t, validTransaction.Equals(transaction))
}

func TestFromCSVRow_MissingInputs(t *testing.T) {
	t.Parallel()

	invalidRow := []string{"20/03/2025", "500.00"}

	transaction, err := transactions.FromCSVRow(invalidRow)
	assert.Nil(t, transaction)
	assert.ErrorContains(t, err, transactions.ErrMissingInputsMsg)
}

func TestFromCSVRow_InvalidDate(t *testing.T) {
	t.Parallel()

	invalidRow := []string{"invalid-date", "500.00", "VENDOR ONE"}

	transaction, err := transactions.FromCSVRow(invalidRow)
	assert.Nil(t, transaction)
	assert.ErrorContains(t, err, transactions.ErrFailedToParseDateMsg)
}

func TestFromCSVRow_InvalidAmount(t *testing.T) {
	t.Parallel()

	invalidRow := []string{"20/03/2025", "invalid-amount", "VENDOR ONE"}

	transaction, err := transactions.FromCSVRow(invalidRow)
	assert.Nil(t, transaction)
	assert.ErrorContains(t, err, transactions.ErrFailedToParseAmountMsg)
}
