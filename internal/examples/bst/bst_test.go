package bst

import (
	"reflect"
	"testing"
)


func TestBST(t *testing.T) {
	t.Run("Insert and search", func(t *testing.T) {
		bst := NewBST[int]()
		values := []int{10, 5, 15, 3, 8, 12, 18}
		for _, v := range values {
			bst.Insert(v)
		}

		for _, v := range values {
			if _, err := bst.Search(v); err != nil {
				t.Errorf("Search failed for existing value %d: %v", v, err)
			}
		}

		if _, err := bst.Search(99); err == nil {
			t.Error("Search succeeded for non-existent value 99, but it should have failed")
		}


		want := []int{3, 5, 8, 10, 12, 15, 18}
		got := bst.InOrder() 
		if !reflect.DeepEqual(got, want) {
			t.Errorf("InOrder traversal is incorrect: got %v, want %v", got, want)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		testCases := []struct {
			name          string
			initialValues []int
			valueToDelete int
			expectedOrder []int
			expectError   bool
		}{
			{
				name:          "delete leaf node",
				initialValues: []int{10, 5, 15, 3},
				valueToDelete: 3,
				expectedOrder: []int{5, 10, 15},
				expectError:   false,
			},
			{
				name:          "delete node with one right child",
				initialValues: []int{10, 5, 15, 3, 4},
				valueToDelete: 3,
				expectedOrder: []int{4, 5, 10, 15},
				expectError:   false,
			},
			{
				name:          "delete node with one left child",
				initialValues: []int{10, 5, 15, 3, 2},
				valueToDelete: 3,
				expectedOrder: []int{2, 5, 10, 15},
				expectError:   false,
			},
			{
				name:          "delete node with two children",
				initialValues: []int{10, 5, 15, 3, 8, 7, 9},
				valueToDelete: 5,
				expectedOrder: []int{3, 7, 8, 9, 10, 15},
				expectError:   false,
			},
			{
				name:          "delete root with two children",
				initialValues: []int{10, 5, 15, 3, 8, 12, 18},
				valueToDelete: 10,
				expectedOrder: []int{3, 5, 8, 12, 15, 18},
				expectError:   false,
			},
			{
				name:          "delete root with only right child",
				initialValues: []int{10, 15, 18},
				valueToDelete: 10,
				expectedOrder: []int{15, 18},
				expectError:   false,
			},
			{
				name:          "delete root with only left child",
				initialValues: []int{10, 5, 3},
				valueToDelete: 10,
				expectedOrder: []int{3, 5},
				expectError:   false,
			},
			{
				name:          "delete root when it is the only node",
				initialValues: []int{10},
				valueToDelete: 10,
				expectedOrder: []int{},
				expectError:   false,
			},
			{
				name:          "delete non-existent value",
				initialValues: []int{10, 5, 15},
				valueToDelete: 99,
				expectedOrder: []int{5, 10, 15},
				expectError:   true,
			},
			{
				name:          "delete from an empty tree",
				initialValues: []int{},
				valueToDelete: 10,
				expectedOrder: []int{},
				expectError:   true,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				bst := NewBST[int]()
				for _, v := range tc.initialValues {
					bst.Insert(v)
				}

				err := bst.Delete(tc.valueToDelete)

				if tc.expectError {
					if err == nil {
						t.Errorf("Expected an error but got nil")
					}
				} else {
					if err != nil {
						t.Errorf("Did not expect an error but got: %v", err)
					}
					got := bst.InOrder()
					if !reflect.DeepEqual(got, tc.expectedOrder) {
						t.Errorf("Tree state is incorrect after deletion: got %v, want %v", got, tc.expectedOrder)
					}
				}
			})
		}
	})
}
