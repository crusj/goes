package leetcode

import (
	"fmt"
	"github.com/crusj/goes/tree"
)

/**
给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

例如：
给定二叉树 [3,9,20,null,null,15,7] ,

3
/
9 20
/
15 7

返回其自底向上的层次遍历为：

[
[15,7],
[9,20],
[3]
]
*/
type Question107 struct {
	Input *tree.Node
	Queue []*tree.Node
}

func (it *Question107) Deal() {
	it.Queue = append(it.Queue, it.Input)

	for i := 0; i <= len(it.Queue)-1; i++ {
		if it.Queue[i].LChild != nil {
			it.Queue = append(it.Queue, it.Queue[i].LChild)
		}
		if it.Queue[i].RChild != nil {
			it.Queue = append(it.Queue, it.Queue[i].RChild)
		}
	}
}
func (it *Question107) String() string {
	rsl := ""
	for i := len(it.Queue) - 1; i >= 0; i-- {
		rsl += fmt.Sprintf("%s ", it.Queue[i].Value)
	}
	return rsl
}
