package StrStr

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	if len(haystack) == 0 || len(haystack) < len(needle) {
		return -1
	}

	for i, _ := range haystack {
		for j := 0; j < len(needle); j++ {
			if len(haystack) == i+j || haystack[i+j] != needle[j] {
				break
			}
			if j == len(needle)-1 {
				return i
			}
		}
	}

	return -1
}
