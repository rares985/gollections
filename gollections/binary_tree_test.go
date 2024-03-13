package gollections

import (
	"reflect"
	"testing"
)

func TestInorder(t *testing.T) {
	testCases := map[string]struct {
		tree *BinaryTree
		want []any
	}{
		"nil tree": {
			tree: nil,
			want: []any{},
		},
		"empty tree": {
			tree: &BinaryTree{Root: nil},
			want: []any{},
		},
		"non empty tree": {
			tree: &BinaryTree{Root: &BinaryTreeNode{2,
				&BinaryTreeNode{3, nil, nil},
				&BinaryTreeNode{4, nil, nil}}},
			want: []any{3, 2, 4},
		},
		"more complex tree": {
			tree: &BinaryTree{Root: &BinaryTreeNode{1,
				&BinaryTreeNode{2,
					&BinaryTreeNode{4, nil, nil},
					&BinaryTreeNode{5, nil, nil},
				},
				&BinaryTreeNode{3, nil, nil},
			}},
			want: []any{4, 2, 5, 1, 3},
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			got := tc.tree.Inorder()

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("Got(%+v) != Want(%+v)", got, tc.want)
			}
		})
	}
}

func TestPreorder(t *testing.T) {
	testCases := map[string]struct {
		tree *BinaryTree
		want []any
	}{
		"nil tree": {
			tree: nil,
			want: []any{},
		},
		"empty tree": {
			tree: &BinaryTree{Root: nil},
			want: []any{},
		},
		"non empty tree": {
			tree: &BinaryTree{Root: &BinaryTreeNode{2,
				&BinaryTreeNode{3, nil, nil},
				&BinaryTreeNode{4, nil, nil}}},
			want: []any{2, 3, 4},
		},
		"more complex tree": {
			tree: &BinaryTree{Root: &BinaryTreeNode{1,
				&BinaryTreeNode{2,
					&BinaryTreeNode{4, nil, nil},
					&BinaryTreeNode{5, nil, nil},
				},
				&BinaryTreeNode{3, nil, nil},
			}},
			want: []any{1, 2, 4, 5, 3},
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			got := tc.tree.Preorder()

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("Got(%+v) != Want(%+v)", got, tc.want)
			}
		})
	}
}

func TestPostorder(t *testing.T) {
	testCases := map[string]struct {
		tree *BinaryTree
		want []any
	}{
		"nil tree": {
			tree: nil,
			want: []any{},
		},
		"empty tree": {
			tree: &BinaryTree{Root: nil},
			want: []any{},
		},
		"non empty tree": {
			tree: &BinaryTree{Root: &BinaryTreeNode{2,
				&BinaryTreeNode{3, nil, nil},
				&BinaryTreeNode{4, nil, nil}}},
			want: []any{3, 4, 2},
		},
		"more complex tree": {
			tree: &BinaryTree{Root: &BinaryTreeNode{1,
				&BinaryTreeNode{2,
					&BinaryTreeNode{4, nil, nil},
					&BinaryTreeNode{5, nil, nil},
				},
				&BinaryTreeNode{3, nil, nil},
			}},
			want: []any{4, 5, 2, 3, 1},
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			got := tc.tree.Postorder()

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("Got(%+v) != Want(%+v)", got, tc.want)
			}
		})
	}
}
