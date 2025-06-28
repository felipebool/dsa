package heap_test

import (
	"github.com/felipebool/dsa/ds/element"
	"testing"

	"github.com/felipebool/dsa/ds/heap"
	"github.com/stretchr/testify/assert"
)

type item struct {
	key int
}

func (e item) GetKey() int {
	return e.key
}

func TestHeap(t *testing.T) {
	testCases := map[string]struct {
		elements            []element.Getter
		elementsToPop       []element.Getter
		compareType         heap.CompareType
		expectedPeekElement element.Getter
		expectedString      string
	}{
		"max heap random elements": {
			elements: []element.Getter{
				item{key: 17},
				item{key: 2},
				item{key: 15},
				item{key: 23},
				item{key: 4},
				item{key: 9},
				item{key: 0},
			},
			elementsToPop: []element.Getter{
				item{key: 23},
				item{key: 17},
				item{key: 15},
				item{key: 9},
				item{key: 4},
				item{key: 2},
				item{key: 0},
			},
			compareType:         heap.MaxHeap,
			expectedPeekElement: item{key: 23},
			expectedString:      "[23] -> [17] -> [15] -> [2] -> [4] -> [9] -> [0]",
		},
		"max heap ascending elements": {
			elements: []element.Getter{
				item{key: 0},
				item{key: 2},
				item{key: 4},
				item{key: 9},
				item{key: 15},
				item{key: 17},
				item{key: 23},
			},
			elementsToPop: []element.Getter{
				item{key: 23},
				item{key: 17},
				item{key: 15},
				item{key: 9},
				item{key: 4},
				item{key: 2},
				item{key: 0},
			},
			compareType:         heap.MaxHeap,
			expectedPeekElement: item{key: 23},
			expectedString:      "[23] -> [9] -> [17] -> [0] -> [4] -> [2] -> [15]",
		},
		"max heap descending elements": {
			elements: []element.Getter{
				item{key: 23},
				item{key: 17},
				item{key: 15},
				item{key: 9},
				item{key: 4},
				item{key: 2},
				item{key: 0},
			},
			elementsToPop: []element.Getter{
				item{key: 23},
				item{key: 17},
				item{key: 15},
				item{key: 9},
				item{key: 4},
				item{key: 2},
				item{key: 0},
			},
			compareType:         heap.MaxHeap,
			expectedPeekElement: item{key: 23},
			expectedString:      "[23] -> [17] -> [15] -> [9] -> [4] -> [2] -> [0]",
		},
		"max heap no elements": {
			elements:            []element.Getter{},
			elementsToPop:       []element.Getter{},
			compareType:         heap.MaxHeap,
			expectedPeekElement: nil,
			expectedString:      "[]",
		},
		"max heap duplicated elements": {
			elements: []element.Getter{
				item{key: 23},
				item{key: 17},
				item{key: 2},
				item{key: 15},
				item{key: 23},
				item{key: 4},
				item{key: 9},
				item{key: 0},
				item{key: 15},
			},
			elementsToPop: []element.Getter{
				item{key: 23},
				item{key: 23},
				item{key: 17},
				item{key: 15},
				item{key: 15},
				item{key: 9},
				item{key: 4},
				item{key: 2},
				item{key: 0},
			},
			compareType:         heap.MaxHeap,
			expectedPeekElement: item{key: 23},
			expectedString:      "[23] -> [23] -> [9] -> [15] -> [17] -> [2] -> [4] -> [0] -> [15]",
		},
		"max heap single element": {
			elements: []element.Getter{
				item{key: 23},
			},
			elementsToPop: []element.Getter{
				item{key: 23},
			},
			compareType:         heap.MaxHeap,
			expectedPeekElement: item{key: 23},
			expectedString:      "[23]",
		},
		"min heap random elements": {
			elements: []element.Getter{
				item{key: 17},
				item{key: 2},
				item{key: 15},
				item{key: 23},
				item{key: 4},
				item{key: 9},
				item{key: 0},
			},
			elementsToPop: []element.Getter{
				item{key: 0},
				item{key: 2},
				item{key: 4},
				item{key: 9},
				item{key: 15},
				item{key: 17},
				item{key: 23},
			},
			compareType:         heap.MinHeap,
			expectedPeekElement: item{key: 0},
			expectedString:      "[0] -> [4] -> [2] -> [23] -> [17] -> [15] -> [9]",
		},
		"min heap ascending elements": {
			elements: []element.Getter{
				item{key: 0},
				item{key: 2},
				item{key: 4},
				item{key: 9},
				item{key: 15},
				item{key: 17},
				item{key: 23},
			},
			elementsToPop: []element.Getter{
				item{key: 0},
				item{key: 2},
				item{key: 4},
				item{key: 9},
				item{key: 15},
				item{key: 17},
				item{key: 23},
			},
			compareType:         heap.MinHeap,
			expectedPeekElement: item{key: 0},
			expectedString:      "[0] -> [2] -> [4] -> [9] -> [15] -> [17] -> [23]",
		},
		"min heap descending elements": {
			elements: []element.Getter{
				item{key: 23},
				item{key: 17},
				item{key: 15},
				item{key: 9},
				item{key: 4},
				item{key: 2},
				item{key: 0},
			},
			elementsToPop: []element.Getter{
				item{key: 0},
				item{key: 2},
				item{key: 4},
				item{key: 9},
				item{key: 15},
				item{key: 17},
				item{key: 23},
			},
			compareType:         heap.MinHeap,
			expectedPeekElement: item{key: 0},
			expectedString:      "[0] -> [9] -> [2] -> [23] -> [15] -> [17] -> [4]",
		},
		"min heap no elements": {
			elements:            []element.Getter{},
			elementsToPop:       []element.Getter{},
			compareType:         heap.MinHeap,
			expectedPeekElement: nil,
			expectedString:      "[]",
		},
		"min heap duplicated elements": {
			elements: []element.Getter{
				item{key: 23},
				item{key: 17},
				item{key: 2},
				item{key: 15},
				item{key: 23},
				item{key: 4},
				item{key: 9},
				item{key: 0},
				item{key: 15},
			},
			elementsToPop: []element.Getter{
				item{key: 0},
				item{key: 2},
				item{key: 4},
				item{key: 9},
				item{key: 15},
				item{key: 15},
				item{key: 17},
				item{key: 23},
				item{key: 23},
			},
			compareType:         heap.MinHeap,
			expectedPeekElement: item{key: 0},
			expectedString:      "[0] -> [2] -> [4] -> [15] -> [23] -> [17] -> [9] -> [23] -> [15]",
		},
		"min heap single element": {
			elements: []element.Getter{
				item{key: 23},
			},
			elementsToPop: []element.Getter{
				item{key: 23},
			},
			compareType:         heap.MinHeap,
			expectedPeekElement: item{key: 23},
			expectedString:      "[23]",
		},
	}

	for label := range testCases {
		tc := testCases[label]
		t.Run(label, func(t *testing.T) {
			t.Parallel()

			mHeap1 := heap.NewHeap(tc.compareType)
			mHeap1.Heapify(tc.elements)
			assert.Equal(t, tc.expectedPeekElement, mHeap1.Peek())

			mHeap2 := heap.NewHeap(tc.compareType)
			for _, e := range tc.elements {
				mHeap2.Push(e)
			}
			assert.Equal(t, tc.expectedPeekElement, mHeap2.Peek())

			assert.Equal(t, tc.expectedString, mHeap1.String())
			assert.Equal(t, tc.expectedString, mHeap2.String())
			for i := range tc.elements {
				assert.Equal(t, tc.elementsToPop[i], mHeap1.Pop())
				assert.Equal(t, tc.elementsToPop[i], mHeap2.Pop())
			}

			assert.Equal(t, "[]", mHeap1.String())
			assert.Equal(t, "[]", mHeap2.String())

			assert.Nil(t, mHeap1.Peek())
			assert.Nil(t, mHeap2.Peek())

			assert.True(t, mHeap1.IsEmpty())
			assert.True(t, mHeap2.IsEmpty())
		})
	}
}
