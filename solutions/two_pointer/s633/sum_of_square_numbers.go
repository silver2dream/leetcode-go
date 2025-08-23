package s633

import "math"

func judgeSquareSum(c int) bool {
	if c < 0 {
		return false
	}
	l, r := 0, int(math.Sqrt(float64(c)))
	C := int64(c)
	for l <= r {
		s := int64(l)*int64(l) + int64(r)*int64(r)
		if s == C {
			return true
		}
		if s < C {
			l++
		} else {
			r--
		}
	}
	return false
}
