package linkedList

import (
	. "github.com/crusj/goes/print"
)

type CircleDuLinkedList struct {
	Value *string
	Pre   *CircleDuLinkedList
	Next  *CircleDuLinkedList
}

func (*CircleDuLinkedList) Push(head *CircleDuLinkedList, value *string) {
	node := &CircleDuLinkedList{value, nil, nil}
	if nil == head {
		node.Pre, node.Next = node, node
	} else {
		pre := head
		next := head.Next
		for next != head {
			pre = next
			next = next.Next
		}
		pre.Next, node.Pre, node.Next, head.Pre = node, pre, head, node
	}

}

func (*CircleDuLinkedList) Traverse(head *CircleDuLinkedList) {
	p := head
	for p.Next != head {
		P(*p.Value)
		p = p.Next
	}
	P(*p.Value)
}
func (*CircleDuLinkedList) Unshift(head *CircleDuLinkedList, value *string) *CircleDuLinkedList {
	node := &CircleDuLinkedList{value, nil, nil}
	if head == nil {
		node.Next, node.Pre = node, node
		return node
	} else {
		head.Pre.Next = node
		node.Pre = head.Pre
		node.Next = head
		head.Pre = node
		return node

	}

}
func (*CircleDuLinkedList) Pop(head *CircleDuLinkedList) *CircleDuLinkedList {
	if head == nil || head.Next == head {
		return nil
	} else {
		pre := head
		next := head.Next

		for next.Next != head {
			pre = next
			next = next.Next
		}

		pre.Next = head
		head.Pre = pre
		return head
	}
}

func (*CircleDuLinkedList) Shift(head *CircleDuLinkedList) *CircleDuLinkedList {
	if head.Next == head {
		return nil
	} else {
		head.Pre.Next = head.Next
		head.Next.Pre = head.Pre
		return head.Next
	}
}
func (*CircleDuLinkedList) Insert(head *CircleDuLinkedList, index int, value *string) *CircleDuLinkedList {
	node := &CircleDuLinkedList{value, nil, nil}
	if head == nil {
		return node
	} else {
		if index == 1 {
			head.Pre.Next = node
			node.Pre = head.Pre

			node.Next = head
			head.Pre = node
			return node
		} else {
			step := 2
			pre := head
			next := head.Next
			for next != head && index != step {
				pre = next
				next = next.Next
				step++
			}
			if next != head {
				pre.Next = node
				node.Pre = pre
				node.Next = next
				next.Pre = node
			} else {
				node.Pre = next
				node.Next = head
				next.Next = node
				head.Pre = node
			}
			return head

		}
	}
}
func (*CircleDuLinkedList) Delete(head *CircleDuLinkedList, index int) *CircleDuLinkedList {
	if nil == head || head.Next == head {
		return nil
	} else {
		if 1 == index {
			head.Pre.Next = head.Next
			head.Next.Pre = head.Pre
			return head.Next
		} else {
			step := 2
			pre := head
			next := head.Next
			for next != head && index != step {
				pre = next
				next = next.Next
				step++
			}
			if next == head {
			} else {
				pre.Next = next.Next
				next.Next.Pre = next.Pre
			}
			return head
		}
	}
}
