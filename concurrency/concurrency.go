package concurrency

import (
	"sync"
)

func ExecuteInParallel(funcs ...func()) {
	var wg sync.WaitGroup
	wg.Add(len(funcs))

	for _, fn := range funcs {
		go func(f func()) {
			defer wg.Done()
			f()
		}(fn)
	}

	wg.Wait()
}

// Semaphore provides a way to limit the number of goroutines running concurrently
type Semaphore struct {
	ch chan struct{}
}

// NewSemaphore creates a new semaphore with the given limit
func NewSemaphore(limit int) *Semaphore {
	return &Semaphore{ch: make(chan struct{}, limit)}
}

// Acquire acquires the semaphore, blocking if necessary
func (s *Semaphore) Acquire() {
	s.ch <- struct{}{}
}

// Release releases the semaphore
func (s *Semaphore) Release() {
	<-s.ch
}
