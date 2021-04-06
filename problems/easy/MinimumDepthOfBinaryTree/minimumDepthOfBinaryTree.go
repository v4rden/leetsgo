package MinimumDepthOfBinaryTree

import (
	"leetsgo/internal/binarytree"
)

type TreeNode = binarytree.TreeNode

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil {
		return minDepth(root.Right) + 1
	}

	if root.Right == nil {
		return minDepth(root.Left) + 1
	}

	leftLength := minDepth(root.Left)
	rightLength := minDepth(root.Right)
	return min(leftLength, rightLength) + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
