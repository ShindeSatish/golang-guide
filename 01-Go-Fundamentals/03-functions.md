# Go Fundamentals: Functions

This section covers functions in Go.

## Basic Syntax

**Syntax:**
`func <functionName>(<params>) <return types> { ... }`

**Example:**
```go
package main

import "fmt"

// A simple function with two int params and one int return value.
func add(x int, y int) int {
    return x + y
}

// When two or more consecutive named function parameters share a type,
// you can omit the type from all but the last.
func subtract(x, y int) int {
    return x - y
}

func main() {
    fmt.Println(add(42, 13))       // 55
    fmt.Println(subtract(42, 13)) // 29
}
```

## Multiple Return Values

A function can return any number of results.

**Example:**
```go
package main

import "fmt"

func swap(x, y string) (string, string) {
    return y, x
}

func main() {
    a, b := swap("hello", "world")
    fmt.Println(a, b) // "world hello"
}
```
**Use Case:** This is idiomatic in Go for returning a result and an error value (e.g., `result, err := someFunc()`).

## Named Return Values

Return values may be named. They are treated as variables defined at the top of the function. A `return` statement without arguments returns the named return values. This is known as a "naked" return.

**Example:**
```go
package main

import "fmt"

func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

func main() {
    fmt.Println(split(17)) // 7 10
}
```
**When to use:** Named return values are useful for documenting the meaning of return values. Naked returns are best used in short functions, as they can harm readability in longer ones.

## Variadic Functions

A function that can be called with any number of trailing arguments.

**Syntax:**
`func <functionName>(<params>...<type>)`

**Example:**
```go
package main

import "fmt"

func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {
    sum(1, 2)
    sum(1, 2, 3)

    nums := []int{1, 2, 3, 4}
    sum(nums...) // Slices can be passed with the ... suffix.
}
```

## Functions as Values

Functions are first-class citizens in Go. They can be treated like any other value.

**Example:**
```go
package main

import (
    "fmt"
    "math"
)

func compute(fn func(float64, float64) float64) float64 {
    return fn(3, 4)
}

func main() {
    hypot := func(x, y float64) float64 {
        return math.Sqrt(x*x + y*y)
    }
    fmt.Println(hypot(5, 12))      // 13
    fmt.Println(compute(hypot))   // 5
    fmt.Println(compute(math.Pow)) // 81
}
```

## Closures

A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables.

**Example:**
```go
package main

import "fmt"

func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func main() {
    pos, neg := adder(), adder()
    for i := 0; i < 10; i++ {
        fmt.Println(
            pos(i),
            neg(-2*i),
        )
    }
}
```
**Internals:** The `sum` variable is "closed over" by the inner anonymous function. Each call to `adder` returns a new closure, each with its own `sum` variable. This is a powerful way to manage state. 