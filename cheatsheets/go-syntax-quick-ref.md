# Cheatsheet: Go Syntax Quick Reference

A quick reference for common Go syntax.

## Variables
```go
// Long form
var x int = 10

// Type is inferred
var y = 20

// Short declaration (inside functions only)
z := 30
```

## Basic Types
`int`, `float64`, `bool`, `string`, `byte` (alias for `uint8`), `rune` (alias for `int32`).

## Composite Types

### Structs
```go
type Point struct {
    X, Y int
}
p := Point{1, 2}
p.X = 10
```

### Arrays (fixed size)
```go
var a [5]int
b := [...]int{1, 2, 3}
```

### Slices (dynamic size)
```go
// Create with make
s := make([]string, 3)

// Slice literal
t := []string{"a", "b", "c"}

// Slicing
u := t[1:3] // {"b", "c"}

// Appending
t = append(t, "d")
```

### Maps
```go
// Create with make
m := make(map[string]int)

// Map literal
n := map[string]int{"one": 1, "two": 2}

// Accessing
val, ok := n["one"] // ok is true if "one" exists

// Deleting
delete(n, "one")
```

## Control Flow

### `if-else`
```go
if x > 0 {
    // ...
} else if x < 0 {
    // ...
} else {
    // ...
}

// With a short statement
if v := someFunc(); v > 10 {
    // ... v is scoped to the if block
}
```

### `for` loop
```go
// "C-style" for
for i := 0; i < 10; i++ {
    // ...
}

// "While" style
for x < 10 {
    // ...
}

// Infinite loop
for {
    // ...
}

// For-range loop
for index, value := range mySlice {
    // ...
}
```

### `switch`
```go
switch os := runtime.GOOS; os {
case "darwin":
    // ...
case "linux":
    // ...
default:
    // ...
}
```

### `defer`
Executes a function call just before the surrounding function returns.
```go
f, _ := os.Open("file.txt")
defer f.Close()
```

## Functions
```go
func add(x, y int) int {
    return x + y
}

// Multiple return values
func swap(x, y string) (string, string) {
    return y, x
}
```

## Methods & Interfaces

### Method
```go
type MyType struct{}

// Method with a pointer receiver
func (t *MyType) MyMethod() {
    // ...
}
```

### Interface
```go
type MyInterface interface {
    MyMethod()
}
```

## Concurrency

### Goroutine
Starts a new lightweight thread of execution.
```go
go myFunction()
```

### Channel
```go
// Unbuffered channel
ch := make(chan int)

// Buffered channel
bufCh := make(chan int, 10)

// Send
ch <- 10

// Receive
val := <-ch
```

### `select`
Waits on multiple channel operations.
```go
select {
case msg1 := <-ch1:
    // ...
case ch2 <- msg2:
    // ...
default:
    // non-blocking operation
}
``` 