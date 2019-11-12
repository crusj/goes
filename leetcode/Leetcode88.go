package leetcode

import (
	"github.com/crusj/goes/linkedList"
	"strconv"
)

/*

说明:

初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n ）来保存 nums2 中的元素。
示例:

输入:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6], n = 3

输出: [1,2,2,3,5,6]

解决方案：考虑到可能移动次数过多，链表，两个数组都只需要遍历一次就好，将nums1存储为链表可能还需要遍历一次，最后写入还需要遍历一次
*/
type Question88 struct {
	ArrOne  []int
	ArrTwo  []int
	Combine []int
}

func (it *Question88) Deal() {
	if len(it.ArrOne) > 0 {
		root := it.CreateLinkList()
		println("创建的双向链表为：")
		root.Traverse(root)
		p := root
		for _, v := range it.ArrTwo {
			for p != nil {
				tmp := &linkedList.DuLinkedList{Value: strconv.Itoa(v)}
				pv, _ := strconv.Atoi(p.Value) //当前节点值
				if pv > v {
					if p.Pre == nil {
						tmp.Next = p
						p.Pre = tmp
						root = tmp
					} else {
						tmp.Pre = p.Pre
						tmp.Next = p
						p.Pre.Next = tmp
						p.Pre = tmp
					}
					break
				} else if pv == v {
					if p.Next == nil {
						tmp.Pre = p
						p.Next = tmp
					} else {
						tmp.Next = p.Next
						tmp.Pre = p
						p.Next.Pre = tmp
						p.Next = tmp
					}
					p = p.Next
					break
				} else if pv < v {
					if p.Next == nil {
						tmp.Pre = p
						p.Next = tmp
					}
					p = p.Next
				}
			}
		}
		println("插入后链表为：")
		root.Traverse(root)
	} else {
		it.Combine = it.ArrTwo
	}
}

//创建链表
func (it *Question88) CreateLinkList() *linkedList.DuLinkedList {
	head := &linkedList.DuLinkedList{Value: strconv.Itoa(it.ArrOne[0])}
	point := head
	for i := 1; i < len(it.ArrOne); i++ {
		point.Push(point, strconv.Itoa(it.ArrOne[i]))
	}
	return head
}
