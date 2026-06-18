package _206

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := head

	var newHead *ListNode
	for cur != nil {
		if newHead == nil {
			newHead = &ListNode{Val: cur.Val}
		} else {
			newHead = &ListNode{Val: cur.Val, Next: newHead}
		}
		cur = cur.Next
	}

	return newHead
}
