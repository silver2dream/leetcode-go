package test

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	ans := 0

	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	ans = sum

	for r := k; r < len(nums); r++ {
		sum += nums[r]
		if r >= k {
			sum -= nums[r-k]
		}

		if ans < sum {
			ans = sum
		}
	}

	return float64(ans) / float64(k)
}
