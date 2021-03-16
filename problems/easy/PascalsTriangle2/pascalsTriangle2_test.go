package PascalsTriangle2

import "testing"

func TestPascalsTriangleTableDriven(t *testing.T) {
	var tests = []struct {
		id     string
		input  int
		result []int
	}{
		{
			"T1",
			4,
			[]int{1, 4, 6, 4, 1},
		},
		{
			"T2",
			0,
			[]int{1},
		},
		{
			"T3",
			1,
			[]int{1, 1},
		},
		{
			"T4",
			3,
			[]int{1, 3, 3, 1},
		},
	}
	for _, test := range tests {
		t.Run(test.id, func(t *testing.T) {
			ans := getRow(test.input)
			if !areSame(ans, test.result) {
				t.Errorf("want %v got %v", test.result, ans)
			}
		})
	}
}

func areSame(left []int, right []int) bool {
	if len(left) != len(right) {
		return false
	}
	for i := 0; i < len(left); i++ {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}
