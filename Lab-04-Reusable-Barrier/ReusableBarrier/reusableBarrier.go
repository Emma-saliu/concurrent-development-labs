package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// barrier represents a reusable barrier.
//
// Two turnstiles are used:
// turnstile1 controls entry through the barrier.
// turnstile2 controls the reset phase.
type barrier struct {
	turnstile1 chan bool
	turnstile2 chan bool
	theLock    sync.Mutex
	total      int
	count      int
}

// createBarrier creates and initialises a reusable barrier.
// N is the number of goroutines that must reach the barrier
// before any of them can continue.
//
// A pointer is returned so the mutex is not copied.
func createBarrier(N int) *barrier {
	theBarrier := &barrier{
		turnstile1: make(chan bool, 1),
		turnstile2: make(chan bool, 1),
		total:      N,
		count:      0,
	}

	// Turnstile 1 starts closed.
	// Turnstile 2 starts open.
	theBarrier.turnstile2 <- true

	return theBarrier
}

// wait blocks each goroutine until all goroutines have reached
// the barrier.
//
// After all goroutines pass through, the barrier resets so
// that the same barrier can be used again.
func (b *barrier) wait() {

	// -------------------------
	// Phase 1 - Arrival
	// -------------------------

	b.theLock.Lock()

	// Record that another goroutine has reached the barrier.
	b.count++

	// The final goroutine to arrive closes turnstile 2
	// and opens turnstile 1.
	if b.count == b.total {
		<-b.turnstile2
		b.turnstile1 <- true
	}

	b.theLock.Unlock()

	// Pass through turnstile 1.
	//
	// Each goroutine takes the token and immediately returns it
	// so the next goroutine can also pass through.
	<-b.turnstile1
	b.turnstile1 <- true

	// -------------------------
	// Phase 2 - Departure
	// -------------------------

	b.theLock.Lock()

	// Record that one goroutine has left the barrier.
	b.count--

	// The final goroutine to leave closes turnstile 1
	// and reopens turnstile 2.
	//
	// This resets the barrier for the next use.
	if b.count == 0 {
		<-b.turnstile1
		b.turnstile2 <- true
	}

	b.theLock.Unlock()

	// Pass through turnstile 2.
	<-b.turnstile2
	b.turnstile2 <- true
}

// WorkWithRendezvous demonstrates that the same barrier
// can be used more than once.
func WorkWithRendezvous(
	wg *sync.WaitGroup,
	num int,
	theBarrier *barrier,
) bool {

	// Ensure the WaitGroup is updated when this goroutine finishes.
	defer wg.Done()

	// Random delay so that goroutines reach the barrier
	// at different times.
	delay := time.Duration(rand.IntN(5))
	time.Sleep(delay * time.Second)

	fmt.Println("Part A", num)

	// First use of the reusable barrier.
	theBarrier.wait()

	fmt.Println("Part B", num)

	// Add another random delay before using
	// the same barrier again.
	delay = time.Duration(rand.IntN(5))
	time.Sleep(delay * time.Second)

	// Second use of the SAME barrier.
	theBarrier.wait()

	fmt.Println("Part C", num)

	return true
}

func main() {

	threadCount := 5

	var wg sync.WaitGroup

	// Create one reusable barrier shared by all 5 goroutines.
	theBarrier := createBarrier(threadCount)

	wg.Add(threadCount)

	// Create the goroutines.
	for n := range threadCount {
		go WorkWithRendezvous(
			&wg,
			n,
			theBarrier,
		)
	}

	// Wait until all goroutines have finished.
	wg.Wait()
}