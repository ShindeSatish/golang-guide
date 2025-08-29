# Advanced Topics: Generics

Generics were introduced in Go 1.18. They provide a way to write functions and data structures that can work with any of a set of types, while preserving compile-time type safety.

## What Problem Do Generics Solve?

Before generics, if you wanted to write a function that operated on different types (e.g., a `Min` function for both `int` and `float64`), you had a few options, none of them ideal:
- Write multiple functions (`MinInt`, `MinFloat64`). This leads to code duplication.
- Use `interface{}` (or `any`). This sacrifices type safety, as you need type assertions at runtime, and can be less performant.

Generics solve this by allowing you to write a single function that is type-safe and efficient.

## Type Parameters

Generics introduce the concept of **type parameters**. A type parameter is a placeholder for a type that will be provided by the calling code.

**Syntax:**
A function with type parameters includes them in square brackets `[]` before the regular function arguments.

```go
func MyGenericFunc[T any](arg T) {
    // ...
}
```
Here, `T` is a type parameter. `any` is a **constraint** that means `T` can be any type.

**Example: A Generic `Print` function**
```go
package main

import "fmt"

func Print[T any](s []T) {
    for _, v := range s {
        fmt.Printf("%v ", v)
    }
    fmt.Println()
}

func main() {
    Print([]int{1, 2, 3})
    Print([]string{"a", "b", "c"})
}
```
The compiler *instantiates* a version of the `Print` function for each type it is called with (`int` and `string`).

## Constraints

A constraint is an interface that specifies the methods that a type argument must have. This allows your generic function to operate on the type argument in a type-safe way.

The `comparable` constraint is a built-in constraint that includes all types that can be compared with `==` and `!=`. The `any` constraint (an alias for `interface{}`) allows any type.

You can define your own constraints as interfaces.

**Example: A Generic `Sum` function**
Let's write a function that sums the values in a map. The map's values could be `int64` or `float64`.

```go
package main

import "fmt"

// Define a constraint that includes int64 and float64
type Number interface {
    int64 | float64
}

// SumNumbers sums the values of map m.
// It supports both int64 and float64 as value types.
func SumNumbers[K comparable, V Number](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}

func main() {
    ints := map[string]int64{"a": 1, "b": 2}
    floats := map[string]float64{"a": 1.1, "b": 2.2}

    fmt.Println("Sum of ints:", SumNumbers(ints))
    fmt.Println("Sum of floats:", SumNumbers(floats))
}
```
- `K comparable`: The map key type `K` can be any comparable type.
- `V Number`: The map value type `V` must be either `int64` or `float64`, as defined by our `Number` interface. The `|` syntax is called a *union* and is new with generics.
- The `+` operator is allowed on `s` and `v` because the constraint guarantees that `V` will be a type that supports addition.

## Generic Data Structures

You can also define generic types, like a linked list or a binary tree, that can hold values of any type.

```go
// A generic stack
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
    // ... implementation ...
}
```

## Lead Engineer Perspective

- **Embrace Generics:** Encourage the team to use generics where they can reduce code duplication without sacrificing clarity. Common data structures and utility functions are prime candidates.
- **Don't Overuse:** Generics are not a replacement for interfaces. If you have a set of types that share common *behavior*, an interface is still the right tool. Generics are for when you want the same *algorithm* to operate on different types.
- **Write Clear Constraints:** A good, clear constraint is key to a useful generic function. It documents what the function needs and ensures type safety. 