package _349

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test819_EmptyBanned(t *testing.T) {
	req := require.New(t)

	req.Equal([]int{2}, intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	req.Equal([]int{4, 9}, intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))

}
