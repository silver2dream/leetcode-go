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

	})
}
