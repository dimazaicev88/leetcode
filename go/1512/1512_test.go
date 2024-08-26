package _512

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNumIdenticalPairs(t *testing.T) {
	req := require.New(t)
	req.Equal(4, numIdenticalPairs([]int{1, 2, 3, 1, 1, 3}))
	req.Equal(6, numIdenticalPairs([]int{1, 1, 1, 1}))
}
