package leetcode

import (
	"fmt"
	"github.com/crusj/goes/linkedList"
)

/**
编写一个程序，找到两个单链表相交的起始节点。

例如，下面的两个链表 ：

A: a1 → a2
↘
c1 → c2 → c3
↗
B: b1 → b2 → b3

在节点 c1 开始相交。

注意：

如果两个链表没有交点，返回 null .
在返回结果后，两个链表仍须保持原有的结构。
可假定整个链表结构中没有循环。
程序尽量满足 O( n ) 时间复杂度，且仅用 O( 1 ) 内存。
致谢 :
特别感谢 @stellari 添加此问题并创建所有测试用例。
*/
type Question160 struct {
	LinkOne *linkedList.SingleNode
	LinkTwo *linkedList.SingleNode
	Output  string
}

func (it *Question160) Deal() {
	a, b := it.LinkOne, it.LinkTwo
	linkedToB, linkedToA := false, false
	for a != nil && b != nil {
		if a.Value == b.Value {
			it.Output = a.Value
			return
		}
		a, b = a.Next, b.Next
		if a == nil && linkedToB == false {
			a = it.LinkTwo
			linkedToB = true
		}
		if b == nil && linkedToA == false {
			b = it.LinkOne
			linkedToA = true
		}
	}
	it.Output = ""
}

func (it *Question160) String() string {
	return fmt.Sprintf("两条链表相交于%s", it.Output)
}
