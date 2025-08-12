package heaps

import (
	"fmt"
)



type MinHeap struct {
	values []int
}


type MinHeaper interface {
	
	Insert(value int)
	ExtractMin() (int, error)
}

func (mh *MinHeap) lastNonLeafNode() int {
	return  (len(mh.values) / 2) - 1
}

func (mh *MinHeap) parentIndex(i int) int {
	return (i - 1)/2
}

func (mh *MinHeap) leftChildIndex(i int) int {
	return (2 * i) + 1
}

func (mh *MinHeap) rightChildIndex(i int) int {
	return (2 * i) + 2
}

func (mh *MinHeap) swap(i, j int) {
	mh.values[i], mh.values[j] = mh.values[j], mh.values[i]
}

func (mh *MinHeap) Insert(value int) {
	mh.values = append(mh.values, value)
	idx := len(mh.values) - 1

	for idx > 0{
		parentIdx := mh.parentIndex(idx)
		if mh.values[idx] < mh.values[parentIdx] {
			mh.swap(idx, parentIdx)
			idx = parentIdx
		} else {
			break
		}
		
	}
}
func (mh *MinHeap) heapifyDown(index int) {
	currentIndex := index
    lastIndex := len(mh.values) - 1

    for mh.leftChildIndex(currentIndex) <= lastIndex {
        leftChildIndex := mh.leftChildIndex(currentIndex)
        smallestChildIndex := leftChildIndex 


		rightChildIndex := mh.rightChildIndex(currentIndex)
        if rightChildIndex <= lastIndex && mh.values[rightChildIndex] < mh.values[leftChildIndex] {
            smallestChildIndex = rightChildIndex
        }

        if mh.values[currentIndex] <= mh.values[smallestChildIndex] {
            break
        }

        mh.swap(currentIndex, smallestChildIndex)
        
        currentIndex = smallestChildIndex
    }
} 


func (mh *MinHeap) ExtractMin() (int, error) {
	lastIndex := len(mh.values) - 1
	if lastIndex == -1 {
		return -1, fmt.Errorf("heap is empty")
	}
	rootValue := mh.values[0]

	// both are same index
	if len(mh.values) == 1 {
		mh.values = []int{}
		return rootValue, nil

	}

	mh.swap(lastIndex, 0)
	// lets remove last element
	mh.values = mh.values[:lastIndex]

	// heapify down the new root until we find location for it
	mh.heapifyDown(0)
	return rootValue, nil
}



func (mh *MinHeap) Peek() (int, error) {
	if len(mh.values) == 0 {
		return -1, fmt.Errorf("heap is empty")
	}
	return mh.values[0], nil
}

func (mh *MinHeap) BuildHeap(slice []int) {
	mh.values = slice
	lastParentIndex := mh.lastNonLeafNode()
	for i := lastParentIndex; i >= 0; i-- {
		mh.heapifyDown(i)
	}
}