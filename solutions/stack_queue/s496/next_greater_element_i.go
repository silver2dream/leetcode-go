package s496

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := make([]int, 0, len(nums2))
	next := make(map[int]int)

	for _, v := range nums2 {
		for len(stack) > 0 && v > stack[len(stack)-1] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			next[top] = v
		}
		stack = append(stack, v)
	}

	for _, v := range stack {
		next[v] = -1
	}

	res := []int{}
	for _, v := range nums1 {
		nextV, _ := next[v]
		res = append(res, nextV)
	}

	return res
}
