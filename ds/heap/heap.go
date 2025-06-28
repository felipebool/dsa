package heap

import (
	"fmt"
	"sync"

	"github.com/felipebool/dsa/ds/element"
)

const (
	MinHeap CompareType = iota
	MaxHeap
)

// CompareType is the type of the heap,
// 0 for MinHeap
// 1 for MaxHeap
type CompareType int

type compareFn func(x, y int) bool

// Heap is the structure that holds the elements in the Heap,
// it has a slice of Element, which is an interface that defines
// a method called GetKey(), used to get the key to place the
// elements in the Heap, a comparer which is a function
// that receives two integers and returns a boolean, if the Heap
// is a MaxHeap, returns true when the first element is larger than
// the second, and if it is a MinHeap, returns true when the first
// element is smaller than the second one. It has also a sync.Mutex
// to ensure goroutine safety.
type Heap struct {
	mu       sync.Mutex
	elements []element.Getter
	comparer compareFn
}

// Peek returns the element with the smallest key in a MinHeap
// and the element with the largest key in a MaxHeap, without
// removing it from the Heap. It returns nil if the Heap is empty.
func (h *Heap) Peek() element.Getter {
	h.mu.Lock()
	defer h.mu.Unlock()

	if len(h.elements) == 0 {
		return nil
	}
	return h.elements[0]
}

// Pop removes the element with the smallest key in a MinHeap
// and the element with the largest key in a MaxHeap, and returns it.
// It also restores the Heap condition by calling siftDown and will
// return nil if the Heap is empty.
func (h *Heap) Pop() element.Getter {
	h.mu.Lock()
	defer h.mu.Unlock()

	if len(h.elements) == 0 {
		return nil
	}

	// remove root element
	el := h.elements[0]

	// place the last element in the root position
	h.elements[0] = h.elements[len(h.elements)-1]

	// resize the list by removing the last element
	h.elements = h.elements[:len(h.elements)-1]

	h.siftDown(0)

	return el
}

// Push inserts a new element in the Heap, it does
// by adding the element in the last position of the
// Heap and then calling siftUp to restore the Heap condition.
func (h *Heap) Push(x element.Getter) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if len(h.elements) == 0 {
		h.elements = append(h.elements, x)
		return
	}

	// add element to the end of the heap
	h.elements = append(h.elements, x)

	h.siftUp(len(h.elements) - 1)
}

// IsEmpty returns true when there are no elements in the
// Heap, false otherwise.
func (h *Heap) IsEmpty() bool {
	h.mu.Lock()
	defer h.mu.Unlock()

	return len(h.elements) == 0
}

// Heapify receives a slice of elements and pushes them to the
// Heap by calling Push to each element. If Heapify is called
// in a non-empty Heap, the current elements will be preserved
// and the new ones will be pushed to the Heap.
func (h *Heap) Heapify(elements []element.Getter) {
	for _, el := range elements {
		h.Push(el)
	}
}

// String returns a string representation of the Heap,
// it is useful for debugging and visualization. It returns
// "[]" if the Heap is empty.
func (h *Heap) String() string {
	h.mu.Lock()
	defer h.mu.Unlock()

	result := ""
	if len(h.elements) == 0 {
		return "[]"
	}

	for i := range h.elements {
		el := h.elements[i]
		if i == len(h.elements)-1 {
			result += fmt.Sprintf("[%d]", el.GetKey())
			continue
		}
		result += fmt.Sprintf("[%d] -> ", el.GetKey())
	}
	return result
}

func (h *Heap) siftDown(x int) {
	if len(h.elements) == 0 {
		return
	}

	left, right := h.getChildren(x)
	if left < 0 && right < 0 {
		return
	}

	current := h.elements[x]
	if left >= 0 && right >= 0 {
		leftChild := h.elements[left]
		rightChild := h.elements[right]
		if h.comparer(leftChild.GetKey(), rightChild.GetKey()) {
			if h.comparer(leftChild.GetKey(), current.GetKey()) {
				h.swap(x, left)
				h.siftDown(left)
				return
			}
		}
		if h.comparer(rightChild.GetKey(), current.GetKey()) {
			h.swap(x, right)
			h.siftDown(right)
			return
		}
		return
	}

	if left >= 0 {
		leftChild := h.elements[left]
		if h.comparer(leftChild.GetKey(), current.GetKey()) {
			h.swap(x, left)
			h.siftDown(left)
			return
		}
	}

	if right >= 0 {
		rightChild := h.elements[right]
		if h.comparer(rightChild.GetKey(), current.GetKey()) {
			h.swap(x, right)
			h.siftDown(right)
			return
		}
	}
}

func (h *Heap) siftUp(x int) {
	if x == 0 {
		return
	}
	parent := h.getParent(x)
	current := h.elements[x]
	if h.comparer(current.GetKey(), h.elements[parent].GetKey()) {
		h.swap(x, parent)
		h.siftUp(parent)
	}
}

func (h *Heap) swap(x, y int) {
	h.elements[x], h.elements[y] = h.elements[y], h.elements[x]
}

func (h *Heap) setComparer(cType CompareType) {
	switch cType {
	case MinHeap:
		h.comparer = minHeap
	case MaxHeap:
		h.comparer = maxHeap
	default:
		h.comparer = minHeap
	}
}

func (h *Heap) getChildren(x int) (int, int) {
	leftChild := 2*x + 1
	if leftChild >= len(h.elements) {
		leftChild = -1
	}
	rightChild := 2*x + 2
	if rightChild >= len(h.elements) {
		rightChild = -1
	}
	return leftChild, rightChild
}

func (h *Heap) getParent(x int) int {
	return (x - 1) / 2
}

func maxHeap(x, y int) bool {
	return x > y
}

func minHeap(x, y int) bool {
	return x < y
}

// NewHeap returns a new Heap with no elements.
func NewHeap(cType CompareType) *Heap {
	heap := &Heap{}
	heap.setComparer(cType)
	heap.elements = make([]element.Getter, 0)
	return heap
}
