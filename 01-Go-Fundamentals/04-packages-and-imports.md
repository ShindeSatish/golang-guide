# Go Fundamentals: Packages and Imports

This section covers how Go code is organized into packages.

## Packages

Every Go program is made up of packages. The program starts running in package `main`.

**Key Concepts:**
- A package is a directory of Go files.
- All files in a directory must belong to the same package.
- The package name is declared at the top of each file with `package <name>`.
- By convention, the package name is the same as the last element of the import path. For example, the files in `"math/rand"` start with `package rand`.

## Imports

The `import` statement is used to import packages.

**Syntax:**
```go
import "fmt"
import "math"
```
Or, a parenthesized "factored" import statement:
```go
import (
    "fmt"
    "math"
)
```

## Exported Names

In Go, a name is exported if it begins with a capital letter.

- **Exported (Public):** `Pi` in the `math` package. Can be accessed from other packages (e.g., `math.Pi`).
- **Unexported (Private):** `pi`. Cannot be accessed from outside its own package.

This applies to variables, functions, types, and struct fields.

**Example:**
```go
// in package "mymath"
package mymath

// Pi is an exported name
const Pi = 3.14159

// not exported
func secretFormula() {}
```

```go
// in package "main"
package main

import (
    "fmt"
    "path/to/mymath"
)

func main() {
    fmt.Println(mymath.Pi)
    // mymath.secretFormula() // This would be a compile-time error
}
```

## Package Initialization

Go has an `init()` function that can be defined in any file of a package.

**Features:**
- An `init()` function cannot be called or referenced.
- It is executed when the package is imported.
- A package can have multiple `init()` functions (in one or multiple files). They are executed in the order they appear to the compiler.

**Execution Order:**
1. Imported packages are initialized first.
2. Package-level variables are initialized.
3. `init()` functions are executed.
4. `main()` function is executed.

**Example:**
```go
// in package "mypkg"
package mypkg

import "fmt"

var MyVar int

func init() {
    fmt.Println("Initializing mypkg")
    MyVar = 100
}
```
```go
// in package "main"
package main

import (
    "fmt"
    "path/to/mypkg"
)

func main() {
    fmt.Println("main function starts")
    fmt.Println("MyVar from mypkg:", mypkg.MyVar)
}

// Output:
// Initializing mypkg
// main function starts
// MyVar from mypkg: 100
```
**Why use `init()`?**: It's used for setup tasks that must be done before the program starts, such as setting up database connections or initializing data structures.

## Package Aliases and Blank Identifier

### Aliases
You can use an alias for an imported package.
`import f "fmt"`

### Blank Identifier
Importing a package with the blank identifier `_` executes the package's `init()` function without making its exported names available.
`import _ "net/http/pprof"`

**Use Case:** This is often used for packages that have side effects upon initialization, such as registering a database driver or an HTTP handler. 