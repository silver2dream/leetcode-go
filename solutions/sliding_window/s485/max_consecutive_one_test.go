package test

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMaxConsecutiveOne(t *testing.T) {
	convey.Convey("TestMaxConsecutiveOne", t, func() {
		convey.Convey("Case1", func() {
			nums := []int{1, 1, 0, 1, 1, 1}
			result := findMaxConsecutiveOnes(nums)
			expect := 3
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case2", func() {
			nums := []int{1, 0, 1, 1, 0, 1}
			result := findMaxConsecutiveOnes(nums)
			expect := 2
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case3", func() {
			nums := []int{1}
			result := findMaxConsecutiveOnes(nums)
			expect := 1
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case4", func() {
			nums := []int{0}
			result := findMaxConsecutiveOnes(nums)
			expect := 0
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case5", func() {
			nums := []int{0, 0}
			result := findMaxConsecutiveOnes(nums)
			expect := 0
			convey.So(result, convey.ShouldResemble, expect)
		})
	})
}
