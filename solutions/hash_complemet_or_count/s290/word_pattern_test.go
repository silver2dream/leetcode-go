package test

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestWordPattern(t *testing.T) {
	convey.Convey("TestWordPattern", t, func() {
		convey.Convey("Case1", func() {
			s := "abba"
			t := "dog cat cat dog"
			result := wordPattern(s, t)
			expect := true
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case2", func() {
			s := "abba"
			t := "dog cat cat fish"
			result := wordPattern(s, t)
			expect := false
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case3", func() {
			s := "aaaa"
			t := "dog cat cat dog"
			result := wordPattern(s, t)
			expect := false
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case4", func() {
			s := "ac"
			t := "dog dog"
			result := wordPattern(s, t)
			expect := false
			convey.So(result, convey.ShouldResemble, expect)
		})

	})
}
