package linkedlist

import (
	"reflect"
	"testing"
)

func TestDoublyLinkedList(t *testing.T) {
	t.Run("Append and prepend should build the list correctly", func(t *testing.T) {
		// initialize
		dll := NewDoubly[uint8]()
		// append
		dll.Append(4)
		dll.Append(5)
		dll.Append(6)

		// prepend
		dll.Prepend(1)
		dll.Prepend(2)
		dll.Prepend(3)
		// outcome
		want := []uint8{3, 2, 1, 4, 5, 6}
		got := dll.ToSlice()

		if reflect.DeepEqual(want, got) == false {
			t.Errorf("Failed deep equality check, expected: %v, received: %v", want, got)
		}
	})

	t.Run("Find should locate values correctly", func(t *testing.T) {
		dll := NewDoubly[uint8]()

		dll.Append(4)
		dll.Append(5)
		dll.Append(6)

		var want uint8 = 5
		got, err := dll.Find(want)
		if err != nil {
			t.Errorf("Failed to find value that should exist: searched: %d, error: %v", want, err)
		}
		if got.Value != want {
			t.Errorf("Failed to find correct value: searched: %d, received: %d", want, got.Value)
		}

		want = 0
		_, err = dll.Find(want)
		if err == nil {
			t.Errorf("Failed to return error for non existing value: %d", want)
		}

	})

	t.Run("Delete should delete values correctly and handle edge cases", func(t *testing.T) {

		testCases := []struct {
			name        string
			initial     []uint8
			toDelete    uint8
			expected    []uint8
			expectedErr bool
		}{
			{
				name:        "delete head",
				initial:     []uint8{1, 2, 3},
				toDelete:    1,
				expected:    []uint8{2, 3},
				expectedErr: false,
			},
			{
				name:        "delete tail",
				initial:     []uint8{1, 2, 3},
				toDelete:    3,
				expected:    []uint8{1, 2},
				expectedErr: false,
			},
			{
				name:        "delete inside",
				initial:     []uint8{1, 2, 3, 4, 5},
				toDelete:    4,
				expected:    []uint8{1, 2, 3, 5},
				expectedErr: false,
			},
			{
				name:        "delete non existent",
				initial:     []uint8{1, 2, 3},
				toDelete:    0,
				expected:    []uint8{1, 2, 3},
				expectedErr: true,
			},
			{
				name:        "delete only",
				initial:     []uint8{1},
				toDelete:    1,
				expected:    []uint8{},
				expectedErr: false,
			},
			{
				name:        "delete on empty",
				initial:     []uint8{},
				toDelete:    1,
				expected:    []uint8{},
				expectedErr: true,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func (t *testing.T) {
				dll := NewDoubly[uint8]()
				for _, val := range tc.initial {
					dll.Append(val)
				}

				err := dll.Delete(tc.toDelete)
				
				if tc.expectedErr == true {
					if err == nil {
						t.Errorf("Expected to receive error from delete")
					}
				} 

				if err != nil {
					t.Errorf("Received unexpected error %v", err)
				}

				got := dll.ToSlice()
				if reflect.DeepEqual(tc.expected, got) == false {
					t.Errorf("Failed deep equality check after delete, received: %v, expected: %v", got, tc.expected)
				}



			})
		}
	})
}
