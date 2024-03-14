package gollections

import (
	"reflect"
	"testing"
)

func TestEnqueue(t *testing.T) {
	testCases := map[string]struct {
		elem  int
		queue Queue[int]
		want  Queue[int]
	}{
		"enqueue empty": {
			elem:  2,
			queue: Queue[int]{elements: []int{}},
			want:  Queue[int]{elements: []int{2}},
		},
		"enqueue ok": {
			elem:  2,
			queue: Queue[int]{elements: []int{1, 3}},
			want:  Queue[int]{elements: []int{1, 3, 2}},
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
		elem  *int
		queue Queue[int]
		want  Queue[int]
	}{
		"dequeue empty": {
			elem:  nil,
			queue: Queue[int]{},
			want:  Queue[int]{},
		},
		"dequeue one element": {
			elem:  &two,
			queue: Queue[int]{elements: []int{2}},
			want:  Queue[int]{elements: []int{}},
		},
		"dequeue ok": {
			elem:  &one,
			queue: Queue[int]{elements: []int{1, 3, 2}},
			want:  Queue[int]{elements: []int{3, 2}},
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
		elem  *int
		queue Queue[int]
		want  Queue[int]
	}{
		"peek empty": {
			queue: Queue[int]{},
			want:  Queue[int]{},
		},
		"peek one": {
			elem:  &two,
			queue: Queue[int]{elements: []int{2}},
		},
		"peek ok": {
			elem:  &one,
			queue: Queue[int]{elements: []int{1, 3, 2}},
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
		queue    Queue[int]
	}{
		"is empty": {
			expected: true,
			queue:    Queue[int]{elements: []int{}},
		},
		"is not empty": {
			queue: Queue[int]{elements: []int{1, 3, 2}},
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
func TestSize(t *testing.T) {
	testCases := map[string]struct {
		expected int
		queue    Queue[int]
	}{
		"0": {
			expected: 0,
			queue:    Queue[int]{elements: []int{}},
		},
		"non 0": {
			expected: 3,
			queue:    Queue[int]{elements: []int{1, 3, 2}},
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			size := tc.queue.Size()

			if !reflect.DeepEqual(tc.expected, size) {
				t.Fatalf("Got(%+v) != Want(%+v)", size, tc.expected)
			}
		})
	}
}
