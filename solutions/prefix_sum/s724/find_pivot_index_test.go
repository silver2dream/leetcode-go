package test

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFindPivotIndex(t *testing.T) {
	convey.Convey("FindPivotIndex", t, func() {
		convey.Convey("Case1", func() {
			s := []int{2, 1, -1}
			result := pivotIndex(s)
			expect := 0
			convey.So(result, convey.ShouldResemble, expect)
		})

	})
}
