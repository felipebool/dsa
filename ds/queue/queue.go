package queue

import (
	"fmt"
	"sync"
)

// Queue is the structure that holds the elements in the Queue.
// It has also a sync.Mutex to ensure goroutine safety.
type Queue struct {
	mu       sync.Mutex
	elements []any
}

// Enqueue adds a new element to the end of the Queue.
func (q *Queue) Enqueue(element any) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.elements = append(q.elements, element)
}

// Dequeue removes the first element from the Queue and resize it.
// It returns nil if the Queue is empty.
func (q *Queue) Dequeue() any {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.elements) == 0 {
		return nil
	}
	element := q.elements[0]
	q.elements = q.elements[1:]
	return element
}

// Peek returns the first element without removing
// it from the Queue. It returns nil if the Queue is empty.
func (q *Queue) Peek() any {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.elements) == 0 {
		return nil
	}
	return q.elements[0]
}

// IsEmpty returns true if the Queue has no elements,
// false otherwise.
func (q *Queue) IsEmpty() bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	return len(q.elements) == 0
}

// String returns a string representation of the Queue,
// it is useful for debugging and visualization. It returns
// "[]" if the Queue is empty.
func (q *Queue) String() string {
	q.mu.Lock()
	defer q.mu.Unlock()

	result := ""
	if len(q.elements) == 0 {
		return "[]"
	}

	for i := range q.elements {
		if i == len(q.elements)-1 {
			result += fmt.Sprintf("[%+v]", q.elements[i])
			break
		}
		result += fmt.Sprintf("[%+v] -> ", q.elements[i])
	}
	return result
}

// NewQueue returns a new Queue with no elements.
func NewQueue() *Queue {
	return &Queue{
		elements: make([]any, 0),
	}
}
