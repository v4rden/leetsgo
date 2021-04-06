package LinkedListCycle

import (
	"leetsgo/internal/linkedlist"
	"testing"
)

func TestLinkedListCycleTableDriven(t *testing.T) {
	tests := []struct {
		id     string
		input  *ListNode
		result bool
	}{
		{"T1", getLinked(), true},
		{"T2", linkedlist.NewNode(1, nil), false},
	}

	for _, test := range tests {
		t.Run(test.id, func(t *testing.T) {
			ans := hasCycle(test.input)
			if ans != test.result {
				t.Errorf("got %v wan't %v", ans, test.result)
			}
		})
	}
}

func getLinked() *linkedlist.ListNode {
	n := linkedlist.NewNode(3, linkedlist.NewNode(2, linkedlist.NewNode(0, linkedlist.NewNode(-4, nil))))
	n.Next.Next.Next.Next = n.Next
	return n
}
