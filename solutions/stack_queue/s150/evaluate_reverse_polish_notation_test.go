package s150

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestEvalRPN(t *testing.T) {
	convey.Convey("EvalRPN", t, func() {
		convey.Convey("Case1", func() {
			nums := []string{"2", "1", "+", "3", "*"}
			result := evalRPN(nums)
			expect := 9
			convey.So(result, convey.ShouldResemble, expect)
		})

	})
}
