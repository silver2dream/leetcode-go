package test

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	convey.Convey("TestContainsDuplicate", t, func() {
		convey.Convey("Case1", func() {
			nums := []int{1, 2, 3, 1}
			result := containsNearbyDuplicate(nums, 3)
			expect := true
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case2", func() {
			nums := []int{1, 0, 1, 1}
			result := containsNearbyDuplicate(nums, 1)
			expect := true
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case3", func() {
			nums := []int{1, 2, 3, 1, 2, 3}
			result := containsNearbyDuplicate(nums, 2)
			expect := false
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Cas4", func() {
			nums := []int{99, 99}
			result := containsNearbyDuplicate(nums, 2)
			expect := true
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Cas5", func() {
			nums := []int{1, 4, 2, 3, 1, 2}
			result := containsNearbyDuplicate(nums, 3)
			expect := true
			convey.So(result, convey.ShouldResemble, expect)
		})

	})
}
