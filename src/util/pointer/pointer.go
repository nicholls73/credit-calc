package pointer

func String(s string) *string {
	return &s
}

func Error(err error) *string {
	errorMessage := err.Error()

	return &errorMessage
}