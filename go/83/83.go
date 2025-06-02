package _83

type ListNode struct {
	Val  int
	Next *ListNode
}

// 83. Remove Duplicates from Sorted List
func deleteDuplicates(head *ListNode) *ListNode {
	mapNum := map[int]int{}
	currentNode := head

	for currentNode.Next != nil {
		count, ok := mapNum[currentNode.Val]
		if ok && count > 0 {
			currentNode.Next = currentNode.Next.Next
		}
		mapNum[currentNode.Val]++
		currentNode = currentNode.Next
	}

	return head
}
