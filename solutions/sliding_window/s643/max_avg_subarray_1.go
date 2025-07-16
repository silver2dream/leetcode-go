package test

func findMaxAverage(nums []int, k int) float64 {
	n := len(nums)
	// 1) 算出第 1 個窗口的和，並設 maxSum = sum
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxSum := sum

	// 2) 窗口從 i=k…n-1 滑動：減掉離開的 + 加上進來的 → 更新 maxSum
	for i := k; i < n; i++ {
		sum += nums[i] - nums[i-k]
		if sum > maxSum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)
}
