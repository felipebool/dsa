package binary_test

import (
	"github.com/felipebool/dsa/ds/element"
	"github.com/felipebool/dsa/ds/tree/binary"
	"github.com/stretchr/testify/assert"
	"testing"
)

type item struct {
	key int
}

func (e item) GetKey() int {
	return e.key
}

func (e item) SetKey(key int) {
	e.key = key
}

func TestBinaryTreeInsertion(t *testing.T) {
	testCases := map[string]struct {
		elements          []element.GetterSetter
		expectedInOrder   string
		expectedPreOrder  string
		expectedPostOrder string
	}{
		"random elements": {
			elements: []element.GetterSetter{
				item{key: 8},
				item{key: 3},
				item{key: 10},
				item{key: 1},
				item{key: 6},
				item{key: 14},
				item{key: 4},
				item{key: 7},
				item{key: 13},
			},
			expectedInOrder:   "[1] [3] [4] [6] [7] [8] [10] [13] [14] ",
			expectedPreOrder:  "[8] [3] [1] [6] [4] [7] [10] [14] [13] ",
			expectedPostOrder: "[1] [4] [7] [6] [3] [13] [14] [10] [8] ",
		},
	}

	for label := range testCases {
		tc := testCases[label]
		t.Run(label, func(t *testing.T) {
			t.Parallel()

			bst := binary.NewTree()
			for _, i := range tc.elements {
				bst.Insert(i)
			}

			assert.Equal(t, tc.expectedInOrder, bst.Traverse(binary.InOrder))
			assert.Equal(t, tc.expectedPreOrder, bst.Traverse(binary.PreOrder))
			assert.Equal(t, tc.expectedPostOrder, bst.Traverse(binary.PostOrder))
		})
	}
}

func TestBinaryTreeRemove(t *testing.T) {
	testCases := map[string]struct {
		elements          []element.GetterSetter
		elementToRemove   item
		expectedInOrder   string
		expectedPreOrder  string
		expectedPostOrder string
	}{
		"removing leaf node": {
			elements: []element.GetterSetter{
				item{key: 8},
				item{key: 3},
				item{key: 10},
				item{key: 1},
				item{key: 6},
				item{key: 14},
				item{key: 4},
				item{key: 7},
				item{key: 13},
			},
			elementToRemove:   item{key: 4},
			expectedInOrder:   "[1] [3] [6] [7] [8] [10] [13] [14] ",
			expectedPreOrder:  "[8] [3] [1] [6] [7] [10] [14] [13] ",
			expectedPostOrder: "[1] [7] [6] [3] [13] [14] [10] [8] ",
		},
		"removing node with a single child": {
			elements: []element.GetterSetter{
				item{key: 8},
				item{key: 3},
				item{key: 10},
				item{key: 1},
				item{key: 6},
				item{key: 14},
				item{key: 4},
				item{key: 7},
				item{key: 13},
			},
			elementToRemove:   item{key: 14},
			expectedInOrder:   "[1] [3] [4] [6] [7] [8] [10] [13] ",
			expectedPreOrder:  "[8] [3] [1] [6] [4] [7] [10] [13] ",
			expectedPostOrder: "[1] [4] [7] [6] [3] [13] [10] [8] ",
		},
		"removing node with two children - right child has no right subtree": {
			elements: []element.GetterSetter{
				item{key: 8},
				item{key: 3},
				item{key: 10},
				item{key: 1},
				item{key: 6},
				item{key: 14},
				item{key: 4},
				item{key: 7},
				item{key: 13},
			},
			elementToRemove:   item{key: 6},
			expectedInOrder:   "[1] [3] [4] [7] [8] [10] [13] [14] ",
			expectedPreOrder:  "[8] [3] [1] [7] [4] [10] [14] [13] ",
			expectedPostOrder: "[1] [4] [7] [3] [13] [14] [10] [8] ",
		},
		"removing node with two children - right child has right subtree": {
			elements: []element.GetterSetter{
				item{key: 8},
				item{key: 3},
				item{key: 10},
				item{key: 1},
				item{key: 6},
				item{key: 14},
				item{key: 4},
				item{key: 7},
				item{key: 13},
			},
			elementToRemove:   item{key: 3},
			expectedInOrder:   "[1] [4] [6] [7] [8] [10] [13] [14] ",
			expectedPreOrder:  "[8] [4] [1] [6] [7] [10] [14] [13] ",
			expectedPostOrder: "[1] [7] [6] [4] [13] [14] [10] [8] ",
		},
		"removing root node": {
			elements: []element.GetterSetter{
				item{key: 8},
				item{key: 3},
				item{key: 10},
				item{key: 1},
				item{key: 6},
				item{key: 14},
				item{key: 4},
				item{key: 7},
				item{key: 13},
			},
			elementToRemove:   item{key: 8},
			expectedInOrder:   "[1] [3] [4] [6] [7] [10] [13] [14] ",
			expectedPreOrder:  "[10] [3] [1] [6] [4] [7] [14] [13] ",
			expectedPostOrder: "[1] [4] [7] [6] [3] [13] [14] [10] ",
		},
	}

	for label := range testCases {
		tc := testCases[label]
		t.Run(label, func(t *testing.T) {
			t.Parallel()

			bst := binary.NewTree()
			for _, i := range tc.elements {
				bst.Insert(i)
			}

			bst.Remove(tc.elementToRemove.GetKey())

			assert.Equal(t, tc.expectedInOrder, bst.Traverse(binary.InOrder))
			assert.Equal(t, tc.expectedPreOrder, bst.Traverse(binary.PreOrder))
			assert.Equal(t, tc.expectedPostOrder, bst.Traverse(binary.PostOrder))
		})
	}
}
