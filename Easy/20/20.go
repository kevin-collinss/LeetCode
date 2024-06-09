package Easy

func isValid(s string) bool {

	stack := []rune{}

	for _, val := range s {

		if val == '(' || val == '{' || val == '[' {
			stack = append(stack, val)
		} else if val == ')' || val == '}' || val == ']' {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if (val == ')' && top != '(') || (val == '}' && top != '{') || (val == ']' && top != '[') {
				return false // mismatched pair
			}
		}
	}

	return len(stack) == 0
}
