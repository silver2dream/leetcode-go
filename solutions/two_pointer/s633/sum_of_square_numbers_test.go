package s633

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestJudgeSquareSum(t *testing.T) {
	convey.Convey("JudgeSquareSum", t, func() {
		convey.Convey("Case1", func() {
			input := 3
			result := judgeSquareSum(input)
			expect := false
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case2", func() {
			input := 11
			result := judgeSquareSum(input)
			expect := false
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case3", func() {
			input := 25
			result := judgeSquareSum(input)
			expect := true
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case4", func() {
			input := 5
			result := judgeSquareSum(input)
			expect := true
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case6", func() {
			input := 6
			result := judgeSquareSum(input)
			expect := false
			convey.So(result, convey.ShouldResemble, expect)
		})

	})
}
