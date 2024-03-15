package gollections

import (
	"errors"
	"reflect"
	"testing"
)

var (
	thirty = 30
	twenty = 20
)

func TestHeapInsert(t *testing.T) {
	less := func(a, b int) bool { return a < b }
	testCases := map[string]struct {
		toPush []int
		before *Heap[int]
		want   *Heap[int]
	}{
		"insert one in empty": {
			toPush: []int{10},
			before: &Heap[int]{less: less, elements: []int{}},
			want:   &Heap[int]{less: less, elements: []int{10}},
		},
		"insert one": {
			toPush: []int{20},
			before: &Heap[int]{less: less, elements: []int{10}},
			want:   &Heap[int]{less: less, elements: []int{20, 10}},
		},
		"insert two": {
			toPush: []int{20, 15},
			before: &Heap[int]{less: less, elements: []int{10}},
			want:   &Heap[int]{less: less, elements: []int{20, 10, 15}},
		},
		"insert three": {
			toPush: []int{20, 15, 30},
			before: &Heap[int]{less: less, elements: []int{10}},
			want:   &Heap[int]{less: less, elements: []int{30, 20, 15, 10}},
		},
		"insert four": {
			toPush: []int{20, 15, 30, 5},
			before: &Heap[int]{less: less, elements: []int{10}},
			want:   &Heap[int]{less: less, elements: []int{30, 20, 15, 10, 5}},
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {

			for _, elem := range tc.toPush {
				tc.before.Insert(elem)
			}

			if !tc.before.Equals(tc.want) {
				t.Fatalf("Got(%+v) != Want(%+v)", tc.before, tc.want)
			}
		})
	}
}

func TestHeapSize(t *testing.T) {
	less := func(a, b int) bool { return a < b }
	testCases := map[string]struct {
		heap *Heap[int]
		want int
	}{
		"empty heap": {
			want: 0,
			heap: &Heap[int]{less: less, elements: []int{}},
		},
		"non-empty heap": {
			want: 3,
			heap: &Heap[int]{less: less, elements: []int{1, 2, 3}},
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {

			got := tc.heap.Size()
			if got != tc.want {
				t.Fatalf("Got(%+v) != Want(%+v)", got, tc.want)
			}
		})
	}
}

func TestHeapIsEmpty(t *testing.T) {
	less := func(a, b int) bool { return a < b }
	testCases := map[string]struct {
		heap *Heap[int]
		want bool
	}{
		"empty heap": {
			want: true,
			heap: &Heap[int]{less: less, elements: []int{}},
		},
		"non-empty heap": {
			want: false,
			heap: &Heap[int]{less: less, elements: []int{1, 2, 3}},
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {

			got := tc.heap.IsEmpty()
			if got != tc.want {
				t.Fatalf("Got(%+v) != Want(%+v)", got, tc.want)
			}
		})
	}
}

func TestHeapEquals(t *testing.T) {
	less := func(a, b int) bool { return a < b }
	testCases := map[string]struct {
		heap1 *Heap[int]
		heap2 *Heap[int]
		want  bool
	}{
		"nil vs nil": {
			want:  true,
			heap1: nil,
			heap2: nil,
		},
		"nil vs non nil": {
			want:  false,
			heap1: nil,
			heap2: &Heap[int]{},
		},
		"non nil vs nil": {
			want:  false,
			heap1: &Heap[int]{},
			heap2: nil,
		},
		"size differs": {
			want:  false,
			heap1: &Heap[int]{less: less, elements: []int{1}},
			heap2: &Heap[int]{less: less, elements: []int{1, 2}},
		},
		"same size diff elems": {
			want:  false,
			heap1: &Heap[int]{less: less, elements: []int{1, 2}},
			heap2: &Heap[int]{less: less, elements: []int{1, 3}},
		},
		"equal": {
			want:  true,
			heap1: &Heap[int]{less: less, elements: []int{1, 2, 3, 4}},
			heap2: &Heap[int]{less: less, elements: []int{1, 2, 3, 4}},
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {

			got := tc.heap1.Equals(tc.heap2)
			if got != tc.want {
				t.Fatalf("Got(%+v) != Want(%+v)", got, tc.want)
			}
		})
	}
}

func TestHeapDelete(t *testing.T) {
	less := func(a, b int) bool { return a < b }
	testCases := map[string]struct {
		expectedResults []*int
		expectedErrors  []error
		before          *Heap[int]
		after           *Heap[int]
	}{
		"empty": {
			expectedResults: []*int{nil},
			expectedErrors:  []error{errors.New("can not delete from empty heap")},
			before:          &Heap[int]{less: less, elements: []int{}},
			after:           &Heap[int]{less: less, elements: []int{}},
		},
		"from one to empty": {
			expectedResults: []*int{&one},
			before:          &Heap[int]{less: less, elements: []int{1}},
			after:           &Heap[int]{less: less, elements: []int{}},
		},
		"multi delete": {
			expectedResults: []*int{&thirty, &twenty},
			before:          &Heap[int]{less: less, elements: []int{30, 20, 15, 10, 5}},
			after:           &Heap[int]{less: less, elements: []int{15, 10, 5}},
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			for idx, wantResult := range tc.expectedResults {
				gotResult, gotErr := tc.before.Delete()

				if !reflect.DeepEqual(gotResult, wantResult) {
					t.Fatalf("Got(%+v) != Want(%+v)", gotResult, wantResult)
				}

				var wantErr error
				wantErr = nil
				if idx < len(tc.expectedErrors) {
					wantErr = tc.expectedErrors[idx]
				}
				compareErrors(t, wantErr, gotErr)
			}

			if !tc.before.Equals(tc.after) {
				t.Fatalf("Expected heap (%v) but have (%v)", tc.after, tc.before)
			}
		})
	}
}
