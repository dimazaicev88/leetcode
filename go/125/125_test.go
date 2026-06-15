package _25

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsPalindrome(t *testing.T) {
	req := require.New(t)

	//req.Equal(true, isPalindrome("A man, a plan, a canal: Panama"))
	//req.Equal(false, isPalindrome("race a car"))
	//req.Equal(true, isPalindrome("a."))
	//req.Equal(true, isPalindrome("ab@a"))
	req.Equal(false, isPalindrome("0P"))
}
