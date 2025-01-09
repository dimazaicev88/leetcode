package _2744

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestMaximumNumberOfStringPairs(t *testing.T) {
	req := require.New(t)
	req.Equal(2, maximumNumberOfStringPairs([]string{"cd", "ac", "dc", "ca", "zz"}))
}

func GetPointer[T any](x T) *T {
	if reflect.ValueOf(x).IsZero() == true {
		return nil
	}
	return &x
}

func BenchmarkReflect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetPointer(1)
	}
}
