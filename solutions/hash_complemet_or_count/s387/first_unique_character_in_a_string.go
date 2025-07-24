package test

func firstUniqChar(s string) int {
	// 1. 統計所有字母頻次
	var cnt [26]int
	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
	}

	// 2. 再次掃描，第一個頻次為 1 的即為答案
	for i := 0; i < len(s); i++ {
		if cnt[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
