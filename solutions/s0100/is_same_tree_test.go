package test

import (
	"github.com/smartystreets/goconvey/convey"
	"main/lib/tree"
	"testing"
)

func TestSameTree(t *testing.T) {
	convey.Convey("TestSameTree", t, func() {
		convey.Convey("IsSame", func() {
			p := []any{2, 1, 3}
			q := []any{2, 1, 3}
			pt := tree.NewListTree(p)
			qt := tree.NewListTree(q)
			result := isSameTree(pt.GetRoot(), qt.GetRoot())
			expect := true
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("NotSame", func() {
			p := []any{1, 2}
			q := []any{1, nil, 2}
			pt := tree.NewListTree(p)
			qt := tree.NewListTree(q)
			result := isSameTree(pt.GetRoot(), qt.GetRoot())
			expect := false
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("Empty", func() {
			p := []any{}
			q := []any{}
			pt := tree.NewListTree(p)
			qt := tree.NewListTree(q)
			result := isSameTree(pt.GetRoot(), qt.GetRoot())
			expect := true
			convey.So(result, convey.ShouldResemble, expect)
		})

		convey.Convey("OnlyRoot", func() {
			p := []any{1}
			q := []any{1}
			pt := tree.NewListTree(p)
			qt := tree.NewListTree(q)
			result := isSameTree(pt.GetRoot(), qt.GetRoot())
			expect := true
			convey.So(result, convey.ShouldResemble, expect)
		})

	})
}
