# Advanced Topics: Testing and Benchmarking

Go has a built-in testing framework provided by the `testing` package. It's lightweight but powerful, and writing tests is a core part of idiomatic Go development.

## Basic Testing

- Test files must end with `_test.go` (e.g., `main_test.go`).
- Test functions must start with `Test` and take a single argument: `t *testing.T`.
- The test file should be in the same package as the code it is testing.

**Example:**
File `mymath/mymath.go`:
```go
package mymath

func Add(a, b int) int {
    return a + b
}
```
File `mymath/mymath_test.go`:
```go
package mymath

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}
```

**Running Tests:**
```bash
# Run all tests in the current directory and subdirectories
go test ./...

# Run with verbose output
go test -v

# Run tests that match a specific pattern
go test -v -run TestAdd
```
The `t.Errorf` function reports a failure but allows the test function to continue. `t.Fatalf` reports a failure and stops the current test function immediately.

## Table-Driven Tests

Table-driven tests are a common and effective pattern in Go for testing multiple cases. You define a slice of test cases (the "table") and iterate over it.

```go
func TestAddTable(t *testing.T) {
    var tests = []struct {
        a, b     int
        expected int
    }{
        {1, 2, 3},
        {0, 0, 0},
        {-1, 1, 0},
        {10, -5, 5},
    }

    for _, tt := range tests {
        testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {
            ans := Add(tt.a, tt.b)
            if ans != tt.expected {
                t.Errorf("got %d, want %d", ans, tt.expected)
            }
        })
    }
}
```
Using `t.Run` creates sub-tests, which gives you more granular control and better output when a specific case fails.

## Mocks and Interfaces

To test a piece of code in isolation, you often need to "mock" its dependencies (like a database or a network service). This is where interfaces shine. If your code depends on an interface, you can provide a simple, fake implementation of that interface in your test.

## Benchmarking

Go's testing package also includes support for benchmarking.
- Benchmark functions must start with `Benchmark` and take `b *testing.B` as an argument.
- The code to be benchmarked should be run inside a `for` loop that runs `b.N` times.

**Example:**
```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(100, 200)
    }
}
```

**Running Benchmarks:**
```bash
go test -bench=.
```
The framework will automatically adjust `b.N` until the benchmark runs for a reasonable amount of time, and then report the time per operation.

## Code Coverage

You can generate a code coverage report to see which parts of your code are exercised by your tests.
```bash
# Generate a coverage profile
go test -coverprofile=coverage.out

# View the report in your browser
go tool cover -html=coverage.out
```

## Lead Engineer Perspective

- **Mandate Tests:** All new features and bug fixes must be accompanied by tests.
- **Promote Table-Driven Tests:** Encourage this pattern as it makes tests clearer and easier to extend.
- **Coverage is a Guide, Not a Goal:** A high coverage number is good, but it doesn't guarantee your tests are meaningful. The goal is to test the *behavior* of your code, not just execute lines.
- **CI Integration:** Your continuous integration system must run all tests on every commit. Failed tests should block a merge.
- **Benchmark Critical Code:** For performance-sensitive parts of the system, benchmarks should be written to prevent regressions. 