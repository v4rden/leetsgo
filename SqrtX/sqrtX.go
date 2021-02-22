package SqrtX

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	left, right, res := 1, x, 0

	for left <= right {
		mid := left + ((right - left) / 2)

		if mid < x/mid {
			left = mid + 1
			res = mid
		} else if mid == x/mid {
			return mid
		} else {
			right = mid - 1
		}
	}

	return res
}
