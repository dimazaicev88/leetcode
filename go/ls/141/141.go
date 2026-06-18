package _141

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slowPointer, fastPointer := head, head

	for fastPointer != nil && fastPointer.Next != nil {
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next.Next
		if slowPointer == fastPointer {
			return true
		}
	}

	return false
}

//func hasCycle(head *ListNode) bool {
//	if head == nil || head.Next == nil {
//		return false
//	}
//	mapNodes := make(map[*ListNode]bool)
//	mapNodes[head] = true
//	cur := head
//
//	for cur.Next != nil {
//		if mapNodes[cur.Next] {
//			return true
//		}
//		mapNodes[cur.Next] = true
//		cur = cur.Next
//	}
//
//	return false
//}
