package leetsgo

func Reverse(x int) int {
	var buff int

	for x != 0 {
		buff *= 10
		buff += x % 10
		x /= 10

		if buff > 1<<31-1 || buff < -(1<<31) {
			return 0
		}
	}

	return buff
}
