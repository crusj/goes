package linkedList

import (
	. "goes/print"
)

type DuLinkedList struct {
	Value string
	Pre   *DuLinkedList
	Next  *DuLinkedList
}

func (*DuLinkedList) Push(head *DuLinkedList, value string) {
	node := &DuLinkedList{value, nil, nil}
	if nil == head {
		head = node
	} else {
		for nil != head.Next {
			head = head.Next
		}
		head.Next = node
		node.Pre = head
	}
}

//遍历
func (*DuLinkedList) Traverse(head *DuLinkedList) {
	for nil != head {
		P(head.Value, "=>")
		head = head.Next
	}
}

//
func (*DuLinkedList) Unshift(head *DuLinkedList, value string) *DuLinkedList {
	node := &DuLinkedList{value, nil, nil}
	if nil != head {
		node.Next = head
		head.Pre = node
	}
	return node
}

func (*DuLinkedList) Insert(head *DuLinkedList, index int, value string) *DuLinkedList {
	node := &DuLinkedList{value, nil, nil}
	if nil == head {
		return node
	} else {
		if index == 1 {
			node.Next = head
			head.Pre = node
			return node
		} else {
			pre := head
			next := head.Next
			step := 2
			for next != nil && step != index {
				pre = next
				next = next.Next
				step++
			}
			if nil != next {
				pre.Next = node
				node.Pre = pre
				node.Next = next
				next.Pre = node
			}
			return head

		}
	}
}
func (*DuLinkedList) Shift(head *DuLinkedList) *DuLinkedList {
	next := head.Next
	next.Pre = nil
	return next
}
func (*DuLinkedList) Pop(head *DuLinkedList) {
	var pre *DuLinkedList
	for head.Next != nil {
		pre = head
		head = head.Next
	}
	pre.Next = nil
	head.Pre = nil
}
func (*DuLinkedList) Delete(head *DuLinkedList, index int) *DuLinkedList {
	if nil == head {
		return nil
	} else {
		if index == 1 {
			if head.Next != nil {
				head.Pre = nil
			}
			return head.Next

		} else {
			pre := head
			next := head.Next
			step := 2
			for next != nil && step != index {
				pre = next
				next = next.Next
				step++
			}
			if nil != next && nil != next.Next {
				pre.Next = next.Next
			}
			return head
		}
	}
}
