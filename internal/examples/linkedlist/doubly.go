package linkedlist

import "fmt"

type DoublyNode[T comparable] struct {
	Value    T
	next     *DoublyNode[T]
	previous *DoublyNode[T]
}

type Doubler[T comparable] interface {
	Append(value T)
	Prepend(value T)
	Find(value T) (*DoublyNode[T], error)
	Delete(value T) error
	Length() int
	ToSlice() []T
}

type Doubly[T comparable] struct {
	head   *DoublyNode[T]
	tail   *DoublyNode[T]
	length int
}

func NewDoubly[T comparable]() *Doubly[T] {
	return &Doubly[T]{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (dl *Doubly[T]) Append(value T) {
	node := &DoublyNode[T]{Value: value}
	dl.length++
	if dl.head == nil {
		dl.head = node
		dl.tail = node
		return
	}

	dl.tail.next = node
	node.previous = dl.tail
	dl.tail = node
}

func (dl *Doubly[T]) Prepend(value T) {
	node := &DoublyNode[T]{Value: value}
	dl.length++

	if dl.head == nil {
		dl.head = node
		dl.tail = node
		return
	}

	dl.head.previous = node
	node.next = dl.head
	dl.head = node
}

func (dl *Doubly[T]) Find(value T) (*DoublyNode[T], error) {

	curr := dl.head
	for curr != nil {
		// if found return
		if curr.Value == value {
			return curr, nil
		}

		curr = curr.next
	}

	return nil, fmt.Errorf("value not found")
}

func (dl *Doubly[T]) Delete(value T) error {
	if dl.head == nil {
		return fmt.Errorf("value not found")
	}

	
	// use find to search the node
	nodeToDelete, err := dl.Find(value)
	if err != nil {
		return err
	}

	if nodeToDelete.previous == nil {
		dl.head = nodeToDelete.next
	} else {
		nodeToDelete.previous.next = nodeToDelete.next
	}

	if nodeToDelete.next == nil {
		dl.tail = nodeToDelete.previous

	} else {
		nodeToDelete.next.previous = nodeToDelete.previous
	}
	dl.length--
	return nil
}

func (dl *Doubly[T]) Length() int {
	return dl.length
}

func (dl *Doubly[T]) ToSlice() []T {
	slice := make([]T, 0, dl.length)
	curr := dl.head
	for curr != nil {
		slice = append(slice, curr.Value)
		curr = curr.next
	}
	return slice
}
