package gollections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsListEmpty(t *testing.T) {
	testCases := map[string]struct {
		list *List
		want bool
	}{
		"empty list": {
			list: nil,
			want: true,
		},
		"empty head": {
			list: &List{head: nil},
			want: true,
		},
		"non empty": {
			list: &List{head: &Node{2, nil}},
			want: false,
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			got := tc.list.IsEmpty()

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("Got(%+v) != Want(%+v)", got, tc.want)
			}
		})
	}
}

func TestListLength(t *testing.T) {
	testCases := map[string]struct {
		list *List
		want int
	}{
		"empty list": {
			list: nil,
			want: 0,
		},
		"empty head": {
			list: &List{head: nil},
			want: 0,
		},
		"non empty": {
			list: &List{head: &Node{2, nil}},
			want: 1,
		},
		"non empty longer": {
			list: &List{head: &Node{2, &Node{3, nil}}},
			want: 2,
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			got := tc.list.Length()

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("Got(%+v) != Want(%+v)", got, tc.want)
			}
		})
	}
}

func TestListFront(t *testing.T) {
	testCases := map[string]struct {
		list *List
		want *Node
	}{
		"empty list": {
			list: nil,
			want: nil,
		},
		"empty head": {
			list: &List{head: nil},
			want: nil,
		},
		"non empty": {
			list: &List{head: &Node{2, nil}},
			want: &Node{2, nil},
		},
		"non empty longer": {
			list: &List{head: &Node{2, &Node{3, nil}}},
			want: &Node{2, &Node{3, nil}},
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			got := tc.list.Front()

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("Got(%+v) != Want(%+v)", got, tc.want)
			}
		})
	}
}

func TestListBack(t *testing.T) {
	testCases := map[string]struct {
		list *List
		want *Node
	}{
		"empty list": {
			list: nil,
			want: nil,
		},
		"empty head": {
			list: &List{head: nil},
			want: nil,
		},
		"non empty": {
			list: &List{head: &Node{2, nil}},
			want: &Node{2, nil},
		},
		"non empty longer": {
			list: &List{head: &Node{2, &Node{3, nil}}},
			want: &Node{3, nil},
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			got := tc.list.Back()

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("Got(%+v) != Want(%+v)", got, tc.want)
			}
		})
	}
}

func TestListHasCycle(t *testing.T) {
	head := &Node{1, &Node{2, &Node{3, nil}}}
	head.Next.Next.Next = head

	testCases := map[string]struct {
		head *Node
		list *List
		want bool
	}{
		"empty list": {
			list: nil,
			want: false,
		},
		"empty head": {
			list: &List{head: nil},
			want: false,
		},
		"non empty non cycle": {
			list: &List{head: &Node{2, nil}},
			want: false,
		},
		"non empty longer no cycle": {
			list: &List{head: &Node{2, &Node{3, nil}}},
			want: false,
		},
		"non empty longer with cycle": {
			list: &List{head: head},
			want: true,
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			got := tc.list.HasCycle()

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("Got(%+v) != Want(%+v)", got, tc.want)
			}
		})
	}
}

func TestListPushFront(t *testing.T) {
	testCases := map[string]struct {
		toPush any
		list   *List
		want   *List
	}{
		"empty list": {
			toPush: 2,
			list:   nil,
			want:   nil,
		},
		"empty head": {
			toPush: 3,
			list:   &List{head: nil},
			want:   &List{head: &Node{3, nil}},
		},
		"non empty longer": {
			toPush: 4,
			list:   &List{head: &Node{1, &Node{2, &Node{3, nil}}}},
			want:   &List{head: &Node{4, &Node{1, &Node{2, &Node{3, nil}}}}},
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			tc.list.PushFront(tc.toPush)

			if !reflect.DeepEqual(tc.list, tc.want) {
				t.Fatalf("Got(%+v) != Want(%+v)", tc.list, tc.want)
			}
		})
	}
}

func TestListPushBack(t *testing.T) {
	testCases := map[string]struct {
		toPush any
		list   *List
		want   *List
	}{
		"empty list": {
			toPush: 2,
			list:   nil,
			want:   nil,
		},
		"empty head": {
			toPush: 3,
			list:   &List{head: nil},
			want:   &List{head: &Node{3, nil}},
		},
		"non empty longer": {
			toPush: 4,
			list:   &List{head: &Node{1, &Node{2, &Node{3, nil}}}},
			want:   &List{head: &Node{1, &Node{2, &Node{3, &Node{4, nil}}}}},
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			tc.list.PushBack(tc.toPush)

			if !reflect.DeepEqual(tc.list, tc.want) {
				t.Fatalf("Got(%+v) != Want(%+v)", tc.list, tc.want)
			}
		})
	}
}

