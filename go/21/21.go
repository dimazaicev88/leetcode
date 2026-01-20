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

	var newList *ListNode

	for list1 != nil || list2 != nil {
		if list1 != nil {
			listNode := &ListNode{Val: list1.Val, Next: newList}
			newList = listNode
			list1 = list1.Next
		}

		if list2 != nil {
			listNode2 := &ListNode{Val: list2.Val, Next: newList}
			newList = listNode2
			list2 = list2.Next
		}
	}
	return newList
}
