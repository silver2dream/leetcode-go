package s739

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	convey.Convey("DailyTemperatures", t, func() {
		convey.Convey("Case1", func() {
			nums := []int{73, 74, 75, 71, 69, 72, 76, 73}
			result := dailyTemperatures(nums)
			expect := []int{1, 1, 4, 2, 1, 1, 0, 0}
			convey.So(result, convey.ShouldResemble, expect)
		})

	})
}
