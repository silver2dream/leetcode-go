package test

func findLengthOfLCIS(nums []int) int {

	left := 0
	ans := 1
	for right := 1; right < len(nums); right++ {
		if nums[right] <= nums[right-1] {
			left = right
		}

		if nums[right] > nums[right-1] {
			count := right - left + 1
			if ans < count {
				ans = count
			}
		}
	}

	return ans
}
