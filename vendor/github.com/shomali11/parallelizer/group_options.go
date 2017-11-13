package parallelizer

const (
	defaultPoolSize     = 10
	defaultJobQueueSize = 100
)

// GroupOption an option for Groups
type GroupOption func(*GroupOptions)

// WithPoolSize sets pool size
func WithPoolSize(size int) GroupOption {
	return func(options *GroupOptions) {
		options.PoolSize = size
	}
}

// WithJobQueueSize sets job size
func WithJobQueueSize(size int) GroupOption {
	return func(options *GroupOptions) {
		options.JobQueueSize = size
	}
}

// GroupOptions configuration for the Group
type GroupOptions struct {
	PoolSize     int
	JobQueueSize int
}

func newGroupOptions(options ...GroupOption) *GroupOptions {
	config := &GroupOptions{
		PoolSize:     defaultPoolSize,
		JobQueueSize: defaultJobQueueSize,
	}

	for _, option := range options {
		option(config)
	}
	return config
}
