package leetcode

import (
	"fmt"
	"github.com/crusj/goes/linkedList"
)

/**
给定一个链表，判断链表中是否有环。

进阶：
你能否不使用额外空间解决此题？
*/
type Question141 struct {
	Head   *linkedList.SingleNode
	Output bool
}

func (it *Question141) Deal() {
}

//哈希表法
func (it *Question141) Way1() {
	set := make(map[*linkedList.SingleNode]Void)
	point := it.Head
	for point != nil {
		if _, ok := set[point]; ok {
			it.Output = true
			return
		} else {
			set[point] = Void{}
		}
		point = point.Next
	}
	it.Output = false
}

//快慢指针
func (it Question141) Way2() {
	slow := it.Head
	fast := it.Head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			it.Output = false
			break
		} else {
			slow = slow.Next
			fast = fast.Next.Next
		}
	}
	it.Output = true
}
func (it *Question141) String() string {
	return fmt.Sprintf("%b\n", it.Output)
}
