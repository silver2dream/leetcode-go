package test

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	convey.Convey("TestSingleNumber", t, func() {
		convey.Convey("Case1", func() {
			input := []int{2, 2, 1}
			result := singleNumber(input)
			expect := 1
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case2", func() {
			input := []int{4, 1, 2, 1, 2}
			result := singleNumber(input)
			expect := 4
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case3", func() {
			input := []int{1}
			result := singleNumber(input)
			expect := 1
			convey.So(result, convey.ShouldResemble, expect)
		})

	})

}
