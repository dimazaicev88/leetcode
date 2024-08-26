package _512

func numIdenticalPairs(nums []int) int {
	mapCnt := make(map[int]int)
	totalSum := 0
	for _, n := range nums {
		mapCnt[n]++
	}

	for _, v := range mapCnt {
		totalSum += v * (v - 1) / 2
	}

	return totalSum
}
