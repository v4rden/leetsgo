package PathSum

import "testing"
import "leetsgo/internal"

func TestPathSumTableDriven(t *testing.T) {
	var tests = []struct {
		id     string
		input  *TreeNode
		sum    int
		result bool
	}{
		{"T1",
			internal.ArrayToBinaryTree([]int{5, 4, 8, 11, internal.Null, 13, 4, 7, 2, internal.Null, internal.Null, internal.Null, 1}),
			22,
			true,
		},
		{"T2",
			internal.ArrayToBinaryTree([]int{1, 2, 3}),
			5,
			false,
		},
		{"T3",
			internal.ArrayToBinaryTree([]int{1, 2}),
			50,
			false,
		},
		{"T4",
			internal.ArrayToBinaryTree([]int{}),
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
