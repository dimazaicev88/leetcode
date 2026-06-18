package _234

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head *ListNode) *ListNode {
	var reverseList *ListNode
	for head != nil {
		if reverseList == nil {
			reverseList = &ListNode{Val: head.Val}
		} else {
			reverseList = &ListNode{Val: head.Val, Next: reverseList}
		}
		head = head.Next
	}

	return reverseList
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	} else if head.Next == nil {
		return true
	}

	cur := head
	var slow, fast = cur, cur

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	reverseList := reverse(slow)
	cur = head
	for reverseList != nil {
		if cur.Val != reverseList.Val {
			return false
		}
		cur = cur.Next
		reverseList = reverseList.Next
	}

	return true
}
