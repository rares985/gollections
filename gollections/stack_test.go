package gollections

import (
	"reflect"
	"testing"
)

func TestPush(t *testing.T) {
	testCases := map[string]struct {
		elem   any
		before Stack
		after  Stack
	}{
		"push ok": {
			elem:   2,
			before: Stack{elements: []any{}},
			after:  Stack{elements: []any{2}},
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
		expectedPop any
		before      Stack
		after       Stack
	}{
		"push ok": {
			expectedPop: 2,
			before:      Stack{elements: []any{2}},
			after:       Stack{elements: []any{}},
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
		expectedPeek any
		before       Stack
		after        Stack
	}{
		"push ok": {
			expectedPeek: 2,
			before:       Stack{elements: []any{2}},
			after:        Stack{elements: []any{2}},
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
