package queue_test

import (
	"reflect"
	"testing"

	"github.com/felipebool/dsa/ds/queue"
	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	type element struct {
		value string
	}

	testCases := map[string]struct {
		elementsToEnqueue   []any
		elementsToDequeue   []any
		expectedString      string
		expectedPeekElement any
	}{
		"only integers": {
			elementsToEnqueue:   []any{1, 2, 3},
			elementsToDequeue:   []any{1, 2, 3},
			expectedString:      "[1] -> [2] -> [3]",
			expectedPeekElement: 1,
		},
		"only strings": {
			elementsToEnqueue:   []any{"a", "b", "c"},
			elementsToDequeue:   []any{"a", "b", "c"},
			expectedString:      "[a] -> [b] -> [c]",
			expectedPeekElement: "a",
		},
		"only structs": {
			elementsToEnqueue: []any{
				element{value: "a"},
				element{value: "b"},
				element{value: "c"},
			},
			elementsToDequeue: []any{
				element{value: "a"},
				element{value: "b"},
				element{value: "c"},
			},
			expectedString:      "[{value:a}] -> [{value:b}] -> [{value:c}]",
			expectedPeekElement: element{value: "a"},
		},
		"mixed types": {
			elementsToEnqueue: []any{
				1,
				"b",
				element{value: "c"},
			},
			elementsToDequeue: []any{
				1,
				"b",
				element{value: "c"},
			},
			expectedString:      "[1] -> [b] -> [{value:c}]",
			expectedPeekElement: 1,
		},
	}

	for label := range testCases {
		tc := testCases[label]
		t.Run(label, func(t *testing.T) {
			t.Parallel()

			q := queue.NewQueue()
			for i := range tc.elementsToEnqueue {
				q.Enqueue(tc.elementsToEnqueue[i])
			}

			assert.Equal(t, tc.expectedString, q.String())
			assert.Equal(t, tc.expectedPeekElement, q.Peek())
			assert.False(t, q.IsEmpty())

			for i := range tc.elementsToDequeue {
				el := q.Dequeue()
				assert.Equal(t, tc.elementsToDequeue[i], el)
				assert.True(t, reflect.DeepEqual(el, tc.elementsToDequeue[i]))
			}

			assert.Equal(t, "[]", q.String())
			assert.Nil(t, q.Dequeue())
			assert.Nil(t, q.Peek())
			assert.True(t, q.IsEmpty())
		})
	}
}
