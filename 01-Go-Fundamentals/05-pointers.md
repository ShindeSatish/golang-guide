# Go Fundamentals: Pointers

This section covers pointers in Go.

## What is a Pointer?

A pointer holds the memory address of a value.

- The type `*T` is a pointer to a `T` value. Its zero value is `nil`.
- The `&` operator generates a pointer to its operand.
- The `*` operator *dereferences* a pointer, giving access to the underlying value.

**Example:**
```go
package main

import "fmt"

func main() {
    i, j := 42, 2701

    p := &i         // p points to i
    fmt.Println(*p) // read i through the pointer p
    *p = 21         // set i through the pointer p
    fmt.Println(i)  // see the new value of i

    p = &j         // p now points to j
    *p = *p / 37   // divide j through the pointer
    fmt.Println(j) // see the new value of j
}
```

## Why Use Pointers?

1.  **Modify arguments passed to a function:** Go passes arguments by value. To modify a variable inside a function, you must pass a pointer to it.
2.  **Efficiency:** To avoid copying large data structures (like large structs). Passing a pointer is much cheaper than copying the entire structure.
3.  **Indicate absence of a value:** A pointer can be `nil`, which can be used to signify that a value is missing or not applicable. This is common for optional fields in structs.

### Example: Modifying a function argument
```go
package main

import "fmt"

func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
    i := 1
    fmt.Println("initial:", i)

    zeroval(i)
    fmt.Println("zeroval:", i)

    zeroptr(&i)
    fmt.Println("zeroptr:", i)
    fmt.Println("pointer:", &i)
}
// Output:
// initial: 1
// zeroval: 1
// zeroptr: 0
// pointer: 0x...
```
`zeroval` gets a copy of `i`, so the change is not reflected in `main`. `zeroptr` gets a pointer to `i`, so it can modify the original variable.

## Pointers to Structs

It is common to use pointers with structs.

**Example:**
```go
type Vertex struct {
    X, Y int
}

v := Vertex{1, 2}
p := &v
p.X = 1e9 // equivalent to (*p).X
```
Go provides the `p.X` syntax as a convenience; you don't need to explicitly dereference with `(*p).X`.

## Pointers and Garbage Collection

Go has a garbage collector, so you don't need to manually free memory. When a value is no longer referenced by any pointers, its memory will be reclaimed by the garbage collector.

## Internals: Stack vs. Heap

- **Stack:** Used for static memory allocation. Local variables in a function are typically allocated on the stack. Memory is allocated and de-allocated automatically as functions are called and return.
- **Heap:** Used for dynamic memory allocation. Objects that need to exist beyond a single function call are allocated on the heap.

**Escape Analysis:** The Go compiler performs *escape analysis* to determine if a variable should be allocated on the stack or the heap. If a variable's address is taken and might be used after the function returns, it "escapes" to the heap.

**Example:**
```go
func newInt() *int {
    var i int = 42
    return &i // 'i' escapes to the heap
}
```
In this case, `i` is allocated on the heap because its address is returned and will be used by the caller. If `&i` were only used within `newInt`, it would likely be allocated on the stack. This is a key optimization in Go that reduces the pressure on the garbage collector. 