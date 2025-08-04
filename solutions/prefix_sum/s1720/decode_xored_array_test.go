package s1720

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDecodeXoRed_Array(t *testing.T) {
	convey.Convey("DecodeXoRed_Array", t, func() {
		convey.Convey("Case1", func() {
			s := []int{1, 2, 3}
			first := 1
			result := decode(s, first)
			expect := []int{1, 0, 2, 1}
			convey.So(result, convey.ShouldResemble, expect)
		})

	})
}
