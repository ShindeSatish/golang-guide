# Object-Oriented Concepts in Go: Methods

Go does not have classes. However, you can define methods on types. A method is a function with a special *receiver* argument.

## Method Declaration

The receiver appears in its own argument list between the `func` keyword and the method name.

**Syntax:**
`func (v <ReceiverType>) <MethodName>() <ReturnTypes> { ... }`

**Example:**
```go
package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    X, Y float64
}

// A method on the Vertex type
func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
    v := Vertex{3, 4}
    fmt.Println(v.Abs()) // Call the method
}
```
Here, the `Abs` method has a receiver of type `Vertex` named `v`.

## Pointer Receivers vs. Value Receivers

You can declare methods with either a pointer receiver (`*T`) or a value receiver (`T`).

1.  **Value Receiver (`func (v T) Method()`):**
    - The method operates on a *copy* of the original value.
    - Modifications made inside the method will **not** affect the original value.

2.  **Pointer Receiver (`func (v *T) Method()`):**
    - The method operates on a pointer to the original value.
    - Modifications made inside the method **will** affect the original value.
    - It avoids copying the value on each method call, which can be more efficient for large structs.

**Example:**
```go
package main

import "fmt"

type Vertex struct {
    X, Y float64
}

// Value receiver - operates on a copy
func (v Vertex) ScaleCopy(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

// Pointer receiver - operates on the original
func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func main() {
    v := Vertex{3, 4}
    v.ScaleCopy(10)
    fmt.Println(v) // {3 4} - unchanged

    v.Scale(10)
    fmt.Println(v) // {30 40} - changed
}
```

## When to Use Pointer vs. Value Receivers

The two main reasons to use a pointer receiver are:
1.  **To modify the receiver:** If the method needs to change the state of the receiver, it *must* be a pointer receiver.
2.  **To avoid copying:** If the receiver is a large struct, a pointer receiver is more efficient.

**General Rule of Thumb:**
- If in doubt, use a pointer receiver.
- Be consistent. If a type has some methods with pointer receivers, the other methods should probably have pointer receivers too.

## Methods and Pointer Indirection

Go is helpful when it comes to methods. If you have a value `v` of type `T`, and `T` has a method with a `*T` receiver, you can call `v.Method()` and Go will automatically interpret it as `(&v).Method()`.

Conversely, if you have a pointer `p` of type `*T`, and `T` has a method with a `T` receiver, you can call `p.Method()` and Go will interpret it as `(*p).Method()`.

This convenience means you don't have to constantly worry about whether your variable is a value or a pointer when calling a method. 