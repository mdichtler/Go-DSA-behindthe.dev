package linkedlist

import (
	"reflect"
	"testing"
)


func TestPrependSingly(t *testing.T) {
	sll := NewSingly[int8]()

	sll.Prepend(1)
	const prevValue int8 = 2
	sll.Prepend(prevValue)

	const targetValue int8 = 3
	sll.Prepend(int8(targetValue))



	got := sll.ToSlice()
	want := []int8{3, 2, 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Incorrect items or order in list: got %v, want %v", got, want)
	}

	if sll.Length() != len(want) {
		t.Errorf("Incorrect length: got %d, want %d", sll.Length(), len(want))
	}

}

func TestAppendSingly(t *testing.T) {
	sll := NewSingly[int8]()

	sll.Append(1)
	const nextValue int8 = 2
	sll.Append(nextValue)

	const targetValue int8 = 3
	sll.Append(int8(targetValue))


	got := sll.ToSlice()
	want := []int8{1, 2, 3}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Incorrect items or order in list: got %v, want %v", got, want)
	}

	if sll.Length() != len(want) {
		t.Errorf("Incorrect length: got %d, want %d", sll.Length(), len(want))
	}

}


func TestFindSingly(t *testing.T) {
	sll := NewSingly[int8]()

	sll.Append(1)
	sll.Append(2)

	var want int8 = 2

	node, err := sll.Find(want)
	if err != nil {
		t.Errorf("Failed to find: want %d error: %v", want, err )
	}
	if node.Value != want {
		t.Errorf("Found incorrect node: got %d, want: %d", node.Value, want)
	}

	want = 4
	_, err = sll.Find(want)
	if err == nil {
		t.Errorf("Failed to return an error for value not in the list: want %d", want)
	}
}

func TestDeleteSingly(t *testing.T) {
	sll := NewSingly[int8]()

	var valToDelete int8 = 1
	err := sll.Delete(valToDelete)
	if err == nil {
		t.Errorf("Failed to return an error when deleting non existing value %d", valToDelete)
	}

	sll.Append(1)

	err = sll.Delete(valToDelete)
	if err != nil {
		t.Errorf("Failed to delete value expected in the list: %d, error: %v", valToDelete, err)
	}
	sll.Append(2)

	// reach end of not empty list without finding
	err = sll.Delete(10) 
	if err == nil {
		t.Errorf("Failed to handle non existing value in non empty list")
	}


	sll.Append(3)
	valToDelete = 9
	sll.Append(valToDelete)

	sll.Append(7)
	err = sll.Delete(valToDelete) 
	if err != nil {
		t.Errorf("Failed to delete value expected in the list: %d, error: %v", valToDelete, err)
	}

	if sll.Length() != 3 {
		t.Errorf("Length mismatch after deletion, expected: %d, got: %d", 3, sll.Length())
	}
	// test tail deletion
	err = sll.Delete(7)
	
	if err != nil {
		t.Errorf("Failed at tail deletion: %v", err)
	}

}