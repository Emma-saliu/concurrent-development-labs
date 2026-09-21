//Barrier.go Template Code
//Copyright (C) 2024 Dr. Joseph Kehoe

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

//--------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Emoshoke Saliu
// Issues: None
// A simple barrier has been implemented using a mutex and semaphore
//--------------------------------------------

package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"sync"
	"time"
)

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(
	goNum int,
	wg *sync.WaitGroup,
	theLock *sync.Mutex,
	sem *semaphore.Weighted,
	counter *int,
	totalRoutines int,
	ctx context.Context,
) bool {

	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)

	// Lock while changing the shared counter.
	theLock.Lock()
	*counter++

	// If this is the last goroutine to arrive,
	// open the barrier.
	if *counter == totalRoutines {
		sem.Release(1)
	}
	theLock.Unlock()

	// Barrier:
	// wait here until all goroutines have completed Part A.
	sem.Acquire(ctx, 1)
	sem.Release(1)

	fmt.Println("Part B", goNum)

	wg.Done()
	return true
}

func main() {
	totalRoutines := 10

	var wg sync.WaitGroup
	wg.Add(totalRoutines)

	ctx := context.TODO()

	// Mutex protects the shared counter.
	var theLock sync.Mutex

	// Counts how many goroutines have reached the barrier.
	counter := 0

	// Semaphore acts like a gate.
	sem := semaphore.NewWeighted(1)

	// Close the gate before starting the goroutines.
	sem.Acquire(ctx, 1)

	for i := range totalRoutines {
		go doStuff(
			i,
			&wg,
			&theLock,
			sem,
			&counter,
			totalRoutines,
			ctx,
		)
	}

	wg.Wait()
}
