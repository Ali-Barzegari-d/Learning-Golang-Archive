# Loops in Go

The basic loop in Go is written in standard C-like syntax:

```go
for INITIAL; CONDITION; AFTER{
  // do something
}
```

`INITIAL` is run once at the beginning of the loop and can create
variables within the scope of the loop.

`CONDITION` is checked before each iteration. If the condition doesn't pass
then the loop breaks.

`AFTER` is run after each iteration.

For example:

```go
for i := 0; i < 10; i++ {
  fmt.Println(i)
}
// Prints 0 through 9
```

