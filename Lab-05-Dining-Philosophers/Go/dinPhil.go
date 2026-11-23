// Dining Philosophers
// Original template by Joseph Kehoe
// Modified by: Emoshoke Saliu
//
// This program demonstrates the Dining Philosophers concurrency problem.
// A deadlock is prevented by changing the order in which one philosopher
// picks up their forks.

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

const philosopherCount = 5

// think simulates a philosopher thinking for a random amount of time.
func think(index int) {
	delay := time.Duration(rand.IntN(5))
	time.Sleep(delay * time.Second)

	fmt.Println("Phil:", index, "was thinking")
}

// eat simulates a philosopher eating for a random amount of time.
func eat(index int) {
	delay := time.Duration(rand.IntN(5))
	time.Sleep(delay * time.Second)

	fmt.Println("Phil:", index, "was eating")
}

// getForks attempts to acquire the two forks required by a philosopher.
//
// Normally a philosopher takes their own fork first and then the fork
// to their right.
//
// The final philosopher takes the forks in the opposite order.
// This breaks the circular wait condition and prevents deadlock.
func getForks(index int, forks map[int]chan bool) {

	leftFork := index
	rightFork := (index + 1) % philosopherCount

	if index == philosopherCount-1 {

		// Final philosopher takes the right fork first.
		forks[rightFork] <- true
		forks[leftFork] <- true

	} else {

		// Other philosophers take the left fork first.
		forks[leftFork] <- true
		forks[rightFork] <- true
	}
}

// putForks releases both forks after the philosopher has eaten.
func putForks(index int, forks map[int]chan bool) {

	leftFork := index
	rightFork := (index + 1) % philosopherCount

	<-forks[leftFork]
	<-forks[rightFork]
}

// doPhilStuff represents the behaviour of one philosopher.
func doPhilStuff(
	index int,
	wg *sync.WaitGroup,
	forks map[int]chan bool,
) {

	// Signal that this philosopher has finished when the function exits.
	defer wg.Done()

	// Repeat the think/eat cycle a few times.
	// Using a finite loop allows the program to eventually finish.
	for round := 0; round < 3; round++ {

		think(index)

		getForks(index, forks)

		eat(index)

		putForks(index, forks)
	}

	fmt.Println("Phil:", index, "has finished")
}

func main() {

	var wg sync.WaitGroup

	philCount := philosopherCount

	wg.Add(philCount)

	// Each buffered channel represents one fork.
	// A capacity of 1 means only one philosopher can hold a fork at a time.
	forks := make(map[int]chan bool)

	for k := range philCount {
		forks[k] = make(chan bool, 1)
	}

	// Start the philosophers.
	for n := range philCount {
		go doPhilStuff(n, &wg, forks)
	}

	// Wait until every philosopher has finished.
	wg.Wait()

	fmt.Println("All philosophers have finished")
}