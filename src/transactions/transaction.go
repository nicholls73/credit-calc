package transactions

import (
	c "credit-calc/config"
	"fmt"
	"strconv"
	"strings"
	"time"

	v "github.com/go-ozzo/ozzo-validation/v4"
)

type Transaction struct {
	Date        time.Time
	Amount      float64
	Description string
}

func (t Transaction) Equals(other *Transaction) bool {
	return t.Date == other.Date && t.Amount == other.Amount && t.Description == other.Description
}

func (t Transaction) validate() error {
	return v.ValidateStruct(&t,
		v.Field(&t.Date, v.Required),
		v.Field(&t.Amount, v.Required),
		v.Field(&t.Description, v.Required),
	)
}

func FromCSVRow(row []string) (*Transaction, error) {
	if len(row) != 3 {
		return nil, ErrMissingInputs(fmt.Errorf("expected 3 fields, got %d", len(row)))
	}

	date, err := time.Parse("02/01/2006", row[0])
	if err != nil {
		return nil, ErrFailedToParseDate(err)
	}
	date = date.In(time.Local)

	amount, err := strconv.ParseFloat(row[1], 64)
	if err != nil {
		return nil, ErrFailedToParseAmount(err)
	}

	transaction := Transaction{
		Date:        date,
		Amount:      amount,
		Description: row[2],
	}

	if err := transaction.validate(); err != nil {
		return nil, ErrFailedValidation(err)
	}

	return &transaction, nil
}

func (t Transaction) IsEligible(config *c.Config) bool {
	for _, pattern := range config.IneligiblePatterns {
		if strings.Contains(strings.ToUpper(t.Description), pattern) {
			return false
		}
	}

	return true
}
