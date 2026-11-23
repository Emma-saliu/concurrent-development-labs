# Dining Philosophers

## Description

This lab implements the Dining Philosophers concurrency problem using Go.

Five philosophers share five forks. Each philosopher alternates between thinking and eating. To eat, a philosopher must acquire both the fork on their left and the fork on their right.

The original implementation could deadlock if every philosopher acquired one fork and then waited indefinitely for the second fork.

## Deadlock Prevention

The program prevents deadlock by changing the order in which the final philosopher acquires their forks.

Philosophers 0 to 3 acquire their left fork first and then their right fork.

Philosopher 4 acquires the forks in the opposite order.

This removes the circular wait condition required for deadlock.

## Concurrency

The program uses:

- Goroutines to represent philosophers
- Buffered channels to represent forks
- A WaitGroup to wait for all philosophers to finish
- Random delays to simulate thinking and eating

Each philosopher performs three think/eat cycles before finishing.

## Running the Program

From the root of the repository, run:

```bash
go run ./Lab-05-Dining-Philosophers/Go
```

The philosophers will think and eat concurrently.

The program should eventually finish with:

```text
All philosophers have finished
```

The order of the output may change each time because the philosophers execute concurrently.

## Author

Modified by Emoshoke Saliu.

Original Dining Philosophers template created by Joseph Kehoe.