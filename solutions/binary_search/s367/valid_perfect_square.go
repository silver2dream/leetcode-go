package s367

func isPerfectSquare(num int) bool {
	if num < 2 { // 0,1 都是完全平方
		return true
	}
	l, r := 1, num/2 // 對 num>=2，平方根一定 <= num/2
	for l <= r {
		mid := l + (r-l)/2
		q := num / mid // 等價於 floor(num/mid)
		if mid == q && num%mid == 0 {
			return true // mid*mid == num，但用除法避免溢位
		}
		if mid < q {
			l = mid + 1 // mid*mid < num
		} else {
			r = mid - 1 // mid*mid > num
		}
	}
	return false
}
