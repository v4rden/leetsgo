package MergeSortedArray

import (
	"fmt"
	"testing"
)

func TestMergeSortedArrayTableDriven(t *testing.T) {
	var tests = []struct {
		nums1  []int
		m      int
		nums2  []int
		n      int
		result []int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{1, 2, 3}, 3, []int{1, 1, 2, 2, 3, 3}},
		{[]int{1, 1, 3, 0, 0, 0}, 3, []int{2, 2, 2}, 3, []int{1, 1, 2, 2, 2, 3}},
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{4, 5, 6}, 3, []int{1, 2, 3, 4, 5, 6}},
		{[]int{-3, -1, 2, 0, 0, 0}, 3, []int{-2, 1, 3}, 3, []int{-3, -2, -1, 1, 2, 3}},
		{[]int{1}, 1, []int{}, 0, []int{1}},
		{[]int{0}, 0, []int{1}, 1, []int{1}},
	}

	for i, test := range tests {
		testName := fmt.Sprintf("Runing %d : %v with %v", i, test.nums1, test.nums2)
		t.Run(testName, func(t *testing.T) {
			merge(test.nums1, test.m, test.nums2, test.n)
			if !areSame(test.nums1, test.result) {
				t.Errorf("got %v want %v", test.nums1, test.result)
			}
		})
	}
}

func areSame(left []int, right []int) bool {
	if len(left) != len(right) {
		return false
	}
	for i, el := range left {
		if el != right[i] {
			return false
		}
	}
	return true
}
