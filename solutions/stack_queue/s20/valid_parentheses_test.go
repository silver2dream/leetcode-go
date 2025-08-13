package s20

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIsValid(t *testing.T) {
	convey.Convey("isValid", t, func() {
		convey.Convey("Case1", func() {
			nums := "()[]{}"
			result := isValid(nums)
			expect := true
			convey.So(result, convey.ShouldResemble, expect)
		})

	})
}
