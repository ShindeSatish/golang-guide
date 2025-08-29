# Cheatsheet: Common Go Commands

A quick reference for common commands used in Go development.

## `go run`
Compiles and runs one or more Go source files.
```bash
go run main.go
```

## `go build`
Compiles a package and its dependencies, but doesn't run it.
```bash
# Build an executable for the current directory
go build

# Build and name the output
go build -o myapp

# Cross-compile for a different OS/architecture
GOOS=linux GOARCH=amd64 go build -o myapp-linux .
```

## `go test`
Runs tests and benchmarks.
```bash
# Run all tests in the current directory and subdirectories
go test ./...

# Run with verbose output
go test -v

# Run tests matching a regex
go test -run "TestMyFunction"

# Run benchmarks
go test -bench .

# Get a code coverage profile
go test -coverprofile=c.out
go tool cover -html=c.out
```

## `go fmt`
Formats Go source code according to the standard Go style. Run this often!
```bash
go fmt ./...
```

## `go vet`
A static analysis tool that reports suspicious constructs in Go code.
```bash
go vet ./...
```

## `go mod`
The command for managing dependencies (Go modules).

### `go mod init <module-path>`
Initializes a new module in the current directory.
```bash
go mod init github.com/my-org/my-repo
```

### `go mod tidy`
Ensures `go.mod` matches the source code. Adds missing dependencies and removes unused ones.
```bash
go mod tidy
```

### `go get <path@version>`
Adds or updates a dependency.
```bash
# Get the latest version
go get github.com/gin-gonic/gin

# Get a specific version
go get github.com/gin-gonic/gin@v1.7.4

# Update all dependencies
go get -u ./...
```

### `go mod why <path>`
Explains why a specific package is in your module's dependency graph.
```bash
go mod why github.com/gin-gonic/gin
```

## `go doc`
Shows documentation for a package or symbol.
```bash
# Show docs for the http package
go doc net/http

# Show docs for the Get function in net/http
go doc net/http.Get
```

## `go install`
Compiles and installs a package or command. The executable is placed in your `$GOPATH/bin` or `$HOME/go/bin` directory.
```bash
# Install a tool
go install golang.org/x/tools/cmd/goimports@latest
``` 