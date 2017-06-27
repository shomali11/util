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
	options := &parallelizer.Options{Timeout: timeout}
	group := parallelizer.NewGroup(options)
	for _, function := range functions {
		group.Add(function)
	}
	group.Run()
}
