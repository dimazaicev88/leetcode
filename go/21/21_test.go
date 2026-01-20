package _21

import (
	"fmt"
	"testing"
)

func Test21(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {
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
		for ml != nil {
			fmt.Println(ml.Val)
			ml = ml.Next
		}
	})

	t.Run("case 2", func(t *testing.T) {
		list1 := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
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
		for ml != nil {
			fmt.Println(ml.Val)
			ml = ml.Next
		}
	})

	t.Run("case 3", func(t *testing.T) {
		list1 := &ListNode{
			Val: 2,
		}

		list2 := &ListNode{
			Val: 1,
		}

		ml := mergeTwoLists(list1, list2)
		for ml != nil {
			fmt.Println(ml.Val)
			ml = ml.Next
		}
	})

}
