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

- for e.g. validation when you need to return an error for each field use `errors.Join`

```go
type Person struct {
    FirstName string
    LastName  string
    Age       int
}

func ValidatePerson(p Person) error {
    var errs []error
    if len(p.FirstName) == 0 {
        errs = append(errs, errors.New("field FirstName cannot be empty"))
    }
    if len(p.LastName) == 0 {
        errs = append(errs, errors.New("field LastName cannot be empty"))
    }
    if p.Age < 0 {
        errs = append(errs, errors.New("field Age cannot be negative"))
    }
    if len(errs) > 0 {
        return errors.Join(errs...)
    }
    return nil
}
```

- another way to merge errors is using multiple `%w`

```go
err1 := errors.New("first error")
err2 := errors.New("second error")
err3 := errors.New("third error")
err := fmt.Errorf("first: %w, second: %w, third: %w", err1, err2, err3)
```

***

###### Reference

- [Learning Go, 2nd Edition](https://learning.oreilly.com/library/view/learning-go-2nd/9781098139285)
- [Go in Action](https://learning.oreilly.com/library/view/go-in-action/9781617291784)
- [Go in Practice](https://learning.oreilly.com/library/view/go-in-practice/9781633430075/)
