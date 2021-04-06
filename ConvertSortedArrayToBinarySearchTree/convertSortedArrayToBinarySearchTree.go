package ConvertSortedArrayToBinarySearchTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return createNode(nums[0], nil, nil)
	}

	if len(nums) == 2 {
		return createNode(nums[1], sortedArrayToBST(nums[0:1]), nil)
	}

	if len(nums) == 3 {
		return createNode(nums[1], sortedArrayToBST(nums[:1]), sortedArrayToBST(nums[2:]))
	}

	pos := len(nums)/2 + 1

	node := createNode(nums[pos], sortedArrayToBST(nums[:pos]), sortedArrayToBST(nums[pos+1:]))

	return node
}

func createNode(val int, left *TreeNode, right *TreeNode) *TreeNode {
	node := TreeNode{Val: val, Left: left, Right: right}
	return &node
}
