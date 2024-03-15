package gollections

import "fmt"

type Heap[T comparable] struct {
	elements []T
	less     func(a, b T) bool
}

func (heap *Heap[T]) Insert(val T) {
	heap.elements = append(heap.elements, val)

	idx := len(heap.elements) - 1

	heap.heapifyUp(idx)
}

func (heap *Heap[T]) heapifyUp(startIndex int) {
	childIdx := startIndex

	for {
		parentIdx := (childIdx - 1) / 2

		heapPropertyViolated := heap.less(heap.elements[parentIdx], heap.elements[childIdx])
		if !heapPropertyViolated {
			break
		}

		aux := heap.elements[parentIdx]
		heap.elements[parentIdx] = heap.elements[childIdx]
		heap.elements[childIdx] = aux

		childIdx = parentIdx
	}
}

func (heap *Heap[T]) Peek() T {
	return heap.elements[0]
}

func (heap *Heap[T]) Delete() (*T, error) {
	if heap.IsEmpty() {
		return nil, fmt.Errorf("can not delete from empty heap")
	}

	last := len(heap.elements) - 1

	heap.swap(0, last)

	toReturn := heap.elements[last]
	heap.elements = heap.elements[:last]

	fmt.Printf("Deleting %v\n", toReturn)
	fmt.Printf("before heapify down heap = %v\n", heap.elements)
	heap.heapifyDown(0)
	fmt.Printf("after hepify down heap = %v\n", heap.elements)

	return &toReturn, nil
}

func (heap *Heap[T]) swap(first, second int) {
	if first >= len(heap.elements) || second >= len(heap.elements) {
		return
	}

	aux := heap.elements[first]
	heap.elements[first] = heap.elements[second]
	heap.elements[second] = aux
}

func (heap *Heap[T]) heapifyDown(parent int) {
	left := 2*parent + 1
	right := 2*parent + 2

	for left < len(heap.elements) || right < len(heap.elements) {
		leftChildValid := left < len(heap.elements)
		heapPropertyLeftOk := leftChildValid && heap.less(heap.elements[left], heap.elements[parent])

		rightChildValid := right < len(heap.elements)
		heapPropertyRightOk := rightChildValid && heap.less(heap.elements[right], heap.elements[parent])

		if heapPropertyLeftOk && heapPropertyRightOk {
			break
		}

		switch {
		// has both, swap parent with max/min
		case leftChildValid && rightChildValid:
			if heap.less(heap.elements[left], heap.elements[right]) {
				heap.swap(parent, right)
				parent = right
			} else {
				heap.swap(parent, left)
				parent = left
			}
		// has only right, swap if doesn't respect property
		case rightChildValid && heap.less(heap.elements[parent], heap.elements[right]):
			heap.swap(parent, right)
			parent = right
		// has only left, swap if doesn't respect property
		case leftChildValid && heap.less(heap.elements[parent], heap.elements[left]):
			heap.swap(parent, left)
			parent = left
		}

		left = 2*parent + 1
		right = 2*parent + 2
	}
}

func (heap *Heap[T]) Size() int {
	return len(heap.elements)
}

func (heap *Heap[T]) IsEmpty() bool {
	return len(heap.elements) == 0
}

func (heap *Heap[T]) Equals(other *Heap[T]) bool {
	switch {
	case heap == nil && other == nil:
		return true
	case heap != nil && other == nil:
		return false
	case heap == nil && other != nil:
		return false
	case heap.Size() != other.Size():
		return false
	}

	for i := 0; i < heap.Size(); i++ {
		if heap.elements[i] != other.elements[i] {
			return false
		}
	}

	return true
}
