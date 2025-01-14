package _217

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test217_ContainsDuplicate(t *testing.T) {
	req := require.New(t)

	req.Equal(true, containsDuplicate([]int{1, 2, 3, 1}))
	req.Equal(false, containsDuplicate([]int{1, 2, 3}))
}
