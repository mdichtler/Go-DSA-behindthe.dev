package stack

import "fmt"

// Node shape definition
type Node[T any] struct {
	Value T
	previous *Node[T]
}


// Interface providing list of functions
type Stacker[T any] interface {
	Peek() (T, error)
	Push(value T)
	Pop() (T, error)
	Size() int32
}


type Stack[T any] struct {
	head *Node[T]
	size int32
}


func (s *Stack[T]) Peek() (T, error) {
	if s.head == nil {
		var zero T
		return zero, fmt.Errorf("head is not set")
	}
	return s.head.Value, nil
}

func (s *Stack[T]) Push(value T) {

	s.size++;

	s.head = &Node[T]{
		Value: value,
		previous: s.head,
	}
}

func (s *Stack[T]) Pop() (T, error) {
	if s.head == nil {
		var zero T
		return zero, fmt.Errorf("you cannot pop from an empty stack")
	}

	s.size--;
	previous := s.head
	s.head = s.head.previous
	return previous.Value, nil

}

func (s *Stack[T]) Size() int32 {
	return s.size
}

func NewStack[T any]() Stacker[T] {
	return &Stack[T]{
		head: nil,
		size: 0,
	}
}