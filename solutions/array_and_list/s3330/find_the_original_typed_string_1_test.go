package test

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFindTheOriginalTypedString_1(t *testing.T) {
	convey.Convey("TestFindTheOriginalTypedString_1", t, func() {
		convey.Convey("Case1", func() {
			s := "abbcccc"
			result := possibleStringCount(s)
			expect := 5
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case2", func() {
			s := "abcd"
			result := possibleStringCount(s)
			expect := 1
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case3", func() {
			s := "aaaa"
			result := possibleStringCount(s)
			expect := 4
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case4", func() {
			s := "aaaabb"
			result := possibleStringCount(s)
			expect := 4
			convey.So(result, convey.ShouldResemble, expect)
		})

	})
}
