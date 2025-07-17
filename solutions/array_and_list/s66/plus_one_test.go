package test

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBestPlusOne(t *testing.T) {
	convey.Convey("TestBestPlusOne", t, func() {
		convey.Convey("Case1", func() {
			input := []int{1, 2, 3}
			result := plusOne(input)
			expect := []int{1, 2, 4}
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case2", func() {
			input := []int{4, 3, 2, 1}
			result := plusOne(input)
			expect := []int{4, 3, 2, 2}
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case3", func() {
			input := []int{9}
			result := plusOne(input)
			expect := []int{1, 0}
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case4", func() {
			input := []int{9, 0, 9}
			result := plusOne(input)
			expect := []int{9, 1, 0}
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case5", func() {
			input := []int{9, 9, 9}
			result := plusOne(input)
			expect := []int{1, 0, 0, 0}
			convey.So(result, convey.ShouldResemble, expect)
		})

	})

}
