package _349

func intersection(nums1 []int, nums2 []int) []int {
	mapNums := map[int]bool{}

	var result []int
	for i := 0; i < len(nums2); i++ {
		mapNums[nums2[i]] = false
	}

	for i := 0; i < len(nums1); i++ {
		if _, ok := mapNums[nums1[i]]; ok {
			if !mapNums[nums1[i]] {
				result = append(result, nums1[i])
				mapNums[nums1[i]] = true
			}
		}
	}
	return result
}
