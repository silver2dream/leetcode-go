package test

func largeGroupPositions(s string) [][]int {
	var group [][]int
	l := 0
	n := len(s)
	for r := 1; r <= n; r++ {
		if r == n || s[r] != s[l] {
			if r-l >= 3 {
				group = append(group, []int{l, r - 1})
			}
			l = r
		}
	}
	return group
}
