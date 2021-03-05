package BalancedBinaryTree

import "testing"

func TestBalancedBinaryTreeTableDriven(t *testing.T) {
	var tests = []struct {
		id     string
		input  *TreeNode
		result bool
	}{
		{"T1",
			getNode(0, getNode(-3, getNode(-10, nil, nil), nil),
				getNode(5, getNode(9, nil, nil), nil)),
			true,
		},
	}

	for _, test := range tests {
		t.Run(test.id, func(t *testing.T) {
			ans := isBalanced(test.input)
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
