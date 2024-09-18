# Computed Constants

Constants must be known at compile time. More often than not they will be declared with a static value:

```go
const myInt = 15
```

However, constants *can be computed* so long as the computation can happen at *compile time*.

For example, this is valid:

```go
const firstName = "Lane"
const lastName = "Wagner"
const fullName = firstName + " " + lastName
```

That said, you *cannot* declare a constant that can only be computed at run-time.


