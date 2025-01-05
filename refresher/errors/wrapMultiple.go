package errors

import (
	"errors"
	"fmt"
)

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

func WrapMultipleErrors() error {
	err1 := errors.New("first error")
	err2 := errors.New("second error")
	err3 := errors.New("third error")
	err := fmt.Errorf("first: %w, second: %w, third: %w", err1, err2, err3)
	return err
}
