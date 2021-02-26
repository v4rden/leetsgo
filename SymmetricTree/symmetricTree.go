package SymmetricTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return areSymmetric(root.Left, invert(root.Right))
}

func areSymmetric(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil || left.Val != right.Val {
		return false
	}

	return areSymmetric(left.Left, right.Left) && areSymmetric(left.Right, right.Right)
}

func invert(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}

	invert(node.Left)
	invert(node.Right)
	node.Left, node.Right = node.Right, node.Left

	return node
}
