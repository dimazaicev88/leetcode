package _83

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// 83. Remove Duplicates from Sorted List
func deleteDuplicates(head *ListNode) *ListNode {
	mapNum := map[int]int{}
	for node := head; node != nil; node = node.Next {
		if _, ok := mapNum[node.Val]; ok {
			node.Next = &ListNode{Val: node.Next.Val, Next: node.Next}
			continue
		}
		mapNum[node.Val]++
	}

	return head
}
