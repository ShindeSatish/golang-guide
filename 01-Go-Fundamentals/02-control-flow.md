# Go Fundamentals: Control Flow

This section covers control flow statements in Go.

## `if/else`

Go's `if` statements are similar to other languages, but with a few distinctions.

**Syntax:**
```go
if <condition> {
    // ...
} else if <condition> {
    // ...
} else {
    // ...
}
```

**Features:**
- No parentheses `()` around conditions.
- Braces `{}` are always required.
- A short statement can be executed before the condition, its scope limited to the `if` block.

**Example with a short statement:**
```go
package main

import (
    "fmt"
    "math"
)

func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    }
    return lim
}

func main() {
    fmt.Println(
        pow(3, 2, 10), // 9
        pow(3, 3, 20), // 20
    )
}
```
**When to use:** Use the short statement to declare variables that are only needed within the `if/else` branches, improving code clarity and managing variable scope.

## `for`

Go has only one looping construct, the `for` loop.

### 1. The complete `for` statement (like C's `for`)
**Syntax:**
`for <init>; <condition>; <post> { ... }`

**Example:**
```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

### 2. The `while` loop (condition-only `for`)
**Syntax:**
`for <condition> { ... }`

**Example:**
```go
sum := 1
for sum < 1000 {
    sum += sum
}
```

### 3. The infinite loop
**Syntax:**
`for { ... }`

### 4. `for-range` loop
Used to iterate over slices, arrays, strings, maps, and channels.

**Example (Slice):**
```go
s := []string{"a", "b", "c"}
for index, value := range s {
    fmt.Printf("index: %d, value: %s\n", index, value)
}
```
If you only need the value, you can use the blank identifier `_`:
```go
for _, value := range s {
    fmt.Println(value)
}
```

## `switch`

Go's `switch` is more flexible than in C or Java.

**Features:**
- Cases can be non-constants.
- Cases are evaluated from top to bottom, and the first match is chosen.
- `break` is provided automatically (no "fallthrough" by default). Use the `fallthrough` keyword if you need it.
- A `switch` without an expression is an alternate way to express `if/else` logic.

**Example:**
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }
}
```

**"Tagless" Switch Example:**
```go
t := time.Now()
switch {
case t.Hour() < 12:
    fmt.Println("Good morning!")
case t.Hour() < 17:
    fmt.Println("Good afternoon.")
default:
    fmt.Println("Good evening.")
}
```

## `defer`

A `defer` statement defers the execution of a function until the surrounding function returns.

**Use Cases:**
- Closing files or network connections.
- Unlocking mutexes.
- Printing debug information.

**Example:**
```go
package main

import "fmt"

func main() {
    defer fmt.Println("world")
    fmt.Println("hello")
    // Output:
    // hello
    // world
}
```

**Internals:**
- Deferred calls are pushed onto a stack. When the function returns, the calls are popped and executed in LIFO (Last-In, First-Out) order.
- Arguments to deferred functions are evaluated when the `defer` statement is executed, not when the call is executed.
```go
func a() {
    i := 0
    defer fmt.Println(i) // Prints 0
    i++
    return
}
``` 