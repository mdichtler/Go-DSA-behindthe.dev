package queue

import "fmt"


type Node[T any] struct {
	Value T
	next *Node[T]
}

type Queuer[T any] interface {
	Enqueue(T)
	Dequeue() (T, error)
	Size() int32
	Peek() (T, error)
}


type Queue[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int32
}


func NewQueue[T any]() Queuer[T]{
	return &Queue[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (qs *Queue[T]) Enqueue(val T) {

	qs.size++;
	node := &Node[T]{Value: val}
	// if head doesn't exist, assume no tail either
	if qs.head == nil {
		qs.tail = node
		qs.head = node
		return
	}

	// add to the tail
	// set pointer of current tail to our node
	qs.tail.next = node

	// set new tail of the current node
	qs.tail = node

}

func (qs *Queue[T]) Dequeue() (T, error) {
	if qs.head == nil {
		// queue is empty
		var zero T
		return zero, fmt.Errorf("Queue is empty")
	}

	// decrement size only after validating queue is not empty
	qs.size--
	curr := qs.head
	qs.head = qs.head.next
	
	if qs.head == nil {
		qs.tail = nil
	}

	return  curr.Value, nil

}

func (qs *Queue[T]) Size() int32 {
	return qs.size
}

func (qs *Queue[T]) Peek() (T, error) {
	if qs.head == nil {
		var zero T
		return zero, fmt.Errorf("Queue is empty")
	}
	return qs.head.Value, nil
}