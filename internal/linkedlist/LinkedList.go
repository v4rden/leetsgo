package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewNode(val int, nexNode *ListNode) *ListNode {
	node := ListNode{Val: val, Next: nexNode}
	return &node
}
