# Go Fundamentals: Functions

This section covers functions in Go.

## Basic Syntax

**Syntax:**
`func <functionName>(<params>) <return types> { ... }`

**Example:**
```go
package main

import "fmt"

// A simple function with two int params and one int return value.
func add(x int, y int) int {
    return x + y
}

// When two or more consecutive named function parameters share a type,
// you can omit the type from all but the last.
func subtract(x, y int) int {
    return x - y
}

func main() {
    fmt.Println(add(42, 13))       // 55
    fmt.Println(subtract(42, 13)) // 29
}
```

## Multiple Return Values

A function can return any number of results.

**Example:**
```go
package main

import "fmt"

func swap(x, y string) (string, string) {
    return y, x
}

func main() {
    a, b := swap("hello", "world")
    fmt.Println(a, b) // "world hello"
}
```
**Use Case:** This is idiomatic in Go for returning a result and an error value (e.g., `result, err := someFunc()`).

## Named Return Values

Return values may be named. They are treated as variables defined at the top of the function. A `return` statement without arguments returns the named return values. This is known as a "naked" return.

**Example:**
```go
package main

import "fmt"

func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

func main() {
    fmt.Println(split(17)) // 7 10
}
```
**When to use:** Named return values are useful for documenting the meaning of return values. Naked returns are best used in short functions, as they can harm readability in longer ones.

## Variadic Functions

A function that can be called with any number of trailing arguments.

**Syntax:**
`func <functionName>(<params>...<type>)`

**Example:**
```go
package main

import "fmt"

func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {
    sum(1, 2)
    sum(1, 2, 3)

    nums := []int{1, 2, 3, 4}
    sum(nums...) // Slices can be passed with the ... suffix.
}
```

## Functions as Values

Functions are first-class citizens in Go. They can be treated like any other value.

**Example:**
```go
package main

import (
    "fmt"
    "math"
)

func compute(fn func(float64, float64) float64) float64 {
    return fn(3, 4)
}

func main() {
    hypot := func(x, y float64) float64 {
        return math.Sqrt(x*x + y*y)
    }
    fmt.Println(hypot(5, 12))      // 13
    fmt.Println(compute(hypot))   // 5
    fmt.Println(compute(math.Pow)) // 81
}
```

## Closures

A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables.

**Example:**
```go
package main

import "fmt"

func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func main() {
    pos, neg := adder(), adder()
    for i := 0; i < 10; i++ {
        fmt.Println(
            pos(i),
            neg(-2*i),
        )
    }
}
```
**Internals:** The `sum` variable is "closed over" by the inner anonymous function. Each call to `adder` returns a new closure, each with its own `sum` variable. This is a powerful way to manage state.

### Benefits of Closures

Closures provide several key advantages in Go programming:

#### 1. **State Encapsulation** 🔒
Closures allow you to create private state that persists between function calls:

```go
func adder() func(int) int {
    sum := 0  // This variable is "closed over" - private to each closure
    return func(x int) int {
        sum += x  // Each closure has its own 'sum' variable
        return sum
    }
}
```

The `sum` variable is private - you can't access it directly from outside the closure.

#### 2. **Data Persistence** 💾
Unlike regular function calls, closures maintain state between invocations:

```go
pos := adder()
fmt.Println(pos(5))  // 5 (sum = 0 + 5)
fmt.Println(pos(3))  // 8 (sum = 5 + 3)
fmt.Println(pos(2))  // 10 (sum = 8 + 2)
```

#### 3. **Factory Pattern** 🏭
Create specialized functions with pre-configured behavior:

```go
func multiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

double := multiplier(2)
triple := multiplier(3)
fmt.Println(double(5))  // 10
fmt.Println(triple(5))  // 15
```

### Real-World Use Cases

#### 1. **Configuration and Middleware** ⚙️
```go
func withLogging(logger *log.Logger) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            logger.Printf("Request: %s %s", r.Method, r.URL.Path)
            next.ServeHTTP(w, r)
        })
    }
}

// Usage
logMiddleware := withLogging(myLogger)
handler := logMiddleware(myHandler)
```

