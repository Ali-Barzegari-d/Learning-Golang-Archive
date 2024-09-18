# Append

The built-in append function is used to dynamically add elements to a slice:

```go
func append(slice []Type, elems ...Type) []Type
```

If the underlying array is not large enough, `append()` will create a new underlying array and point the slice to it.

Notice that `append()` is variadic, the following are all valid:

```go
slice = append(slice, oneThing)
slice = append(slice, firstThing, secondThing)
slice = append(slice, anotherSlice...)
```


### Example

Input in `day, cost` format:

```go
[]cost{
    {0, 4.0},
    {1, 2.1},
    {1, 3.1},
    {5, 2.5},
}
```

I would expect this result:

```go
[]float64{
    4.0,
    5.2,
    0.0,
    0.0,
    0.0,
    2.5,
}
```
