package Remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	var lastUnique, lastIndex int
	lastUnique = nums[0]

	for _, number := range nums {

		if number == lastUnique {
			continue
		}

		lastUnique = number
		lastIndex += 1
		nums[lastIndex] = number
	}

	return lastIndex + 1
}
