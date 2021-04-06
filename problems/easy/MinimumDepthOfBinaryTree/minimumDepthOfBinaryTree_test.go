package MinimumDepthOfBinaryTree

import (
	"leetsgo/internal/binarytree"
	"testing"
)

//type TreeNode = internal.TreeNode

func TestMinimumDepthBinaryTreeTableDriven(t *testing.T) {
	var tests = []struct {
		id     string
		input  *TreeNode
		result int
	}{
		{"T1",
			binarytree.ArrayToBinaryTree([]int{3, 9, 20, binarytree.Null, binarytree.Null, 15, 7}),
			2,
		},
		{
			"T2",
			binarytree.ArrayToBinaryTree([]int{2, binarytree.Null, 3, binarytree.Null, 4, binarytree.Null, 5, binarytree.Null, 6}),
			5,
		},
	}

	for _, test := range tests {
		t.Run(test.id, func(t *testing.T) {
			ans := minDepth(test.input)
			if ans != test.result {
				t.Errorf("want %v got %v", test.result, ans)
			}
		})
	}
}
