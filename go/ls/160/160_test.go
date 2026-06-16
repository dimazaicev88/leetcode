package _60

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntersectionNode(t *testing.T) {

	t.Run("Case1", func(t *testing.T) {
		req := require.New(t)
		node3 := &ListNode{
			Val:  3,
			Next: nil,
		}

		headA := &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: node3,
				},
			},
		}

		headB := &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: node3,
				},
			},
		}

		req.Equal(node3, getIntersectionNode(headA, headB))
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
						Val:  3,
						Next: nil,
					},
				},
			},
		}

		headB := &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		}
		val := getIntersectionNode(headA, headB)
		req.Nil(val)
	})
}
