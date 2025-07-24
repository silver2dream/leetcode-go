package test

func largeGroupPositions(s string) [][]int {
	var res [][]int
	l := 0
	for r := 0; r <= len(s); r++ {
		if r == len(s) || s[r] != s[l] {
			if r-l >= 3 && s[r-1] == s[l] {
				res = append(res, []int{l, r - 1})
			}
			l = r
		}
	}

	return res
}
