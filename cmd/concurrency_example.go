package main

import (
	"fmt"
	"goLibrary/concurrency"
	"sync"
)

func runConcurrencyExamples() {
	// Example of ExecuteInParallel
	fmt.Println("Running ExecuteInParallel example")

	task1 := func() {
		fmt.Println("Task 1 executed")
	}
	task2 := func() {
		fmt.Println("Task 2 executed")
	}
	task3 := func() {
		fmt.Println("Task 3 executed")
	}

	concurrency.ExecuteInParallel(task1, task2, task3)

	// Example of Semaphore
	fmt.Println("Running Semaphore example")

	sem := concurrency.NewSemaphore(2)
	var wg sync.WaitGroup

	limitedTask := func(taskID int) {
		sem.Acquire()
		defer sem.Release()
		defer wg.Done()
		fmt.Printf("Task %d executed\n", taskID)
	}

	wg.Add(3)
	go limitedTask(1)
	go limitedTask(2)
	go limitedTask(3)
	wg.Wait()
}
