package _21

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val > list2.Val {
		return mergeTwoLists(list2, list1)
	}

	var head *ListNode
	var tail *ListNode

	for list1 != nil || list2 != nil {
		if list1 != nil {
			listNode := &ListNode{Val: list1.Val}
			if head == nil {
				head = listNode
				tail = listNode
			} else {
				tail.Next = listNode
				tail = tail.Next
			}
			list1 = list1.Next
		}

		if list2 != nil {
			listNode := &ListNode{Val: list2.Val}
			if head == nil {
				head = listNode
				tail = listNode
			} else {
				tail.Next = listNode
				tail = tail.Next
			}
			list2 = list2.Next
		}
	}

	return head
}
