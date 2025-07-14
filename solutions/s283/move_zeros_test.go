package test

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMoveZeros(t *testing.T) {
	convey.Convey("TestMoveZeros", t, func() {
		convey.Convey("Case1", func() {
			input := []int{0, 3, 0, 7, 11}
			moveZeroes(input)
			expect := []int{3, 7, 11, 0, 0}
			convey.So(input, convey.ShouldResemble, expect)
		})

	})

}
