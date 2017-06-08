package errors

// DefaultErrorIfNil checks if the err is nil, if true returns the default message otherwise err.Error()
func DefaultErrorIfNil(err error, defaultMessage string) string {
	if err != nil {
		return err.Error()
	}
	return defaultMessage
}
