package xconcurrency

import (
	"github.com/shomali11/parallelizer"
	"time"
)

// Parallelize parallelizes the function calls
func Parallelize(functions ...func()) error {
	return ParallelizeTimeout(0, functions...)
}

// ParallelizeTimeout parallelizes the function calls with a timeout
func ParallelizeTimeout(timeout time.Duration, functions ...func()) error {
	group := parallelizer.NewGroup()
	for _, function := range functions {
		group.Add(function)
	}
	return group.Wait(parallelizer.WithTimeout(timeout))
}
