package gollections

// Queue represents a queue data structure.
type Queue struct {
	elements []interface{}
}

// Enqueue adds an element to the end of the queue.
func (q *Queue) Enqueue(value interface{}) {
	q.elements = append(q.elements, value)
}

// Dequeue removes and returns the front element from the queue.
func (q *Queue) Dequeue() interface{} {
	if len(q.elements) == 0 {
		return nil
	}
	value := q.elements[0]
	q.elements = q.elements[1:]
	return value
}

// Peek returns the front element of the queue without removing it.
func (q *Queue) Peek() interface{} {
	if len(q.elements) == 0 {
		return nil
	}
	return q.elements[0]
}

// Size returns the number of elements in the queue.
func (q *Queue) Size() int {
	return len(q.elements)
}

// IsEmpty returns true if the queue is empty, otherwise false.
func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}
