package test

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	convey.Convey("TestMyAtoi", t, func() {
		convey.Convey("Case1", func() {
			input := []int{3, 0, 1}
			result := missingNumber(input)
			expect := 2
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case2", func() {
			input := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
			result := missingNumber(input)
			expect := 8
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case3", func() {
			input := []int{0, 1}
			result := missingNumber(input)
			expect := 2
			convey.So(result, convey.ShouldResemble, expect)
		})

	})

}
