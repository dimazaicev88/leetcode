package _819

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test819_EmptyBanned(t *testing.T) {
	req := require.New(t)

	req.Equal("a", mostCommonWord("a.", []string{""}))
	req.Equal("ball", mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}))
	req.Equal("bob", mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"ball", "hit"}))
	req.Equal("bob", mostCommonWord("Bob", []string{}))
}
