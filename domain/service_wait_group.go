package domain

import (
	"sync"
)

var wg = &sync.WaitGroup{}

// AddTaskCount add one task in the general wait group service
func AddTaskCount() {
	wg.Add(1)
}

// DoneTask done task in the general wait group service
func DoneTask() {
	wg.Done()
}

// WaitUntilAllTasksDone lock current goroutine until
// all tasks done from general wait group service
func WaitUntilAllTasksDone() {
	wg.Wait()
}

// CleanupWaitGroup override wait group variable with a new wait group.
// should use for tests purposes
func CleanupWaitGroup() {
	wg = new(sync.WaitGroup)
}
