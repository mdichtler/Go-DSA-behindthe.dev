package stack

import "fmt"

// Node shape definition
type Node[T any] struct {
	Value T
	Previous *Node[T]
}


// Interface providing list of functions
type Stacker[T any] interface {
	Peek() (T, error)
	Push(value T)
	Pop() (T, error)
	Size() int32
}


type Stack[T any] struct {
	Head *Node[T]
	size int32
}


func (s *Stack[T]) Peek() (T, error) {
	if s.Head == nil {
		var zero T
		return zero, fmt.Errorf("head is not set")
	}
	return s.Head.Value, nil
}

func (s *Stack[T]) Push(value T) {

	s.size++;

	s.Head = &Node[T]{
		Value: value,
		Previous: s.Head,
	}
}

func (s *Stack[T]) Pop() (T, error) {
	if s.Head == nil {
		var zero T
		return zero, fmt.Errorf("you cannot pop from an empty stack")
	}

	s.size--;
	previous := s.Head
	s.Head = s.Head.Previous
	return previous.Value, nil

}

func (s *Stack[T]) Size() int32 {
	return s.size
}

func NewStack[T any]() Stacker[T] {
	return &Stack[T]{
		Head: nil,
		size: 0,
	}
}