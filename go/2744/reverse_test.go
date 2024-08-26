package main

import "testing"

func BenchmarkReverse(b *testing.B) {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < b.N; i++ {
		for _, v := range sl {
			_ = v
		}
	}
}
