package domain

import "sync"

var wg = &sync.WaitGroup{}

func AddTaskCount() {
	wg.Add(1)
}

func DoneTask() {
	wg.Done()
}

func WaitUntilAllTasksDone() {
	wg.Wait()
}
