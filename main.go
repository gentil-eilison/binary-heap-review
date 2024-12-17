package main

import (
	"github.com/gentil-eilison/binary-heap-review/binary_heap"
)

func main() {
	heap := binary_heap.New(10, 20, 30, 40, 60)
	heap.Heapify()
	heap.Display()
}
