package s441

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestArrangeCoins(t *testing.T) {
	convey.Convey("ArrangeCoins", t, func() {
		convey.Convey("Case1", func() {
			input := 5
			result := arrangeCoins(input)
			expect := 2
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case2", func() {
			input := 8
			result := arrangeCoins(input)
			expect := 3
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case4", func() {
			input := 10
			result := arrangeCoins(input)
			expect := 4
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case5", func() {
			input := 3
			result := arrangeCoins(input)
			expect := 2
			convey.So(result, convey.ShouldResemble, expect)
		})

	})
}
