package gollections

import (
	"reflect"
	"testing"
)

func TestIsLeaf(t *testing.T) {
	testCases := map[string]struct {
		tree *BinaryTree[int]
		want bool
	}{
		"nil tree": {
			tree: nil,
			want: false,
		},
		"leaf": {
			tree: &BinaryTree[int]{2, nil, nil},
			want: true,
		},
		"non leaf": {
			tree: &BinaryTree[int]{1,
				&BinaryTree[int]{2,
					&BinaryTree[int]{4, nil, nil},
					&BinaryTree[int]{5, nil, nil},
				},
				&BinaryTree[int]{3, nil, nil},
			},
			want: false,
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			got := tc.tree.IsLeaf()

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("Got(%+v) != Want(%+v)", got, tc.want)
			}
		})
	}
}

func TestInorder(t *testing.T) {
	testCases := map[string]struct {
		tree *BinaryTree[int]
		want []int
	}{
		"nil tree": {
			tree: nil,
			want: []int{},
		},
		"non empty tree": {
			tree: &BinaryTree[int]{2,
				&BinaryTree[int]{3, nil, nil},
				&BinaryTree[int]{4, nil, nil}},
			want: []int{3, 2, 4},
		},
		"more complex tree": {
			tree: &BinaryTree[int]{1,
				&BinaryTree[int]{2,
					&BinaryTree[int]{4, nil, nil},
					&BinaryTree[int]{5, nil, nil},
				},
				&BinaryTree[int]{3, nil, nil},
			},
			want: []int{4, 2, 5, 1, 3},
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
		tree *BinaryTree[int]
		want []int
	}{
		"nil tree": {
			tree: nil,
			want: []int{},
		},
		"non empty tree": {
			tree: &BinaryTree[int]{2,
				&BinaryTree[int]{3, nil, nil},
				&BinaryTree[int]{4, nil, nil}},
			want: []int{2, 3, 4},
		},
		"more complex tree": {
			tree: &BinaryTree[int]{1,
				&BinaryTree[int]{2,
					&BinaryTree[int]{4, nil, nil},
					&BinaryTree[int]{5, nil, nil},
				},
				&BinaryTree[int]{3, nil, nil},
			},
			want: []int{1, 2, 4, 5, 3},
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
		tree *BinaryTree[int]
		want []int
	}{
		"nil tree": {
			tree: nil,
			want: []int{},
		},
		"non empty tree": {
			tree: &BinaryTree[int]{2,
				&BinaryTree[int]{3, nil, nil},
				&BinaryTree[int]{4, nil, nil}},
			want: []int{3, 4, 2},
		},
		"more complex tree": {
			tree: &BinaryTree[int]{1,
				&BinaryTree[int]{2,
					&BinaryTree[int]{4, nil, nil},
					&BinaryTree[int]{5, nil, nil},
				},
				&BinaryTree[int]{3, nil, nil},
			},
			want: []int{4, 5, 2, 3, 1},
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

func TestDepth(t *testing.T) {
	testCases := map[string]struct {
		tree *BinaryTree[int]
		want int
	}{
		"nil tree": {
			tree: nil,
			want: 0,
		},
		"non empty tree": {
			tree: &BinaryTree[int]{2,
				&BinaryTree[int]{3, nil, nil},
				&BinaryTree[int]{4, nil, nil}},
			want: 2,
		},
		"more complex tree": {
			tree: &BinaryTree[int]{1,
				&BinaryTree[int]{2,
					&BinaryTree[int]{4, nil, nil},
					&BinaryTree[int]{5, nil, nil},
				},
				&BinaryTree[int]{3, nil, nil},
			},
			want: 3,
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			got := tc.tree.Depth()

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("Got(%+v) != Want(%+v)", got, tc.want)
			}
		})
	}
}

func TestNbNodes(t *testing.T) {
	testCases := map[string]struct {
		tree *BinaryTree[int]
		want int
	}{
		"nil tree": {
			tree: nil,
			want: 0,
		},
		"non empty tree": {
			tree: &BinaryTree[int]{2,
				&BinaryTree[int]{3, nil, nil},
				&BinaryTree[int]{4, nil, nil}},
			want: 3,
		},
		"more complex tree": {
			tree: &BinaryTree[int]{1,
				&BinaryTree[int]{2,
					&BinaryTree[int]{4, nil, nil},
					&BinaryTree[int]{5, nil, nil},
				},
				&BinaryTree[int]{3, nil, nil},
			},
			want: 5,
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			got := tc.tree.NbNodes()

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("Got(%+v) != Want(%+v)", got, tc.want)
			}
		})
	}
}

func TestHasLeft(t *testing.T) {
	testCases := map[string]struct {
		node *BinaryTree[int]
		want bool
	}{
		"nil node": {
			node: nil,
			want: false,
		},
		"has right": {
			node: &BinaryTree[int]{2, nil, &BinaryTree[int]{3, nil, nil}},
			want: false,
		},
		"has left": {
			node: &BinaryTree[int]{2, &BinaryTree[int]{3, nil, nil}, nil},
			want: true,
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			got := tc.node.HasLeft()

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("Got(%+v) != Want(%+v)", got, tc.want)
			}
		})
	}
}

