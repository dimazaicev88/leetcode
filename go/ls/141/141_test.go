package _141

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntersectionNode(t *testing.T) {
	req := require.New(t)

	node4 := &ListNode{
		Val: 4,
	}

	node3 := &ListNode{
		Val:  3,
		Next: node4,
	}

	node2 := &ListNode{
		Val:  2,
		Next: node3,
	}

	node1 := &ListNode{
		Val:  1,
		Next: node2,
	}

	node := &ListNode{
		Val:  0,
		Next: node1,
	}

	node2.Next = node

	req.True(hasCycle(node))
}
