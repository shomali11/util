package parallelizer

import (
	"errors"
	"sync"
	"time"
)

const (
	timeoutError     = "timeout"
	nilFunctionError = "nil function"
)

// DefaultGroup create a new group with default options
func DefaultGroup() *Group {
	return &Group{}
}

// NewGroup create a new group with options
func NewGroup(options *Options) *Group {
	return &Group{
		timeout:        options.Timeout,
		workerPoolSize: options.WorkerPoolSize,
	}
}

// Group allows you to parallelize function
type Group struct {
	timeout        time.Duration
	workerPoolSize int
	functions      []func()
	waitGroup      *sync.WaitGroup
}

// Add adds function to list of functions to parallelize
func (g *Group) Add(function func()) error {
	if function == nil {
		return errors.New(nilFunctionError)
	}

	g.functions = append(g.functions, function)
	return nil
}

// Run parallelizes the function calls
func (g *Group) Run() error {
	g.waitGroup = &sync.WaitGroup{}
	g.waitGroup.Add(len(g.functions))

	jobs := make(chan func(), len(g.functions))

	poolSize := g.getPoolSize()
	for i := 1; i <= poolSize; i++ {
		go g.worker(jobs)
	}

	for _, function := range g.functions {
		jobs <- function
	}
	close(jobs)

	return g.wait()
}

func (g *Group) getPoolSize() int {
	// If no worker pool size was provided
	if g.workerPoolSize <= 0 {
		return len(g.functions)
	}
	return g.workerPoolSize
}

func (g *Group) worker(jobs <-chan func()) {
	for job := range jobs {
		defer g.waitGroup.Done()
		job()
	}
}

func (g *Group) wait() error {
	// If no timeout was provided
	if g.timeout <= 0 {
		g.waitGroup.Wait()
		return nil
	}

	channel := make(chan struct{})

	go func() {
		g.waitGroup.Wait()
		close(channel)
	}()

	select {
	case <-channel:
		return nil
	case <-time.After(g.timeout):
		return errors.New(timeoutError)
	}
}
