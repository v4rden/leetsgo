package Remove_element

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	var valCount int

	for i := len(nums) - 1; i > -1; i-- {
		if nums[i] == val {
			valCount++
			nums[len(nums)-valCount], nums[i] = nums[i], nums[len(nums)-valCount]
		}
	}

	return len(nums) - valCount
}
