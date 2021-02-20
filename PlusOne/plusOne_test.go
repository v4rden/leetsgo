package PlusOne

import (
	"fmt"
	"testing"
)

func TestPlusONeTableDriven(t *testing.T) {
	var tests = []struct {
		n      []int
		result []int
	}{
		{[]int{0}, []int{1}},
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{[]int{9}, []int{1, 0}},
		{[]int{9, 9}, []int{1, 0, 0}},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%v", tt.n)
		t.Run(testName, func(t *testing.T) {
			ans := plusOne(tt.n)
			if !areEqual(ans, tt.result) {
				t.Errorf("got %v, want %v", ans, tt.result)
			}
		})
	}
}

func areEqual(left []int, right []int) bool {
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
