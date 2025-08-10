package hashmap

import "fmt"

type Node[T any] struct {
	Key   string
	Value T
	next  *Node[T]
}

type HashMap[T any] struct {
	buckets  []*Node[T]
	capacity int
}

func NewHashMap[T any](capacity int) *HashMap[T] {
	return &HashMap[T]{
		buckets:  make([]*Node[T], capacity),
		capacity: capacity,
	}
}

func (h *HashMap[T]) hash(key string) int {
	total := 0
	for _, char := range key {
		total += int(char)
	}
	return total % h.capacity
}

func (h *HashMap[T]) Set(key string, value T) {
	idx := h.hash(key)

	node := &Node[T]{
		Value: value,
		Key:   key,
	}
	if h.buckets[idx] == nil {
		h.buckets[idx] = node
		return
	} else {
		// collision
		curr := h.buckets[idx]
		for curr != nil {
			if curr.Key == key {
				// keys matching handle update
				curr.Value = value
				return
			}
			if curr.next == nil {
				curr.next = node
				return
			}

			curr = curr.next

		}

	}
}

func (h *HashMap[T]) Get(key string) (T, error) {
	bucketIdx := h.hash(key)

	curr := h.buckets[bucketIdx]
	for curr != nil {
		// found
		if curr.Key == key {
			return curr.Value, nil
		}

		curr = curr.next

	}
	var zero T
	return zero, fmt.Errorf("value not found")

}

func (h *HashMap[T]) Remove(key string) (T, error) {
	bucketIdx := h.hash(key)
	var prev *Node[T]
	curr := h.buckets[bucketIdx]

	for curr != nil {

		if key == curr.Key {
			if prev == nil {
				h.buckets[bucketIdx] = curr.next
			} else {
				prev.next = curr.next
			}
			return curr.Value, nil
		}

		prev = curr
		curr = curr.next
	}

	var zero T
	return zero, fmt.Errorf("key not found")
}
