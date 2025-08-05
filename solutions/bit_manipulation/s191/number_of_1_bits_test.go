package s191

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNumberOf1Bits(t *testing.T) {
	convey.Convey("NumberOf1Bits", t, func() {
		convey.Convey("Case1", func() {
			s := 11
			result := hammingWeight(s)
			expect := 3
			convey.So(result, convey.ShouldResemble, expect)
		})

	})
}
