package gollections

type Less[T any] func(a, b T) bool

type BinarySearchTree[T comparable] struct {
	Root *BinaryTree[T]
	less Less[T]
}

func (bst *BinarySearchTree[T]) Find(value T) *BinaryTree[T] {
	if bst == nil || bst.Root == nil {
		return nil
	}
	return bst.findHelper(bst.Root, value)
}

func (bst *BinarySearchTree[T]) findHelper(node *BinaryTree[T], value T) *BinaryTree[T] {
	if node == nil || node.Val == value {
		return node
	}

	if bst.less(value, node.Val) {
		return bst.findHelper(node.Left, value)
	}

	return bst.findHelper(node.Right, value)
}

func (bst *BinarySearchTree[T]) Insert(value T) {
	bst.Root = bst.insertHelper(bst.Root, value)
}

func (bst *BinarySearchTree[T]) insertHelper(node *BinaryTree[T], value T) *BinaryTree[T] {
	if node == nil {
		return &BinaryTree[T]{Val: value}
	}
	if bst.less(value, node.Val) {
		node.Left = bst.insertHelper(node.Left, value)
	} else {
		node.Right = bst.insertHelper(node.Right, value)
	}
	return node
}

func (bst *BinarySearchTree[T]) Delete(value T) {
	if bst == nil || bst.Root == nil {
		return
	}
	bst.Root = bst.deleteHelper(bst.Root, value)
}

func (bst *BinarySearchTree[T]) deleteHelper(node *BinaryTree[T], value T) *BinaryTree[T] {
	if node == nil {
		return nil
	}

	if value == node.Val {
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}
		minRight := bst.findMin(node.Right)
		node.Val = minRight.Val
		node.Right = bst.deleteHelper(node.Right, minRight.Val)
	} else if bst.less(value, node.Val) {
		node.Left = bst.deleteHelper(node.Left, value)
	} else {
		node.Right = bst.deleteHelper(node.Right, value)
	}
	return node
}

func (bst *BinarySearchTree[T]) findMin(node *BinaryTree[T]) *BinaryTree[T] {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func (bst *BinarySearchTree[T]) Equals(other *BinarySearchTree[T]) bool {
	if bst == nil && other == nil {
		return true
	}
	if bst == nil && other != nil || bst != nil && other == nil {
		return false
	}

	if bst.Root == nil && other.Root == nil {
		return true
	}

	if bst.Root != nil && other.Root == nil || bst.Root == nil && other.Root != nil {
		return false
	}

	return bst.Root.Equals(other.Root)
}
