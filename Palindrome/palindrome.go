package Palindrome

func IsPalindrome(x int) bool {

	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	var rev int
	var buff int = x
	for buff != 0 {
		rev *= 10
		rev += buff % 10
		buff /= 10
	}

	return rev == x
}
