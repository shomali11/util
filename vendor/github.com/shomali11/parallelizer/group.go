package parallelizer

import (
	"errors"
	"sync"
)

const (
	nilFunctionError = "nil function"
)

// NewGroup create a new group of workers
func NewGroup(options ...GroupOption) *Group {
	groupOptions := newGroupOptions(options...)

	group := &Group{
		jobsChannel: make(chan func(), groupOptions.JobQueueSize),
		waitGroup:   &sync.WaitGroup{},
	}

	for i := 1; i <= groupOptions.PoolSize; i++ {
		go group.worker()
	}
	return group
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

	channel := make(chan bool)
	go func() {
		g.waitGroup.Wait()
		close(channel)
	}()

	select {
	case <-waitOptions.Context.Done():
		return waitOptions.Context.Err()
	case <-channel:
		return nil
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
