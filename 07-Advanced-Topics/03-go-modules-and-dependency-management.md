# Advanced Topics: Go Modules & Dependency Management

Go modules are the standard for dependency management in Go. They were introduced in Go 1.11 and are the successor to older systems like `GOPATH` and `dep`.

## Key Concepts

- **Module:** A collection of related Go packages that are versioned together as a single unit. Each module is defined by a `go.mod` file at its root.
- **`go.mod` file:** The manifest for a module. It defines:
    - The module's path (its unique identifier).
    - The version of Go it was built with.
    - The list of direct dependencies (other modules) required by the module.
- **`go.sum` file:** A checksum file that contains the expected cryptographic checksums of the content of specific module versions. This ensures that the dependencies you download are bit-for-bit identical to the ones the author intended, protecting against tampering. It is automatically generated and should be checked into version control.
- **Semantic Versioning (SemVer):** Go modules use the SemVer standard (`vMajor.Minor.Patch`) for versioning.
    - `v1.2.3`
    - A change in the major version (e.g., `v2.0.0`) indicates a breaking API change.

## Common Commands

### `go mod init <module-path>`
Initializes a new module in the current directory. The module path is typically the URL where your repository is hosted (e.g., `github.com/my-org/my-project`).
```bash
go mod init github.com/satish/my-app
```

### `go get <package-path>`
Adds a new dependency or updates an existing one.
```bash
# Add the latest version of a package
go get github.com/gin-gonic/gin

# Get a specific version
go get github.com/gin-gonic/gin@v1.7.4
```

### `go mod tidy`
This is a crucial command. It ensures that your `go.mod` file matches the source code in your module.
- It removes dependencies that are no longer used.
- It adds any dependencies that are being used but are missing from `go.mod`.
You should run `go mod tidy` before committing changes.

### `go build`, `go test`
These commands, along with others like `go run`, will automatically download the necessary dependencies as specified in `go.mod` if they are not already in the local module cache.

## Module Proxy

By default, the `go` command downloads modules directly from their source repositories (e.g., GitHub). However, it can be configured to use a module proxy.

- **`GOPROXY` Environment Variable:** Specifies the URL of the proxy. The default is `https://proxy.golang.org,direct`.
- **Benefits:**
    - **Reliability:** The proxy provides a highly-available cache, so you can still download dependencies even if the source repository (e.g., GitHub) is down.
    - **Immutability:** Once a version of a module is in the proxy, it can't be changed. This protects against authors force-pushing changes to a version tag, which could break builds.
    - **Performance:** Proxies can be faster to download from.

## Private Repositories

To use private repositories as dependencies, you need to configure Go to bypass the public proxy for them.
- **`GOPRIVATE` Environment Variable:** A comma-separated list of glob patterns for repositories that should be accessed directly.
```bash
# Example for a private GitHub repository
export GOPRIVATE=github.com/my-private-org/*
```

## Lead Engineer Perspective

- **Enforce Module Usage:** All new projects must be Go modules.
- **Version Pinning:** `go.mod` effectively "pins" your dependency versions. This ensures reproducible builds. A `go build` today will produce the same result as a `go build` a year from now.
- **CI/CD:** Your continuous integration pipeline should run `go mod tidy` to check that the `go.mod` file is up-to-date. It should also run tests and build the application, which implicitly verifies that all dependencies are downloadable and correct.
- **Managing Major Versions:** When a dependency releases a new major version (e.g., `v2`), it is treated as a completely separate module. You can even import `v1` and `v2` of the same package in the same program if necessary. Upgrading a major version requires a code change (updating the import path) and should be planned carefully. 