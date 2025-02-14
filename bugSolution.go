```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var count int
        var mu sync.Mutex // Add a mutex
        for i := 0; i < 1000; i++ {
                wg.Add(1)
                go func() {
                        defer wg.Done()
                        mu.Lock() // Acquire the mutex
                        count++
                        mu.Unlock() // Release the mutex
                }()
        }
        wg.Wait()
        fmt.Println(count)
}
```