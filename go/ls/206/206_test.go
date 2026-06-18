package _206

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntersectionNode(t *testing.T) {

	t.Run("Case1", func(t *testing.T) {
		req := require.New(t)

		headA := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		}

		val := reverseList(headA)
		req.NotNil(val)
	})
}
