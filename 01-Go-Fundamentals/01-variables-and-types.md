# Go Fundamentals: Variables and Types

This section covers the basics of variables and types in Go.

## Declaration and Initialization

### `var` keyword
You can declare variables using the `var` keyword.

**Syntax:**
`var <variableName> <type>`

**Example:**
```go
package main

import "fmt"

func main() {
    // Declare a variable 'i' of type int
    var i int
    fmt.Println("i:", i) // Prints "i: 0" (zero value)

    // Declare and initialize
    var s string = "hello"
    fmt.Println("s:", s)
}
```

### Short Variable Declaration `:=`
The `:=` syntax is shorthand for declaring and initializing a variable. The type is inferred. It can only be used inside functions.

**Syntax:**
`<variableName> := <value>`

**Example:**
```go
package main

import "fmt"

func main() {
    // Shorthand declaration
    f := 3.14
    fmt.Printf("f is of type %T\n", f) // Prints "f is of type float64"
}
```

## Basic Types

Go has a rich set of built-in types.

- **Numeric Types:**
  - `int`, `int8`, `int16`, `int32`, `int64`
  - `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`
  - `float32`, `float64`
  - `complex64`, `complex128`
  - `byte` (alias for `uint8`)
  - `rune` (alias for `int32`, represents a Unicode code point)

- **Boolean Type:**
  - `bool` (`true` or `false`)

- **String Type:**
  - `string`

## Zero Values

In Go, variables declared without an explicit initial value are given their *zero value*.

- `0` for numeric types
- `false` for the boolean type
- `""` (the empty string) for strings
- `nil` for pointers, functions, interfaces, slices, channels, and maps.

## Type Conversion

Go requires explicit type conversion.

**Syntax:**
`<type>(<value>)`

**Example:**
```go
package main

import "fmt"

func main() {
    var i int = 42
    var f float64 = float64(i)
    var u uint = uint(f)
    fmt.Println(i, f, u) // "42 42 42"
}
```

## Internals and "Why"

- **Static Typing:** Go is statically typed. This means variable types are known at compile time, which allows the compiler to catch errors early and optimize code execution. This prevents a whole class of runtime errors.
- **Why `:=`?**: The short variable declaration was introduced to make Go code more concise and readable, especially for local variables within a function. It reduces verbosity.
- **Memory Allocation**: When you declare a variable, Go allocates memory for it on the stack (for local variables) or in the heap (for variables that escape the function's scope). The zero value ensures that this memory has a predictable state, preventing bugs from uninitialized variables. 