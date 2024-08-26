package _2744

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMaximumNumberOfStringPairs(t *testing.T) {
	req := require.New(t)
	req.Equal(2, maximumNumberOfStringPairs([]string{"cd", "ac", "dc", "ca", "zz"}))
}
