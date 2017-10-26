package xerrors

import "errors"

// New returns an error with the given text.
func New(text string) error {
	return errors.New(text)
}

// DefaultErrorIfNil checks if the err is nil, if true returns the default message otherwise err.Error()
func DefaultErrorIfNil(err error, defaultMessage string) string {
	if err != nil {
		return err.Error()
	}
	return defaultMessage
}
