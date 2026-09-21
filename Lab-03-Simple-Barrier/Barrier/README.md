# Simple Barrier

## Description

This program demonstrates a simple barrier implemented in Go using goroutines, a mutex and a semaphore.

Ten goroutines are created. Each goroutine completes Part A before reaching the barrier. The barrier prevents any goroutine from continuing to Part B until all ten goroutines have completed Part A.

A shared counter keeps track of how many goroutines have reached the barrier. A mutex protects this counter so that only one goroutine can modify it at a time. When the final goroutine reaches the barrier, the semaphore is released and the waiting goroutines are allowed to continue to Part B.

## Running the Program

From the root directory of the project, run:

```bash
go run ./Lab-03-Simple-Barrier/Barrier
```

The program should display all Part A messages before displaying any Part B messages.

Example:

```text
Part A 4
Part A 1
Part A 8
...
Part A 0
Part B 4
Part B 8
...
```

The order of the goroutines may change each time because they execute concurrently.

## Synchronisation Used

* **Goroutines** are used to execute tasks concurrently.
* **Mutex** protects the shared counter from concurrent access.
* **Semaphore** acts as the barrier and prevents goroutines from continuing until everyone has arrived.
* **WaitGroup** prevents the main function from terminating before all goroutines have finished.

## Author

Modified by Emoshoke Saliu.

Original template created by Dr. Joseph Kehoe.

## Licence

The original template and this modification are distributed under the GNU General Public License version 3 or later.
