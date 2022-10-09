package test

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestEatingSpeed(t *testing.T) {
	convey.Convey("TestEatingSpeed", t, func() {
		convey.Convey("Case1", func() {
			piles := []int{3, 6, 7, 11}
			h := 8
			result := minEatingSpeed(piles, h)
			expect := 4
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case2", func() {
			piles := []int{30, 11, 23, 4, 20}
			h := 5
			result := minEatingSpeed(piles, h)
			expect := 30
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case3", func() {
			piles := []int{30, 11, 23, 4, 20}
			h := 6
			result := minEatingSpeed(piles, h)
			expect := 23
			convey.So(result, convey.ShouldResemble, expect)
		})

	})
}
