package test

func possibleStringCount(word string) int {
	n := len(word)
	if n == 0 {
		return 0
	}
	ans := 1    // 「一次都沒按太久」這種情況
	runLen := 1 // 當前連續段長度

	// 掃描 word，找到所有連續段長 runLen
	for i := 1; i < n; i++ {
		if word[i] == word[i-1] {
			runLen++
		} else {
			// 一段結束，若 runLen>1，累加 (runLen-1)
			if runLen > 1 {
				ans += runLen - 1
			}
			runLen = 1
		}
	}
	// 別忘了最後一段
	if runLen > 1 {
		ans += runLen - 1
	}
	return ans
}
