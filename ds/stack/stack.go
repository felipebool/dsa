package stack

import (
	"fmt"
	"sync"
)

// Stack is the structure that holds the elements in the Stack,
// it has also a mutex to ensure goroutine safety.
type Stack struct {
	mu       sync.Mutex
	elements []any
}

// Push pushes a new element on the top of the Stack.
func (s *Stack) Push(element any) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.elements = append(s.elements, element)
}

// Pop pops the element on top of the Stack,
// resizes the Stack and returns the element.
// It returns nil if the Stack is empty.
func (s *Stack) Pop() any {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.elements) == 0 {
		return nil
	}

	last := len(s.elements) - 1
	element := s.elements[last]
	s.elements = s.elements[:last]
	return element
}

// Peek returns the element on top of the Stack
// without removing it. It returns nil if the
// Stack is empty.
func (s *Stack) Peek() any {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.elements) == 0 {
		return nil
	}
	return s.elements[len(s.elements)-1]
}

// IsEmpty returns true if the Stack has no elements,
// false otherwise.
func (s *Stack) IsEmpty() bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	return len(s.elements) == 0
}

// String returns a string representation of the Stack,
// it is useful for debugging and visualization. It returns
// "[]" if the Stack is empty.
func (s *Stack) String() string {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := ""
	if len(s.elements) == 0 {
		return "[]"
	}
	for i := len(s.elements) - 1; i >= 0; i-- {
		if i == 0 {
			result += fmt.Sprintf("[%+v]", s.elements[i])
			break
		}
		result += fmt.Sprintf("[%+v] -> ", s.elements[i])
	}
	return result
}

// NewStack returns a new Stack with no elements
func NewStack() *Stack {
	return &Stack{
		elements: make([]any, 0),
	}
}
