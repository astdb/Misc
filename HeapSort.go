package main

import(
	"fmt"
)

func main() {
	
}

func Parent(i int) int {
	return i/2
}

func Left(i int) int {
	return 2*i
}

func Right(i int) int {
	return 2*i + 1
}

func MaxHeapify(heap *Heap, i int) []int  {
	l = Left(i)
	r = Right(i)

	var largest int

	if l <= heap.HeapSize && heap.Data[l] > heap.Data[i] {
		largest = l
} else {
	largest = i
}

if r <= heap.HeapSize && heap.Data[r] > heap.Data[largest] {
	largest = r
}

if largest != i {
	temp := heap.Data[i]
	heap.Data[i] = heap.Data[largest]
	heap.Data[largest] = temp

	MaxHeapify(heap, largest)
}
}

func BuildMaxHeap(heap *Heap) {
	heap.HeapSize = len(heap.Data)

	for i := len(heap.Data)/2; i >= 1; i-- {
		MaxHeapify(heap, i)
	}
}

func HeapSort(heap *Heap) *Heap {
	BuildMaxHeap(heap)
	for i := len(heap.Data); i >= 2; i-- {
		temp := heap.Data[1]
		heap.Data[1] = heap.Data[i]
		heap.Data[i] = temp
	}

return heap
}

type Heap struct {
	Data []int
	HeapSize int
}

func NewHeap(data []int) *Heap {
	var h Heap
	h.Data = data
	HeapSize = len(h.Data)
	return &h
}
