package PascalsTriangle2

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}

	row := make([]int, rowIndex+1)
	row[0] = 1

	for i := 1; i < len(row); i++ {
		row[i] = row[i-1] * (rowIndex - i + 1) / i
	}

	return row
}
