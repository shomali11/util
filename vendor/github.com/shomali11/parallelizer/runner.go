package parallelizer

import (
	"sync"
	"time"
	"errors"
)

const (
	timeoutError = "timeout"
)

// Runner allows you to parallelize function calls with an optional timeout
type Runner struct {
	Timeout time.Duration
}

// Run parallelizes the function calls
func (p *Runner) Run(functions ...func()) error {
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

func (p *Runner) wait(waitGroup *sync.WaitGroup) error {
	// If no timeout was provided
	if p.Timeout <= 0 {
		waitGroup.Wait()
		return nil
	}

	channel := make(chan struct{})

	go func() {
		waitGroup.Wait()
		close(channel)
	}()

	select {
	case <-channel:
		return nil
	case <-time.After(p.Timeout):
		return errors.New(timeoutError)
	}
}
