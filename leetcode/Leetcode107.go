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
	Queue [][]*tree.Node
}

func (it *Question107) Deal() {
	it.Queue = append(it.Queue, []*tree.Node{it.Input})

	for i := 0; i <= len(it.Queue)-1; i++ { //第一层
		if len(it.Queue[i]) > 0{
			it.Queue = append(it.Queue, []*tree.Node{}) //增加一层
		}
		for j := 0; j <= len(it.Queue[i])-1; j++ {
			if it.Queue[i][j].LChild != nil {
				it.Queue[i+1] = append(it.Queue[i+1], it.Queue[i][j].LChild)
			}
			if it.Queue[i][j].RChild != nil {
				it.Queue[i+1] = append(it.Queue[i+1], it.Queue[i][j].RChild)
			}
		}
	}
}
func (it *Question107) String() string {
	rsl := ""
	for i := len(it.Queue) - 1; i >= 0; i-- {
		for j := 0; j < len(it.Queue[i]); j++ {
			rsl += fmt.Sprintf("%s ", it.Queue[i][j].Value)
		}
	}
	return rsl
}
