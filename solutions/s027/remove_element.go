package test

func removeElement(nums []int, val int) int {
	tail := len(nums) - 1
	k := 0
	for head := 0; head <= tail; {
		if nums[head] == val {
			k++
			nums[head], nums[tail] = nums[tail], nums[head]
			tail--
		} else {
			head++
		}
	}
	return len(nums) - k
}
