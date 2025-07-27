package test

func pivotIndex(nums []int) int {
	n := len(nums)
	pre := make([]int, n+1)

	for i := 0; i < n; i++ {
		pre[i+1] = pre[i] + nums[i]
	}

	for x := 0; x < n; x++ {
		if pre[x] == pre[n]-pre[x+1] {
			return x
		}
	}

	return -1

}
