package _2022

func construct2DArray(original []int, m int, n int) [][]int {
	if m*n != len(original) {
		return [][]int{}
	}

	resultArray := make([][]int, m)
	for i := range resultArray {
		resultArray[i] = make([]int, n)
	}

	col := 0
	row := 0
	for _, val := range original {
		resultArray[row][col] = val
		col++
		if col >= n {
			row++
			col = 0
		}
	}
	return resultArray
}
