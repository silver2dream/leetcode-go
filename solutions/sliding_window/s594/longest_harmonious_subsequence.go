package test

func findLHS(nums []int) int {
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}

	ans := 0
	for x, cnt := range freq {
		if cnt2, ok := freq[x+1]; ok {
			if total := cnt + cnt2; total > ans {
				ans = total
			}
		}
	}
	return ans
}
