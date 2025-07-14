package atoi

func myAtoi(s string) int {
	const (
		intMax = 1<<31 - 1 //  2147483647
		intMin = -1 << 31  // -2147483648
	)

	n, i := len(s), 0

	// 1. 跳過前導空白
	for i < n && s[i] == ' ' {
		i++
	}
	if i == n { // 只包含空白
		return 0
	}

	// 2. 處理正負號
	sign := 1
	if s[i] == '+' {
		i++
	} else if s[i] == '-' {
		sign = -1
		i++
	}

	// 3. 讀取數字並即時檢查溢出
	val := 0
	for i < n && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')

		// 若下一步會溢出則直接回傳邊界值
		if sign == 1 && (val > (intMax-digit)/10) {
			return intMax
		}
		if sign == -1 && (val > (-(intMin + digit))/10) {
			return intMin
		}

		val = val*10 + digit
		i++
	}

	return val * sign
}
