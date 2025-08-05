package s338

func countBits(n int) []int {
	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		arr[i] = arr[i&(i-1)] + 1
	}
	return arr
}
