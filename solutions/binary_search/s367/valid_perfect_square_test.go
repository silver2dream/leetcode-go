package s367

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIsPerfectSquare(t *testing.T) {
	convey.Convey("IsPerfectSquare", t, func() {
		//convey.Convey("Case1", func() {
		//	input := 17
		//	result := isPerfectSquare(input)
		//	expect := false
		//	convey.So(result, convey.ShouldResemble, expect)
		//})
		//
		//convey.Convey("Case2", func() {
		//	input := 100
		//	result := isPerfectSquare(input)
		//	expect := true
		//	convey.So(result, convey.ShouldResemble, expect)
		//})
		//
		//convey.Convey("Case3", func() {
		//	input := 4
		//	result := isPerfectSquare(input)
		//	expect := true
		//	convey.So(result, convey.ShouldResemble, expect)
		//})

		convey.Convey("Case4", func() {
			input := 100000001
			result := isPerfectSquare(input)
			expect := false
			convey.So(result, convey.ShouldResemble, expect)
		})

	})
}
