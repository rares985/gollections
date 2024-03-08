package gollections

type Stack struct {
	elements []interface{}
}

// Push adds an element to the top of the stack.
func (s *Stack) Push(value interface{}) {
	s.elements = append(s.elements, value)
}

// Pop removes and returns the top element from the stack.
func (s *Stack) Pop() interface{} {
	if len(s.elements) == 0 {
		return nil
	}
	lastIndex := len(s.elements) - 1
	value := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return value
}

// Peek returns the top element from the stack without removing it.
func (s *Stack) Peek() interface{} {
	if len(s.elements) == 0 {
		return nil
	}
	return s.elements[len(s.elements)-1]
}

// IsEmpty returns true if the stack is empty, false otherwise.
func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}
