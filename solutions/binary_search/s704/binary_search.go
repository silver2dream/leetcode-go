package test

func search(nums []int, target int) int {
	ans := -1
	n := len(nums)
	mid := 0
	for i := 0; i < n; i++ {
		mid = (mid + n) / 2
		if nums[mid] > target {
			n = mid
		} else {
			i = mid
		}

		if nums[i] == target {
			ans = i
			break
		}
	}
	return ans
}
