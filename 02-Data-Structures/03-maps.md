# Data Structures: Maps

This section covers maps in Go.

## Definition

A map is an unordered collection of key-value pairs. It is Go's implementation of a hash table.

**Syntax:** `map[KeyType]ValueType`

**Key Characteristics:**
- **Unordered:** The iteration order over a map is not specified and is not guaranteed to be the same from one iteration to the next.
- **Reference Type:** Like slices, maps are reference types. When you pass a map to a function, a copy of the pointer to the underlying data structure is passed. Changes made to the map inside the function will be visible to the caller.
- The zero value of a map is `nil`. A `nil` map has no keys, nor can keys be added.

## Declaration and `make`

```go
package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

var m map[string]Vertex

func main() {
    // A map must be created with `make` before use.
    m = make(map[string]Vertex)
    m["Bell Labs"] = Vertex{
        40.68433, -74.39967,
    }
    fmt.Println(m["Bell Labs"])

    // Map literal
    var n = map[string]Vertex{
        "Bell Labs": {40.68433, -74.39967},
        "Google":    {37.42202, -122.08408},
    }
    fmt.Println(n)
}
```

## Manipulating Maps

```go
m := make(map[string]int)

// Insert or update an element
m["Answer"] = 42

// Retrieve an element
v := m["Answer"]

// Delete an element
delete(m, "Answer")

// Test if a key is present with a two-value assignment.
// This is a common idiom.
elem, ok := m["Answer"]
if ok {
    fmt.Println("The value:", elem)
} else {
    fmt.Println("The key is not present")
}
```
If a key is not in the map, `elem` will be the zero value for the map's value type. The `ok` boolean is how you distinguish between a key that is absent and a key that is present with a zero value.

## `for-range` on Maps

Iterating over a map yields key-value pairs.

```go
m := map[string]int{"one": 1, "two": 2, "three": 3}
for key, value := range m {
    fmt.Printf("key: %s, value: %d\n", key, value)
}
```
Again, the order of iteration is not guaranteed.

## Internals: Hash Tables

A map is a pointer to a `runtime.hmap` struct. This struct contains information about the map, including a pointer to an array of `runtime.bmap` structs (buckets).

1.  **Hashing:** When you insert a key-value pair, Go hashes the key to determine which bucket it should go into.
2.  **Buckets:** Each bucket is a small, fixed-size array that can hold a few key-value pairs.
3.  **Collisions:** If multiple keys hash to the same bucket (a collision), they are stored in the same bucket. If a bucket becomes full, a new "overflow" bucket is allocated and linked to it.
4.  **Growth:** When the map becomes too full (the "load factor" exceeds a certain threshold, currently 6.5), the map is grown. A new, larger array of buckets is allocated (typically double the size), and the existing key-value pairs are gradually moved to the new buckets. This is an incremental process to avoid long pauses.

**Why is iteration order random?**
To ensure that programmers do not rely on a specific iteration order, the starting point of the iteration is randomized. This was a deliberate choice by the Go team after observing that some programmers were unintentionally depending on the (then-stable) iteration order of older Go versions. 