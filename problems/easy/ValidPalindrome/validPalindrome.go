package ValidPalindrome

func isPalindrome(s string) bool {
	var i, j int
	input := make([]byte, len(s))

	for idx := range s {
		if ('a' <= s[idx] && s[idx] <= 'z') ||
			('0' <= s[idx] && s[idx] <= '9') {
			input[i] = s[idx]
			i++
			continue
		}
		if 'A' <= s[idx] && s[idx] <= 'Z' {
			input[i] = s[idx] + 32
			i++
		}
	}
	i--

	for i > j {
		if input[i] != input[j] {
			return false
		}
		i--
		j++
	}
	return true
}
