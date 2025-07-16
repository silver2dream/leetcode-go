package test

func containsNearbyDuplicate(nums []int, k int) bool {
	left := 0
	window := make(map[int]struct{})
	for right, v := range nums {
		if _, exist := window[v]; exist {
			return true
		}
		window[v] = struct{}{}

		// 收斂
		if right-left+1 > k {
			delete(window, nums[left])
			left++
		}
	}

	return false
}
