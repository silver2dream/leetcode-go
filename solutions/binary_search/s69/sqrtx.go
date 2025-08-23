package s69

func mySqrt(x int) int {
	if x < 2 { // 0→0, 1→1
		return x
	}
	l, r := 1, x/2
	ans := 0
	for l <= r {
		mid := l + (r-l)/2
		if mid <= x/mid { // 等價於 mid*mid <= x，但不會溢位
			ans = mid // mid 可能是答案，往右找更大但仍可平方不超過 x 的
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}
