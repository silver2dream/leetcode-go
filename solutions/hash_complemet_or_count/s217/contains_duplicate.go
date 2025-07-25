package test

func containsDuplicate(nums []int) bool {
	if len(nums) == 1 {
		return false
	}

	res := make(map[int]struct{})
	for _, v := range nums {
		if _, e := res[v]; e {
			return true
		}
		res[v] = struct{}{}
	}

	return false
}
