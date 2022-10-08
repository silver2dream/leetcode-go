package test

func searchRange(nums []int, target int) []int {
	length := len(nums)

	if length < 2 {
		if length < 1 {
			return []int{-1, -1}
		}
		if nums[0] == target {
			return []int{0, 0}
		}
	}

	var result []int
	search(nums, target, 0, length-1, &result)
	resultLen := len(result)
	if resultLen < 1 {
		return []int{-1, -1}
	}
	return []int{result[0], result[resultLen-1]}
}

func search(nums []int, target int, start int, end int, result *[]int) {
	pivot := (end + start) / 2
	if pivot < start || pivot > end {
		return
	}

	right := (pivot + end) / 2
	left := (pivot + start) / 2

	if nums[left] <= nums[pivot] {
		search(nums, target, start, pivot-1, result)
	}

	if nums[pivot] == target {
		*result = append(*result, pivot)
	}

	if nums[right] >= nums[pivot] {
		search(nums, target, pivot+1, end, result)
	}
}
