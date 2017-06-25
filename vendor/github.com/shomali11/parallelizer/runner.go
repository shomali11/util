package parallelizer

import (
	"sync"
	"time"
)

// Runner allows you to parallelize function calls with an optional timeout
type Runner struct {
	Timeout time.Duration
}

// Run parallelizes the function calls and returns a boolean that determines whether the functions completed or timed out
func (p *Runner) Run(functions ...func()) bool {
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(functions))

	for _, function := range functions {
		go func(copy func()) {
			defer waitGroup.Done()
			copy()
		}(function)
	}

	return p.wait(&waitGroup)
}

func (p *Runner) wait(waitGroup *sync.WaitGroup) bool {
	// If no timeout was provided
	if p.Timeout <= 0 {
		waitGroup.Wait()
		return true
	}

	channel := make(chan struct{})

	go func() {
		waitGroup.Wait()
		close(channel)
	}()

	select {
	case <-channel:
		return true
	case <-time.After(p.Timeout):
		return false
	}
}
