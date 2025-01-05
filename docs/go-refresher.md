# Go Refresher

## Tooling

- Coverage with browser output: [link](https://go.dev/blog/cover),

```bash
go tool cover -html=coverage.out
```

## Interfaces

- A more typical pattern in Go is to return concrete structs, and use interfaces for parameter types received by functions.
- Interfaces are typically used for return types only where different implementation structs might be returned.

## Errors

Learning Go, 2nd Edition, Ch [9. Errors](https://learning.oreilly.com/library/view/learning-go-2nd/9781098139285/ch09.html#id109)

### Basics

- use `error` interface type as function return types
- use `nil` for the error return if none happened
- never return custom errors

### Sentinel Errors

- use Sentinel errors sparingly and write them starting `ErrSomeName`
- use `errors.As` instead of type assertion
- used in std packages and libraries, e.g. `os.ErrNotExist`

### Wrapping Errors

- *error tree* is a series of wrapped errors
- wrap errors using `fmt.Errorf` and `%w`, e.g. `fmt.Errorf("in fileChecker: %w", err)`
    - see `refresher/errors/wrap.go`
- instead of `errors.Unwrap` use `errors.Is` and `errors.As`
- implement `Unwrap` to wrap errors with your custom error
- if you want to create a new error that contains the message from another error, but don't want to wrap it use `%v`, e.g.

```go
err := internalFunction()
if err != nil {
    return fmt.Errorf("internal failure: %v", err)
}
```

### Wrapping Multiple Errors

- for e.g. validation when you need to return an error for each field use `errors.Join`, see `refresher/errors/wrapMultiple.go`
- another way to merge errors is using multiple `%w`, see `refresher/errors/wrapMultiple.go`

### Is and As

- use `errors.Is` to check for sentinel errors in the error tree, it uses `==` internally, see `refresher/errors/is.go`
- for noncomparable types implement the `Is` method, see `refresher/errors/isImpl.go`
- for errors that aren't identical instances which you want to pattern match, see `refresher/errors/isPatternMatch.go`
- `error.As` allows to check if the error matches a specific type which you have to declare with `var` and then use the pointer of, see `refresher/errors/as.go`
- use `errors.Is` when you look for an *instance* or specific *values* and `errors.As` when you are looking for a specific *type*

***

###### Reference

- [Learning Go, 2nd Edition](https://learning.oreilly.com/library/view/learning-go-2nd/9781098139285)
- [Go in Action](https://learning.oreilly.com/library/view/go-in-action/9781617291784)
- [Go in Practice](https://learning.oreilly.com/library/view/go-in-practice/9781633430075/)
