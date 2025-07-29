package test

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMinimumValueToGetPositiveStepByStepSum(t *testing.T) {
	convey.Convey("MinimumValueToGetPositiveStepByStepSum", t, func() {
		convey.Convey("Case1", func() {
			s := []int{2, 3, 5, -5, -1}
			result := minStartValue(s)
			expect := 1
			convey.So(result, convey.ShouldResemble, expect)
		})

	})
}
