package parallelizer

import "time"

type Options struct {
	Timeout        time.Duration
	WorkerPoolSize int
}
