package MaximumSubArray

import (
	"fmt"
	"testing"
)

func TestMaxSubArrayTableDriven(t *testing.T) {
	var tests = []struct {
		n      []int
		result int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{-1}, -1},
		{[]int{0}, 0},
		{[]int{1}, 1},
		{[]int{-10000}, -10000},
		{[]int{1, 2, 3}, 6},
		{[]int{10, 20, 30}, 60},
		{[]int{-1, -2, -3}, -1},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d", tt.n)
		t.Run(testName, func(t *testing.T) {
			ans := maxSubArray(tt.n)
			if ans != tt.result {
				t.Errorf("got %d, want %d", ans, tt.result)
			}
		})
	}
}