func TestListPopFront(t *testing.T) {
	testCases := map[string]struct {
		list      *List
		wantElem  any
		wantList  *List
		wantError error
	}{
		"empty list": {
			list:      nil,
			wantList:  nil,
			wantError: fmt.Errorf("can not pop empty list"),
		},
		"empty head": {
			list:      &List{head: nil},
			wantList:  &List{head: nil},
			wantError: fmt.Errorf("can not pop empty list"),
		},
		"one element": {
			list:     &List{head: &Node{1, nil}},
			wantList: &List{head: nil},
			wantElem: 1,
		},
		"more elements": {
			list:     &List{head: &Node{1, &Node{2, &Node{3, nil}}}},
			wantList: &List{head: &Node{2, &Node{3, nil}}},
			wantElem: 1,
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			gotElem, gotErr := tc.list.PopFront()

			if !reflect.DeepEqual(tc.wantList, tc.list) {
				t.Fatalf("Got(%+v) != Want(%+v)", tc.list, tc.wantList)
			}
			if !reflect.DeepEqual(tc.wantElem, gotElem) {
				t.Fatalf("Pop got elem %v, but wants %v", gotElem, tc.wantElem)
			}
			compareErrors(t, tc.wantError, gotErr)
		})
	}
}

func TestListPopBack(t *testing.T) {
	testCases := map[string]struct {
		list      *List
		wantElem  any
		wantList  *List
		wantError error
	}{
		"empty list": {
			list:      nil,
			wantList:  nil,
			wantError: fmt.Errorf("can not pop empty list"),
		},
		"empty head": {
			list:      &List{head: nil},
			wantList:  &List{head: nil},
			wantError: fmt.Errorf("can not pop empty list"),
		},
		"one element": {
			list:     &List{head: &Node{1, nil}},
			wantList: &List{head: nil},
			wantElem: 1,
		},
		"more elements": {
			list:     &List{head: &Node{1, &Node{2, &Node{3, nil}}}},
			wantList: &List{head: &Node{1, &Node{2, nil}}},
			wantElem: 3,
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			gotElem, gotErr := tc.list.PopBack()

			if !reflect.DeepEqual(tc.wantList, tc.list) {
				t.Fatalf("Got(%+v) != Want(%+v)", tc.list, tc.wantList)
			}
			if !reflect.DeepEqual(tc.wantElem, gotElem) {
				t.Fatalf("Pop got elem %v, but wants %v", gotElem, tc.wantElem)
			}
			compareErrors(t, tc.wantError, gotErr)
		})
	}
}

func TestListReverse(t *testing.T) {
	testCases := map[string]struct {
		list     *List
		wantList *List
	}{
		"empty list": {
			list:     nil,
			wantList: nil,
		},
		"empty head": {
			list:     &List{head: nil},
			wantList: &List{head: nil},
		},
		"one element": {
			list:     &List{head: &Node{1, nil}},
			wantList: &List{head: &Node{1, nil}},
		},
		"more elements": {
			list:     &List{head: &Node{1, &Node{2, &Node{3, nil}}}},
			wantList: &List{head: &Node{3, &Node{2, &Node{1, nil}}}},
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {

			tc.list.Reverse()

			if !reflect.DeepEqual(tc.wantList, tc.list) {
				t.Fatalf("Got(%+v) != Want(%+v)", tc.list, tc.wantList)
			}
		})
	}
}

func TestString(t *testing.T) {
	testCases := map[string]struct {
		list *List
		want string
	}{
		"empty list": {
			list: nil,
			want: "nil",
		},
		"empty head": {
			list: &List{head: nil},
			want: "nil",
		},
		"one element": {
			list: &List{head: &Node{1, nil}},
			want: "1->nil",
		},
		"more elements": {
			list: &List{head: &Node{1, &Node{2, &Node{3, nil}}}},
			want: "1->2->3->nil",
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			got := tc.list.String()

			if got != tc.want {
				t.Fatalf("Expected %v, but got %v", tc.want, got)
			}
		})
	}
}

func TestListEqual(t *testing.T) {
	testCases := map[string]struct {
		list1 *List
		list2 *List
		equal bool
	}{
		"nil vs nil": {
			list1: nil,
			list2: nil,
			equal: true,
		},
		"nil vs head nil": {
			list1: nil,
			list2: &List{head: nil},
			equal: false,
		},
		"head nil vs head nil": {
			list1: &List{head: nil},
			list2: &List{head: nil},
			equal: true,
		},
		"non nil vs non nil equal": {
			list1: &List{head: &Node{1, &Node{2, &Node{3, nil}}}},
			list2: &List{head: &Node{1, &Node{2, &Node{3, nil}}}},
			equal: true,
		},
		"non nil vs non nil non equal": {
			list1: &List{head: &Node{1, &Node{2, &Node{3, nil}}}},
			list2: &List{head: &Node{1, &Node{2, &Node{4, nil}}}},
			equal: false,
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			equal := tc.list1.Equals(tc.list2)

			if equal != tc.equal {
				t.Fatalf("Got %v, want %v", equal, tc.equal)
			}
		})
	}
}
