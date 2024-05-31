package concurrency

import (
	"sync"
	"testing"
)

func TestExecuteInParallel(t *testing.T) {
	var count int
	increment := func() {
		count++
	}

	ExecuteInParallel(increment, increment, increment)

	if count != 3 {
		t.Errorf("Expected count to be 3, but got %d", count)
	}
}

func TestSemaphore(t *testing.T) {
	sem := NewSemaphore(2)
	var wg sync.WaitGroup
	var count int

	increment := func() {
		sem.Acquire()
		defer sem.Release()
		defer wg.Done()
		count++
	}

	wg.Add(3)
	go increment()
	go increment()
	go increment()
	wg.Wait()

	if count != 3 {
		t.Errorf("Expected count to be 3, but got %d", count)
	}
}
