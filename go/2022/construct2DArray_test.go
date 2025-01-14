package _2022

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConstruct2DArray(t *testing.T) {
	req := require.New(t)
	req.Equal([][]int{
		{1, 2},
		{3, 4},
	}, construct2DArray([]int{1, 2, 3, 4}, 2, 2))

	req.Equal([][]int{
		{1, 2, 3},
	}, construct2DArray([]int{1, 2, 3}, 1, 3))

	fmt.Println(construct2DArray([]int{1, 2}, 1, 1))
	req.Equal([][]int{}, construct2DArray([]int{1, 2}, 1, 1))
}
