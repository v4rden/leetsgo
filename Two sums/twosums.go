package leetsgo

func twoSum(nums []int, target int) []int {
	result := []int{-1, -1}
	numberToIndex := make(map[int]int)

	for iNum, num := range nums {
		if index, ok := numberToIndex[target-num]; ok {
			result[0] = iNum
			result[1] = index
			break
		}
		numberToIndex[num] = iNum
	}

	return result
}
