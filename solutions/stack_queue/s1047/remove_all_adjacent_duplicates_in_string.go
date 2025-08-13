package s1047

func removeDuplicates(s string) string {
	stack := make([]rune, 0, len(s))
	for _, v := range s {
		n := len(stack)
		if n > 0 && v == stack[n-1] {
			stack = stack[:n-1]
		} else {
			stack = append(stack, v)
		}
	}
	return string(stack)
}
