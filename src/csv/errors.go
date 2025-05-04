package csv

import (
	"fmt"
)

const (
	ErrFailedToOpenFileMsg = "failed to open file"
	ErrFileEmptyMsg        = "file is empty"
)

func ErrFailedToOpenFile(err error) error {
	return fmt.Errorf("%s: %w", ErrFailedToOpenFileMsg, err)
}

func ErrFileEmpty(err error) error {
	return fmt.Errorf("%s: %w", ErrFileEmptyMsg, err)
}
