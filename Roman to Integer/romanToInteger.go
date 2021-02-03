package Roman_to_Integer

var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	if len(s) < 1 {
		return 0
	}

	var previous, current, result int

	for i := len(s); i > 0; i-- {

		current = roman[s[i-1:i]]

		if current < previous {
			result -= current
		} else {
			result += current
		}

		previous = current
	}

	return result
}
