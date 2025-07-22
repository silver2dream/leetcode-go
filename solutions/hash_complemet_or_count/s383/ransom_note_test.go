package test

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestValidAnagram(t *testing.T) {
	convey.Convey("TestValidAnagram", t, func() {
		convey.Convey("Case1", func() {
			s := "a"
			t := "b"
			result := canConstruct(s, t)
			expect := false
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case2", func() {
			s := "aa"
			t := "ab"
			result := canConstruct(s, t)
			expect := false
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case3", func() {
			s := "aa"
			t := "aab"
			result := canConstruct(s, t)
			expect := true
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case4", func() {
			s := "abbbc"
			t := "aab"
			result := canConstruct(s, t)
			expect := false
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case5", func() {
			s := "bg"
			t := "efjbdfbdgfjhhaiigfhbaejahgfbbgbjagbddfgdiaigdadhcfcj"
			result := canConstruct(s, t)
			expect := true
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case6", func() {
			s := "abc"
			t := "ab"
			result := canConstruct(s, t)
			expect := false
			convey.So(result, convey.ShouldResemble, expect)
		})
	})
}
