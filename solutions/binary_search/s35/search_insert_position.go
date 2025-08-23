package s35

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums) // 使用半開區間 [l, r)
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l // l 即為目標位置或應插入的位置
}
