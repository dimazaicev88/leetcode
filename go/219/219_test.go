package _219

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test217_ContainsDuplicate(t *testing.T) {
	req := require.New(t)

	//req.Equal(true, containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	req.Equal(true, containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	//req.Equal(false, containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}
