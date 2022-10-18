package test

import (
	"github.com/smartystreets/goconvey/convey"
	list2 "main/lib/linked-list"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	convey.Convey("TestMergeKLists", t, func() {
		convey.Convey("k=3", func() {
			a := list2.NewList()
			a.Add(1)
			a.Add(4)
			a.Add(5)
			b := list2.NewList()
			b.Add(1)
			b.Add(3)
			b.Add(4)
			c := list2.NewList()
			c.Add(2)
			c.Add(6)
			list := []*ListNode{a.GetRoot(), b.GetRoot(), c.GetRoot()}
			result := mergeKLists(list)
			expect := []int{1, 1, 2, 3, 4, 4, 5, 6}
			convey.So(result.CovertArray(), convey.ShouldResemble, expect)
		})

		convey.Convey("empty", func() {
			list := []*ListNode{}
			result := mergeKLists(list)
			expect := []int{}
			convey.So(result.CovertArray(), convey.ShouldResemble, expect)
		})

		convey.Convey("k=1", func() {
			a := list2.NewList()
			list := []*ListNode{a.GetRoot()}
			result := mergeKLists(list)
			expect := []int{}
			convey.So(result.CovertArray(), convey.ShouldResemble, expect)
		})

	})
}
