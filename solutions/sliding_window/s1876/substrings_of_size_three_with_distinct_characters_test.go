package test

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCountGoodSubstrings(t *testing.T) {
	convey.Convey("CountGoodSubstrings", t, func() {
		convey.Convey("Case1", func() {
			nums := "xyzzaz"
			result := countGoodSubstrings(nums)
			expect := 1
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case2", func() {
			nums := "aababcabc"
			result := countGoodSubstrings(nums)
			expect := 4
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case3", func() {
			nums := "owuxoelszb"
			result := countGoodSubstrings(nums)
			expect := 8
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case4", func() {
			nums := "npdrlvffzefb"
			result := countGoodSubstrings(nums)
			expect := 8
			convey.So(result, convey.ShouldResemble, expect)
		})

	})
}
