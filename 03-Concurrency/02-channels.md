# Concurrency: Channels

Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another.

## Definition

**Syntax:** `chan T` is a channel of type `T`.

**Key Characteristics:**
- **Typed:** A channel can only transport data of its specified type.
- **Communication and Synchronization:** Channels are the primary way to communicate between goroutines. The send and receive operations *block* until the other side is ready, which provides a powerful synchronization mechanism. "Do not communicate by sharing memory; instead, share memory by communicating."

## Creating and Using Channels

Use `make(chan T)` to create a channel.

```go
package main

import "fmt"

func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum // send sum to c
}

func main() {
    s := []int{7, 2, 8, -9, 4, 0}
    c := make(chan int)

    go sum(s[:len(s)/2], c)
    go sum(s[len(s)/2:], c)
    x, y := <-c, <-c // receive from c

    fmt.Println(x, y, x+y)
}
```
Sends (`ch <- v`) and receives (`v := <-ch`) are blocking. This guarantees that in the example above, the `main` goroutine waits for both `sum` goroutines to finish their work before printing the final result.

## Buffered Channels

You can create a channel with a buffer using `make(chan T, capacity)`.
- Sends to a buffered channel block only when the buffer is full.
- Receives block only when the buffer is empty.

```go
ch := make(chan int, 2)
ch <- 1
ch <- 2
// ch <- 3 // This would block until a value is received.
```
**When to use:** Buffered channels can be useful for managing throughput, such as a pool of workers processing jobs. They can reduce the amount of time goroutines spend waiting for each other. However, unbuffered channels are often preferred as they force you to think about synchronization points more clearly.

## `range` and `close`

A sender can `close` a channel to indicate that no more values will be sent. A receiver can test whether a channel has been closed by using the two-value form of the receive operation.

```go
v, ok := <-ch
```
`ok` is `false` if there are no more values to receive and the channel is closed.

You can also use a `for-range` loop to receive values from a channel until it is closed.

```go
package main

import (
    "fmt"
)

func fibonacci(n int, c chan int) {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        c <- x
        x, y = y, x+y
    }
    close(c) // Close the channel when done
}

func main() {
    c := make(chan int, 10)
    go fibonacci(cap(c), c)
    for i := range c { // Loop continues until channel is closed
        fmt.Println(i)
    }
}
```
**Important:** Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.

## `select` Statement

The `select` statement lets a goroutine wait on multiple communication operations. A `select` blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()

    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}
```

A `default` case can be added to a `select` to make it non-blocking. It will run if no other case is ready. 