package summary

import "fmt"

const (
	ErrEmptyConfigMsg = "config is empty"
)

func ErrEmptyConfig(err error) error {
	return fmt.Errorf("%s: %w", ErrEmptyConfigMsg, err)
}
