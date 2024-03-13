package gollections

type BinaryTreeNode struct {
	Val   any
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func (node *BinaryTreeNode) IsLeaf() bool {
	return node.Left == nil && node.Right == nil
}

func (node *BinaryTreeNode) Inorder() []any {
	if node == nil {
		return []any{}
	}

	left := node.Left.Inorder()
	right := node.Right.Inorder()
	return append(append(left, node.Val), right...)
}

func (node *BinaryTreeNode) Preorder() []any {
	if node == nil {
		return []any{}
	}

	left := node.Left.Preorder()
	right := node.Right.Preorder()
	return append(append([]any{node.Val}, left...), right...)
}

func (node *BinaryTreeNode) Postorder() []any {
	if node == nil {
		return []any{}
	}

	left := node.Left.Postorder()
	right := node.Right.Postorder()
	return append(append(left, right...), node.Val)
}

type BinaryTree struct {
	Root *BinaryTreeNode
}

func (tree *BinaryTree) Inorder() []any {
	if tree == nil {
		return []any{}
	}
	return tree.Root.Inorder()
}

func (tree *BinaryTree) Postorder() []any {
	if tree == nil {
		return []any{}
	}
	return tree.Root.Postorder()
}

func (tree *BinaryTree) Preorder() []any {
	if tree == nil {
		return []any{}
	}
	return tree.Root.Preorder()
}
