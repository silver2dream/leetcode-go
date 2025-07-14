package test

func moveZeroes(nums []int) {
	slow := 0 // 指向下一個應放非零值的位置
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}
