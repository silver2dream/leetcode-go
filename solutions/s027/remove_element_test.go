package test

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	convey.Convey("TestRemoveElement", t, func() {
		convey.Convey("val=3", func() {
			nums := []int{3, 2, 2, 3}
			result := removeElement(nums, 3)
			expect := 2
			convey.So(result, convey.ShouldResemble, expect)

			finalNums := []int{2, 2}
			convey.So(nums[:2], convey.ShouldResemble, finalNums)
		})

		convey.Convey("empty", func() {
			nums := []int{}
			result := removeElement(nums, 3)
			expect := 0
			convey.So(result, convey.ShouldResemble, expect)

			finalNums := []int{}
			convey.So(nums, convey.ShouldResemble, finalNums)
		})

		convey.Convey("val=2", func() {
			nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
			result := removeElement(nums, 2)
			expect := 5
			convey.So(result, convey.ShouldResemble, expect)

			finalNums := []int{0, 1, 4, 0, 3}
			convey.So(nums[:5], convey.ShouldResemble, finalNums)
		})
	})
}
