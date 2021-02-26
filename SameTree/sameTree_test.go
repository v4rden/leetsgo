package SameTree

import "testing"

func TestIsSameTreeTableDriven(t *testing.T) {
	tests := []struct {
		id         string
		firstNode  *TreeNode
		secondNode *TreeNode
		result     bool
	}{
		{"T1",
			getNode(1, getNode(2, nil, nil), getNode(3, nil, nil)),
			getNode(1, getNode(2, nil, nil), getNode(3, nil, nil)),
			true},
	}

	for _, test := range tests {
		t.Run(test.id, func(t *testing.T) {
			ans := isSameTree(test.firstNode, test.secondNode)
			if ans != test.result {
				t.Errorf("got %v wan't %v", ans, test.result)
			}
		})
	}
}

func getNode(val int, left *TreeNode, right *TreeNode) *TreeNode {
	node := TreeNode{Val: val, Left: left, Right: right}
	return &node
}
