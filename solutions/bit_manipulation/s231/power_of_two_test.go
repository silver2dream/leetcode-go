package s231

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNumberOf1Bits(t *testing.T) {
	convey.Convey("NumberOf1Bits", t, func() {
		convey.Convey("Case1", func() {
			s := 16
			result := isPowerOfTwo(s)
			expect := true
			convey.So(result, convey.ShouldResemble, expect)
		})

	})
}
