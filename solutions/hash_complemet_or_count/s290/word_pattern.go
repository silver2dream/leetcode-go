package test

import "strings"

func wordPattern(pattern string, s string) bool {
	strArray := strings.Split(s, " ")
	if len(pattern) != len(strArray) {
		return false
	}

	record1 := make(map[int32]string)
	record2 := make(map[string]int32)
	for i, ch := range pattern {

		if v, e := record1[ch]; e {
			if v != strArray[i] {
				return false
			}
		}

		if v2, e2 := record2[strArray[i]]; e2 {
			if v2 != ch {
				return false
			}
		}

		record1[ch] = strArray[i]
		record2[strArray[i]] = ch

	}
	return true
}
