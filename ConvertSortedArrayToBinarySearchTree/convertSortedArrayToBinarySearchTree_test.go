package ConvertSortedArrayToBinarySearchTree

import (
	"testing"
)

func TestConvertSortedArrayToBinarySearchTreeTableDriven(t *testing.T) {
	var tests = []struct {
		id     string
		input  []int
		result *TreeNode
	}{
		{"T1",
			[]int{-10, -3, 0, 5, 9},
			getNode(0, getNode(-3, getNode(-10, nil, nil), nil),
				getNode(5, getNode(9, nil, nil), nil))},
	}

	for _, test := range tests {
		t.Run(test.id, func(t *testing.T) {
			ans := sortedArrayToBST(test.input)
			if ans != test.result {
				t.Errorf("want %v got %v", test.result, ans)
			}
		})
	}
}

func getNode(val int, left *TreeNode, right *TreeNode) *TreeNode {
	node := TreeNode{Val: val, Left: left, Right: right}
	return &node
}
