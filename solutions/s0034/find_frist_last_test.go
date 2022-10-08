package test

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSearchRange(t *testing.T) {
	convey.Convey("TestSearchRange", t, func() {
		convey.Convey("Single", func() {
			nums := []int{1, 2, 3, 4, 5, 6}
			result := searchRange(nums, 5)
			expect := []int{4, 4}
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("MultiSame", func() {
			nums := []int{1, 2, 5, 5, 5, 5, 5, 5, 8, 9}
			result := searchRange(nums, 5)
			expect := []int{2, 7}
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("AllSame", func() {
			nums := []int{5, 5, 5, 5, 5, 5}
			result := searchRange(nums, 5)
			expect := []int{0, 5}
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("NotFound", func() {
			nums := []int{1, 2, 3, 4, 4, 6}
			result := searchRange(nums, 5)
			expect := []int{-1, -1}
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("OneElement", func() {
			nums := []int{1}
			result := searchRange(nums, 1)
			expect := []int{0, 0}
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Empty", func() {
			var nums []int
			result := searchRange(nums, 1)
			expect := []int{-1, -1}
			convey.So(result, convey.ShouldResemble, expect)
		})
	})
}
