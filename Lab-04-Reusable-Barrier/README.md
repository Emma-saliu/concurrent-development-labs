# Reusable Barrier

## Description

This lab implements a reusable barrier in Go.

The barrier synchronises multiple goroutines so that no goroutine can continue until all goroutines have reached the same point.

Unlike a simple barrier, this barrier resets after all goroutines pass through it. This allows the same barrier to be used multiple times.

The program demonstrates this by using the same barrier between:

- Part A and Part B
- Part B and Part C

## Implementation

The barrier uses:

- Goroutines for concurrent execution
- A mutex to protect the shared counter
- Two channels that act as turnstiles
- A counter to record how many goroutines have reached or left the barrier
- A WaitGroup to wait for all goroutines to finish

The first turnstile opens when every goroutine has arrived.

The second phase resets the barrier after every goroutine has passed through, allowing the barrier to be used again.

## Running the Program

From the root directory of the repository, run:

```bash
go run ./Lab-04-Reusable-Barrier/ReusableBarrier
```

## Expected Behaviour

All Part A messages should appear before any Part B messages.

All Part B messages should appear before any Part C messages.

The order of the goroutine numbers may change because the goroutines execute concurrently.

Example:

```text
Part A 3
Part A 0
Part A 4
Part A 1
Part A 2

Part B 2
Part B 4
Part B 0
Part B 3
Part B 1

Part C 1
Part C 0
Part C 3
Part C 4
Part C 2
```

## Author

Modified by Emoshoke Saliu.

The implementation is based on the barrier examples supplied as part of the Concurrent Development lab material.