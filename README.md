# Minimal Go errors

This package is a drop-in replacement for the standard library's `errors` package.

It then the functionalities from the `github.com/pkg/errors` package that are missing from the standard library. The difference from this package to `github.com/pkg/errors` is that this package does not automatically add a stack trace to the error, but the user can add one by calling `errors.WithStack()` and passing in another error.

Creating the stacktrace is a costly operation, so it is not done by default. This package is meant to be used in situations where the stacktrace is not needed, or where the stacktrace is added manually.

All the functions in this package that are the same as the standard library's `errors` package are actually just `var`s that are set to the same value as the standard library's `errors` package. This means that this package can be used as a drop-in replacement for the standard library's `errors` package and has minimal implementation and stable API.

## Usage

```go
package main

import "github.com/TheDSCPL/minimal-errors"

func main() {
    err := errors.New("this is an error")
    err = errors.WithStack(err)
}
```
