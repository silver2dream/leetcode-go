package test

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPositionsOfLargeGroups(t *testing.T) {
	convey.Convey("PositionsOfLargeGroups", t, func() {
		convey.Convey("Case1", func() {
			nums := "abbxxxxzzy"
			result := largeGroupPositions(nums)
			expect := [][]int{{3, 6}}
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case2", func() {
			nums := "abc"
			result := largeGroupPositions(nums)
			var expect [][]int
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case3", func() {
			nums := "abcdddeeeeaabbbcd"
			result := largeGroupPositions(nums)
			expect := [][]int{{3, 5}, {6, 9}, {12, 14}}
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case4", func() {
			nums := "abcdddeeeeaabbbcdeeeeee"
			result := largeGroupPositions(nums)
			expect := [][]int{{3, 5}, {6, 9}, {12, 14}, {17, 22}}
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case5", func() {
			nums := "aa"
			result := largeGroupPositions(nums)
			var expect [][]int
			convey.So(result, convey.ShouldResemble, expect)
		})

	})
}
