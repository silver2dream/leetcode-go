package test

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestFindTheDifference(t *testing.T) {
	convey.Convey("FindTheDifference", t, func() {
		convey.Convey("Case1", func() {
			s := "abcd"
			ts := "abcde"
			result := findTheDifference(s, ts)
			expect := "e"
			convey.So(string(result), convey.ShouldResemble, expect)
		})

	})
}
