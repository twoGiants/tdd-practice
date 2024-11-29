package storage_test

import (
	"tdd/examples/storage"
	"testing"
)

func Test_CanCreateStack(t *testing.T) {
	stack := storage.NewStack()

	if stack == nil {
		t.Fatalf("expected a stack but got %v", stack)
	}

	if !stack.IsEmpty() {
		t.Fatalf("expected empty check to be 'true' but got '%t'", stack.IsEmpty())
	}
}

func Test_afterOnePush_isNotEmpty(t *testing.T) {
	stack := storage.NewStack()

	stack.Push(0)

	if stack.IsEmpty() {
		t.Fatal("stack should not be empty but it is")
	}

	if stack.Size() != 1 {
		t.Fatalf("stack size should be 1 but it is: %d", stack.Size())
	}
}

func Test_afterOnePushAndPop_isEmpty(t *testing.T) {
	stack := storage.NewStack()

	stack.Push(0)
	stack.Pop()

	if !stack.IsEmpty() {
		t.Fatal("stack should be empty but it isn't")
	}

	if stack.Size() != 0 {
		t.Fatalf("stack size should be 0 but it is: %d", stack.Size())
	}
}

func Test_afterTwoPushes_sizeIsTwo(t *testing.T) {
	stack := storage.NewStack()

	stack.Push(0)
	stack.Push(0)

	if stack.Size() != 2 {
		t.Fatalf("stack size should be 2 but is: %d", stack.Size())
	}
}

func Test_poppingEmptyStack_panicsWithUnderflow(t *testing.T) {
	defer func() {
		if v := recover(); v.(string) != "underflow" {
			t.Fatal("should panic with an 'underflow' message but did with: ", v)
		}
	}()

	stack := storage.NewStack()

	stack.Pop()

	t.Fatal("should have panicked but didn't")
}

func Test_afterPushingX_willPopX(t *testing.T) {
	stack := storage.NewStack()

	stack.Push(99)

	popped := stack.Pop()
	if popped != 99 {
		t.Fatal("should have popped 99 but popped: ", popped)
	}

	stack.Push(88)
	popped = stack.Pop()
	if popped != 88 {
		t.Fatal("should have popped 88 but popped: ", popped)
	}
}

func Test_afterPushingXandY_willPopYthenX(t *testing.T) {
	stack := storage.NewStack()

	stack.Push(99)
	stack.Push(88)

	popped := stack.Pop()
	if popped != 88 {
		t.Fatal("first pop should've been 88 but was: ", popped)
	}

	popped = stack.Pop()
	if popped != 99 {
		t.Fatal("last pop should've been 99 but was: ", popped)
	}
}
