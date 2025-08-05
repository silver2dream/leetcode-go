package s338

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCountingBits(t *testing.T) {
	convey.Convey("CountingBits", t, func() {
		convey.Convey("Case1", func() {
			s := 5
			result := countBits(s)
			expect := []int{0, 1, 1, 2, 1, 2}
			convey.So(result, convey.ShouldResemble, expect)
		})

	})
}
