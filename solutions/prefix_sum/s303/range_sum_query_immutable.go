package test

type NumArray struct {
	pre []int
}

func Constructor(nums []int) NumArray {
	pre := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		pre[i+1] = pre[i] + nums[i]
	}
	return NumArray{pre}
}

func (a *NumArray) SumRange(i, j int) int {
	return a.pre[j+1] - a.pre[i]
}
