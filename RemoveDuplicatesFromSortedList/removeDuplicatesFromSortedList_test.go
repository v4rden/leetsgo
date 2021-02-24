package RemoveDuplicatesFromSortedList

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicatesTableDriven(t *testing.T) {

	tests := []struct {
		input  []ListNode
		result []ListNode
	}{
		{
			makeList([]int{1, 1, 2, 3, 3, 4, 5, 5}),
			makeList([]int{1, 2, 3, 4, 5}),
		},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%v", test.input)
		t.Run(testName, func(t *testing.T) {
			ans := deleteDuplicates(&test.input[0])
			if !compareValues(ans, &test.result[0]) {

			}
		})
	}
}

func makeList(arr []int) []ListNode {
	result := make([]ListNode, len(arr))
	for i, el := range arr {
		result[i] = ListNode{
			Val: el,
		}
		if i > 0 {
			result[i-1].Next = &result[i]
		}
	}
	return result
}

func compareValues(left *ListNode, right *ListNode) bool {
	for left.Next != nil && right.Next != nil {
		if left.Val != right.Val {
			return false
		}
		left, right = left.Next, right.Next //messes up array, but ima ok with it
	}
	return true
}
