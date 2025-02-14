# Go Race Condition Example

This repository demonstrates a common race condition in Go programs that use goroutines to update shared variables.  The program attempts to increment a counter 1000 times using concurrent goroutines, but the result is often less than 1000 due to race conditions.

The `bug.go` file contains the buggy code.  The `bugSolution.go` file provides a corrected version using a mutex to protect the shared counter.

## Running the Code

1. Clone the repository:
   ```bash
git clone https://github.com/yourusername/go-race-condition.git
```
2. Navigate to the repository directory:
   ```bash
cd go-race-condition
```
3. Run the buggy code:
   ```bash
go run bug.go
```
4. Run the corrected code:
   ```bash
go run bugSolution.go
```

## Solution

The solution involves using a mutex to protect access to the shared `count` variable.  This ensures that only one goroutine can update the counter at a time, preventing race conditions.