func TestHasRight(t *testing.T) {
	testCases := map[string]struct {
		node *BinaryTree[int]
		want bool
	}{
		"nil node": {
			node: nil,
			want: false,
		},
		"has right": {
			node: &BinaryTree[int]{2, nil, &BinaryTree[int]{3, nil, nil}},
			want: true,
		},
		"has left": {
			node: &BinaryTree[int]{2, &BinaryTree[int]{3, nil, nil}, nil},
			want: false,
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			got := tc.node.HasRight()

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("Got(%+v) != Want(%+v)", got, tc.want)
			}
		})
	}
}

func TestInsertLeft(t *testing.T) {
	testCases := map[string]struct {
		tree     *BinaryTree[int]
		wantTree *BinaryTree[int]
		toInsert int
	}{
		"nil tree": {
			tree:     nil,
			wantTree: nil,
		},
		"non empty tree": {
			toInsert: 3,
			tree:     &BinaryTree[int]{2, nil, nil},
			wantTree: &BinaryTree[int]{2, &BinaryTree[int]{3, nil, nil}, nil},
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			tc.tree.InsertLeft(tc.toInsert)

			if !reflect.DeepEqual(tc.tree, tc.wantTree) {
				t.Fatalf("Got(%+v) != Want(%+v)", tc.tree, tc.wantTree)
			}
		})
	}
}
func TestInsertRight(t *testing.T) {
	testCases := map[string]struct {
		tree     *BinaryTree[int]
		wantTree *BinaryTree[int]
		toInsert int
	}{
		"nil tree": {
			tree:     nil,
			wantTree: nil,
		},
		"non empty tree": {
			toInsert: 3,
			tree:     &BinaryTree[int]{2, nil, nil},
			wantTree: &BinaryTree[int]{2, nil, &BinaryTree[int]{3, nil, nil}},
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			tc.tree.InsertRight(tc.toInsert)

			if !reflect.DeepEqual(tc.tree, tc.wantTree) {
				t.Fatalf("Got(%+v) != Want(%+v)", tc.tree, tc.wantTree)
			}
		})
	}
}

func TestInvert(t *testing.T) {
	testCases := map[string]struct {
		tree     *BinaryTree[int]
		wantTree *BinaryTree[int]
		toInsert int
	}{
		"nil tree": {
			tree:     nil,
			wantTree: nil,
		},
		"non empty simple tree": {
			toInsert: 3,
			tree:     &BinaryTree[int]{2, &BinaryTree[int]{3, nil, nil}, &BinaryTree[int]{4, nil, nil}},
			wantTree: &BinaryTree[int]{2, &BinaryTree[int]{4, nil, nil}, &BinaryTree[int]{3, nil, nil}},
		},
		"non empty complex tree": {
			toInsert: 3,
			tree: &BinaryTree[int]{1,
				&BinaryTree[int]{2,
					&BinaryTree[int]{4, nil, nil},
					&BinaryTree[int]{5, nil, nil},
				},
				&BinaryTree[int]{3, nil, nil},
			},
			wantTree: &BinaryTree[int]{1,
				&BinaryTree[int]{3, nil, nil},
				&BinaryTree[int]{2,
					&BinaryTree[int]{5, nil, nil},
					&BinaryTree[int]{4, nil, nil},
				},
			},
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			tc.tree.Invert()

			if !reflect.DeepEqual(tc.tree, tc.wantTree) {
				t.Fatalf("Got(%+v) != Want(%+v)", tc.tree, tc.wantTree)
			}
		})
	}
}

func TestEquals(t *testing.T) {
	testCases := map[string]struct {
		tree1 *BinaryTree[int]
		tree2 *BinaryTree[int]
		want  bool
	}{
		"nil vs nil": {
			tree1: nil,
			tree2: nil,
			want:  true,
		},
		"nil vs non nil": {
			tree1: nil,
			tree2: &BinaryTree[int]{2, nil, nil},
			want:  false,
		},
		"non equal": {
			tree1: &BinaryTree[int]{4, nil, nil},
			tree2: &BinaryTree[int]{5, nil, nil},
		},
		"equal complex": {
			tree1: &BinaryTree[int]{5,
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
			tree2: &BinaryTree[int]{5,
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
			want: true,
		},
	}

	for description, tc := range testCases {
		t.Run(description, func(t *testing.T) {
			got := tc.tree1.Equals(tc.tree2)

			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("Got(%+v) != Want(%+v)", got, tc.want)
			}
		})
	}
}
