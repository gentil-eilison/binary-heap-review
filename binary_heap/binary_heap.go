package binary_heap

import (
	"fmt"
	"math"
	"slices"
)

type BinaryHeap struct {
	values []int
}

func (binaryHeap BinaryHeap) getSize() int {
	return len(binaryHeap.values) - 1
}

func New(values ...int) BinaryHeap {
	return BinaryHeap{
		values: slices.Insert(values, 0, int(math.Inf(0))),
	}
}

func (binaryHeap *BinaryHeap) GoDown(elementIdx int) {
	j := 2 * elementIdx

	if j <= binaryHeap.getSize() {
		if j < binaryHeap.getSize() {
			if binaryHeap.values[j+1] > binaryHeap.values[j] {
				j = j + 1
			}
		}
		if binaryHeap.values[j] > binaryHeap.values[elementIdx] {
			aux := binaryHeap.values[elementIdx]
			binaryHeap.values[elementIdx] = binaryHeap.values[j]
			binaryHeap.values[j] = aux
			binaryHeap.GoDown(j)
		}
	}
}

func (binaryHeap *BinaryHeap) Heapify() {
	for i := binaryHeap.getSize() / 2; i >= 1; i-- {
		binaryHeap.GoDown(i)
	}
}

func (binaryHeap BinaryHeap) Display() {
	fmt.Println(binaryHeap.values[1:])
}
