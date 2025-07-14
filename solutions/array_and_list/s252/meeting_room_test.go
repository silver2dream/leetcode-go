package test

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMeetingRoom(t *testing.T) {
	convey.Convey("TestMeetingRoom", t, func() {
		convey.Convey("no intervals", func() {
			input := [][]int{}
			result := canAttendMeetings(input)
			expect := true
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("single interval", func() {
			input := [][]int{{5, 10}}
			result := canAttendMeetings(input)
			expect := true
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case1: overlap", func() {
			input := [][]int{{0, 30}, {5, 10}, {15, 20}}
			result := canAttendMeetings(input)
			expect := false
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case2: no overlap", func() {
			input := [][]int{{7, 10}, {2, 4}}
			result := canAttendMeetings(input)
			expect := true
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("touching endpoints", func() {
			input := [][]int{{1, 2}, {2, 3}, {3, 4}}
			result := canAttendMeetings(input)
			expect := true
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("unsorted input no conflict", func() {
			input := [][]int{{5, 8}, {1, 4}, {4, 5}}
			result := canAttendMeetings(input)
			expect := true
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("unsorted input conflict", func() {
			input := [][]int{{5, 8}, {1, 6}}
			result := canAttendMeetings(input)
			expect := false
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("max bounds", func() {
			input := [][]int{{0, 1000000}, {1000000, 1000001}}
			result := canAttendMeetings(input)
			expect := true
			convey.So(result, convey.ShouldResemble, expect)
		})

	})

}
