package IntersectionOfTwoLinkedLists

//https://leetcode.com/problems/intersection-of-two-linked-lists/

import "leetsgo/internal/linkedlist"

type ListNode = linkedlist.ListNode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	var aLength, bLength int
	a, b := headA, headB

	for a != nil || b != nil {
		if a == b {
			return a
		}
		if a != nil {
			a = a.Next
			aLength++
		}
		if b != nil {
			b = b.Next
			bLength++
		}
	}

	if aLength > bLength {
		for i := 0; i < aLength-bLength; i++ {
			headA = headA.Next
		}
	}

	if bLength > aLength {
		for i := 0; i < bLength-aLength; i++ {
			headB = headB.Next
		}
	}

	for headB != headA {
		headA = headA.Next
		headB = headB.Next
	}

	return headB
}
