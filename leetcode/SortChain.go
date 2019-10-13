package leetcode

/**
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

**输入：**1->2->4, 1->3->4
**输出：**1->1->2->3->4->4
*/

type Chain struct {
	Value int
	Next  *Chain
}

type SortChain struct {
	ChainOne *Chain
	ChainTwo *Chain
	OutPut   *Chain
}

func (it *SortChain) Deal() {
	point2 := it.ChainTwo
	for point2 != nil {
		point1 := it.ChainOne
		for point1 != nil {
			{
				if point2.Value <= point1.Value {
					point := &Chain{point2.Value, nil}
					point.Next = point1
					it.ChainOne = point
				} else if point1.Value < point2.Value && point1.Next == nil {
					point1.Next = point2
					return
				} else if point1.Value < point2.Value && point1.Next.Value >= point2.Value {
					point := &Chain{point2.Value, nil}
					point.Next = point1.Next
					point1.Next = point
					break
				}
			}
			point1 = point1.Next
		}
		point2 = point2.Next
	}
}

//创建链表
func (it *SortChain) CreateChain(arr []int) *Chain {
	if len(arr) == 0 {
		return nil
	}
	var head *Chain
	tail := head
	for _, v := range arr {
		if head == nil {
			head = &Chain{v, nil}
			tail = head
		} else {
			tail.Next = &Chain{v, nil}
			tail = tail.Next
		}
	}
	return head
}
