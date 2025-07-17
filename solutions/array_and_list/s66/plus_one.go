package test

func plusOne(digits []int) []int {
	digit := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += digit
		if digits[i] >= 10 {
			digit = digits[i] / 10
			digits[i] = 0
		} else {
			digit = 0
			break
		}
	}

	if digit > 0 {
		result := []int{}
		result = append(result, digit)
		for _, v := range digits {
			result = append(result, v)
		}
		return result
	} else {
		return digits
	}

	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	res := make([]int, n+1)
	res[0] = 1
	return res
}
