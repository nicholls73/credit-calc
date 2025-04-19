package errors

import (
	"errors"
)

var (
	ErrFileNotFound = errors.New("file not found")
	ErrFileEmpty = errors.New("file is empty")
	ErrInvalidRow = errors.New("invalid row format, expected: DATE,AMOUNT,VENDOR")
	ErrInvalidDate = errors.New("invalid date format, expected: DD/MM/YYYY")
	ErrInvalidAmount = errors.New("invalid amount format, expected: \"1234.56\" or \"-1234.56\"")
)