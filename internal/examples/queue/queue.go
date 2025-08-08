package queue

import "fmt"


type Node[T any] struct {
	Value T
	Next *Node[T]
}

type Queuer[T any] interface {
	Enqueue(T)
	Dequeue() (T, error)
	Size() int32
	Peek() (T, error)
}


type Queue[T any] struct {
	Head *Node[T]
	Tail *Node[T]
	size int32
}


func NewQueue[T any]() Queuer[T]{
	return &Queue[T]{
		Head: nil,
		Tail: nil,
		size: 0,
	}
}

func (qs *Queue[T]) Enqueue(val T) {

	qs.size++;
	node := &Node[T]{Value: val}
	// if head doesn't exist, assume no tail either
	if qs.Head == nil {
		qs.Tail = node
		qs.Head = node
		return
	}

	// add to the tail
	// set pointer of current tail to our node
	qs.Tail.Next = node

	// set new tail of the current node
	qs.Tail = node

}

func (qs *Queue[T]) Dequeue() (T, error) {
	if qs.Head == nil {
		// queue is empty
		var zero T
		return zero, fmt.Errorf("Queue is empty")
	}

	// decrement size only after validating queue is not empty
	qs.size--
	curr := qs.Head
	qs.Head = qs.Head.Next
	
	if qs.Head == nil {
		qs.Tail = nil
	}

	return  curr.Value, nil

}

func (qs *Queue[T]) Size() int32 {
	return qs.size
}

func (qs *Queue[T]) Peek() (T, error) {
	if qs.Head == nil {
		var zero T
		return zero, fmt.Errorf("Queue is empty")
	}
	return qs.Head.Value, nil
}