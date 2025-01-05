package errors

import (
	"fmt"
	"slices"
)

type MyErr struct {
	Codes []int
}

func (me MyErr) Error() string {
	return fmt.Sprintf("codes: %v", me.Codes)
}

func (me MyErr) Is(target error) bool {
	if me2, ok := target.(MyErr); ok {
		return slices.Equal(me.Codes, me2.Codes)
	}
	return false
}
