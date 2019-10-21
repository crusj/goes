package leetcode

import "fmt"

/**
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

你可以假设数组中无重复元素。

示例 1:

输入: [1,3,5,6], 5
输出: 2

示例 2:

输入: [1,3,5,6], 2
输出: 1

示例 3:

输入: [1,3,5,6], 7
输出: 4

示例 4:

输入: [1,3,5,6], 0
输出: 0
*/
type SearchSortArray struct {
	Input  []int
	Search int
	Index  int
}
type Node struct {
	Value int
	Index int
}

func (it *SearchSortArray) Deal() {
	nodes := make([]*Node, 0)
	for k, v := range it.Input {
		nodes = append(nodes, &Node{v, k})
	}
	it.Index = it.binarySearch(nodes)
}
func (it *SearchSortArray) binarySearch(nodes []*Node) int {
	len := len(nodes)
	if len == 1 {
		if nodes[0].Value == it.Search {
			return nodes[0].Index
		} else if nodes[0].Value > it.Search {
			return nodes[0].Index
		} else {
			return nodes[0].Index + 1
		}
	}
	half := len / 2
	if nodes[half].Value == it.Search {
		return nodes[half].Index
	} else if nodes[half].Value < it.Search {
		if len-1 == half {
			return len
		} else {
			return it.binarySearch(nodes[half+1:])
		}
	} else {
		return it.binarySearch(nodes[0:half])
	}
}
func (it *SearchSortArray) String() string {
	return fmt.Sprintf("数组为%v,查找的值为%d,下标为%d", it.Input, it.Search, it.Index)
}
