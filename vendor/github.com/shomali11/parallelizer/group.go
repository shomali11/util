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

// NewGroup create a new group of workers
func NewGroup(options ...GroupOption) *Group {
	groupOptions := newGroupOptions(options...)

	pool := &Group{
		jobsChannel: make(chan func(), groupOptions.JobQueueSize),
		waitGroup:   &sync.WaitGroup{},
	}

	for i := 1; i <= groupOptions.PoolSize; i++ {
		go pool.worker()
	}
	return pool
}

// Group a group of workers executing functions concurrently
type Group struct {
	jobsChannel chan func()
	waitGroup   *sync.WaitGroup
}

// Add adds function to queue of jobs to execute
func (g *Group) Add(function func()) error {
	if function == nil {
		return errors.New(nilFunctionError)
	}

	g.jobsChannel <- function
	g.waitGroup.Add(1)
	return nil
}

// Wait waits until workers finished the jobs in the queue
func (g *Group) Wait(options ...WaitOption) error {
	waitOptions := newWaitOptions(options...)

	// If no timeout was provided
	if waitOptions.Timeout <= 0 {
		g.waitGroup.Wait()
		return nil
	}

	channel := make(chan bool)
	go func() {
		g.waitGroup.Wait()
		close(channel)
	}()

	select {
	case <-channel:
		return nil
	case <-time.After(waitOptions.Timeout):
		return errors.New(timeoutError)
	}
}

// Close closes resources
func (g *Group) Close() {
	close(g.jobsChannel)
}

func (g *Group) worker() {
	for job := range g.jobsChannel {
		job()
		g.waitGroup.Done()
	}
}
