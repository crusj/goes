package linkedList

import (
	. "goes/print"
)

type CircleSingleList struct {
	Value string
	Next  *CircleSingleList
}

//增加
func (*CircleSingleList) Push(head *CircleSingleList, value string) {
	node := &CircleSingleList{value, nil}
	if nil != head {
		step := head
		for head != step.Next && nil != step.Next {
			step = step.Next
		}
		step.Next, node.Next = node, head
	}
}

//遍历
func (*CircleSingleList) Traverse(head *CircleSingleList) {
	if nil != head {
		P(head.Value, "=>")
		step := head.Next
		for step != head {
			P(step.Value, "=>")
			step = step.Next
		}
	}
}

//头部插入一个元素
func (*CircleSingleList) Unshift(head *CircleSingleList, value string) *CircleSingleList {
	if nil != head {
		node := &CircleSingleList{value, nil}
		tail := head
		for tail.Next != head {
			tail = tail.Next
		}
		tail.Next = node
		node.Next = head
		return node
	} else {
		return nil
	}

}

//删除
func (*CircleSingleList) Delete(head *CircleSingleList, index int) *CircleSingleList {
	if nil == head {
		return nil
	} else if index == 1 {
		tail := head
		for tail.Next != head {
			tail = tail.Next
		}
		tail.Next = head.Next
		return head.Next
	} else {
		step := 2
		pre := head
		pointer := head.Next
		for pointer != head && index != step {
			step++
			pre = pointer
			pointer = pointer.Next
		}
		if index == step {
			P("find it")
			pre.Next = pointer.Next
		} else {
			P("not find it")
		}
		return head
	}
}
