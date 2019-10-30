package linkedList

import (
	. "github.com/crusj/goes/print"
)

type SingleNode struct {
	Value string
	Next  *SingleNode
}

//追加节点
func (*SingleNode) Push(head *SingleNode, value string) {
	node := &SingleNode{value, nil}
	for head.Next != nil {
		head = head.Next
	}
	head.Next = node
}

//遍历
func (*SingleNode) Traverse(head *SingleNode) {
	for nil != head {
		P(head.Value, "=>")
		head = head.Next
	}
}

//表头插入
func (*SingleNode) Unshift(head *SingleNode, value string) *SingleNode {
	node := &SingleNode{value, nil}
	node.Next = head
	return node
}

//随机位置插入
func (*SingleNode) Rand(head *SingleNode, index int, value string) *SingleNode {
	node := &SingleNode{value, nil}
	if nil == head {
		return node
	} else {
		if 1 == index {
			node.Next = head
			return node
		} else {
			pre := head
			next := pre.Next
			step := 2

			for nil != next && step != index {
				step++
				pre = next
				next = next.Next
			}
			pre.Next = node
			if nil != next {
				node.Next = next
			}
			return head
		}
	}
}

//删除
func (*SingleNode) Delete(head *SingleNode, index int) *SingleNode {
	if index == 1 {
		return head.Next
	} else {
		step := 2
		pre := head
		next := pre.Next

		for nil != next {
			if step == index {
				pre.Next = next.Next
			}
			pre = next
			next = next.Next
			step++
		}
		return head
	}
}
