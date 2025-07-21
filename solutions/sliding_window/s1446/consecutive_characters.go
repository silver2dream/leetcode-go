package test

func maxPower(s string) int {
	n := len(s)
	if n == 1 {
		return 1
	}

	l := 0
	ans := 0
	for r := 0; r < len(s); r++ {
		if s[l] != s[r] {
			l = r
			continue
		}

		if s[l] == s[r] {
			count := r - l + 1
			if count > ans {
				ans = count
			}
		}
	}
	return ans
}
