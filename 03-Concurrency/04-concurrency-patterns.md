# Concurrency: Concurrency Patterns in Go

This section covers some common and powerful concurrency patterns used in Go. Understanding these patterns is key to writing robust and scalable concurrent programs.

## 1. Fan-out, Fan-in

This pattern involves starting multiple goroutines to do some work (fan-out) and then collecting their results into a single channel (fan-in). It's a great way to parallelize a task.

**Example: Parallel computation**
```go
package main

import (
    "fmt"
    "sync"
)

// A function that does some work
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        // ... do some work ...
        results <- j * 2
    }
}

func main() {
    numJobs := 10
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    // Fan-out: Start 3 workers
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // Send jobs
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)

    // Fan-in: Collect results
    for a := 1; a <= numJobs; a++ {
        <-results
    }
}
```

## 2. Worker Pools

A worker pool is a collection of goroutines that are available to process tasks from a queue. This is a very common pattern for controlling the number of concurrent operations, preventing resource exhaustion.

The "Fan-out, Fan-in" example above is a simple implementation of a worker pool.

## 3. Rate Limiting

Rate limiting is an important mechanism for controlling resource utilization and maintaining service quality. Go's channels and `time.Ticker` make this easy to implement.

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    requests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)

    // A ticker that fires every 200 milliseconds
    limiter := time.NewTicker(200 * time.Millisecond)

    for req := range requests {
        <-limiter.C // Block until the ticker fires
        fmt.Println("request", req, time.Now())
    }
}
```
This ensures that we only process one request every 200ms.

## 4. Bounded Parallelism

This pattern is used when you want to run a number of tasks in parallel, but limit how many are running at any given time. This is useful for tasks like web scraping, where you don't want to overwhelm a server with too many concurrent requests.

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    var wg sync.WaitGroup
    concurrency := 3
    sem := make(chan struct{}, concurrency) // A semaphore

    tasks := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    for _, task := range tasks {
        wg.Add(1)
        go func(taskID int) {
            defer wg.Done()
            sem <- struct{}{} // Acquire a slot
            defer func() { <-sem }() // Release the slot

            // Do the work
            fmt.Printf("Processing task %d\n", taskID)
            time.Sleep(1 * time.Second)
        }(task)
    }

    wg.Wait()
}
```
A buffered channel is used as a semaphore to limit the number of active goroutines to the channel's capacity.

## 5. Pipelines

A pipeline is a series of "stages" connected by channels, where each stage is a group of goroutines running the same function. In each stage, the goroutines:
- Receive values from an upstream channel.
- Perform some function on that data, producing new values.
- Send the new values to a downstream channel.

```go
package main

import "fmt"

// Stage 1: Generates numbers
func gen(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

// Stage 2: Squares the numbers
func sq(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

func main() {
    // Set up the pipeline.
    c := gen(2, 3)
    out := sq(c)

    // Consume the output.
    fmt.Println(<-out) // 4
    fmt.Println(<-out) // 9
}
```
Pipelines are a powerful way to structure concurrent code and are often very efficient. 