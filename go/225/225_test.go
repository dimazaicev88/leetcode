package _225

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test217_ContainsDuplicate(t *testing.T) {
	req := require.New(t)
	ms := Constructor()
	ms.Push(1)
	ms.Push(2)

	req.Equal(2, ms.Top())
	req.Equal(2, ms.Pop())
	req.Equal(false, ms.Empty())
}
