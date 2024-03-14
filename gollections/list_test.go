package gollections

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	one   = 1
	three = 3
)

func TestIsListEmpty(t *testing.T) {
	testCases := map[string]struct {
		list *List[int]
		want bool
	}{
		"empty list": {
			list: nil,
			want: true,
		},
		"empty head": {
			list: &List[int]{head: nil},
			want: true,
		},
		"non empty": {
			list: &List[int]{head: &ListNode[int]{2, nil}},
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
		list *List[int]
		want int
	}{
		"empty list": {
			list: nil,
			want: 0,
		},
		"empty head": {
			list: &List[int]{head: nil},
			want: 0,
		},
		"non empty": {
			list: &List[int]{head: &ListNode[int]{2, nil}},
			want: 1,
		},
		"non empty longer": {
			list: &List[int]{head: &ListNode[int]{2, &ListNode[int]{3, nil}}},
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
		list *List[int]
		want *ListNode[int]
	}{
		"empty list": {
			list: nil,
			want: nil,
		},
		"empty head": {
			list: &List[int]{head: nil},
			want: nil,
		},
		"non empty": {
			list: &List[int]{head: &ListNode[int]{2, nil}},
			want: &ListNode[int]{2, nil},
		},
		"non empty longer": {
			list: &List[int]{head: &ListNode[int]{2, &ListNode[int]{3, nil}}},
			want: &ListNode[int]{2, &ListNode[int]{3, nil}},
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
		list *List[int]
		want *ListNode[int]
	}{
		"empty list": {
			list: nil,
			want: nil,
		},
		"empty head": {
			list: &List[int]{head: nil},
			want: nil,
		},
		"non empty": {
			list: &List[int]{head: &ListNode[int]{2, nil}},
			want: &ListNode[int]{2, nil},
		},
		"non empty longer": {
			list: &List[int]{head: &ListNode[int]{2, &ListNode[int]{3, nil}}},
			want: &ListNode[int]{3, nil},
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
	head := &ListNode[int]{1, &ListNode[int]{2, &ListNode[int]{3, nil}}}
	head.Next.Next.Next = head

	testCases := map[string]struct {
		head *ListNode[int]
		list *List[int]
		want bool
	}{
		"empty list": {
			list: nil,
			want: false,
		},
		"empty head": {
			list: &List[int]{head: nil},
			want: false,
		},
		"non empty non cycle": {
			list: &List[int]{head: &ListNode[int]{2, nil}},
			want: false,
		},
		"non empty longer no cycle": {
			list: &List[int]{head: &ListNode[int]{2, &ListNode[int]{3, nil}}},
			want: false,
		},
		"non empty longer with cycle": {
			list: &List[int]{head: head},
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
		toPush int
		list   *List[int]
		want   *List[int]
	}{
		"empty list": {
			toPush: 2,
			list:   nil,
			want:   nil,
		},
		"empty head": {
			toPush: 3,
			list:   &List[int]{head: nil},
			want:   &List[int]{head: &ListNode[int]{3, nil}},
		},
		"non empty longer": {
			toPush: 4,
			list:   &List[int]{head: &ListNode[int]{1, &ListNode[int]{2, &ListNode[int]{3, nil}}}},
			want:   &List[int]{head: &ListNode[int]{4, &ListNode[int]{1, &ListNode[int]{2, &ListNode[int]{3, nil}}}}},
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
		toPush int
		list   *List[int]
		want   *List[int]
	}{
		"empty list": {
			toPush: 2,
			list:   nil,
			want:   nil,
		},
		"empty head": {
			toPush: 3,
			list:   &List[int]{head: nil},
			want:   &List[int]{head: &ListNode[int]{3, nil}},
		},
		"non empty longer": {
			toPush: 4,
			list:   &List[int]{head: &ListNode[int]{1, &ListNode[int]{2, &ListNode[int]{3, nil}}}},
			want:   &List[int]{head: &ListNode[int]{1, &ListNode[int]{2, &ListNode[int]{3, &ListNode[int]{4, nil}}}}},
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
		list      *List[int]
		wantElem  *int
		wantList  *List[int]
		wantError error
	}{
		"empty list": {
			list:      nil,
			wantList:  nil,
			wantError: fmt.Errorf("can not pop empty list"),
		},
		"empty head": {
			list:      &List[int]{head: nil},
			wantList:  &List[int]{head: nil},
			wantError: fmt.Errorf("can not pop empty list"),
		},
		"one element": {
			list:     &List[int]{head: &ListNode[int]{1, nil}},
			wantList: &List[int]{head: nil},
			wantElem: &one,
		},
		"more elements": {
			list:     &List[int]{head: &ListNode[int]{1, &ListNode[int]{2, &ListNode[int]{3, nil}}}},
			wantList: &List[int]{head: &ListNode[int]{2, &ListNode[int]{3, nil}}},
			wantElem: &one,
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
		list      *List[int]
		wantElem  *int
		wantList  *List[int]
		wantError error
	}{
		"empty list": {
			list:      nil,
			wantList:  nil,
			wantError: fmt.Errorf("can not pop empty list"),
		},
		"empty head": {
			list:      &List[int]{head: nil},
			wantList:  &List[int]{head: nil},
			wantError: fmt.Errorf("can not pop empty list"),
		},
		"one element": {
			list:     &List[int]{head: &ListNode[int]{1, nil}},
			wantList: &List[int]{head: nil},
			wantElem: &one,
		},
		"more elements": {
			list:     &List[int]{head: &ListNode[int]{1, &ListNode[int]{2, &ListNode[int]{3, nil}}}},
			wantList: &List[int]{head: &ListNode[int]{1, &ListNode[int]{2, nil}}},
			wantElem: &three,
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
		list     *List[int]
		wantList *List[int]
	}{
		"empty list": {
			list:     nil,
			wantList: nil,
		},
		"empty head": {
			list:     &List[int]{head: nil},
			wantList: &List[int]{head: nil},
		},
		"one element": {
			list:     &List[int]{head: &ListNode[int]{1, nil}},
			wantList: &List[int]{head: &ListNode[int]{1, nil}},
		},
		"more elements": {
			list:     &List[int]{head: &ListNode[int]{1, &ListNode[int]{2, &ListNode[int]{3, nil}}}},
			wantList: &List[int]{head: &ListNode[int]{3, &ListNode[int]{2, &ListNode[int]{1, nil}}}},
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
		list *List[int]
		want string
	}{
		"empty list": {
			list: nil,
			want: "nil",
		},
		"empty head": {
			list: &List[int]{head: nil},
			want: "nil",
		},
		"one element": {
			list: &List[int]{head: &ListNode[int]{1, nil}},
			want: "1->nil",
		},
		"more elements": {
			list: &List[int]{head: &ListNode[int]{1, &ListNode[int]{2, &ListNode[int]{3, nil}}}},
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
		list1 *List[int]
		list2 *List[int]
		equal bool
	}{
		"nil vs nil": {
			list1: nil,
			list2: nil,
			equal: true,
		},
		"nil vs head nil": {
			list1: nil,
			list2: &List[int]{head: nil},
			equal: false,
		},
		"head nil vs head nil": {
			list1: &List[int]{head: nil},
			list2: &List[int]{head: nil},
			equal: true,
		},
		"non nil vs non nil equal": {
			list1: &List[int]{head: &ListNode[int]{1, &ListNode[int]{2, &ListNode[int]{3, nil}}}},
			list2: &List[int]{head: &ListNode[int]{1, &ListNode[int]{2, &ListNode[int]{3, nil}}}},
			equal: true,
		},
		"non nil vs non nil non equal": {
			list1: &List[int]{head: &ListNode[int]{1, &ListNode[int]{2, &ListNode[int]{3, nil}}}},
			list2: &List[int]{head: &ListNode[int]{1, &ListNode[int]{2, &ListNode[int]{4, nil}}}},
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
