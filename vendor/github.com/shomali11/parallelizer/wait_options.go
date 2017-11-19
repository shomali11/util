package parallelizer

import (
	"context"
)

// WaitOptions configuration for the Wait
type WaitOptions struct {
	Context context.Context
}

// WaitOption an option for Waiting
type WaitOption func(*WaitOptions)

// WithContext sets context
func WithContext(ctx context.Context) WaitOption {
	return func(options *WaitOptions) {
		options.Context = ctx
	}
}

func newWaitOptions(options ...WaitOption) *WaitOptions {
	config := &WaitOptions{
		Context: context.Background(),
	}

	for _, option := range options {
		option(config)
	}
	return config
}
