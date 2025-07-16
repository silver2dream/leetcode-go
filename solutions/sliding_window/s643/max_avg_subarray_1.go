package test

import "math"

func findMaxAverage(nums []int, k int) float64 {
	result := math.MinInt
	sum := 0
	left := 0
	for right, val := range nums {
		sum += val
		if right-left+1 > k {
			sum -= nums[left]
			left++
		}

		if sum > result && right-left+1 == k {
			result = sum
		}
	}

	return float64(result) / float64(k)
}
