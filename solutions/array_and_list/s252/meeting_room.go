package test

import "sort"

func canAttendMeetings(intervals [][]int) bool {
	if len(intervals) < 2 {
		return true
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	prevEnd := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		start, end := intervals[i][0], intervals[i][1]
		if start < prevEnd {
			return false
		}
		prevEnd = end
	}

	return true
}
