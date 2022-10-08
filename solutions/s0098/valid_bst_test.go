package test

import (
	"github.com/smartystreets/goconvey/convey"
	"main/lib/tree"
	"testing"
)

func TestValidBst(t *testing.T) {
	convey.Convey("TestValidBst", t, func() {
		convey.Convey("IsValid", func() {
			nums := []any{2, 1, 3}
			bt := tree.NewListTree(nums)
			result := isValidBST(bt.GetRoot())
			expect := true
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("IsNotValid", func() {
			nums := []any{5, 1, 4, nil, nil, 3, 6}
			bt := tree.NewListTree(nums)
			result := isValidBST(bt.GetRoot())
			expect := false
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("SameNode", func() {
			nums := []any{2, 2, 2}
			bt := tree.NewListTree(nums)
			result := isValidBST(bt.GetRoot())
			expect := false
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("SameTwoNode", func() {
			nums := []any{1, 1}
			bt := tree.NewListTree(nums)
			result := isValidBST(bt.GetRoot())
			expect := false
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("OnlyRoot", func() {
			nums := []any{0}
			bt := tree.NewListTree(nums)
			result := isValidBST(bt.GetRoot())
			expect := true
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("RightSideNegative", func() {
			nums := []any{0, -1}
			bt := tree.NewListTree(nums)
			result := isValidBST(bt.GetRoot())
			expect := true
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Negative", func() {
			nums := []any{-1, -9}
			bt := tree.NewListTree(nums)
			result := isValidBST(bt.GetRoot())
			expect := true
			convey.So(result, convey.ShouldResemble, expect)
		})
	})
}
