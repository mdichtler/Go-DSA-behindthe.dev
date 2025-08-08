package queue_todo

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

}

func (qs *Queue[T]) Dequeue() (T, error) {

}

func (qs *Queue[T]) Size() int32 {

}

func (qs *Queue[T]) Peek() (T, error) {

}