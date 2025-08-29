# Error Handling: Modern Error Handling

Go 1.13 introduced new functions in the standard `errors` package that provide better ways to inspect and handle errors: `errors.Is` and `errors.As`.

## Error Wrapping

Often, when an error occurs in a lower-level function, you want to add more context before returning it up the call stack. This is called *wrapping* the error.

The `fmt.Errorf` function with the `%w` verb is the standard way to wrap an error.

```go
package main

import (
    "errors"
    "fmt"
    "os"
)

func readFile() error {
    err := os.ErrNotExist // Simulate a file not found error
    if err != nil {
        // Wrap the original error with more context
        return fmt.Errorf("failed to read config file: %w", err)
    }
    return nil
}

func main() {
    err := readFile()
    if err != nil {
        fmt.Println(err)
    }
}
```
The `%w` verb ensures that the original error is "embedded" within the new error, allowing you to inspect the error chain.

## `errors.Is`

The `errors.Is` function compares an error to a target value. It will "unwrap" the error, if it has been wrapped, to see if any error in the chain matches the target.

**Use Case:** To check for a specific, known sentinel error value (like `io.EOF` or `os.ErrNotExist`).

```go
package main

import (
    "errors"
    "fmt"
    "os"
)

func main() {
    err := fmt.Errorf("access denied: %w", os.ErrPermission)

    if errors.Is(err, os.ErrPermission) {
        fmt.Println("Permission denied error.")
    } else {
        fmt.Println("Some other error.")
    }
}
```
This is the modern, preferred way over a simple `err == os.ErrPermission` check, because it correctly handles wrapped errors.

## `errors.As`

The `errors.As` function checks whether any error in the chain matches a specific type. If it finds a match, it sets a target variable to that error value and returns `true`.

**Use Case:** To access the methods or fields of a specific custom error type.

```go
package main

import (
    "errors"
    "fmt"
    "time"
)

// A custom error type
type MyError struct {
    When time.Time
    What string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
    return &MyError{time.Now(), "it didn't work"}
}

func main() {
    err := run()
    var myErr *MyError
    if errors.As(err, &myErr) {
        fmt.Println("It was a MyError!")
        fmt.Println("Time:", myErr.When)
        fmt.Println("Message:", myErr.What)
    } else {
        fmt.Println("Not a MyError.")
    }
}
```
This is the modern replacement for using a type assertion (`if e, ok := err.(*MyError); ok`), because it correctly handles wrapped errors.

## Summary: Lead Engineer Perspective

As a lead engineer, you should enforce the use of these modern error handling patterns:
- **Wrap errors with context** using `fmt.Errorf` and `%w` to create informative error messages that aid debugging.
- **Use `errors.Is`** to check for specific sentinel error values.
- **Use `errors.As`** to check for and interact with specific error types.
- Discourage direct equality checks (`err == os.ErrNotExist`) and type assertions, as they are not robust to error wrapping. 