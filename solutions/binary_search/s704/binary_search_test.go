package test

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSearch(t *testing.T) {
	convey.Convey("Search", t, func() {
		convey.Convey("Case1", func() {
			input := []int{-1, 0, 3, 5, 9, 12}
			result := search(input, 9)
			expect := 4
			convey.So(result, convey.ShouldResemble, expect)
		})
		convey.Convey("Case2", func() {
			input := []int{-1, 0, 3, 5, 9, 12}
			result := search(input, 2)
			expect := -1
			convey.So(result, convey.ShouldResemble, expect)
		})

	})
}
