package csv

import (
	"fmt"
)

const (
	ErrFailedToOpenFileMsg    = "failed to open file"
	ErrFileEmptyMsg           = "file is empty"
	ErrFailedToParseDateMsg   = "failed to parse date"
	ErrFailedToParseAmountMsg = "failed to parse amount"
)

func ErrFailedToOpenFile(err error) error {
	return fmt.Errorf("%s: %w", ErrFailedToOpenFileMsg, err)
}

func ErrFileEmpty(err error) error {
	return fmt.Errorf("%s: %w", ErrFileEmptyMsg, err)
}

func ErrFailedToParseDate(err error) error {
	return fmt.Errorf("%s: %w", ErrFailedToParseDateMsg, err)
}

func ErrFailedToParseAmount(err error) error {
	return fmt.Errorf("%s: %w", ErrFailedToParseAmountMsg, err)
}
