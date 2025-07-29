package test

func minStartValue(nums []int) int {
	n := len(nums)
	sum := 0
	ans := 100
	for i := 0; i < n; i++ {
		sum += nums[i]
		if ans > sum {
			ans = sum
		}
	}

	if ans < 0 {
		return 1 - ans
	}

	return 1
}
