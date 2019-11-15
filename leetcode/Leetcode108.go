package leetcode

import "github.com/crusj/goes/tree"

/**
将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树 每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:

给定有序数组: [-10,-3,0,5,9],

一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

  0
 / \
-3 9
/ /
-10 5
*/
type Question108 struct {
	Input []int
	Tree  *tree.Node
}

func (it *Question108) Deal() {
	it.Tree = it.Generate(it.Input)
	t := &tree.BinaryTree{}
	t.PreOrder(it.Tree)
}
func (it *Question108) Generate(arr []int) *tree.Node {
	arrLen := len(arr)
	if arrLen == 0 {
		return nil
	}
	if arrLen == 1 {
		return &tree.Node{arr[0], nil, nil}
	}
	node := &tree.Node{arr[arrLen/2], nil, nil}
	node.LChild = it.Generate(arr[0 : arrLen/2])
	if arrLen == 2 {
		node.RChild = nil
	}else{
		node.RChild = it.Generate(arr[arrLen/2 + 1:])
	}
	return node
}
