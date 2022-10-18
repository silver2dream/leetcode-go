package test

import (
	"github.com/smartystreets/goconvey/convey"
	"main/lib/tree"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	convey.Convey("TestInorderTraversal", t, func() {
		convey.Convey("case1", func() {
			nums := []any{1, nil, 2, nil, nil, 3}
			tt := tree.NewListTree(nums)
			result := inorderTraversal(tt.GetRoot())
			expect := []int{1, 3, 2}
			convey.So(result, convey.ShouldResemble, expect)
		})

	})
}
