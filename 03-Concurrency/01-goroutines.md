# Concurrency: Goroutines

This section covers goroutines, Go's lightweight approach to concurrency.

## Definition

A goroutine is a lightweight thread of execution. They are managed by the Go runtime.

**Key Characteristics:**
- **Lightweight:** Goroutines are much cheaper than traditional OS threads. They start with a small stack (a few KB) which can grow and shrink as needed. It's common to have thousands or even hundreds of thousands of goroutines in a single program.
- **Multiplexed:** The Go runtime multiplexes goroutines onto a small number of OS threads. If a goroutine blocks on I/O or a channel operation, the runtime will automatically switch to another runnable goroutine on the same OS thread, without blocking the thread itself.

## Creating a Goroutine

Starting a goroutine is simple: use the `go` keyword before a function call.

```go
package main

import (
    "fmt"
    "time"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

func main() {
    go say("world")
    say("hello")
}
// Possible output (order is not guaranteed):
// hello
// world
// world
// hello
// hello
// world
// hello
// world
// hello
// world
```
The `main` function runs in its own goroutine. The new goroutine for `say("world")` runs concurrently. The program exits when the `main` goroutine finishes, even if other goroutines are still running.

## Goroutines and Closures

It's common to start a goroutine with an anonymous function (a closure). Be mindful of how variables are captured.

**Common Pitfall:**
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    s := []string{"a", "b", "c"}
    for _, v := range s {
        go func() {
            fmt.Println(v)
        }()
    }
    time.Sleep(1 * time.Second)
}
// Incorrect Output (likely):
// c
// c
// c
```
**Why?** The goroutines share the *same* `v` variable from the loop. By the time they execute, the loop has finished and `v` holds its last value ("c").

**Correct Way:**
```go
for _, v := range s {
    v := v // Create a new variable 'v' for each iteration.
    go func() {
        fmt.Println(v)
    }()
}
// Or, pass the value as an argument:
for _, v := range s {
    go func(val string) {
        fmt.Println(val)
    }(v)
}
```
Passing the value as an argument is the cleanest way to ensure each goroutine gets the correct value for that specific iteration.

## Internals: Go Scheduler (M, P, G)

The Go scheduler is responsible for managing goroutines. It uses a model called M:P:G.
- **M (Machine):** An OS thread.
- **P (Processor):** A logical processor, representing the resources needed to execute Go code. It's a context for scheduling. There is a `P` for each `GOMAXPROCS` value.
- **G (Goroutine):** A goroutine.

**How it works:**
1. Each `P` has a local queue of runnable goroutines.
2. An `M` (OS thread) executes goroutines from its assigned `P`'s queue.
3. If a goroutine in `M` blocks (e.g., on a syscall), the `P` and its queue are handed off to another `M` so that execution can continue without blocking the thread.
4. If a `P`'s local queue runs out of goroutines, it will try to "steal" goroutines from other `P`s' queues to keep the work balanced.

This model allows Go to achieve high concurrency with a small number of OS threads, making it very efficient for I/O-bound and highly concurrent workloads. It avoids the high overhead of creating and managing OS threads directly. 