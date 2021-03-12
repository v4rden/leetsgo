package PascalsTriangle

import "testing"

func TestPascalsTriangleTableDriven(t *testing.T) {
	var tests = []struct {
		id     string
		input  int
		result [][]int
	}{
		{
			"T1",
			5,
			[][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
		{
			"T2",
			1,
			[][]int{{1}},
		},
	}
	for _, test := range tests {
		t.Run(test.id, func(t *testing.T) {
			ans := generate(test.input)
			if !areSame(ans, test.result) {
				t.Errorf("want %v got %v", test.result, ans)
			}
		})
	}
}

func areSame(left [][]int, right [][]int) bool {
	if len(left) != len(right) {
		return false
	}
	for i := 0; i < len(left); i++ {
		if len(left[i]) != len(right[i]) {
			return false
		}
		for j := 0; j < len(left[i]); j++ {
			if left[i][j] != right[i][j] {
				return false
			}
		}
	}
	return true
}
