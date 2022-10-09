package test

import (
	"github.com/smartystreets/goconvey/convey"
	"main/lib/tree"
	"testing"
)

func TestFindMode(t *testing.T) {
	convey.Convey("TestFindMode", t, func() {
		//convey.Convey("Case1", func() {
		//	nums := []any{1, 2, 2}
		//	bt := tree.NewBinarySearchTree(nums)
		//	result := findMode(bt.GetRoot())
		//	expect := []int{2}
		//	convey.So(result, convey.ShouldResemble, expect)
		//})
		//
		//convey.Convey("Case2", func() {
		//	nums := []any{0}
		//	bt := tree.NewBinarySearchTree(nums)
		//	result := findMode(bt.GetRoot())
		//	expect := []int{0}
		//	convey.So(result, convey.ShouldResemble, expect)
		//})

		convey.Convey("Case3", func() {
			nums := []any{1, 2, 2, 3, 3}
			bt := tree.NewBinarySearchTree(nums)
			result := findMode(bt.GetRoot())
			expect := []int{2, 3}
			convey.So(result, convey.ShouldResemble, expect)
		})
	})

}
