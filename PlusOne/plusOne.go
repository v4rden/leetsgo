package PlusOne

func plusOne(digits []int) []int {
	buf := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+buf > 9 {
			digits[i] = 0
			buf = 1
		} else {
			digits[i] = digits[i] + buf
			buf = 0
		}
	}
	if buf == 0 {
		return digits
	}
	return prepend(digits)
}

func prepend(input []int) []int {
	result := make([]int, len(input)+1)
	result[0] = 1
	for i, d := range input {
		result[i+1] = d
	}
	return result
}
