package Valid_parentheses

func isValid(s string) bool {
	if len(s) < 2 {
		return false
	}

	stack := make([]rune, 0)

	for _, rune := range s {
		if rune == '{' || rune == '[' || rune == '(' {
			stack = append(stack, rune)
		} else if (rune == '}' && len(stack) > 0 && stack[len(stack)-1] == '{') ||
			(rune == ']' && len(stack) > 0 && stack[len(stack)-1] == '[') ||
			(rune == ')' && len(stack) > 0 && stack[len(stack)-1] == '(') {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}
