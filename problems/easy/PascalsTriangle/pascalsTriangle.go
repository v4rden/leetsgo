package PascalsTriangle

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	result[0] = []int{1}

	for i := 1; i < numRows; i++ {
		result[i] = getRow(result[i-1])
	}
	return result
}

func getRow(row []int) []int {
	result := make([]int, len(row)+1)
	result[0] = 1

	for i := 0; i < len(row)-1; i++ {
		result[i+1] = row[i] + row[i+1]
	}

	result[len(row)] = 1
	return result
}
