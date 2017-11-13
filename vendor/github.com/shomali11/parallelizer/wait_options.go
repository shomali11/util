package parallelizer

import "time"

const (
	defaultTimeout = 0
)

// WaitOptions configuration for the Wait
type WaitOptions struct {
	Timeout time.Duration
}

// WaitOption an option for Waiting
type WaitOption func(*WaitOptions)

// WithTimeout sets timeout
func WithTimeout(timeout time.Duration) WaitOption {
	return func(options *WaitOptions) {
		options.Timeout = timeout
	}
}

func newWaitOptions(options ...WaitOption) *WaitOptions {
	config := &WaitOptions{
		Timeout: defaultTimeout,
	}

	for _, option := range options {
		option(config)
	}
	return config
}
