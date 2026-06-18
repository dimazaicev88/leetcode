package _234

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	} else if head.Next == nil {
		return true
	}

	var reverseList *ListNode
	cur := head
	left, right := 0, 0
	for cur != nil {
		if reverseList == nil {
			reverseList = &ListNode{Val: cur.Val}
		} else {
			reverseList = &ListNode{Val: cur.Val, Next: reverseList}
		}
		cur = cur.Next
		right++
	}

	cur = head
	curReverse := reverseList
	for left < right {
		if cur.Val != curReverse.Val {
			return false
		}
		cur = cur.Next
		curReverse = curReverse.Next
		left++
		right--
	}

	return true
}
