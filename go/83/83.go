package _83

type ListNode struct {
	Val  int
	Desc string
	Next *ListNode
}

// 83. Remove Duplicates from Sorted List
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	currentNode := head
	for currentNode.Next != nil {
		if currentNode.Val == currentNode.Next.Val {
			currentNode.Next = currentNode.Next.Next
		} else {
			currentNode = currentNode.Next
		}
	}
	return head
}
