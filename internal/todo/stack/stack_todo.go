package stack_todo

type Node[T any] struct {
	Value T
	previous *Node[T]
}

type Stack[T any] interface {
	Peek() (T, error)
	Push(value T)
	Pop() (T, error)
	Size() int32
}


type NodeStack[T any] struct {
	head *Node[T]
	size int32
}


func (s *NodeStack[T]) Peek() (T, error) {
	
}

func (s *NodeStack[T]) Push(value T) {

}

func (s *NodeStack[T]) Pop() (T, error) {
	
}

func (s *NodeStack[T]) Size() int32 {

}

func NewStack[T any]() Stack[T] {

}