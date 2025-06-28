package stack_test

import (
	"reflect"
	"testing"

	"github.com/felipebool/dsa/ds/stack"
	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	type item struct {
		value string
	}

	testCases := map[string]struct {
		elementsToPush      []any
		elementsToPop       []any
		expectedString      string
		expectedPeekElement any
	}{
		"only integers": {
			elementsToPush:      []any{1, 2, 3},
			elementsToPop:       []any{3, 2, 1},
			expectedString:      "[3] -> [2] -> [1]",
			expectedPeekElement: 3,
		},
		"only strings": {
			elementsToPush:      []any{"a", "b", "c"},
			elementsToPop:       []any{"c", "b", "a"},
			expectedString:      "[c] -> [b] -> [a]",
			expectedPeekElement: "c",
		},
		"only structs": {
			elementsToPush: []any{
				item{value: "a"},
				item{value: "b"},
				item{value: "c"},
			},
			elementsToPop: []any{
				item{value: "c"},
				item{value: "b"},
				item{value: "a"},
			},
			expectedString:      "[{value:c}] -> [{value:b}] -> [{value:a}]",
			expectedPeekElement: item{value: "c"},
		},
		"mixed types": {
			elementsToPush: []any{
				1,
				"b",
				item{value: "c"},
			},
			elementsToPop: []any{
				item{value: "c"},
				"b",
				1,
			},
			expectedString:      "[{value:c}] -> [b] -> [1]",
			expectedPeekElement: item{value: "c"},
		},
	}

	for label := range testCases {
		tc := testCases[label]
		t.Run(label, func(t *testing.T) {
			t.Parallel()

			s := stack.NewStack()
			for i := range tc.elementsToPush {
				s.Push(tc.elementsToPush[i])
			}

			assert.Equal(t, tc.expectedString, s.String())
			assert.Equal(t, tc.expectedPeekElement, s.Peek())
			assert.False(t, s.IsEmpty())

			for i := range tc.elementsToPop {
				el := s.Pop()
				assert.Equal(t, tc.elementsToPop[i], el)
				assert.True(t, reflect.DeepEqual(el, tc.elementsToPop[i]))
			}

			assert.Equal(t, "[]", s.String())
			assert.Nil(t, s.Pop())
			assert.Nil(t, s.Peek())
			assert.True(t, s.IsEmpty())
		})
	}
}
