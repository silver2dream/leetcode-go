package s20

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 { // 奇數長度一定不可能配對完
		return false
	}

	stack := make([]byte, 0, n)
	pairs := map[byte]byte{')': '(', ']': '[', '}': '{'}

	for i := 0; i < n; i++ {
		c := s[i]
		switch c {
		case '(', '[', '{':
			stack = append(stack, c)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[c] {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			// 本題通常只有括號字元，若要嚴謹可直接 return false
			return false
		}
	}
	return len(stack) == 0
}
