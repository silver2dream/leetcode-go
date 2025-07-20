package test

func countGoodSubstrings(s string) int {
	if len(s) < 3 {
		return 0
	}

	freq := make(map[int32]int)
	distinct := 0
	ans := 0
	for right, ch := range s {
		freq[ch]++
		if freq[ch] == 1 {
			distinct++
		}

		if right >= 3 {
			out := rune(s[right-3])
			freq[out]--

			if freq[out] == 0 {
				distinct--
			}
		}

		if right >= 2 && distinct == 3 {
			ans++
		}

	}

	return ans
}
