package BalancedBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return getLength(root) > -1
}

func getLength(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftLength := getLength(node.Left)
	rightLength := getLength(node.Right)

	if leftLength < 0 || rightLength < 0 {
		return -1
	}

	diff := leftLength - rightLength

	if diff < -1 || diff > 1 {
		return -1
	}
	return max(leftLength, rightLength) + 1
}

func max(first, second int) int {
	if first > second {
		return first
	}
	return second
}
