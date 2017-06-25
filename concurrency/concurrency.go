package concurrency

import (
	"github.com/shomali11/parallelizer"
	"time"
)

// Parallelize parallelizes the function calls
func Parallelize(functions ...func()) {
	ParallelizeTimeout(0, functions...)
}

// ParallelizeTimeout parallelizes the function calls with a timeout
func ParallelizeTimeout(timeout time.Duration, functions ...func()) {
	runner := parallelizer.Runner{Timeout: timeout}
	runner.Run(functions...)
}
