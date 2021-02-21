package AddBinary

import (
	"strconv"
	"strings"
)

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	result := make([]string, len(a))
	buf := 0
	ai := len(a) - 1
	bi := len(b) - 1

	for bi >= 0 {
		bInt := toInt(b[bi])
		aInt := toInt(a[ai])

		result[ai] = strconv.Itoa((aInt + bInt + buf) % 2)
		buf = (aInt + bInt + buf) / 2

		bi--
		ai--
	}

	for ai >= 0 {
		aInt := toInt(a[ai])
		result[ai] = strconv.Itoa((aInt + buf) % 2)
		buf = (aInt + buf) / 2
		ai--
	}

	if buf == 1 {
		result = append([]string{"1"}, result...)
	}

	return strings.Join(result, "")
}

func toInt(ui uint8) int {
	if ui == 48 {
		return 0
	}
	return 1
}
