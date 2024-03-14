package gollections

import (
	"reflect"
	"testing"
)

func TestBSTFind(t *testing.T) {
	less := func(a, b int) bool { return a < b }

	testCases := map[string]struct {
		tree   *BinarySearchTree[int]
		want   *BinaryTree[int]
		search int
	}{
		"nil tree": {
			tree:   nil,
			search: 2,
			want:   nil,
		},
		"non empty tree not found": {
			tree: &BinarySearchTree[int]{
				less: less,
				Root: &BinaryTree[int]{
					2,
					&BinaryTree[int]{3, nil, nil},
					&BinaryTree[int]{4, nil, nil}}},
			search: 10,
			want:   nil,
		},
		"more complex tree found": {
			tree: &BinarySearchTree[int]{
				less: less,
				Root: &BinaryTree[int]{1,
					&BinaryTree[int]{2,
						&BinaryTree[int]{4, nil, nil},
						&BinaryTree[int]{5, nil, nil},
					},
					&BinaryTree[int]{3, nil, nil}},
			},
			search: 3,
			want:   &BinaryTree[int]{3, nil, nil},
		},
		"really complex tree found": {
			tree: &BinarySearchTree[int]{
				less: less,
				Root: &BinaryTree[int]{5,
					&BinaryTree[int]{3,
						&BinaryTree[int]{2,
							&BinaryTree[int]{1, nil, nil},
							nil,
						},
						&BinaryTree[int]{4, nil, nil},
					},
					&BinaryTree[int]{8,
						&BinaryTree[int]{6, nil, nil},
						&BinaryTree[int]{9,
							&BinaryTree[int]{7, nil, nil},
							nil,
						},
					},
				},
			},
			search: 1,
			want:   &BinaryTree[int]{1, nil, nil},
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			got := tc.tree.Find(tc.search)

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("Got(%+v) != Want(%+v)", got, tc.want)
			}
		})
	}
}

func TestBSTDelete(t *testing.T) {
	less := func(a, b int) bool { return a < b }

	testCases := map[string]struct {
		tree     *BinarySearchTree[int]
		wantTree *BinarySearchTree[int]
		delete   int
	}{
		"nil tree": {
			tree:     nil,
			delete:   2,
			wantTree: nil,
		},
		"non empty tree not found": {
			tree: &BinarySearchTree[int]{
				less: less,
				Root: &BinaryTree[int]{
					2,
					&BinaryTree[int]{3, nil, nil},
					&BinaryTree[int]{4, nil, nil}}},
			delete: 10,
			wantTree: &BinarySearchTree[int]{
				less: less,
				Root: &BinaryTree[int]{
					2,
					&BinaryTree[int]{3, nil, nil},
					&BinaryTree[int]{4, nil, nil},
				},
			},
		},
		"more complex tree": {
			tree: &BinarySearchTree[int]{
				less: less,
				Root: &BinaryTree[int]{7,
					&BinaryTree[int]{3,
						&BinaryTree[int]{2,
							&BinaryTree[int]{1, nil, nil},
							nil,
						},
						&BinaryTree[int]{5,
							&BinaryTree[int]{4, nil, nil},
							&BinaryTree[int]{6, nil, nil},
						},
					},
					&BinaryTree[int]{10,
						&BinaryTree[int]{8, nil, nil},
						&BinaryTree[int]{12, nil, nil},
					},
				},
			},
			delete: 3,
			wantTree: &BinarySearchTree[int]{
				less: less,
				Root: &BinaryTree[int]{7,
					&BinaryTree[int]{4,
						&BinaryTree[int]{2,
							&BinaryTree[int]{1, nil, nil},
							nil,
						},
						&BinaryTree[int]{5,
							nil,
							&BinaryTree[int]{6, nil, nil},
						},
					},
					&BinaryTree[int]{10,
						&BinaryTree[int]{8, nil, nil},
						&BinaryTree[int]{12, nil, nil},
					},
				},
			},
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			tc.tree.Delete(tc.delete)

			if !tc.tree.Equals(tc.wantTree) {
				t.Fatalf("Got(%+v) != Want(%+v)", tc.tree, tc.wantTree)
			}
		})
	}
}
