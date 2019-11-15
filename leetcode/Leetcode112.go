package leetcode

import (
	"fmt"
	"github.com/crusj/goes/tree"
	"strconv"
)

/**
给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。

说明: 叶子节点是指没有子节点的节点。

示例:
给定如下二叉树，以及目标和 sum = 22 ，

          **5**
         / \
        **4 **  8
       /   / \
      **11 ** 13  4
     /  \      \
    7    **2**      1
返回 true , 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2 。
*/
type Question112 struct {
	Node *tree.Node
	Sum  int
	Path [][]int
}

func (it *Question112) Deal() {
	queue := make([]*tree.Node, 0) //节点队列
	point := it.Node

	queue = append(queue,point)           //根节点入队
	pc := make(map[int]int, 0) //路径点关联表

	v, _ := strconv.Atoi(point.Value.(string)) //根节点加关联表
	pc[v] = 0

	for i := 0; i <= len(queue)-1; i++ {
		parent, _ := strconv.Atoi(queue[i].Value.(string))
		//叶子节点
		if queue[i].LChild == nil && queue[i].RChild == nil {
			if path := it.CheckSum(pc, parent); path != nil {
				it.Path = append(it.Path, path)
			}
		} else {
			if queue[i].LChild != nil {
				queue = append(queue, queue[i].LChild) //左孩子入队
				t, _ := strconv.Atoi(queue[i].LChild.Value.(string))
				pc[t] = parent //左孩子上级映射
			}
			if queue[i].RChild != nil {
				queue = append(queue, queue[i].RChild)                 //右孩子入队
				t, _ := strconv.Atoi(queue[i].RChild.Value.(string))
				pc[t] = parent //左孩子上级映射
			}
		}
	}
}

//检查路径总和
func (it *Question112) CheckSum(set map[int]int, index int) (q []int) {
	q = append(q, index)
	sum := index
	for {
		if _, ok := set[index]; ok {
			q = append(q, set[index])
			index = set[index]
			sum += index
		} else {
			break
		}
	}
	if sum == it.Sum {
		return
	} else {
		return nil
	}
}
func (it *Question112) String() string {
	str := ""
	for i, cate := range it.Path {
		str += fmt.Sprintf("方案%d：", i)
		for j := len(cate) - 1; j >= 0; j-- {
			str += strconv.Itoa(cate[j])+"=>"
		}
		str += "\n"
	}
	return fmt.Sprintf("输入总和为%d,路径方案为：\n%s", it.Sum, str)
}
