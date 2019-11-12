package leetcode

import "github.com/crusj/goes/tree"
/**
检查两棵二叉树是否相等,检查值，和左右孩子节点
 */
type Question100 struct {
}

func (it *Question100) Deal(one *tree.Node, two *tree.Node) bool {
	if one == nil && two == nil {
		return true
	}
	if one == nil || two == nil {
		return false
	}
	if one.Value != two.Value {
		return false
	}
	return it.Deal(one.LChild, two.LChild) && it.Deal(one.RChild, two.RChild)
}
