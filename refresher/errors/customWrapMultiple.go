package errors

import "fmt"

func CustomWrappedMultiError(e error) {
	switch err := e.(type) {
	case interface{ Unwrap() error }:
		// handle single error
		innerErr := err.Unwrap()
		// process innerErr
		fmt.Println(innerErr)
	case interface{ Unwrap() []error }:
		// handle multiple errors
		innerErrs := err.Unwrap()
		for _, innerErr := range innerErrs {
			// process innerErr
			fmt.Println(innerErr)
		}
	default:
		// handle no wrapped error
	}
}
