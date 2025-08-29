# Data Structures: Structs

This section covers structs in Go.

## Definition

A struct is a composite type that groups together zero or more named fields of arbitrary types. They are useful for creating complex data structures.

**Syntax:**
```go
type <StructName> struct {
    <FieldName> <FieldType>
    ...
}
```

**Example:**
```go
type Vertex struct {
    X int
    Y int
}
```

## Creating and Accessing Structs

```go
package main

import "fmt"

type Vertex struct {
    X, Y int
}

func main() {
    // Create a struct
    v := Vertex{1, 2}
    fmt.Println(v)

    // Access fields
    v.X = 4
    fmt.Println(v.X)

    // Pointers to structs
    p := &v
    p.X = 1e9 // Same as (*p).X
    fmt.Println(v)

    // Struct Literals
    var (
        v1 = Vertex{1, 2}  // has type Vertex
        v2 = Vertex{X: 1}  // Y:0 is implicit
        v3 = Vertex{}      // X:0 and Y:0
        p1 = &Vertex{1, 2} // has type *Vertex
    )
    fmt.Println(v1, p1, v2, v3)
}
```

## Structs are Value Types

Like arrays, structs are value types. When you assign a struct to a new variable or pass it to a function, a copy of the entire struct is made.

```go
v1 := Vertex{1, 2}
v2 := v1 // v2 is a copy of v1
v2.X = 100
fmt.Println(v1.X) // Prints 1, not 100
```
For this reason, it is common to pass pointers to structs to functions to avoid copying large amounts of data and to allow functions to modify the original struct.

## Embedded Structs and Anonymous Fields

Go does not have inheritance in the traditional sense, but it does allow for composition through struct embedding. You can "embed" a struct within another struct, and the fields of the inner struct are promoted to the outer one.

**Syntax:**
```go
type MyStruct struct {
    OtherStruct // Anonymous field
    ...
}
```

**Example:**
```go
package main

import "fmt"

type Point struct {
    X, Y int
}

type Circle struct {
    Point  // Anonymous field
    Radius int
}

func main() {
    c := Circle{
        Point: Point{1, 2},
        Radius: 5,
    }
    // Fields of Point are promoted to Circle
    fmt.Println(c.X, c.Y, c.Radius) // "1 2 5"
    c.X = 10
    fmt.Println(c.Point.X) // "10"
}
```
This is a powerful way to build up complex types by composing simpler ones.

## Struct Field Tags

Struct tags are strings of metadata that can be attached to the fields of a struct. They are ignored by the Go compiler but can be accessed using reflection.

**Use Cases:**
- Specifying field names for encoding/decoding (e.g., JSON, XML).
- Database mapping (e.g., ORM field names).
- Validation rules.

**Syntax:**
`<FieldName> <Type> \`key:"value"\``

**Example (JSON):**
```go
import "encoding/json"

type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Password string `json:"-"` // The "-" tag means "don't encode this field"
}
```
When an instance of this `User` struct is marshaled to JSON, the `ID` field will be named `id`, `Username` will be `username`, and `Password` will be omitted. This is crucial for controlling the serialization format of your data. 