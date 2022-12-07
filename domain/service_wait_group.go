package domain

import (
	"fmt"
	"sync"
)

var wg = &sync.WaitGroup{}

func AddTaskCount() {
	fmt.Printf("add task\n")
	wg.Add(1)
}

func DoneTask() {
	fmt.Printf("done task\n")
	wg.Done()
}

func WaitUntilAllTasksDone() {
	wg.Wait()
}

func CleanupWaitGroup() {
	wg = new(sync.WaitGroup)
}
