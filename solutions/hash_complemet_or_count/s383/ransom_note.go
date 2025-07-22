package test

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	var res [26]int
	for idx, val := range magazine {

		res[int(val-'a')]++
		if idx < len(ransomNote) {
			r := ransomNote[idx]
			res[int(r-'a')]--
		}
	}

	for _, v := range res {
		if v < 0 {
			return false
		}
	}

	return true
}
