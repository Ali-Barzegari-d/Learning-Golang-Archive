# Remote Packages


Let's learn how to use an open-source package that's available online.

## A note on how you should publish modules

Be aware that using the "replace" keyword like we did in the last assignment *isn't advised*, but can be useful to get up and running quickly. The *proper* way to create and depend on modules is to publish them to a remote repository. When you do that, the "replace keyword can be dropped from the `go.mod`:

### Bad

This works for local-only development

```go
module github.com/wagslane/hellogo

go 1.20

replace github.com/wagslane/mystrings v0.0.0 => ../mystrings

require (
	github.com/wagslane/mystrings v0.0.0
)
```

### Good

This works if we publish our modules to a remote location like Github as we should.

```go
module github.com/wagslane/hellogo

go 1.20

require (
	github.com/wagslane/mystrings v0.0.0
)
```

