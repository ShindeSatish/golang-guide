# Advanced Topics: Reflection

Reflection is the ability of a program to examine its own structure, particularly through types; it's a form of metaprogramming. The `reflect` package provides this capability in Go.

## Why Use Reflection?

Reflection is a powerful and advanced feature that should be used with care. It's often necessary for:
- Writing generic functions that can work with values of any type.
- Implementing marshaling and unmarshaling for data formats like JSON.
- Building ORMs or other tools that need to dynamically inspect struct fields and tags.

**Warning:** Reflection is slower than normal code execution and can be more complex to write and maintain. It also bypasses static type checking. Always prefer to solve a problem without reflection if possible.

## `reflect.Type` and `reflect.Value`

The `reflect` package is built around two key types:
- `reflect.Type`: Represents a Go type.
- `reflect.Value`: Represents a Go value.

You can get these from a variable using `reflect.TypeOf()` and `reflect.ValueOf()`.

```go
package main

import (
    "fmt"
    "reflect"
)

func main() {
    var x float64 = 3.4
    fmt.Println("type:", reflect.TypeOf(x))
    fmt.Println("value:", reflect.ValueOf(x))
}
```

## Inspecting Structs

A common use case for reflection is to inspect the fields and tags of a struct.

```go
package main

import (
    "fmt"
    "reflect"
)

type MyStruct struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func main() {
    s := MyStruct{ID: 1, Name: "Gopher"}
    t := reflect.TypeOf(s)

    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        fmt.Printf("Field: %s, Type: %s, Tag: %s\n",
            field.Name,
            field.Type,
            field.Tag.Get("json"),
        )
    }
}
// Output:
// Field: ID, Type: int, Tag: id
// Field: Name, Type: string, Tag: name
```

## Modifying Values with Reflection

To modify a value using reflection, you must have a `reflect.Value` that is *settable*. A `reflect.Value` is settable if it was created from a pointer.

- `reflect.ValueOf(&x)` gives you a `Value` of a pointer.
- `.Elem()` on a `Value` of a pointer gives you a settable `Value` of the underlying data.

```go
package main

import (
    "fmt"
    "reflect"
)

func main() {
    var x float64 = 3.4
    v := reflect.ValueOf(&x).Elem() // Get a settable Value

    fmt.Println("Is settable:", v.CanSet())

    v.SetFloat(7.1)
    fmt.Println(x) // Prints 7.1

    // Trying to set a non-pointer value will panic
    // reflect.ValueOf(x).SetFloat(7.1) // This would panic
}
```

## Calling Methods with Reflection

You can also call methods on a type using reflection.

```go
type MyType struct{}
func (t MyType) MyMethod(name string) {
    fmt.Println("Hello,", name)
}

v := reflect.ValueOf(MyType{})
method := v.MethodByName("MyMethod")
args := []reflect.Value{reflect.ValueOf("Gopher")}
method.Call(args) // Calls MyMethod
```

## Lead Engineer Perspective

As a lead, you should be cautious about where reflection is used in your codebase.
- **Approve its use** for framework-level code (like custom encoders, dependency injection containers) where its power is genuinely needed.
- **Discourage its use** in regular application logic where a simpler, type-safe solution (like interfaces) would suffice.
- Ensure that any reflection-heavy code is well-documented, explaining *why* it's necessary, and is accompanied by thorough tests to compensate for the lack of compile-time safety. 