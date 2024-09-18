# Mutations

## Insert and element

```go
m[key] = elem
```

## Get an element

```go
elem = m[key]
```

## Delete an element

```go
delete(m, key)
```

## Check if a key exists

```go
elem, ok := m[key]
```

If `key` is in `m`, then `ok` is `true`. If not, `ok` is `false`.

If `key` is not in the map, then `elem` is the zero value for the map's element type.

## Note on passing maps

Like slices, maps are also passed by reference into functions. This means that when a map is passed into a function we write, we *can* make changes to the original, we don't have a copy.
