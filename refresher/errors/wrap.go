package errors

import (
	"errors"
	"fmt"
	"os"
)

func wrapFileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in wrapFileChecker: %w", err)
	}
	f.Close()
	return nil
}

func RunWrapFileChecker() {
	err := wrapFileChecker("not_here.txt")
	if err != nil {
		fmt.Println(err)
		// will print: in fileChecker: open not_here.txt: no such file or directory
		// just for illustration we unwrap
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Println(wrappedErr)
			// will print: open not_here.txt: no such file or directory
		}
	}
}