#### 2. **Event Handlers with Context** 🎯
```go
func createButtonHandler(userID string, action string) func() {
    return func() {
        fmt.Printf("User %s performed %s\n", userID, action)
        // Handle the specific action for this user
    }
}

// Create specific handlers
saveHandler := createButtonHandler("user123", "save")
deleteHandler := createButtonHandler("user123", "delete")

saveHandler()    // User user123 performed save
deleteHandler()  // User user123 performed delete
```

#### 3. **Database Connection Pools** 🗄️
```go
func createDBHandler(db *sql.DB) func(string) (*sql.Rows, error) {
    return func(query string) (*sql.Rows, error) {
        // The closure "remembers" the db connection
        return db.Query(query)
    }
}

// Each service gets its own query function with the right DB
userService := createDBHandler(userDB)
productService := createDBHandler(productDB)
```

#### 4. **Caching and Memoization** 🧠
```go
func memoize(fn func(int) int) func(int) int {
    cache := make(map[int]int)
    return func(x int) int {
        if result, exists := cache[x]; exists {
            return result
        }
        result := fn(x)
        cache[x] = result
        return result
    }
}

// Expensive function with caching
fibonacci := memoize(func(n int) int {
    if n <= 1 { return n }
    return fibonacci(n-1) + fibonacci(n-2)
})
```

#### 5. **Rate Limiting** 🚦
```go
func createRateLimiter(maxRequests int, window time.Duration) func() bool {
    var requests []time.Time
    return func() bool {
        now := time.Now()
        // Remove old requests outside the window
        for i := 0; i < len(requests); i++ {
            if now.Sub(requests[i]) > window {
                requests = requests[i+1:]
                break
            }
        }
        
        if len(requests) < maxRequests {
            requests = append(requests, now)
            return true // Allow request
        }
        return false // Rate limit exceeded
    }
}

// Each API endpoint gets its own rate limiter
apiLimiter := createRateLimiter(100, time.Minute)
```

#### 6. **Retry Logic with Backoff** 🔄
```go
func createRetryFunc(maxRetries int, backoff time.Duration) func(func() error) error {
    return func(operation func() error) error {
        for i := 0; i < maxRetries; i++ {
            if err := operation(); err == nil {
                return nil
            }
            if i < maxRetries-1 {
                time.Sleep(backoff * time.Duration(i+1))
            }
        }
        return fmt.Errorf("operation failed after %d retries", maxRetries)
    }
}

// Different retry strategies for different operations
quickRetry := createRetryFunc(3, 100*time.Millisecond)
slowRetry := createRetryFunc(5, time.Second)
```

#### 7. **Plugin Architecture** 🔌
```go
func createValidator(rules []string) func(interface{}) error {
    return func(data interface{}) error {
        for _, rule := range rules {
            // Validate based on the specific rules this validator was created with
            if err := validateRule(data, rule); err != nil {
                return err
            }
        }
        return nil
    }
}

// Different validators for different data types
userValidator := createValidator([]string{"required", "email", "min_length:8"})
productValidator := createValidator([]string{"required", "positive_price"})
```

### Key Advantages Summary

1. **Encapsulation**: Private state that can't be accessed directly
2. **Persistence**: State survives between function calls
3. **Customization**: Create specialized functions from generic templates
4. **Clean APIs**: Hide implementation details while providing simple interfaces
5. **Memory Efficiency**: Share code while maintaining separate state
6. **Functional Programming**: Enable higher-order functions and functional patterns

**Best Practices:**
- Use closures when you need to maintain state between function calls
- Prefer closures over global variables for encapsulation
- Be mindful of memory usage - closures keep references to their enclosing scope
- Use closures to create specialized functions from generic templates
- Combine closures with interfaces for powerful abstraction patterns

Closures are particularly powerful in Go because they provide a clean way to manage state without requiring classes or objects, making your code more functional and often easier to reason about. 