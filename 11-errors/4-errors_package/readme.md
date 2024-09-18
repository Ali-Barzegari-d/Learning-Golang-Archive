# The Errors Package

The Go standard library provides an "errors" package that makes it easy to deal with errors.

Read the godoc for the [errors.New()](https://pkg.go.dev/errors#New) function, but here's a simple example:

```go
var err error = errors.New("something went wrong")
```

