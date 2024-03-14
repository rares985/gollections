package gollections

// Queue represents a queue data structure.
type Queue[T any] struct {
	elements []T
}

// Enqueue adds an element to the end of the queue.
func (q *Queue[T]) Enqueue(value T) {
	q.elements = append(q.elements, value)
}

// Dequeue removes and returns the front element from the queue.
func (q *Queue[T]) Dequeue() *T {
	if len(q.elements) == 0 {
		return nil
	}
	value := q.elements[0]
	q.elements = q.elements[1:]
	return &value
}

// Peek returns the front element of the queue without removing it.
func (q *Queue[T]) Peek() *T {
	if len(q.elements) == 0 {
		return nil
	}
	return &q.elements[0]
}

// Size returns the number of elements in the queue.
func (q *Queue[T]) Size() int {
	return len(q.elements)
}

// IsEmpty returns true if the queue is empty, otherwise false.
func (q *Queue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}
