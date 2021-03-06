package PathSum

import (
	"leetsgo/internal/binarytree"
	"testing"
)

func TestPathSumTableDriven(t *testing.T) {
	var tests = []struct {
		id     string
		input  *TreeNode
		sum    int
		result bool
	}{
		{"T1",
			binarytree.ArrayToBinaryTree([]int{5, 4, 8, 11, binarytree.Null, 13, 4, 7, 2, binarytree.Null, binarytree.Null, binarytree.Null, 1}),
			22,
			true,
		},
		{"T2",
			binarytree.ArrayToBinaryTree([]int{1, 2, 3}),
			5,
			false,
		},
		{"T3",
			binarytree.ArrayToBinaryTree([]int{1, 2}),
			50,
			false,
		},
		{"T4",
			binarytree.ArrayToBinaryTree([]int{}),
			0,
			false,
		},
	}

	for _, test := range tests {
		t.Run(test.id, func(t *testing.T) {
			ans := hasPathSum(test.input, test.sum)
			if ans != test.result {
				t.Errorf("want %v got %v", test.result, ans)
			}
		})
	}
}
