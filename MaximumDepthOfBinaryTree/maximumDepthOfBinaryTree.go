package MaximumDepthOfBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftCount := maxDepth(root.Left)
	rightCount := maxDepth(root.Right)

	if leftCount > rightCount {
		leftCount++
		return leftCount
	}

	rightCount++
	return rightCount
}
