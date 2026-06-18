package _234

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsPalindrome(t *testing.T) {

	t.Run("Case1", func(t *testing.T) {
		req := require.New(t)

		headA := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val:  1,
								Next: nil,
							},
						},
					},
				},
			},
		}

		val := isPalindrome(headA)
		req.True(val)
	})
}
