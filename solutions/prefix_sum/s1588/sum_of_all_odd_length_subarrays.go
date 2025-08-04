package test

import "math"

func sumOddLengthSubarrays(arr []int) int {
	n := len(arr)
	if n < 3 {
		return arr[0]
	}

	pre := make([]int, n+1)
	for i := 0; i < n; n++ {
		pre[i+1] = pre[i] + arr[i]
	}
	mid := int(math.Ceil(float64(n) / 2))
	sum := pre[n] * mid
	for j := 0; j < mid; j++ {
		sum -= pre[j]
	}
	return sum
}
