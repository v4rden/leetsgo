package LenghtOfLastWord

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}

	var result int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != 32 {
			result++
		} else {
			if result > 0 {
				return result
			}
		}
	}
	return result
}
