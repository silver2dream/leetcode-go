package atoi

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	convey.Convey("TestMyAtoi", t, func() {
		convey.Convey("Case1", func() {
			input := "42"
			result := myAtoi(input)
			expect := 42
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case2", func() {
			input := "1337c0d3"
			result := myAtoi(input)
			expect := 1337
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case3", func() {
			input := "0-1"
			result := myAtoi(input)
			expect := 0
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case4", func() {
			input := "-042"
			result := myAtoi(input)
			expect := -42
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case5", func() {
			input := "   -042"
			result := myAtoi(input)
			expect := -42
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case6", func() {
			input := "+-12"
			result := myAtoi(input)
			expect := 0
			convey.So(result, convey.ShouldResemble, expect)
		})

	})

}
