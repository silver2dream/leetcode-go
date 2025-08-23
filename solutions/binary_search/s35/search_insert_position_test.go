package s35

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	convey.Convey("SearchInsert", t, func() {
		//convey.Convey("Case1", func() {
		//	input := []int{1, 3, 5, 6}
		//	result := searchInsert(input, 5)
		//	expect := 2
		//	convey.So(result, convey.ShouldResemble, expect)
		//})
		//convey.Convey("Case2", func() {
		//	input := []int{1, 3, 5, 6}
		//	result := searchInsert(input, 2)
		//	expect := 1
		//	convey.So(result, convey.ShouldResemble, expect)
		//})
		convey.Convey("Case3", func() {
			input := []int{1, 3, 5, 6}
			result := searchInsert(input, 7)
			expect := 4
			convey.So(result, convey.ShouldResemble, expect)
		})
		//convey.Convey("Case4", func() {
		//	input := []int{1, 3, 5, 6}
		//	result := searchInsert(input, 0)
		//	expect := 0
		//	convey.So(result, convey.ShouldResemble, expect)
		//})
		//
		//convey.Convey("Case5", func() {
		//	input := []int{1}
		//	result := searchInsert(input, 1)
		//	expect := 0
		//	convey.So(result, convey.ShouldResemble, expect)
		//})

	})
}
