package test

func canAttendMeetings(intervals [][]int) bool {
	hitMap := make(map[int]bool)
	for _, interval := range intervals {
		for i := interval[0]; i < interval[1]; i++ {
			if val := hitMap[i]; val == true {
				return false
			}
			hitMap[i] = true
		}

	}

	return true
}
