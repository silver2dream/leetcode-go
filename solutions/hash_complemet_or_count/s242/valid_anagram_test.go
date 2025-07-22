package test

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestValidAnagram(t *testing.T) {
	convey.Convey("TestValidAnagram", t, func() {
		convey.Convey("Case1", func() {
			s := "anagram"
			t := "nagaram"
			result := isAnagram(s, t)
			expect := true
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case2", func() {
			s := "rat"
			t := "car"
			result := isAnagram(s, t)
			expect := false
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case3", func() {
			s := "a"
			t := "b"
			result := isAnagram(s, t)
			expect := false
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case4", func() {
			s := "a"
			t := "a"
			result := isAnagram(s, t)
			expect := true
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Case5", func() {
			s := "aa"
			t := "bb"
			result := isAnagram(s, t)
			expect := false
			convey.So(result, convey.ShouldResemble, expect)
		})
	})
}
