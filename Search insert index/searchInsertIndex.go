package Search_insert_index

func searchInsert(nums []int, target int) int {
	if len(nums) < 1 {
		return 0
	}

	for i, num := range nums {
		if num >= target {
			return i
		}
	}

	return len(nums)
}
