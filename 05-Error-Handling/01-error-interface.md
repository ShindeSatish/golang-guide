# Error Handling: The `error` Interface

In Go, errors are values. Go's approach to error handling is explicit and checkable, centered around the built-in `error` interface.

## The `error` Type

The `error` type is a simple interface.
```go
type error interface {
    Error() string
}
```
Any type that implements an `Error() string` method fulfills the `error` interface.

## The Idiomatic Go Way

Functions that can fail should return an `error` as their last return value.
```go
func DoSomething() (ResultType, error) {
    // ...
}
```
The caller is then responsible for checking the error.
```go
result, err := DoSomething()
if err != nil {
    // Handle the error
    log.Fatalf("DoSomething failed: %v", err)
}
// Continue with the successful result
```
This pattern makes the control flow for errors explicit and easy to follow. You must check for errors where they happen.

## Creating Custom Errors

You can create your own error types. This is useful for providing more context about what went wrong.

### Using `errors.New`
The simplest way to create an error is with the `errors.New` function.
```go
import "errors"

func Sqrt(x float64) (float64, error) {
    if x < 0 {
        return 0, errors.New("math: square root of negative number")
    }
    // ...
}
```

### Using Custom Structs
For more complex errors that need to carry more information, you can use a custom struct that implements the `error` interface.

```go
package main

import (
    "fmt"
    "time"
)

type MyError struct {
    When time.Time
    What string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("at %v, %s",
        e.When, e.What)
}

func run() error {
    return &MyError{
        time.Now(),
        "it didn't work",
    }
}

func main() {
    if err := run(); err != nil {
        fmt.Println(err)
    }
}
```
This allows callers to use a type assertion or type switch to inspect the details of the error and make more intelligent decisions.

## Why Not Exceptions?

Go deliberately omits exceptions (like `try/catch/finally` in other languages). The Go authors believed that coupling error handling to control flow with constructs like `try` makes code convoluted and encourages programmers to ignore errors that they should be handling.

By making errors explicit return values, Go forces you to confront errors at the point they occur, leading to more robust and reliable code. Errors are not exceptional; they are an expected part of a program's operation. 