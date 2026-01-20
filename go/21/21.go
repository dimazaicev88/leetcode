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

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			newList = &ListNode{
				Val:  list1.Val,
				Next: newList,
			}
		} else if list1.Val > list2.Val {
			newList = &ListNode{
				Val:  list2.Val,
				Next: newList,
			}
		} else {
			listNode := &ListNode{Val: list1.Val, Next: newList}
			newList = listNode
		}
		list1 = list1.Next
		list2 = list2.Next
	}

	return newList
}

func main() {

	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	ml := mergeTwoLists(list1, list2)
	fmt.Println(ml)
}
