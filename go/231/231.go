package _231

import "strconv"

func isPowerOfTwo(n int) bool {
	binNum := strconv.FormatInt(int64(n), 2)
	countOne := 0
	for _, b := range binNum {
		if b == '1' {
			countOne++
		}
	}
	return n > 0 && countOne == 1
}
