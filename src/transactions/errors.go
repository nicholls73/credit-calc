package transactions

import (
	"fmt"
)

const (
	ErrMissingInputsMsg       = "missing inputs"
	ErrFailedValidationMsg    = "failed validation"
	ErrFailedToParseDateMsg   = "failed to parse date"
	ErrFailedToParseAmountMsg = "failed to parse amount"
)

func ErrMissingInputs(err error) error {
	return fmt.Errorf("%s: %w", ErrMissingInputsMsg, err)
}

func ErrFailedValidation(err error) error {
	return fmt.Errorf("%s: %w", ErrFailedValidationMsg, err)
}

func ErrFailedToParseDate(err error) error {
	return fmt.Errorf("%s: %w", ErrFailedToParseDateMsg, err)
}

func ErrFailedToParseAmount(err error) error {
	return fmt.Errorf("%s: %w", ErrFailedToParseAmountMsg, err)
}
