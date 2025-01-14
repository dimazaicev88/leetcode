package _219

func containsNearbyDuplicate(nums []int, k int) bool {
	mapNum := map[int]int{}

	for pos, num := range nums {
		mapNum[num]++
		if val, ok := mapNum[num]; ok {
			if val > 1 && pos <= k {
				return true
			}
		}

		if pos > k {
			break
		}
	}

	return false
}
