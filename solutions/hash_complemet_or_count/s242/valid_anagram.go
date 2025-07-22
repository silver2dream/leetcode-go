package test

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	res := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		res[s[i]]++
		res[t[i]]--
	}

	for _, v := range res {
		if v != 0 {
			return false
		}
	}

	return true
}
