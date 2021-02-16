package MaximumSubArray

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	maxSum, maxCurrent := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		maxCurrent = max(nums[i], maxCurrent+nums[i])
		maxSum = max(maxSum, maxCurrent)
	}

	return maxSum
}

func max(left int, right int) int {
	if left > right {
		return left
	}
	return right
}
