package _217

func containsDuplicate(nums []int) bool {
	mapNum := map[int]int{}

	for _, num := range nums {
		mapNum[num]++
		if val, ok := mapNum[num]; ok {
			if val > 1 {
				return true
			}
		}
	}

	return false
}
