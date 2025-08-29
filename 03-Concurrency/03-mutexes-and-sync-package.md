# Concurrency: Mutexes and the `sync` Package

While channels are the idiomatic way to handle concurrency in Go, sometimes you need to fall back to more traditional methods, especially when managing state that is shared across multiple goroutines. This is where the `sync` package comes in.

## `sync.Mutex`

A `Mutex` (mutual exclusion lock) is used to protect shared data from being accessed by multiple goroutines at the same time.

**Use Case:**
Imagine you have a simple counter that is incremented by multiple goroutines. Without a mutex, you could have a "race condition" where two goroutines read the value at the same time, increment it, and write it back, resulting in a lost increment.

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
    mu sync.Mutex
    v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
    c.mu.Lock()
    // Lock so only one goroutine at a time can access the map c.v.
    c.v[key]++
    c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
    c.mu.Lock()
    // Lock so only one goroutine at a time can access the map c.v.
    defer c.mu.Unlock() // Use defer to ensure the mutex is always unlocked.
    return c.v[key]
}

func main() {
    c := SafeCounter{v: make(map[string]int)}
    for i := 0; i < 1000; i++ {
        go c.Inc("somekey")
    }

    time.Sleep(time.Second)
    fmt.Println(c.Value("somekey")) // Should be 1000
}
```
**Best Practice:** Using `defer c.mu.Unlock()` is a robust pattern. It guarantees that the mutex will be unlocked when the function returns, even if there's a `panic`.

## `sync.RWMutex`

A `RWMutex` (Reader/Writer Mutex) is a more specialized lock that allows for multiple "readers" to access the data simultaneously, but only one "writer".

- `Lock()` and `Unlock()` for writers.
- `RLock()` and `RUnlock()` for readers.

**When to use:** Use a `RWMutex` when you have a resource that is read far more often than it is written to. This can improve performance by allowing concurrent reads.

## `sync.Once`

`sync.Once` is an object that will perform an action exactly once.

**Use Case:**
This is perfect for initialization tasks that should only happen a single time, such as initializing a singleton object or setting up a database connection.

```go
var once sync.Once
var db *Database

func GetDB() *Database {
    once.Do(func() {
        db = setupDatabaseConnection()
    })
    return db
}
```
No matter how many goroutines call `GetDB` simultaneously, `setupDatabaseConnection` will only be called once.

## `sync.WaitGroup`

A `WaitGroup` is used to wait for a collection of goroutines to finish executing.

**How it works:**
1.  The main goroutine calls `Add(n)` to set the number of goroutines to wait for.
2.  Each of the goroutines runs and calls `Done()` when it finishes.
3.  The main goroutine calls `Wait()` to block until all goroutines have called `Done()`.

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }

    wg.Wait()
    fmt.Println("All workers finished")
}
```
`WaitGroup` is a simple and common way to coordinate the completion of multiple goroutines. 