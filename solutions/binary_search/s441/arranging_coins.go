package s441

func arrangeCoins(n int) int {
	l, r := 1, n
	ans := 0
	for l <= r {
		mid := l + (r-l)/2
		total := (1 + mid) * mid / 2
		if total > n {
			r = mid - 1
			ans = r
		} else {
			l = mid + 1
			ans = r
		}
	}
	return ans
}
