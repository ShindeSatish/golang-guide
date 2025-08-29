# Data Structures: Slices

Slices are a key data structure in Go, providing a more powerful interface to sequences than arrays.

## Definition

A slice is a dynamically-sized, flexible view into the elements of an array.

**Key Characteristics:**
- **Flexible Size:** Slices can be resized (within the capacity of the underlying array).
- **Reference Type (mostly):** A slice is a descriptor of an array segment. It consists of a pointer to the array, the length of the segment, and its capacity. When you pass a slice, you are copying this descriptor, not the underlying data. Changes to the elements of the slice will be visible to other slices that share the same underlying array.

## Declaration and `make`

```go
package main

import "fmt"

func main() {
    // Using the `make` built-in function is how you create
    // dynamically-sized slices.
    s := make([]string, 3)
    fmt.Println("emp:", s)

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])
    fmt.Println("len:", len(s))
}
```

## Slicing a Slice

You can create a new slice from an existing one with the `[low:high]` operator.

**Syntax:** `a[low:high]`
- Selects a half-open range which includes the first element, but excludes the last one.

```go
a := [5]int{1, 2, 3, 4, 5}
s := a[1:4] // s is []int{2, 3, 4}
```

## Length and Capacity

- **Length (`len`)**: The number of elements in the slice.
- **Capacity (`cap`)**: The number of elements in the underlying array, counting from the first element in the slice.

```go
s := []int{2, 3, 5, 7, 11, 13}
// len(s) is 6
// cap(s) is 6

s = s[:0]
// len(s) is 0
// cap(s) is 6

s = s[1:4]
// len(s) is 3
// cap(s) is 5
```

## `append`

The built-in `append` function is used to add elements to a slice.

```go
var s []int
s = append(s, 1) // len=1 cap=1
s = append(s, 2, 3, 4) // len=4 cap=4 (or more)
```
**Internals of `append`**: If the capacity of the original slice is not large enough, `append` will create a new, larger underlying array and copy the existing elements to it. The returned slice will point to this new array. This is why you must always assign the result of `append` back to the slice variable (`s = append(s, ...)`).

## `for-range`

The `for-range` loop is the idiomatic way to iterate over a slice.
```go
for index, value := range pow {
    fmt.Printf("2**%d = %d\n", index, value)
}
```

## Internals: Slice Header

A slice is a small struct, called a *slice header*, that contains three fields:
1.  **Pointer:** A pointer to the first element of the underlying array that is reachable through the slice.
2.  **Length:** The number of elements in the slice (`len()`).
3.  **Capacity:** The maximum number of elements the slice can hold without reallocation (`cap()`).

```
type sliceHeader struct {
    Data uintptr
    Len  int
    Cap  int
}
```
When you pass a slice to a function, you are copying this header. Because the `Data` pointer points to the same underlying array, modifications to the slice's elements will be reflected in the original. However, if the function uses `append` and causes a reallocation, the new slice header will point to a *new* array, and changes will no longer affect the original. This is a common source of bugs for new Go programmers. 