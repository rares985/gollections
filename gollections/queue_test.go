package gollections

import (
	"reflect"
	"testing"
)

func TestEnqueue(t *testing.T) {
	testCases := map[string]struct {
		elem  any
		queue Queue
		want  Queue
	}{
		"enqueue empty": {
			elem:  2,
			queue: Queue{elements: []interface{}{}},
			want:  Queue{elements: []interface{}{2}},
		},
		"enqueue ok": {
			elem:  2,
			queue: Queue{elements: []interface{}{1, 3}},
			want:  Queue{elements: []interface{}{1, 3, 2}},
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			got := tc.queue
			got.Enqueue(tc.elem)

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("Got(%+v) != Want(%+v)", got, tc.want)
			}
		})
	}
}

func TestDequeue(t *testing.T) {
	testCases := map[string]struct {
		elem  any
		queue Queue
		want  Queue
	}{
		"dequeue empty": {
			elem:  2,
			queue: Queue{elements: []interface{}{2}},
			want:  Queue{elements: []interface{}{}},
		},
		"dequeue ok": {
			elem:  1,
			queue: Queue{elements: []interface{}{1, 3, 2}},
			want:  Queue{elements: []interface{}{3, 2}},
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			got := tc.queue
			elem := got.Dequeue()

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("Got(%+v) != Want(%+v)", got, tc.want)
			}

			if !reflect.DeepEqual(tc.elem, elem) {
				t.Fatalf("Dequeued(%+v) != Want(%+v)", elem, tc.elem)
			}
		})
	}
}

func TestQueuePeek(t *testing.T) {
	testCases := map[string]struct {
		elem  any
		queue Queue
		want  Queue
	}{
		"dequeue empty": {
			elem:  2,
			queue: Queue{elements: []interface{}{2}},
		},
		"dequeue ok": {
			elem:  1,
			queue: Queue{elements: []interface{}{1, 3, 2}},
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			got := tc.queue
			elem := got.Peek()

			if !reflect.DeepEqual(got, tc.queue) {
				t.Fatalf("Got(%+v) != Want(%+v)", got, tc.want)
			}

			if !reflect.DeepEqual(tc.elem, elem) {
				t.Fatalf("Peeked(%+v) != Want(%+v)", elem, tc.elem)
			}
		})
	}
}
func TestIsEmpty(t *testing.T) {
	testCases := map[string]struct {
		expected bool
		queue    Queue
	}{
		"is empty": {
			expected: true,
			queue:    Queue{elements: []interface{}{}},
		},
		"is not empty": {
			queue: Queue{elements: []interface{}{1, 3, 2}},
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			got := tc.queue
			isEmpty := got.IsEmpty()

			if !reflect.DeepEqual(got, tc.queue) {
				t.Fatalf("Got(%+v) != Want(%+v)", got, tc.queue)
			}

			if !reflect.DeepEqual(tc.expected, isEmpty) {
				t.Fatalf("Got(%+v) != Want(%+v)", isEmpty, tc.expected)
			}
		})
	}
}
