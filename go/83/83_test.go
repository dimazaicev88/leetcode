package _83

import (
	"fmt"
	"testing"
)

func Test83(t *testing.T) {
	head := &ListNode{
		Val:  1,
		Desc: "1",
		Next: &ListNode{
			Val:  3,
			Desc: "2",
			Next: &ListNode{
				Val:  3,
				Desc: "3",
				Next: &ListNode{
					Val:  3,
					Desc: "4",
					Next: &ListNode{
						Val:  4,
						Desc: "5",
					},
				},
			},
		},
	}

	fmt.Println(fmt.Sprintf("%+v", deleteDuplicates(head)))
}
