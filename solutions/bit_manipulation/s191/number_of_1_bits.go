package s191

func hammingWeight(n int) int {
	count := 0
	for n != 0 {
		n &= n - 1 // 每次把最低位變成0
		count++
	}
	return count
}
