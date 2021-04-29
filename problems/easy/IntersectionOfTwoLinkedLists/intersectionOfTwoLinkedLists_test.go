package IntersectionOfTwoLinkedLists

import (
	"leetsgo/internal/linkedlist"
	"testing"
)

func TestIntersectionOfTwoLinkedListsTableDriven(t *testing.T) {
	headA := linkedlist.NewNode(1, linkedlist.NewNode(2, linkedlist.NewNode(3, linkedlist.NewNode(4, nil))))
	headB := linkedlist.NewNode(11, headA.Next.Next)
	tests := []struct {
		id     string
		inputA *ListNode
		inputB *ListNode
		result *ListNode
	}{
		{"T1", headA, headB, headA.Next.Next},
	}

	for _, test := range tests {
		t.Run(test.id, func(t *testing.T) {
			ans := getIntersectionNode(test.inputA, test.inputB)
			if ans != test.result {
				t.Errorf("got %v wan't %v", ans, test.result)
			}
		})
	}
}
