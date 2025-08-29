# Data Structures: Arrays

This section covers arrays in Go.

## Definition

An array is a numbered sequence of elements of a specific length.

**Syntax:**
`[n]T` is an array of `n` values of type `T`.

**Key Characteristics:**
- **Fixed Size:** The length of an array is part of its type. `[5]int` and `[10]int` are distinct types. Because of this, arrays are somewhat inflexible, and slices are almost always used instead.
- **Value Type:** Arrays in Go are value types. When you assign an array to a new variable or pass it to a function, a copy of the entire array is made.

## Declaration and Initialization

```go
package main

import "fmt"

func main() {
    // Declare an array of 5 integers.
    // All elements are initialized to their zero value (0 for int).
    var a [5]int
    fmt.Println("emp:", a)

    // Set a value
    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    // Get the length
    fmt.Println("len:", len(a))

    // Declare and initialize in one line
    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

    // Let the compiler count the elements for you
    c := [...]int{1, 2, 3}
    fmt.Println("dcl:", c, "len:", len(c))

    // Multi-dimensional arrays
    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
```

## Arrays are Values

This is a critical concept for understanding how arrays behave.

```go
package main

import "fmt"

func main() {
    a := [3]int{1, 2, 3}
    b := a // A copy of 'a' is assigned to 'b'
    b[0] = 100

    fmt.Println("a:", a) // a is unchanged
    fmt.Println("b:", b)
}
// Output:
// a: [1 2 3]
// b: [100 2 3]
```

## When to Use Arrays

While slices are more common, arrays have their uses:
- **Strict Memory Layout:** When you need a precise memory layout. Arrays are contiguous blocks of memory.
- **Performance-Critical Code:** To avoid the overhead of slice headers or to prevent heap allocations (if the array is on the stack). For example, a cryptographic function that operates on a fixed-size block of data might use an array.

## Internals

An array is a single, contiguous block of memory. The size is fixed at compile time. For an array `var a [10]int`, the memory layout is just 10 integers, one after the other. This makes them very efficient to access. However, the value semantics and fixed size make them less flexible than slices for many common programming tasks. 