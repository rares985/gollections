package gollections

type Stack[T any] struct {
	elements []T
}

// Push adds an element to the top of the stack.
func (s *Stack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

// Pop removes and returns the top element from the stack.
func (s *Stack[T]) Pop() *T {
	if len(s.elements) == 0 {
		return nil
	}
	lastIndex := len(s.elements) - 1
	value := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return &value
}

// Peek returns the top element from the stack without removing it.
func (s *Stack[T]) Peek() *T {
	if len(s.elements) == 0 {
		return nil
	}
	return &s.elements[len(s.elements)-1]
}

// IsEmpty returns true if the stack is empty, false otherwise.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.elements)
}
