package binarytree

const Null = -1 << 63

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewNode(val int, left *TreeNode, right *TreeNode) *TreeNode {
	node := TreeNode{Val: val, Left: left, Right: right}
	return &node
}

func ArrayToBinaryTree(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: nums[0],
	}

	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && nums[i] != Null {
			node.Left = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && nums[i] != Null {
			node.Right = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func SortedArrayToBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return NewNode(nums[0], nil, nil)
	}

	if len(nums) == 2 {
		return NewNode(nums[1], ArrayToBinaryTree(nums[0:1]), nil)
	}

	if len(nums) == 3 {
		return NewNode(nums[1], ArrayToBinaryTree(nums[:1]), ArrayToBinaryTree(nums[2:]))
	}

	pos := len(nums)/2 + 1

	node := NewNode(nums[pos], ArrayToBinaryTree(nums[:pos]), ArrayToBinaryTree(nums[pos+1:]))

	return node
}

func AreEqual(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return AreEqual(left.Left, right.Left) && AreEqual(left.Right, right.Right)
}
