# Error Handling: Panic and Recover

While Go's main error handling mechanism is based on `error` values, it also has a mechanism for handling truly exceptional conditions: `panic` and `recover`.

## `panic`

`panic` is a built-in function that stops the ordinary flow of control and begins *panicking*. When a function `F` calls `panic`, execution of `F` stops, any deferred functions in `F` are executed, and then `F` returns to its caller. To the caller, `F` then behaves like a call to `panic`. The process continues up the stack until all functions in the current goroutine have returned, at which point the program crashes.

**When to `panic`:**
You should only use `panic` for truly exceptional, unrecoverable errors. A classic example is an index out of bounds on a slice, which indicates a serious programmer error.

**General Rule:** Do not use `panic` for ordinary errors like file-not-found or network connection issues. Use `error` values for those. `panic` is for programming errors that should not happen.

```go
package main

import "fmt"

func main() {
    // This will cause the program to panic and crash
    s := []int{1, 2, 3}
    fmt.Println(s[10])
}
```

## `recover`

`recover` is a built-in function that regains control of a panicking goroutine. `recover` is only useful inside deferred functions. During normal execution, a call to `recover` will return `nil` and have no other effect. If the current goroutine is panicking, a call to `recover` will capture the value given to `panic` and resume normal execution.

```go
package main

import "fmt"

func mayPanic() {
    panic("a problem")
}

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered. Error:\n", r)
        }
    }()

    mayPanic()

    fmt.Println("After mayPanic()") // This line will not be executed
}
// Output:
// Recovered. Error:
//  a problem
```

## Use Case: Robust Servers

A common use for `panic` and `recover` is in a web server. If one request handler panics due to a bug, you don't want the entire server to crash. You can wrap each handler in a deferred function that calls `recover`, logs the error, and cleanly closes the connection for that specific request, allowing the server to continue handling other requests.

```go
func server(workChan <-chan *Work) {
    for work := range workChan {
        go safelyDo(work)
    }
}

func safelyDo(work *Work) {
    defer func() {
        if err := recover(); err != nil {
            log.Printf("work failed: %s", err)
        }
    }()
    do(work)
}
```
This pattern makes the server more robust to unexpected issues in its handler code. It turns a catastrophic failure in one part of the program into a logged error, without taking down the entire system. 