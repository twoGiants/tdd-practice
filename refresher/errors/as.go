package errors

import (
	"errors"
	"fmt"
)

func AsExample(err error) {
	var myErr MyErr
	if errors.As(err, &myErr) {
		fmt.Println(myErr.Codes)
	}

	var coder interface {
		CodeVals() []int
	}
	if errors.As(err, &coder) {
		fmt.Println(coder.CodeVals())
	}
}
