package test

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// mapS 用來記 s → t 的映射，mapT 用來記 t → s 的反向映射
	var mapS [256]byte
	var mapT [256]byte
	// 初始化為 0（byte(0)／'\x00' 表示尚未映射）
	for i := range mapS {
		mapS[i] = 0
		mapT[i] = 0
	}

	for i := 0; i < len(s); i++ {
		cs := s[i]
		ct := t[i]
		if mapS[cs] == 0 && mapT[ct] == 0 {
			// 尚未建立映射，雙向綁定
			mapS[cs] = ct
			mapT[ct] = cs
		} else {
			// 任何一方有映射不符，都不能同構
			if mapS[cs] != ct || mapT[ct] != cs {
				return false
			}
		}
	}
	return true
}
