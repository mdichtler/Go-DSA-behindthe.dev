package queue_todo

import (
	"testing"
)

func TestPeekEmpty(t *testing.T) {
	// should return an error
	queue := NewQueue[int8]()
	got, err := queue.Peek()
	if err == nil {
		t.Errorf("Empty queue Peek didn't return an error, got: %d", got)
	}
}

func TestDequeueEmpty(t *testing.T) {
	queue := NewQueue[int8]()
	got, err := queue.Dequeue()
	if err == nil {
		t.Errorf("Empty queue Deque didn't return an error, got: %d", got)
	}
}

func TestEnqueue(t *testing.T) {
	// initiate empty queue
	queue := NewQueue[int8]()
	queue.Enqueue(2)
	// test case
	got, err:= queue.Peek()
	if err != nil {
		t.Errorf("Failed to peek: %v", err)
	}
	// expected value
	want := 2

	if got != int8(want) {
		t.Errorf("Failed to enqueue, expected: %d, got: %d", want, got)
	}
} 

func TestDequeue(t *testing.T) {
	queue := NewQueue[int8]()
	queue.Enqueue(5)
	queue.Enqueue(3)
	got, err := queue.Dequeue()
	if err != nil {
		t.Errorf("Failed to dequeue, error: %v", err)
	}

	want := 5
	if got != int8(want)  {
		t.Errorf("Failed to dequeue, expected: %d, got: %d", want, got)
	}

	// lets also validate correctly that 3 is now at the front
	got, err = queue.Peek()
	want = 3
	if got != int8(want) {
		t.Errorf("Peek returned wrong value after dequeue, expected: %d, received: %d", want, got)
	}
}


func TestSize(t *testing.T) {
	queue := NewQueue[int8]()
	got := queue.Size()

	want := 0
	if got != int32(want) {
		t.Errorf("Size of an empty queue must be zero, received: %d", got)
	} 

	// lets add items
	queue.Enqueue(2)
	queue.Enqueue(4)
	queue.Enqueue(5)
	queue.Enqueue(2)

	got = queue.Size()

	want = 4

	if got != int32(want) {
		t.Errorf("Size not matching after enqueue, received: %d, expected: %d", got, want)
	}

	_, err := queue.Dequeue()
	if err != nil {
		t.Errorf("Failed to dequeue when testing size: %v, expected 3 more elements after dequeue", err)
	}
	_, err = queue.Dequeue()
	if err != nil {
		t.Errorf("Failed to dequeue when testing size: %v, expected 2 more elements after dequeue", err)
	}
	got = queue.Size()
	want = 2 
	if got != int32(want) {
		t.Errorf("Size not matching after dequeue, received: %d, expected: %d", got, want)
	}
}
