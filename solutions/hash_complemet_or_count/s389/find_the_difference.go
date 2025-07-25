package test

func findTheDifference(s string, t string) byte {
	if len(s) == 0 {
		return t[0]
	}
	var res byte
	for i := 0; i < len(t); i++ {
		if i < len(s) {
			res ^= s[i]
		}
		res ^= t[i]
	}
	return res

}
