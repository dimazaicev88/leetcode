package _83

import (
	"fmt"
	"testing"
)

func Test83(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
		},
	}

	fmt.Println(fmt.Sprintf("%+v", deleteDuplicates(head)))
}
