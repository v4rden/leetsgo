package MaximumDepthOfBinaryTree

import "testing"

func TestMaximumDepthOfBinaryTreeTableDriven(t *testing.T) {
	tests := []struct {
		id        string
		firstNode *TreeNode
		result    int
	}{
		{"T1",
			getNode(1, getNode(2, nil, nil), getNode(2, nil, nil)),
			2},
		{"T2",
			getNode(1, getNode(2, nil, nil), getNode(3, nil, getNode(4, nil, nil))),
			3},
		{"T3",
			getNode(1, nil, nil),
			1},
		{"T4",
			nil,
			0},
	}

	for _, test := range tests {
		t.Run(test.id, func(t *testing.T) {
			ans := maxDepth(test.firstNode)
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
