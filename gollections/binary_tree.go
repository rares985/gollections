package gollections

type BinaryTree[T comparable] struct {
	Val   T
	Left  *BinaryTree[T]
	Right *BinaryTree[T]
}

func NewBinaryTree[T comparable]() *BinaryTree[T] {
	return &BinaryTree[T]{}
}

func (node *BinaryTree[T]) IsLeaf() bool {
	return node != nil && node.Left == nil && node.Right == nil
}

func (node *BinaryTree[T]) Inorder() []T {
	if node == nil {
		return []T{}
	}

	left := node.Left.Inorder()
	right := node.Right.Inorder()
	return append(append(left, node.Val), right...)
}

func (node *BinaryTree[T]) Preorder() []T {
	if node == nil {
		return []T{}
	}

	left := node.Left.Preorder()
	right := node.Right.Preorder()
	return append(append([]T{node.Val}, left...), right...)
}

func (node *BinaryTree[T]) Postorder() []T {
	if node == nil {
		return []T{}
	}

	left := node.Left.Postorder()
	right := node.Right.Postorder()
	return append(append(left, right...), node.Val)
}

func (node *BinaryTree[T]) Depth() int {
	if node == nil {
		return 0
	}
	ld := node.Left.Depth()
	rd := node.Right.Depth()

	return max(ld, rd) + 1
}

func (node *BinaryTree[T]) NbNodes() int {
	if node == nil {
		return 0
	}
	lnb := node.Left.NbNodes()
	rnb := node.Right.NbNodes()

	return lnb + rnb + 1
}

func (node *BinaryTree[T]) HasLeft() bool {
	return node != nil && node.Left != nil
}

func (node *BinaryTree[T]) HasRight() bool {
	return node != nil && node.Right != nil
}

func (node *BinaryTree[T]) InsertLeft(value T) {
	if node == nil {
		return
	}
	newNode := &BinaryTree[T]{value, nil, nil}
	node.Left = newNode
}

func (node *BinaryTree[T]) InsertRight(value T) {
	if node == nil {
		return
	}
	newNode := &BinaryTree[T]{value, nil, nil}
	node.Right = newNode
}

func (node *BinaryTree[T]) Invert() {
	if node == nil {
		return
	}

	aux := node.Left
	node.Left = node.Right
	node.Right = aux

	node.Left.Invert()
	node.Right.Invert()
}

func (node *BinaryTree[T]) Equals(other *BinaryTree[T]) bool {
	if node == nil && other == nil {
		return true
	}

	if node != nil && other == nil || node == nil && other != nil {
		return false
	}

	equalsHere := (node.Val == other.Val)
	equalsLeft := node.Left.Equals(other.Left)
	equalsRight := node.Right.Equals(other.Right)

	return equalsHere && equalsLeft && equalsRight
}
