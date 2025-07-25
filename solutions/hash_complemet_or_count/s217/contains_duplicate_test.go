package test

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	convey.Convey("ContainsDuplicate", t, func() {
		convey.Convey("Case1", func() {
			s := []int{1, 2, 3, 1}
			result := containsDuplicate(s)
			expect := true
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case2", func() {
			s := []int{1, 2, 3, 4}
			result := containsDuplicate(s)
			expect := false
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case3", func() {
			s := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
			result := containsDuplicate(s)
			expect := true
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case4", func() {
			s := []int{-9}
			result := containsDuplicate(s)
			expect := false
			convey.So(result, convey.ShouldResemble, expect)
		})

	})
}
