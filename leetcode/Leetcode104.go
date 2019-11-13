package leetcode

import (
	"fmt"
	"github.com/crusj/goes/tree"
	"math"
)

/**
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7] ，

3
/
9 20
/
15 7
*/
type Question104 struct {
	Tree     *tree.Node
	MaxDepth int
}

func (it *Question104) Deal() {
	it.MaxDepth = int(it.maxDepth(it.Tree, 0))
}
func (it *Question104) maxDepth(node *tree.Node, depth float64) float64 {
	if node == nil {
		return depth
	}
	return math.Max(
		it.maxDepth(node.LChild, float64(depth+1)),
		it.maxDepth(node.RChild, float64(depth+1)),
	)

}
func (it *Question104) String() string {
	return fmt.Sprintf("二叉树最大深度为%d\n", it.MaxDepth)
}
