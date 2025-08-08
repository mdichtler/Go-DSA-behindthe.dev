package linkedlist

import (
	"fmt"
)

type SinglyNode[T comparable] struct {
	Value T
	next *SinglyNode[T]
}


type Singler[T comparable] interface {
	Append(value T)
	Prepend(value T)
	Find(value T) (*SinglyNode[T], error)
	Delete(value T) error
	Length() int
	ToSlice() []T
}

type Singly[T comparable] struct {
	head *SinglyNode[T]
	tail *SinglyNode[T]
	length int
}

func NewSingly[T comparable]() *Singly[T] {
	return &Singly[T]{
		head: nil,
		tail: nil,
		length: 0,
	}
}

func (sl *Singly[T]) Append(value T) {
	sl.length++;
	node := &SinglyNode[T]{Value: value}

	// check if tail exists
	if sl.tail == nil {
		sl.head = node
		sl.tail = node
		return
	}

	// add node to the end we update tail
	sl.tail.next = node
	sl.tail = node
	
}

func (sl *Singly[T]) Prepend(value T) {
	sl.length++;
	node := &SinglyNode[T]{Value: value}

	if sl.head == nil {
		sl.head = node
		sl.tail = node
		return
	}

	node.next = sl.head
	sl.head = node
}

func (sl *Singly[T]) Find(value T) (*SinglyNode[T], error) {
	
	curr := sl.head

	for curr != nil {
		if curr.Value == value {
			return curr, nil
		}
		curr = curr.next
	}

	return nil, fmt.Errorf("value not found")
}

func (sl *Singly[T]) Delete(value T) error {
	// if there is no head, throw an error
	if sl.head == nil {
		return fmt.Errorf("value not found")
	}
	// if value is head, set head to next, if new head is nil, we also set tail to nil
	if sl.head.Value == value {
		sl.length--;
		sl.head = sl.head.next
		if sl.head == nil {
			sl.tail = nil
			
		}
		return nil
	}
	// lets find the value where value of next is not the value
	prev := sl.head
	for prev.next != nil && prev.next.Value != value {
		prev = prev.next
	}
	// reached end without finding
	if prev.next == nil {
		return fmt.Errorf("value not found")
	}

	sl.length--
	// if value is the tail (we checked head before so not head at the same time), we find the previous to tail, and  
	if prev.next == sl.tail {
		sl.tail = prev
	}

	prev.next = prev.next.next
	return nil

}

func (sl *Singly[T]) Length() int {
	return sl.length
}

func (sl *Singly[T]) ToSlice() []T {
	values := make([]T, 0, sl.length)
	curr := sl.head
	for curr != nil {
		values = append(values, curr.Value)
		curr = curr.next
	}
	return values
}


