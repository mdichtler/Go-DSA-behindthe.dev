package stack

import (

	"testing"
)

func TestPushToStack(t *testing.T) {
	stack := NewStack[int8]()
	
	stack.Push(2)

	got, err := stack.Peek()
	if err != nil {
		t.Errorf("Peek threw an unexpected error: %q", err)
	}
	want := 2
	if got != int8(want) {
		t.Errorf("Push failed: received %d, expected %d", got, want)
	}
}

func TestPop(t *testing.T) {
	stack := NewStack[int8]()

	stack.Push(3)
	stack.Push(5)
	// this element will be popped
	stack.Push(1)
	stack.Pop()

	got, err := stack.Peek()
	if err != nil {
		t.Errorf("Peek threw an unexpected error: %q", err)
	}
	want := 5
	if got != int8(want) {
		t.Errorf("Pop failed: received %d, expected %d", got, want)
	}
}

func TestPopEmpty(t *testing.T) {
	stack := NewStack[int8]()
	
	stack.Pop()
	// test the size is handled correctly
	got := stack.Size()
	want := 0

	if got != int32(want) {
		t.Errorf("Pop on empty failed: received size %d, expected size %d", got, want)
	}
}

func TestPeekEmpty(t *testing.T) {
	stack := NewStack[int8]()

	_, err := stack.Peek()
	if err == nil {
		t.Errorf("Failed to throw error when peeking an empty stack.")
	}

}

func TestSize(t *testing.T) {
	stack := NewStack[int8]()

	stack.Push(1)
	stack.Push(2)
	stack.Push(4)
	stack.Push(2)
	poppedValue, err := stack.Pop()
	if err != nil {
		t.Errorf("Pop threw unexpected error: %q", err)
	}
	if poppedValue != 2 {
		t.Errorf("Pop returned incorrect value: %q", poppedValue)
	}
	stack.Push(6)
	stack.Push(3)
	poppedValue, err = stack.Pop()
	if err != nil {
		t.Errorf("Pop threw unexpected error: %q", err)
	}
	if poppedValue != 3 {
		t.Errorf("Pop returned incorrect value: %q", poppedValue)
	}
	
	got := stack.Size()
	want := 4

	// have to use int32 as the size is defined stack level
	if got != int32(want) {
		t.Errorf("Size validation failed: received %d, expected %d", got, want)
	}
}