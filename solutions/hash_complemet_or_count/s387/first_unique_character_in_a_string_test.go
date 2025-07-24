package test

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFirstUniqueCharInAString(t *testing.T) {
	convey.Convey("TestFirstUniqueCharInAString", t, func() {
		convey.Convey("Case1", func() {
			s := "leetcode"
			result := firstUniqChar(s)
			expect := 0
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case2", func() {
			s := "loveleetcode"
			result := firstUniqChar(s)
			expect := 2
			convey.So(result, convey.ShouldResemble, expect)
		})

	})
}
