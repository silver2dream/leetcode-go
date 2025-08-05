package s136

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	convey.Convey("SingleNumber", t, func() {
		convey.Convey("Case1", func() {
			s := []int{2, 2, 1}
			result := singleNumber(s)
			expect := 1
			convey.So(result, convey.ShouldResemble, expect)
		})

	})
}
