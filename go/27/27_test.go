package _27

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test27(t *testing.T) {
	req := require.New(t)

	//req.Equal(0, removeElement([]int{}, 2))
	//req.Equal(0, removeElement([]int{2}, 2))
	//req.Equal(2, removeElement([]int{3, 2, 2, 3}, 3))
	//req.Equal(0, removeElement([]int{2, 2, 2, 2}, 3))
	arr := []int{0, 1, 2, 2, 3, 0, 4, 2}
	req.Equal(0, removeElement(arr, 2))
	fmt.Println(arr)

}
