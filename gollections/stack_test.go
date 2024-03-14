package gollections

import (
	"reflect"
	"testing"
)

var (
	two = 2
)

func TestPush(t *testing.T) {
	testCases := map[string]struct {
		elem   int
		before Stack[int]
		after  Stack[int]
	}{
		"push ok": {
			elem:   2,
			before: Stack[int]{elements: []int{}},
			after:  Stack[int]{elements: []int{2}},
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {

			got := tc.before
			got.Push(tc.elem)

			if !reflect.DeepEqual(got, tc.after) {
				t.Fatalf("Got(%+v) != Want(%+v)", got, tc.after)
			}
		})
	}
}

func TestPop(t *testing.T) {
	testCases := map[string]struct {
		expectedPop *int
		before      Stack[int]
		after       Stack[int]
	}{
		"pop empty": {
			expectedPop: nil,
			before:      Stack[int]{},
			after:       Stack[int]{},
		},
		"pop ok": {
			expectedPop: &two,
			before:      Stack[int]{elements: []int{2}},
			after:       Stack[int]{elements: []int{}},
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {

			got := tc.before
			actualPop := got.Pop()

			if !reflect.DeepEqual(got, tc.after) {
				t.Fatalf("Got(%+v) != Want(%+v)", got, tc.after)
			}

			if !reflect.DeepEqual(actualPop, tc.expectedPop) {
				t.Fatalf("Popped(%+v) != Want(%+v)", actualPop, tc.expectedPop)
			}
		})
	}
}

func TestStackPeek(t *testing.T) {
	testCases := map[string]struct {
		expectedPeek *int
		before       Stack[int]
		after        Stack[int]
	}{
		"peek empty": {
			expectedPeek: nil,
			before:       Stack[int]{},
			after:        Stack[int]{},
		},
		"peek ok": {
			expectedPeek: &two,
			before:       Stack[int]{elements: []int{2}},
			after:        Stack[int]{elements: []int{2}},
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {

			got := tc.before
			actualPeek := got.Peek()

			if !reflect.DeepEqual(got, tc.after) {
				t.Fatalf("Got(%+v) != Want(%+v)", got, tc.after)
			}

			if !reflect.DeepEqual(actualPeek, tc.expectedPeek) {
				t.Fatalf("Peeked(%+v) != Want(%+v)", actualPeek, tc.expectedPeek)
			}
		})
	}
}

func TestStackSize(t *testing.T) {
	testCases := map[string]struct {
		want   int
		before Stack[int]
	}{
		"size empty": {
			want:   0,
			before: Stack[int]{},
		},
		"size positive": {
			want:   1,
			before: Stack[int]{elements: []int{2}},
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {

			size := tc.before.Size()

			if !reflect.DeepEqual(size, tc.want) {
				t.Fatalf("Got(%+v) != Want(%+v)", size, tc.want)
			}
		})
	}
}
func TestStackIsEmpty(t *testing.T) {
	testCases := map[string]struct {
		want   bool
		before Stack[int]
	}{
		"size empty": {
			want:   true,
			before: Stack[int]{},
		},
		"size positive": {
			want:   false,
			before: Stack[int]{elements: []int{2}},
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {

			isEmpty := tc.before.IsEmpty()

			if !reflect.DeepEqual(isEmpty, tc.want) {
				t.Fatalf("Got(%+v) != Want(%+v)", isEmpty, tc.want)
			}
		})
	}
}
