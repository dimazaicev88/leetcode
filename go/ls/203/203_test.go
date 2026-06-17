package _203

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntersectionNode(t *testing.T) {

	t.Run("Case1", func(t *testing.T) {
		req := require.New(t)
		headA := &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		}

		req.NotNil(removeElements(headA, 0))
	})

	t.Run("Case2", func(t *testing.T) {
		req := require.New(t)

		headA := &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		}

		req.NotNil(removeElements(headA, 2))
	})
}
