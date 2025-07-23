package test

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIsomorphicStrings(t *testing.T) {
	convey.Convey("IsomorphicStrings", t, func() {
		convey.Convey("Case1", func() {
			s := "egg"
			t := "add"
			result := isIsomorphic(s, t)
			expect := true
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case2", func() {
			s := "e"
			t := "a"
			result := isIsomorphic(s, t)
			expect := true
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case3", func() {
			s := "ab"
			t := "ba"
			result := isIsomorphic(s, t)
			expect := true
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case4", func() {
			s := "egga"
			t := "addg"
			result := isIsomorphic(s, t)
			expect := true
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case5", func() {
			s := "badc"
			t := "baba"
			result := isIsomorphic(s, t)
			expect := false
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case6", func() {
			s := "foo"
			t := "bar"
			result := isIsomorphic(s, t)
			expect := false
			convey.So(result, convey.ShouldResemble, expect)
		})

	})
}
