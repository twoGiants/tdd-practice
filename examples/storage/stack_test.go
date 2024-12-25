package storage_test

import (
	"tdd/examples/storage"
	"testing"
)

func Test_CanCreateStack(t *testing.T) {
	stack := storage.NewStack()

	if stack == nil {
		t.Fatal("expected a stack but got", stack)
	}

	if !stack.IsEmpty() {
		t.Fatal("expected empty check to be 'true' but got", stack.IsEmpty())
	}
}

func Test_afterOnePush_isNotEmpty(t *testing.T) {
	stack := storage.NewStack()

	stack.Push(0)

	if stack.IsEmpty() {
		t.Fatal("stack should not be empty but it is")
	}

	if stack.Size() != 1 {
		t.Fatal("stack size should be 1 but it is", stack.Size())
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
		t.Fatal("stack size should be 0 but it is", stack.Size())
	}
}

func Test_afterTwoPushes_sizeIsTwo(t *testing.T) {
	stack := storage.NewStack()

	stack.Push(0)
	stack.Push(0)

	if stack.Size() != 2 {
		t.Fatal("stack size should be 2 but is", stack.Size())
	}
}

func Test_poppingEmptyStack_panicsWithUnderflow(t *testing.T) {
	defer recoverWithMessage("underflow", t)

	stack := storage.NewStack()

	stack.Pop()

	t.Fatal("should have panicked but didn't")
}

func Test_afterPushingX_willPopX(t *testing.T) {
	stack := storage.NewStack()

	stack.Push(99)

	popped := stack.Pop()
	if popped != 99 {
		t.Fatal("should have popped 99 but popped", popped)
	}

	stack.Push(88)
	popped = stack.Pop()
	if popped != 88 {
		t.Fatal("should have popped 88 but popped", popped)
	}
}

func Test_afterPushingXandY_willPopYthenX(t *testing.T) {
	stack := storage.NewStack()

	stack.Push(99)
	stack.Push(88)

	popped := stack.Pop()
	if popped != 88 {
		t.Fatal("first pop should've been 88 but was", popped)
	}

	popped = stack.Pop()
	if popped != 99 {
		t.Fatal("last pop should've been 99 but was", popped)
	}
}

func Test_pushingFullStack_panicsWithOverflow(t *testing.T) {
	defer recoverWithMessage("overflow", t)

	stack := storage.NewStack()

	stack.Push(1)
	stack.Push(1)
	stack.Push(1)

	t.Fatal("should have panicked but didn't")
}

func recoverWithMessage(msg string, t *testing.T) {
	if v := recover(); v.(string) != msg {
		t.Fatalf("should panic with an '%s' message but did with %v", msg, v)
	}
}
