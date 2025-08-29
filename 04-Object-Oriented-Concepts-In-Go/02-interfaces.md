# Object-Oriented Concepts in Go: Interfaces

Interfaces in Go provide a way to specify the behavior of an object: if something can do *this*, then it can be used *here*.

## Definition

An interface type is defined as a set of method signatures.

**Syntax:**
```go
type <InterfaceName> interface {
    <MethodName>(<args>) <return types>
    ...
}
```

A value of an interface type can hold any value that implements all the methods of the interface.

**Example:**
```go
type Abser interface {
    Abs() float64
}
```
Any type that has an `Abs() float64` method is said to *implement* the `Abser` interface.

## Implicit Implementation

A type implements an interface by implementing its methods. There is no explicit declaration like `class MyType implements MyInterface`. This is known as *implicit implementation* or *structural typing*.

```go
package main

import (
    "fmt"
    "math"
)

// The interface
type Abser interface {
    Abs() float64
}

// A type that implements the interface
type MyFloat float64

func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}

// Another type that implements the interface
type Vertex struct {
    X, Y float64
}

func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
    var a Abser

    f := MyFloat(-math.Sqrt2)
    v := Vertex{3, 4}

    a = f  // a MyFloat implements Abser
    a = &v // a *Vertex implements Abser

    // a = v // This is a compile error because Vertex (value receiver)
             // does not implement Abser, but *Vertex (pointer receiver) does.

    fmt.Println(a.Abs())
}
```
This decoupling of definition and implementation is a key feature of Go. It allows you to write functions that can operate on types that you didn't even know about when you wrote the function.

## The Empty Interface: `interface{}`

The interface type that specifies zero methods is known as the *empty interface*: `interface{}`.

An empty interface may hold values of any type. They are used by code that handles values of unknown type.

**Example:**
```go
var i interface{}
i = 42
i = "hello"
```
**`any` keyword:** In Go 1.18 and later, `any` is an alias for `interface{}`. It is recommended to use `any` for clarity.

## Type Assertions and Type Switches

To access the underlying concrete value of an interface, you use a *type assertion*.
`t := i.(T)`

This asserts that the interface value `i` holds the concrete type `T` and assigns the underlying `T` value to the variable `t`. If `i` does not hold a `T`, this statement will trigger a panic.

To test whether an interface value holds a specific type, a type assertion can return two values:
`t, ok := i.(T)`

If `ok` is `true`, `t` has the type `T`. If not, `ok` is `false` and `t` is the zero value of type `T`.

A **type switch** is like a regular switch statement, but the cases specify types, not values.
```go
switch v := i.(type) {
case int:
    fmt.Printf("Twice %v is %v\n", v, v*2)
case string:
    fmt.Printf("%q is %v bytes long\n", v, len(v))
default:
    fmt.Printf("I don't know about type %T!\n", v)
}
```

## Internals: Interface Values

An interface value can be thought of as a tuple of a value and a concrete type: `(value, type)`.

1.  **Type:** A pointer to the type information for the stored value. This includes the method set for the type.
2.  **Value:** A pointer to the actual data.

When you call a method on an interface value, the runtime looks up the method in the type's method set and then calls it with the value pointer as the receiver.

A `nil` interface value has both a `nil` type and a `nil` value. A common pitfall is an interface value that holds a `nil` pointer of a concrete type. This interface value is **not nil**.
```go
var p *MyType = nil
var i MyInterface = p
// i is not nil, but i's underlying value is nil
```
Calling a method on `i` might cause a panic if the method doesn't handle a `nil` receiver. 