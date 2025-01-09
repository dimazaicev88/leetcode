package _860

func lemonadeChange(bills []int) bool {
	sum5 := 0
	sum10 := 0
	for _, bill := range bills {
		if bill == 5 {
			sum5 += 1
		} else if bill == 10 {
			if sum5 == 0 {
				return false
			}
			sum5 -= 1
			sum10 += 1
		} else if bill == 20 {
			if sum5 > 0 && sum10 > 0 {
				sum5 -= 1
				sum10 -= 1
			} else if sum5 >= 3 {
				sum5 -= 3
			} else {
				return false
			}
		}
	}

	return true
}
