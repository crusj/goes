package leetcode

import "github.com/crusj/goes/tree"

/**
给定一个二叉树，检查它是否是镜像对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

1
/
2 2
/ \ /
3 4 4 3

但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

1
/
2 2
\
3 3

说明:

如果你可以运用递归和迭代两种方法解决这个问题，会很加分。
*/
type Question101 struct {
	Root *tree.Node
}

func (it *Question101) Deal() {
	println(it.mirror(it.Root.LChild, it.Root.RChild))
}
//是否镜像对称
func (it *Question101) mirror(m *tree.Node, n *tree.Node) bool {
	if m == nil && n == nil {
		return true
	} else if m == nil || n == nil {
		return false
	}
	if m.Value != n.Value {
		return false
	} else {
		return it.mirror(m.LChild, n.RChild) && it.mirror(m.RChild, n.LChild)
	}
}
