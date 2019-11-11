package leetcode

import "github.com/crusj/goes/linkedList"

/**
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

示例 1:

输入: 1->1->2
输出: 1->2

示例 2:

输入: 1->1->2->3->3
输出: 1->2->3
*/
type Question83 struct {
	Input *linkedList.SingleNode
}

func (it *Question83) Deal() {
	head := it.Input
	t := head
	for t != nil {
		if t.Next != nil {
			if t.Value == t.Next.Value {
				t.Next = t.Next.Next
			}
		}
		t = t.Next
	}
	it.Input.Traverse(it.Input)
}
