package test

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestContinuousIncreasingSubsequence(t *testing.T) {
	convey.Convey("ContinuousIncreasingSubsequence", t, func() {
		convey.Convey("Case1", func() {
			nums := []int{1, 3, 5, 4, 7}
			result := findLengthOfLCIS(nums)
			expect := 3
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case2", func() {
			nums := []int{2, 2, 2, 2, 2}
			result := findLengthOfLCIS(nums)
			expect := 1
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case3", func() {
			nums := []int{1, 3, 2, 2, 5, 2, 3, 7}
			result := findLengthOfLCIS(nums)
			expect := 3
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case4", func() {
			nums := []int{1, 1, 2}
			result := findLengthOfLCIS(nums)
			expect := 2
			convey.So(result, convey.ShouldResemble, expect)
		})
	})
}
