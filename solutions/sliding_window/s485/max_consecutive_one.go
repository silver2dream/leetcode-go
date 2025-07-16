package test

func findMaxConsecutiveOnes(nums []int) int {
	left, maxLen := 0, 0
	for right, v := range nums {
		if v == 0 {
			left = right + 1 // window 跳過 0，下一位重新開始計算
			continue
		}
		// 當前 window 長度 = right - left + 1
		if cur := right - left + 1; cur > maxLen {
			maxLen = cur
		}
	}
	return maxLen
}
