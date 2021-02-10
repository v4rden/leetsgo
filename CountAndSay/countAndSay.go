package CountAndSay

import "strings"
import "strconv"

func countAndSay(n int) string {
	result := "1"
	if n == 1 {
		return result
	}

	for i := 0; i < n-1; i++ {
		result = countAndSayM(result)
	}

	return result
}

func countAndSayM(str string) string {
	var buff strings.Builder

	for i := 0; i < len(str); i++ {
		count := 1
		current := str[i]
		for i < len(str)-1 && current == str[i+1] {
			count++
			i++
		}
		buff.WriteString(strconv.Itoa(count))
		buff.WriteString(string(current))
	}
	return buff.String()
}
