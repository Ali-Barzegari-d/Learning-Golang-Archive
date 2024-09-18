# Conditionals

`if` statements in Go don't use parentheses around the condition:

```go
if height > 4 {
    fmt.Println("You are tall enough!")
}
```

`else if` and `else` are supported as you would expect:

```go
if height > 6 {
    fmt.Println("You are super tall!")
} else if height > 4 {
    fmt.Println("You are tall enough!")
} else {
    fmt.Println("You are not tall enough!")
}
```


### Tips

Here are some of the comparison operators in Go:

* `==` equal to
* `!=` not equal to
* `<` less than
* `>` greater than
* `<=` less than or equal to
* `>=` greater than or equal to
