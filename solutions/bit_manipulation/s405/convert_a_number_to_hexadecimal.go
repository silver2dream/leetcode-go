package s405

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	const hexDigits = "0123456789abcdef"
	// 將 num 視作無符號 32 位整數
	u := uint32(num)
	// 最多 8 位，每次處理最低 4 位
	buf := make([]byte, 0, 8)
	for u != 0 {
		d := u & 0xF // 取最低 4 bits
		buf = append(buf, hexDigits[d])
		u >>= 4 // 無符號右移 4 bits
	}
	// 反轉，得到正確順序
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
	return string(buf)
}
