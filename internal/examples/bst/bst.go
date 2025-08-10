package bst

import (
	"fmt"

	"golang.org/x/exp/constraints"

)

// left is always less than or equal to value of parent
// right is always higher than value of parent
type Node[T constraints.Ordered] struct {
	Value T
	left  *Node[T]
	right *Node[T]
}

type BST[T constraints.Ordered] struct {
	items int
	root  *Node[T]
}

type BSTer[T constraints.Ordered] interface {
	Insert(value T)
	Search(value T) (*Node[T], error)
	Delete(value T) error
	InOrder() []T
}

func NewBST[T constraints.Ordered]() *BST[T] {
	return &BST[T]{
		items: 0,
		root:  nil,
	}
}

func (bst *BST[T]) Insert(value T) {
	node := &Node[T]{Value: value}
	bst.items++
	if bst.root == nil {
		bst.root = node
		return
	}

	curr := bst.root
	for {
		// should go on the left
		if value <= curr.Value {
			// check if left is empty
			if curr.left == nil {
				curr.left = node
				return
			} else {
				curr = curr.left
			}

		} else {
			if curr.right == nil {
				curr.right = node
				return
			} else {
				curr = curr.right
			}
		}
	}

}

func (bst *BST[T]) Search(value T) (*Node[T], error) {
	curr := bst.root
	for curr != nil {

		// value was found
		if value == curr.Value {
			return curr, nil
		}

		// value not found
		if value < curr.Value {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}
	return nil, fmt.Errorf("value not found")
}


func (bst *BST[T]) Delete(value T) error {
	if bst.root == nil {
		return fmt.Errorf("value not found")
	}

	var parent *Node[T]
	curr := bst.root

	for curr != nil && curr.Value != value {
		parent = curr
		if value < curr.Value {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}

	if curr == nil {
		return fmt.Errorf("value not found")
	}

	nodeToDelete := curr

	if parent == nil {
		// Case 1: Root is a leaf.
		if nodeToDelete.left == nil && nodeToDelete.right == nil {
			bst.root = nil
		// Case 2: root has only one child
		} else if nodeToDelete.left == nil {
			bst.root = nodeToDelete.right
		} else if nodeToDelete.right == nil {
			bst.root = nodeToDelete.left
		// root has 2 children
		} else {
			successorParent := nodeToDelete
			successor := nodeToDelete.right
			for successor.left != nil {
				successorParent = successor
				successor = successor.left
			}
			nodeToDelete.Value = successor.Value
			if successorParent.left == successor {
				successorParent.left = successor.right
			} else {
				successorParent.right = successor.right
			}
		}
		bst.items--
		return nil 
	}


	// Case 1: leaf node
	if nodeToDelete.left == nil && nodeToDelete.right == nil {
		if parent.left == nodeToDelete {
			parent.left = nil
		} else {
			parent.right = nil
		}
	// Case 2: only one child
	} else if nodeToDelete.left == nil { 
		// only right child
		if parent.left == nodeToDelete {
			parent.left = nodeToDelete.right
		} else {
			parent.right = nodeToDelete.right
		}
	} else if nodeToDelete.right == nil { 
		// only left child
		if parent.left == nodeToDelete {
			parent.left = nodeToDelete.left
		} else {
			parent.right = nodeToDelete.left
		}
	} else {
		successorParent := nodeToDelete
		successor := nodeToDelete.right
		for successor.left != nil {
			successorParent = successor
			successor = successor.left
		}
		nodeToDelete.Value = successor.Value
		if successorParent.left == successor {
			successorParent.left = successor.right
		} else {
			successorParent.right = successor.right
		}
	}

	bst.items--
	return nil
}


func inOrderHelper[T constraints.Ordered](node *Node[T], result *[]T) {
	if node == nil {
		return
	}

	inOrderHelper(node.left, result)
	*result = append(*result, node.Value)
	inOrderHelper(node.right, result)

}

func (bst *BST[T]) InOrder() []T {
	result := make([]T, 0)
	inOrderHelper(bst.root, &result)
	return  result
}