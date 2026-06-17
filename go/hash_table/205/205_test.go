package _205

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumIdenticalPairs(t *testing.T) {
	req := require.New(t)
	req.Equal(true, isIsomorphic("egg", "add"))
	req.Equal(true, isIsomorphic("paper", "title"))
	req.Equal(false, isIsomorphic("badc", "baba"))
	req.Equal(false, isIsomorphic("foo", "bar"))
	req.Equal(false, isIsomorphic("egcd", "adfd"))
}
