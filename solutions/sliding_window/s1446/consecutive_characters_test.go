package test

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestConsecutiveCharacters(t *testing.T) {
	convey.Convey("ConsecutiveCharacters", t, func() {
		convey.Convey("Case1", func() {
			nums := "xyzzaz"
			result := maxPower(nums)
			expect := 2
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case2", func() {
			nums := "a"
			result := maxPower(nums)
			expect := 1
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case3", func() {
			nums := "leetcode"
			result := maxPower(nums)
			expect := 2
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case4", func() {
			nums := "abbcccddddeeeeedcba"
			result := maxPower(nums)
			expect := 5
			convey.So(result, convey.ShouldResemble, expect)
		})

	})
}
