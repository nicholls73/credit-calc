package config

import (
	"fmt"
)

const (
	ErrFailedToReadFileMsg  = "failed to read file"
	ErrFailedToUnmarshalMsg = "failed to unmarshal"
)

func ErrFailedToReadFile(err error) error {
	return fmt.Errorf("%s: %w", ErrFailedToReadFileMsg, err)
}

func ErrFailedToUnmarshal(err error) error {
	return fmt.Errorf("%s: %w", ErrFailedToUnmarshalMsg, err)
}
