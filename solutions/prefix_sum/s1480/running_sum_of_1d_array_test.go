package test

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRunningSumOf1dArray(t *testing.T) {
	convey.Convey("RunningSumOf1dArray", t, func() {
		convey.Convey("Case1", func() {
			s := []int{1, 2, 3, 4}
			result := runningSum(s)
			expect := []int{1, 3, 6, 10}
			convey.So(result, convey.ShouldResemble, expect)
		})

	})
}
